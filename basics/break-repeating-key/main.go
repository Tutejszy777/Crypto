package main

import (
	"fmt"
)

// data file has encrypted(base64) text
// the key is unkown

const input = `this is a test`
const input2 = `wokka wokka!!!`

func main() {
	for keylength := 2; keylength < 40; keylength++ {

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
