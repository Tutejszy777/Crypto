package main

import (
	"fmt"
)

// data file has encrypted(base64) text
// the key is unkown

const input = `this is a test`
const input2 = `wokka wokka!!!`

func main() {
	//normalizedMap := make(map[int]float32)

	for KEYSIZE := 2; KEYSIZE < 40; KEYSIZE++ {
		block := make([]byte, KEYSIZE)
		block2 := make([]byte, KEYSIZE)

		distance := HammingDistance(block, block2) / KEYSIZE
	}

	dist := HammingDistance([]byte(input), []byte(input2))
	fmt.Printf("Hamming distance: %d\n", dist)

}

// The Hamming distance is just the number of differing bits

func HammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		panic("Hamming distance requires equal length strings")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		xor := a[i] ^ b[i]
		for xor > 0 {
			distance += int(xor & 1)
			xor >>= 1
		}
	}
	return distance
}
