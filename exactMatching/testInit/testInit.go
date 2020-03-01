// Initialization Benchmark
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

var (
	caKeySet     KeySetOfCA
	expSize      int              = 350
	serverKeySet []KeySetOfServer = make([]KeySetOfServer, expSize)
	userKeySet   []KeySetOfUser   = make([]KeySetOfUser, expSize)
)

// keyGeneration(): Generate the key set in CA
func keyGenerationFirst() {
	serverKeySet[0].Kb, _ = rand.Int(rand.Reader, bn256.Order) // Kb: one key for server
	userKeySet[0].Ku, _ = rand.Int(rand.Reader, bn256.Order)   // Ku: one key for user
	caKeySet.X = big.NewInt(0)                                 // X: the key held by CA
	caKeySet.X.Add(serverKeySet[0].Kb, userKeySet[0].Ku)       // X=Kb+Ku. for all the users, X is the same

	caKeySet.gx = new(bn256.G2).ScalarBaseMult(caKeySet.X)      // gx: g^x is used for the encryption in user
	userKeySet[0].gx = new(bn256.G2).ScalarBaseMult(caKeySet.X) // gx: g^x is used for the encryption in user

	userKeySet[0].alpha, _ = rand.Int(rand.Reader, bn256.Order)  // alpha: the encryption key for user in encryption
	serverKeySet[0].beta, _ = rand.Int(rand.Reader, bn256.Order) // beta: the encryption key for server in on-chain encryption

	serverKeySet[0].S, _ = rand.Int(rand.Reader, bn256.Order) // S: the encryption key for server in on-chain encryption

	serverKeySet[0].g1Sbeta = new(bn256.G1).ScalarBaseMult(serverKeySet[0].S)
	serverKeySet[0].g1Sbeta.ScalarMult(serverKeySet[0].g1Sbeta, serverKeySet[0].beta) //g1^{S*beta}: the encryption key for server in on-chain encryption
}

func keyGeneration(i int) {
	serverKeySet[i].Kb, _ = rand.Int(rand.Reader, bn256.Order) // Kb: one key for server
	for serverKeySet[i].Kb.Cmp(caKeySet.X) != -1 {
		serverKeySet[i].Kb, _ = rand.Int(rand.Reader, bn256.Order) // Kb: one key for server
	}
	userKeySet[i].Ku = big.NewInt(0).Sub(caKeySet.X, serverKeySet[i].Kb) // Ku: one key for user
	userKeySet[i].gx = new(bn256.G2).ScalarBaseMult(caKeySet.X)          // gx: g^x is used for the encryption in user

	userKeySet[i].alpha, _ = rand.Int(rand.Reader, bn256.Order)  // alpha: the encryption key for user in encryption
	serverKeySet[i].beta, _ = rand.Int(rand.Reader, bn256.Order) // beta: the encryption key for server in on-chain encryption

	serverKeySet[i].S, _ = rand.Int(rand.Reader, bn256.Order) // S: the encryption key for server in on-chain encryption

	serverKeySet[i].g1Sbeta = new(bn256.G1).ScalarBaseMult(serverKeySet[i].S)
	serverKeySet[i].g1Sbeta.ScalarMult(serverKeySet[i].g1Sbeta, serverKeySet[i].beta) //g1^{S*beta}: the encryption key for server in on-chain encryption
}

func main() {
	var exp [10]float64
	for expTime := 0; expTime < 10; expTime++ {
		t := time.Now()
		keyGenerationFirst()
		for i := 1; i < expSize; i++ {
			keyGeneration(i)
		}
		elapsed := time.Since(t)
		exp[expTime] = float64(elapsed.Milliseconds()) / 1000
	}
	fmt.Println(exp)
}
