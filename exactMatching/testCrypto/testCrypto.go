// This file is for the testing of smart contract for experiment

package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/clearmatics/bn256"
)

// the key set of CA
type KeySetOfCA struct {
	X  *big.Int  // X is private
	gx *bn256.G2 // gx is for server
}

// the key set of server
type KeySetOfServer struct {
	Kb      *big.Int
	beta    *big.Int
	S       *big.Int
	g1Sbeta *bn256.G1
}

// the key set of user
type KeySetOfUser struct {
	Ku    *big.Int
	alpha *big.Int
	gx    *bn256.G2
}

type VocabItem struct {
	key string
	id  int
}

type CipherText struct {
	X  *bn256.G1
	Y  *bn256.G2
	id int
}

var (
	caKeySet     KeySetOfCA
	serverKeySet KeySetOfServer
	userKeySet   KeySetOfUser
	taskSize     int          = 800
	cipher       []CipherText = make([]CipherText, taskSize)
	Vocab        [2000]VocabItem
)

// keyGeneration(): Generate the key set in CA
func keyGeneration() {
	serverKeySet.Kb, _ = rand.Int(rand.Reader, bn256.Order) // Kb: one key for server
	userKeySet.Ku, _ = rand.Int(rand.Reader, bn256.Order)   // Ku: one key for user
	caKeySet.X = big.NewInt(0)                              // X: the key held by CA
	caKeySet.X.Add(serverKeySet.Kb, userKeySet.Ku)          // X=Kb+Ku. for all the users, X is the same

	caKeySet.gx = new(bn256.G2).ScalarBaseMult(caKeySet.X)   // gx: g^x is used for the encryption in user
	userKeySet.gx = new(bn256.G2).ScalarBaseMult(caKeySet.X) // gx: g^x is used for the encryption in user

	userKeySet.alpha, _ = rand.Int(rand.Reader, bn256.Order)  // alpha: the encryption key for user in encryption
	serverKeySet.beta, _ = rand.Int(rand.Reader, bn256.Order) // beta: the encryption key for server in on-chain encryption

	serverKeySet.S, _ = rand.Int(rand.Reader, bn256.Order) // S: the encryption key for server in on-chain encryption

	serverKeySet.g1Sbeta = new(bn256.G1).ScalarBaseMult(serverKeySet.S)
	serverKeySet.g1Sbeta.ScalarMult(serverKeySet.g1Sbeta, serverKeySet.beta) //g1^{S*beta}: the encryption key for server in on-chain encryption
}

// userEncryption(w string) (*bn256.G2 C1, *bn256.G2 C2): Input a keyword(string w) and return a ciphertext tuple (C1, C2)
func userEncryption(w string) (*bn256.G2, *bn256.G2) {
	kwByte := []byte(w)                                    // the byte array of keyword w
	hashedKw := sha256.Sum256(kwByte)                      // hashed the keyword w
	numHashedKw := new(big.Int).SetBytes(hashedKw[:])      // tranform the hashed byte array into big.num
	ru, _ := rand.Int(rand.Reader, bn256.Order)            // r_u: the nonce of this encryption procedure
	C1 := new(bn256.G2).ScalarBaseMult(ru)                 // C1=g_2^{r_u}
	C1.Neg(C1)                                             // C1=g_2^{-r_u}
	tempC1 := new(bn256.G2).ScalarBaseMult(numHashedKw)    // tempC1=g_2^{H(w)}
	tempC1.ScalarMult(tempC1, userKeySet.alpha)            // tempC1=g_2^{H(w)*a}
	C1.Add(C1, tempC1)                                     // C1=g_2^{H(w)a-r_u}
	C2 := new(bn256.G2).ScalarMult(userKeySet.gx, ru)      // C2=g_2^{x*r_u}
	temp1C2 := new(bn256.G2).ScalarBaseMult(userKeySet.Ku) // temp1C2=g_2^{Ku}
	temp1C2.Neg(temp1C2)                                   // temp1C2=g_2^{-Ku}
	temp1C2.ScalarMult(temp1C2, ru)                        // temp1C2=g_2^{-Ku*r_u}
	temp2C2 := new(bn256.G2).ScalarBaseMult(userKeySet.Ku) // temp2C2=g_2^{Ku}
	temp2C2.ScalarMult(temp2C2, numHashedKw)               // temp2C2=g_2^{Ku*H(w)}
	temp2C2.ScalarMult(temp2C2, userKeySet.alpha)          // temp2C2=g_2^{Ku*H(w)*a}
	C2 = C2.Add(C2, temp1C2)                               // C2=g_2^{Kb*ru}
	C2 = C2.Add(C2, temp2C2)                               // C2=g_2^{Kb*ru}*g_2^{Ku*H(w)*a}
	return C1, C2
}

