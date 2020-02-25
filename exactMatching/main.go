package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

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

// enc(instance *EMABI, auth *bind.TransactOpts, w string, id uint): Encrypt the keyword and upload the encrypted keyword and id to blockchain
func enc(instance *EMABI, auth *bind.TransactOpts, conn *ethclient.Client, w string, id uint) {
	c1, c2 := userEncryption(w)
	CT := preServerEnc(c1, c2)
	// fmt.Println("CT: ", hex.EncodeToString(CT.Marshal()))
	X, Y := getOnChainCipher(CT)
	X.Neg(X)
	var indexItem Struct0
	indexItem.X = X.Marshal()
	indexItem.Y = Y.Marshal()
	// fmt.Println(hex.EncodeToString(X.Marshal()))
	// fmt.Println()
	// fmt.Println(hex.EncodeToString(Y.Marshal()))
	// fmt.Println()
	bIId := new(big.Int).SetUint64(uint64(id)) // the big.Int form of id
	auth.Nonce = nil
	tx, err := instance.Store(auth, indexItem, bIId)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, conn, tx)
	// fmt.Println("status: ", receipt.Status)
	if err != nil {
		fmt.Println(err)
	}
}

// search(instance *EMABI, auth *bind.TransactOpts, w string) []*big.Int: Search from the index on-chain with keyword w
func search(instance *EMABI, auth *bind.TransactOpts, conn *ethclient.Client, w string) []*big.Int {
	c1, c2 := userEncryption(w)
	CT := preServerEnc(c1, c2)
	X, Y := getOnChainCipher(CT)
	// fmt.Println("CT: ", hex.EncodeToString(CT.Marshal()))
	var indexItem Struct0
	indexItem.X = X.Marshal()
	indexItem.Y = Y.Marshal()
	// fmt.Println(hex.EncodeToString(X.Marshal()))
	// fmt.Println()
	// fmt.Println(hex.EncodeToString(Y.Marshal()))
	// fmt.Println()
	tx, err := instance.Search(auth, indexItem)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, conn, tx)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(tx.Hash().Hex())
	// fmt.Println("Status:", receipt.Status)
	res, _ := instance.GetResult(nil)
	return res
}

// clearResult(instance *EMABI, auth *bind.TransactOpts): Clear last-retrieval's result
func clearResult(instance *EMABI, auth *bind.TransactOpts, conn *ethclient.Client) {
	tx, _ := instance.ClearRes(auth)
	ctx := context.Background()
	_, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		fmt.Println(err)
	}
}

// main(): the main entrance
func main() {
	var url, scAddress, privateKeyStr string
	url = "http://localhost:8545"
	scAddress = "0xFE809aA009dD2f332838eC8799379D3B860fEA04"
	privateKeyStr = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"

	client, err := ethclient.Dial(url)
	fmt.Println("CONNECTED")
	tokenAddress := common.HexToAddress(scAddress)
	instance, err := NewEMABI(tokenAddress, client)
	fmt.Println("Instance Generated")
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Privatekey Settled")
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nil
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(0)

	keyGeneration()
	fmt.Println("Key Generated")
	enc(instance, auth, client, "hello", 1)
	enc(instance, auth, client, "hello", 2)
	fmt.Println("Encryption Completed")
	res := search(instance, auth, client, "hello")
	// fmt.Println(res)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].String())
	}
	clearResult(instance, auth, client)
}
