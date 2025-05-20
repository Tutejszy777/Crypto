package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// A LOT OF OPTIMIZATION CAN BE DONE
func main() {
	// Hex string converted to bytes
	var hexstr string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, er := hex.DecodeString(hexstr)

	if er != nil {
		fmt.Println("Failed to convert")
	}

	// createad a map to key - decrypted string
	var Map map[uint8]string = make(map[uint8]string)
	// string build for efficient string creation
	var strBuilder strings.Builder

	for key := 0; key <= 255; key++ {
		decrypted := bytes

		for i, j := range decrypted {
			decrypted[i] = j ^ byte(key)
		}

		strBuilder.Write(decrypted)
		Map[uint8(key)] = strBuilder.String()
		strBuilder.Reset()
	}

	// can be improved by adding a more key to the map
	// can be improved by uint8
	max := 0
	maxKey := 0 // can be issues with 0

	// iterate through the map
	for k, v := range Map {
		var letter string = "aetoinshr"
		probability := 0

		// count the number of letters in the decrypted string
		for i := 0; i < len(letter); i++ {
			probability += strings.Count(v, string(letter[i]))
		}

		// check if the current key has a higher probability
		if probability >= max {
			max = probability
			if int(k) != maxKey { // check if the current key is not the same as the previous key
				// delete the previous key from the map
				deleteFromMap(Map, uint8(maxKey))
			}
			maxKey = int(k)
		} else {
			delete(Map, k)
		}
	}

	// at current state only one key is left in the map
	for k, v := range Map {
		fmt.Printf("Key: %d, Decrypted: %s\n", k, v)
	}
}

func deleteFromMap(m map[uint8]string, key uint8) {
	delete(m, key)
}
