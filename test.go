package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/clearmatics/bn256"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var url, scAddress, privateKeyStr string
	url = "http://localhost:8545"
	scAddress = "0x751fF0C30D6f589A399461ceF0ca635eD5d31DeD"
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
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = big.NewInt(0)

	a, _ := rand.Int(rand.Reader, bn256.Order)
	b, _ := rand.Int(rand.Reader, bn256.Order)
	pa := new(bn256.G1).ScalarBaseMult(a)
	qa := new(bn256.G2).ScalarBaseMult(a)

	pb := new(bn256.G1).ScalarBaseMult(b)
	qb := new(bn256.G2).ScalarBaseMult(b)
	pb.Neg(pb)
	paByte := pa.Marshal()
	qbByte := qb.Marshal()
	pbByte := pb.Marshal()
	qaByte := qa.Marshal()

	fmt.Println(hex.EncodeToString(paByte))
	fmt.Println(hex.EncodeToString(qbByte))
	fmt.Println(hex.EncodeToString(pbByte))
	fmt.Println(hex.EncodeToString(qaByte))

	var stru Struct0
	stru.X = paByte
	stru.Y = qbByte
	bIId := new(big.Int).SetUint64(uint64(1)) // the big.Int form of id
	tx, _ := instance.Store(auth, stru, bIId)
	ctx := context.Background()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		fmt.Println(err)
	}
	stru.X = pbByte
	stru.Y = qaByte
	fmt.Println(hex.EncodeToString(stru.X))
	tx, _ = instance.Search(auth, stru)
	ctx = context.Background()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		fmt.Println(err)
	}
}
