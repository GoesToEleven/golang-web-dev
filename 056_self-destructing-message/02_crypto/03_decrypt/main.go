package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
	"strings"
)

func main() {
	// unencrypted
	unencrypted := "some message that you want to store / send securely"
	fmt.Println("UNENCRYPTED")
	fmt.Println(unencrypted)
	fmt.Println()

	// encrypt
	var nonce [24]byte
	io.ReadAtLeast(rand.Reader, nonce[:], 24)
	var password [32]byte
	io.ReadAtLeast(rand.Reader, password[:], 32)
	encrypted := secretbox.Seal(nil, []byte(unencrypted), &nonce, &password)
	fmt.Println("ENCRYPTED")
	fmt.Println(encrypted)
	fmt.Println()

	// decrypt
	enHex := fmt.Sprintf("%x:%x", nonce[:], encrypted)
	fmt.Println("NONCE:ENCRYPTED")
	fmt.Println(enHex)
	fmt.Println()

	var nonce2 [24]byte
	parts := strings.SplitN(enHex, ":", 2)
	if len(parts) < 2 {
		fmt.Errorf("expected nonce")
	}
	bs, err := hex.DecodeString(parts[0])
	if err != nil || len(bs) != 24 {
		fmt.Errorf("invalid nonce")
	}
	copy(nonce2[:], bs)
	fmt.Println("NONCE")
	fmt.Println(nonce)
	fmt.Println("NONCE2")
	fmt.Println(nonce2)
	fmt.Println()

	// get message
	bs, err = hex.DecodeString(parts[1])
	if err != nil {
		fmt.Errorf("invalid message")
	}

	// you need the password to open the sealed secret box
	decrypted, ok := secretbox.Open(nil, bs, &nonce2, &password)
	if !ok {
		fmt.Errorf("invalid message")
	}
	fmt.Println("DECRYPTED")
	fmt.Println(string(decrypted))
}
