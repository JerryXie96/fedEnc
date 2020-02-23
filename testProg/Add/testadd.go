package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/clearmatics/bn256"
)

func main() {
	a, _ := rand.Int(rand.Reader, bn256.Order)
	b, _ := rand.Int(rand.Reader, bn256.Order)
	c, _ := rand.Int(rand.Reader, bn256.Order)
	d, _ := rand.Int(rand.Reader, bn256.Order)

	g1a := new(bn256.G1).ScalarBaseMult(a) // g1^a
	g1b := new(bn256.G1).ScalarBaseMult(b)
	g1b = g1b.Neg(g1b) // g1^(-b)

	g2c := new(bn256.G2).ScalarBaseMult(c) // g2^c
	g2d := new(bn256.G2).ScalarBaseMult(d)
	g2d = g2d.Neg(g2d) // g2^(-d)

	g1a = g1a.Add(g1a, g1b) //g1^(a-b)
	g2c = g2c.Add(g2c, g2d) //g2^(c-d)

	ab := a.Sub(a, b)
	cd := c.Sub(c, d)
	g1ab := new(bn256.G1).ScalarBaseMult(ab)
	g2cd := new(bn256.G2).ScalarBaseMult(cd)

	fmt.Println(hex.EncodeToString(g1a.Marshal()))
	fmt.Println()
	fmt.Println(hex.EncodeToString(g1ab.Marshal()))
	fmt.Println()

	fmt.Println(hex.EncodeToString(g2c.Marshal()))
	fmt.Println()
	fmt.Println(hex.EncodeToString(g2cd.Marshal()))
	fmt.Println()
	k1 := bn256.Pair(g1a, g2c)
	k2 := bn256.Pair(g1ab, g2cd)
	k1Byte := k1.Marshal()
	k2Byte := k2.Marshal()

	fmt.Println(hex.EncodeToString(k1Byte))

	fmt.Println(hex.EncodeToString(k2Byte))

	if !bytes.Equal(k1Byte, k2Byte) {
		fmt.Print("keys didn't agree")
	} else {
		fmt.Print("Agreed\n")
	}
}
