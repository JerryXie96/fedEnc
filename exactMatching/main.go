package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/clearmatics/bn256"
)

// the key set of CA
type KeySetOfCA struct {
	X  *big.Int  // X is private
	gx *bn256.G2 // gx is for server
}

// the key set of server
type KeySetOfServer struct {
	Kb   *big.Int
	beta *big.Int
	S    *big.Int
}

// the key set of user
type KeySetOfUser struct {
	Ku    *big.Int
	alpha *big.Int
	gx    *bn256.G2
}

var (
	caKeySet     KeySetOfCA
	serverKeySet KeySetOfServer
	userKeySet   KeySetOfUser
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

// main(): the main entrance
func main() {
	keyGeneration()
	c1, c2 := userEncryption("hello")
	fmt.Println(hex.EncodeToString(c1.Marshal()))
	fmt.Println(hex.EncodeToString(c2.Marshal()))
}
