package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

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
	expSize      int       = 2
	cipher       []Struct0 = make([]Struct0, expSize)
	Vocab        [1800]int // Vocab: the number set
	bin          [32]int

	digit int = 1 // the number of digits in one block
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
func userEncryption(w int64) (*bn256.G2, *bn256.G2) {
	kwByte := []byte(string(w))                            // the byte array of keyword w
	hashedKw := sha256.Sum256(kwByte[:])                   // hashed the keyword w
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

func enc(w int32, id int, index int) {
	wStr := strconv.FormatInt(int64(w), 2) // convert w to binary string
	wStr = fmt.Sprintf("%032s", wStr)      // padding to 32 bits
	for i := 0; i < 32/digit; i++ {
		res, _ := strconv.ParseInt(wStr[i*digit:i*digit+digit], 2, 0)
		c1, c2 := userEncryption(res)
		CT := preServerEnc(c1, c2)
		X, Y := getOnChainCipher(CT)
		X.Neg(X)
		cipher[index].Cipher[i][0] = X.Marshal()
		cipher[index].Cipher[i][1] = Y.Marshal()
	}
	bIId := new(big.Int).SetUint64(uint64(id)) // the big.Int form of id
	cipher[index].Id = bIId
}

func post(instance *TestABI, auth *bind.TransactOpts, conn *ethclient.Client) {
	auth.Nonce = nil
	for i := 0; i < len(cipher)/2; i++ {
		tx, err := instance.Store(auth, cipher[i*2:i*2+2])
		if err != nil {
			fmt.Println(err)
		}
		ctx := context.Background()
		receipt, err := bind.WaitMined(ctx, conn, tx)
		fmt.Println("status: ", receipt.Status)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func search(instance *TestABI, auth *bind.TransactOpts, conn *ethclient.Client, w int32) []*big.Int {
	var query Struct2
	wStr := strconv.FormatInt(int64(w), 2) // convert w to binary string
	wStr = fmt.Sprintf("%032s", wStr)      // padding to 32 bits
	for i := 0; i < 32/digit; i++ {
		res, _ := strconv.ParseInt(wStr[i*digit:i*digit+digit], 2, 0)
		for j := 0; j < int(math.Exp2(float64(digit))); j++ {
			c1, c2 := userEncryption(res + int64(j)) // assume that the relation operator is greater, i.e. >w
			CT := preServerEnc(c1, c2)
			X, Y := getOnChainCipher(CT)
			X.Neg(X)
			query.Cipher[i][j].C1 = X.Marshal()
			query.Cipher[i][j].C2 = Y.Marshal()
		}
	}
	tx, err := instance.Search(auth, query)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, conn, tx)
	fmt.Println("status: ", receipt.Status)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := instance.GetResult(nil)
	return res
}

func clear(instance *TestABI, auth *bind.TransactOpts, conn *ethclient.Client) {
	tx, err := instance.ClearResult(auth)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, conn, tx)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// var t int32
	// t = 200
	// str := strconv.FormatInt(int64(t), 2)
	// str = fmt.Sprintf("%032s", str)
	// res, _ := strconv.ParseInt(str[30:32], 2, 0)
	// fmt.Println(str)
	// fmt.Println(res)

	var url, scAddress, privateKeyStr string
	url = "http://localhost:8545"
	scAddress = "0x56F946E3350E74DcF15cE2c5500921541dEC79B4"
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
	enc(5, 5, 0)
	enc(8, 8, 1)
	post(instance, auth, client)
	fmt.Println("Encryption Completed")
	res := search(instance, auth, client, 6)
	// fmt.Println(res)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].String())
	}
	clear(instance, auth, client)
}
