package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const fileName = "data.txt"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	buffer := bufio.NewScanner(file)
	defer file.Close()

	bestKey := 0
	var bestString string
	bestResult := 0.0

	for {
		bytes, end := readFile(buffer)
		if !end {
			break
		}
		key, bestMatch, probability := searchForBestKey(bytes)
		Print(key, bestMatch, probability)
		if probability > bestResult {
			bestResult = probability
			bestKey = key
			bestString = string(bestMatch)
		}
	}

	fmt.Printf("Best Key: %d\n", bestKey)
	fmt.Printf("Best match: %s\n", bestString)
	fmt.Printf("Length: %d\n", len(bestString))
	fmt.Printf("Best probability: %f\n", bestResult)

}

func Print(key int, bytes []byte, probability float64) {
	fmt.Printf("Key: %d\n", key)
	fmt.Printf("Decrypted: %s\n", string(bytes))
	fmt.Printf("Probability: %f\n", probability)
}

func readFile(scanner *bufio.Scanner) ([]byte, bool) {
	end := scanner.Scan()
	if !end {
		return nil, false
	}
	hexstr := scanner.Text()
	bytes, er := hex.DecodeString(hexstr)
	if er != nil {
		fmt.Println("Failed to convert")
	}

	return bytes, true
}

func searchForBestKey(bytes []byte) (int, []byte, float64) {
	topProbability := 0.0
	topKey := 0
	topBytes := make([]byte, len(bytes))

	for key := 0; key <= 255; key++ {
		// bCopy := bytes does not make a fresh slice—it just makes a new variable pointing at the same underlying array.
		bCopy := make([]byte, len(bytes))
		copy(bCopy, bytes)

		bCopy = xorByKey(bCopy, byte(key))
		probability := evaluate(bCopy)

		if probability > topProbability {
			topProbability = probability
			topKey = key
			topBytes = bCopy
		}
	}

	return topKey, topBytes, topProbability
}

func xorByKey(bytes []byte, key byte) []byte {
	for i, j := range bytes {
		bytes[i] = j ^ key
	}

	return bytes
}

func evaluate(bytes []byte) float64 {
	str := string(bytes)
	frequency := map[rune]float64{
		'e': 12.7, 't': 9.1, 'a': 8.2, 'o': 7.5, 'i': 7.0, 'n': 6.7,
		's': 6.3, 'h': 6.1, 'r': 6.0, 'd': 4.3, 'l': 4.0, 'u': 2.8,
		' ': 13.0,
	}

	score := 0.0
	for _, r := range strings.ToLower(str) {
		if unicode.IsPrint(r) {
			score += frequency[r]
		} else {
			score -= 10
		}
	}
	return score
}
