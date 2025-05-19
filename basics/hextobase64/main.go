package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var hexString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	bytes, er := hex.DecodeString(hexString)

	if er != nil {
		fmt.Println("Failed to convert")
	}

	str64 := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println("Original hex:")
	fmt.Println(hexString)
	fmt.Println("hex to bytes:")
	fmt.Println(bytes)
	fmt.Println("bytes to base64:")
	fmt.Println(str64)
}
