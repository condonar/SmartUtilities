package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input(scanner *bufio.Scanner) ([]int, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("No input")
	}

	str := scanner.Text()
	strSlice := strings.Fields(str)
	intSlice := make([]int, 0, len(strSlice))

	for _, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("Invalid number: %s", s)
		}
		intSlice = append(intSlice, num)
	}

	return intSlice, nil
}

func Intersection(firstSlice []int, secondSlice []int) []int {
	uniqSecondMap := make(map[int]bool)
	for _, v := range secondSlice {
		uniqSecondMap[v] = false
	}

	var result []int
	for _, v := range firstSlice {
		if added, ok := uniqSecondMap[v]; ok && !added {
			result = append(result, v)
			uniqSecondMap[v] = true
		}
	}

	return result
}

func Output(result []int) {
	if len(result) == 0 {
		fmt.Println("Empty intersection")
	} else {
		for i, v := range result {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	firstSlice, err1 := Input(scanner)
	if err1 != nil {
		fmt.Println("Invalid input")
		return
	}
	secondSlice, err2 := Input(scanner)
	if err2 != nil {
		fmt.Println("Invalid input")
		return
	}

	result := Intersection(firstSlice, secondSlice)
	Output(result)
}
