package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Input(scanner *bufio.Scanner) (string, int) {
	scanner.Scan()
	str := scanner.Text()

	scanner.Scan()
	kStr := scanner.Text()
	K, _ := strconv.Atoi(strings.TrimSpace(kStr))

	return str, K
}

func GetTopWords(str string, K int) []string {
	countWords := make(map[string]int)
	words := strings.Fields(str)
	var orderedWords []string
	for _, word := range words {
		countWords[word]++
	}
	for key := range countWords {
		orderedWords = append(orderedWords, key)
	}
	sort.Slice(orderedWords, func(i, j int) bool {
		word1 := orderedWords[i]
		word2 := orderedWords[j]

		if countWords[word1] != countWords[word2] {
			return countWords[word1] > countWords[word2]
		}

		return word1 < word2
	})
	if K > len(orderedWords) {
		K = len(orderedWords)
	}
	return orderedWords[:K]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	str, K := Input(scanner)

	result := GetTopWords(str, K)

	for i, word := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(word)
	}
	fmt.Println()
}
