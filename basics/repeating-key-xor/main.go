package main

import (
	"encoding/hex"
	"fmt"
)

const decrypt = "Burning 'em, if you ain't quick and nimble"
"I go crazy when I hear a cymbal"

const key = "ICE"

func main() {
	bytes, err := hex.DecodeString(decrypt)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}

	decrypted := xorByKey(bytes, []byte(key), len(key))
	fmt.Printf("Decrypted: %s\n", string(decrypted))
}

func xorByKey(bytes []byte, key []byte, keylength int) []byte {
	var keyIndex int = 0
	for i, j := range bytes {
		if keyIndex == keylength-1 {
			keyIndex = 0
		} else {
			keyIndex++
		}
		bytes[i] = j ^ key[keyIndex]
	}

	return bytes
}
