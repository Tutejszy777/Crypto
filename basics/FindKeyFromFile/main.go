package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "data.txt"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	buffer := bufio.NewScanner(file)

	for {
		bytes := readFile(buffer)
		key, bestMatch := searchForBestKey(bytes)
		if key != 0 {
			fmt.Printf("Key: %d\n", key)
			fmt.Printf("Decrypted: %s\n", bestMatch)
		}
	}

}

func readFile(scanner *bufio.Scanner) []byte {
	scanner.Scan()
	bytes := scanner.Bytes()

	return bytes
}

func searchForBestKey(bytes []byte) (int, []byte) {
	topProbability := 0.0
	topKey := 0
	topBytes := make([]byte, len(bytes))

	for key := 0; key <= 255; key++ {
		copy := bytes

		copy = xorByKey(copy, byte(key))
		probability := evaluate(copy)

		if probability > topProbability {
			topProbability = probability
			topKey = key
			topBytes = copy
		}
	}

	return topKey, topBytes
}

func xorByKey(bytes []byte, key byte) []byte {
	for i, j := range bytes {
		bytes[i] = j ^ key
	}

	return bytes
}

func evaluate(bytes []byte) float64 {
	frequency := map[rune]float64{
		'e': 12.7, 't': 9.1, 'a': 8.2, 'o': 7.5, 'i': 7.0, 'n': 6.7,
		's': 6.3, 'h': 6.1, 'r': 6.0, 'd': 4.3, 'l': 4.0, 'u': 2.8,
		' ': 13.0,
	}

	probability := 0.0

	for _, char := range bytes {
		val, ok := frequency[rune(char)]
		if ok {
			probability += val
		}
	}

	return probability
}
