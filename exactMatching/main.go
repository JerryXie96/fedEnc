package main

import (
	"crypto/rand"
	"math/big"

	"github.com/clearmatics/bn256"
)

// the key set of CA
type KeySetOfCA struct {
	X  *big.Int
	g1 *bn256.G1
	g2 *bn256.G2
	gx *bn256.G2
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

// Generate the key set in CA
func keyGeneration() {
	serverKeySet.Kb, _ = rand.Int(rand.Reader, bn256.Order) // Kb: one key for server
	userKeySet.Ku, _ = rand.Int(rand.Reader, bn256.Order)   // Ku: one key for user
	X := big.NewInt(0)                                      // X: the key held by CA
	X = X.Add(serverKeySet.Kb, userKeySet.Ku)               // X=Kb+Ku. for all the users, X is the same

	caKeySet.g1 = new(bn256.G1) // the calculation base in G1
	caKeySet.g2 = new(bn256.G2) // the calculation base in G2

	caKeySet.gx.ScalarMult(caKeySet.g2, X)   // gx: g^x is used for the encryption in user
	userKeySet.gx.ScalarMult(caKeySet.g2, X) // gx: g^x is used for the encryption in user

	userKeySet.alpha, _ = rand.Int(rand.Reader, bn256.Order)  // alpha: the encryption key for user in encryption
	serverKeySet.beta, _ = rand.Int(rand.Reader, bn256.Order) // beta: the encryption key for server in on-chain encryption

	serverKeySet.S, _ = rand.Int(rand.Reader, bn256.Order) // S: the encryption key for server in on-chain encryption
}
func main() {
	keyGeneration()
}
