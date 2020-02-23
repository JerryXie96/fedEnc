package main

import (
	"bytes"
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
	a, _ := rand.Int(rand.Reader, bn256.Order)
	b, _ := rand.Int(rand.Reader, bn256.Order)

	pa := new(bn256.G1).ScalarBaseMult(a)
	qa := new(bn256.G2).ScalarBaseMult(a)

	pb := new(bn256.G1).ScalarBaseMult(b)
	qb := new(bn256.G2).ScalarBaseMult(b)

	npb := new(bn256.G1).Neg(pb)

	paByte := pa.Marshal()
	qbByte := qb.Marshal()
	pbByte := npb.Marshal()
	qaByte := qa.Marshal()

	k1 := bn256.Pair(pa, qb)
	k2 := bn256.Pair(pb, qa)
	k1Byte := k1.Marshal()
	k2Byte := k2.Marshal()

	if !bytes.Equal(k1Byte, k2Byte) {
		fmt.Print("keys didn't agree")
	} else {
		fmt.Print("Agreed\n")
	}

	client, err := ethclient.Dial("http://localhost:8545")
	fmt.Println("CONNECTED")
	tokenAddress := common.HexToAddress("0x605Ac4b7b75b55007D1cBc652107DF9A25a7778b")
	instance, err := NewTestABI(tokenAddress, client)
	fmt.Println("Generated")
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privatekey got")

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nil
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = big.NewInt(0)
	var data Struct0
	data.A = paByte
	data.B = qbByte
	_, err = instance.Store(auth, data)
	fmt.Println("Stored")
	var pairingData Struct0
	pairingData.A = pbByte
	pairingData.B = qaByte
	_, err = instance.Compare(auth, pairingData)
	fmt.Println("Compared")
	isFit, err := instance.ReIsFit(nil)
	fmt.Println(isFit)

	fmt.Println(hex.EncodeToString(paByte))
	fmt.Println()
	fmt.Println(hex.EncodeToString(qbByte))
	fmt.Println()

	fmt.Println(hex.EncodeToString(pbByte))
	fmt.Println()
	fmt.Println(hex.EncodeToString(qaByte))
	fmt.Println()
}
