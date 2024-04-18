package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func main() {
	fmt.Println("Hello this is a test")
	rand := rand.Reader

	privAlice, err := ecdh.X25519().GenerateKey(rand)
	if err != nil {
		panic(err)
	}

	privBob, err := ecdh.X25519().GenerateKey(rand)
	if err != nil {
		panic(err)
	}

	pubAlice := privAlice.PublicKey()
	pubBob := privBob.PublicKey()

	secret, err := privAlice.ECDH(pubBob)
	if err != nil {
		panic(err)
	}

	secret_check, err := privBob.ECDH(pubAlice)
	if err != nil {
		panic(err)
	}

	fmt.Println(secret, secret_check)

}
