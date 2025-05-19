package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var hexString1 string = "1c0111001f010100061a024b53535009181c"
	var hexString2 string = "686974207468652062756c6c277320657965"

	bytes1, err := hex.DecodeString(hexString1)
	if err != nil {
		fmt.Println("Failed to decode")
	}

	bytes2, er := hex.DecodeString(hexString2)
	if er != nil {
		fmt.Println("Failed to decode")
	}

	if len(bytes1) != len(bytes2) {
		fmt.Println("Byte slices have to be same length")
	}

	result := make([]byte, len(bytes1))

	for i := range bytes1 {
		result[i] = bytes1[i] ^ bytes2[i]
	}

	hexResult := hex.EncodeToString(result)

	fmt.Println(hexResult)
}
