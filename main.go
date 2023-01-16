package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"golang.org/x/crypto/curve25519"
)

// KeySize defines the size of the key
const KeySize = 32

func usage() {
	fmt.Printf("wg-util <genkey|genpsk|pubkey>\n")
	os.Exit(0)
}

// genkey generates a new 32-byte key confirming to https://cr.yp.to/ecdh.html
func genkey() (string, error) {
	var k [KeySize]byte
	_, err := rand.Read(k[:])
	if err != nil {
		return "", err
	}

	k[0] &= 248
	k[31] = (k[31] & 127) | 64
	return base64.StdEncoding.EncodeToString(k[:]), nil
}

// pubkey computes the public key matching this curve25519 secret key.
func pubkey(s string) (string, error) {
	k, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	var p [KeySize]byte
	curve25519.ScalarBaseMult(&p, (*[KeySize]byte)(k))
	return base64.StdEncoding.EncodeToString(p[:]), nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	if os.Args[1] == "genkey" || os.Args[1] == "genpsk" {
		k, err := genkey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", k)
	} else if os.Args[1] == "pubkey" {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		pub, err := pubkey(text)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", pub)
	} else {
		usage()
	}
}
