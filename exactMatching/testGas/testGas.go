// This file is for the testing of smart contract for experiment

package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/clearmatics/bn256"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

var (
	caKeySet     KeySetOfCA
	serverKeySet KeySetOfServer
	userKeySet   KeySetOfUser
	taskSize     int       = 1
	cipher       []Struct1 = make([]Struct1, taskSize)
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
	X.Neg(X)
	var indexItem Struct0
	indexItem.X = X.Marshal()
	indexItem.Y = Y.Marshal()
	bIId := new(big.Int).SetUint64(uint64(id)) // the big.Int form of id
	var indexCollectionItem Struct1
	indexCollectionItem.CT = indexItem
	indexCollectionItem.Id = bIId
	cipher[index] = indexCollectionItem
}

func post(instance *TestABI, auth *bind.TransactOpts, conn *ethclient.Client) {
	auth.Nonce = nil
	tx, err := instance.Store(auth, cipher[:])
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, conn, tx)
	fmt.Println(receipt.GasUsed)
	fmt.Println("status: ", receipt.Status)
	if err != nil {
		fmt.Println(err)
	}
	// tx, err = instance.Store(auth, cipher[150:taskSize])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// ctx = context.Background()
	// receipt, err = bind.WaitMined(ctx, conn, tx)
	// fmt.Println(receipt.GasUsed)
	// fmt.Println("status: ", receipt.Status)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func search(instance *TestABI, auth *bind.TransactOpts, conn *ethclient.Client, w string) []*big.Int {
	c1, c2 := userEncryption(w)
	CT := preServerEnc(c1, c2)
	X, Y := getOnChainCipher(CT)
	var indexItem Struct0
	indexItem.X = X.Marshal()
	indexItem.Y = Y.Marshal()
	tx, err := instance.Search(auth, indexItem)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, conn, tx)
	fmt.Println("status: ", receipt.Status)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(receipt.GasUsed)
	res, _ := instance.GetResult(nil)
	return res
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
	var url, scAddress, privateKeyStr string
	url = "http://localhost:8545"
	scAddress = "0x660c2Ae2D2c943b2bcAB77C510300F20cb19aC73"
	privateKeyStr = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"

	client, err := ethclient.Dial(url)
	fmt.Println("CONNECTED")
	tokenAddress := common.HexToAddress(scAddress)
	instance, err := NewTestABI(tokenAddress, client)
	fmt.Println("Instance Generated")
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Privatekey Settled")
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nil
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = big.NewInt(0)
	keyGeneration()
	fmt.Println("Key Generated")
	readVocab()
	for i := 0; i < 1; i++ {
		enc(Vocab[i].key, Vocab[i].id, i)
	}
	post(instance, auth, client)
	fmt.Println("Encryption Completed")
	search(instance, auth, client, "hello")
	// var exp [25]float64
	// for i := 0; i < 25; i++ {
	// 	t := time.Now()
	// 	search(instance, auth, client, "hello")
	// 	elapsed := time.Since(t)
	// 	exp[i] = float64(elapsed.Milliseconds()) / 1000
	// }
	// fmt.Println(exp)
}
