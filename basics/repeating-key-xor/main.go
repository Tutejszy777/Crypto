package main

import (
	"encoding/hex"
	"fmt"
)

const input = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

const key = "ICE"

func main() {
	// encryption
	encrypted := xorByKey([]byte(input), []byte(key))
	fmt.Printf("Encrypted (hex): %s\n", hex.EncodeToString(encrypted))
	// decryption
	decrypted := xorByKey(encrypted, []byte(key))
	fmt.Printf("Decrypted: %s\n", string(decrypted))

}

func xorByKey(data []byte, key []byte) []byte {
	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ key[i%len(key)]
	}
	return result
}