// preServerEnc(c1 *bn256.G2, c2 *bn256.G2) (CT *bn256.G2): Proxy-Encryption in server. Server receives the ciphertext (C1, C2) of the keyword w from user and use this function to change the secret key
func preServerEnc(c1 *bn256.G2, c2 *bn256.G2) *bn256.G2 {
	CT := new(bn256.G2).ScalarMult(c1, serverKeySet.Kb) // CT=C1^{Kb}=g_2^{Kb*H(w)*a-Kb*r_u}
	CT.Add(CT, c2)                                      // CT=C1^{Kb}*C2=g^{x*H(w)*a}
	return CT
}

// getOnChainCipher(CT *bn256.G2) (*bn256.G1, *bn256.G2): Generate the ciphertext on-chain from CT
func getOnChainCipher(CT *bn256.G2) (*bn256.G1, *bn256.G2) {
	hashedCT := sha256.Sum256(CT.Marshal())                 // hash CT
	bigIntCT := new(big.Int).SetBytes(hashedCT[:])          // transform CT in bytes to CT in big.Int
	rb, _ := rand.Int(rand.Reader, bn256.Order)             // r_b: the nonce of this encryption procedure for server
	X := new(bn256.G1).ScalarMult(serverKeySet.g1Sbeta, rb) // X=g1^{S*beta*rb}
	Y := new(bn256.G2).ScalarBaseMult(serverKeySet.beta)    // Y=g2&{beta}
	Y.ScalarMult(Y, bigIntCT)                               // Y=g2^{beta*H(CT)}
	Y.ScalarMult(Y, rb)                                     // Y=g2^{beta*H(CT)*rb}
	return X, Y
}

// func enc(w string, id uint, index int): Encrypt the keyword w and id and add it into the cipher set
func enc(w string, id int, index int) {
	c1, c2 := userEncryption(w)

	CT := preServerEnc(c1, c2)
	X, Y := getOnChainCipher(CT)
	// X.Neg(X) 								// For bn256 in golang, X should not be the negative of original X
	var indexItem CipherText
	indexItem.X = X
	indexItem.Y = Y
	indexItem.id = id
	cipher[index] = indexItem
}

func search(w string) []int {
	c1, c2 := userEncryption(w)
	CT := preServerEnc(c1, c2)
	X, Y := getOnChainCipher(CT)
	var res [100]int
	index := 0
	for i := 0; i < taskSize; i++ {
		k1 := bn256.Pair(cipher[i].X, Y)
		k2 := bn256.Pair(X, cipher[i].Y)
		k1Byte := k1.Marshal()
		k2Byte := k2.Marshal()
		if bytes.Equal(k1Byte, k2Byte) {
			res[index] = cipher[i].id
			index++
		}
	}
	return res[:]
}

func readVocab() {
	f, _ := os.Open("../vocabulary.txt")
	sc := bufio.NewScanner(f)
	for i := 0; i < 2000; i++ {
		sc.Scan()
		t := strings.Split(sc.Text(), " ")
		Vocab[i].key = t[0]
		Vocab[i].id, _ = strconv.Atoi(t[1])
	}
	f.Close()
}

func main() {
	readVocab()
	keyGeneration()
	for i := 0; i < taskSize; i++ {
		enc(Vocab[i].key, Vocab[i].id, i)
	}
	var exp [10]float64
	for test := 0; test < 10; test++ {
		t := time.Now()
		search("hello")
		elapsed := time.Since(t)
		exp[test] = float64(elapsed.Milliseconds()) / 1000
	}
	fmt.Println(exp)
}
