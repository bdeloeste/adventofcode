package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func hasDupChars(str string, numDups int) int {
	charCount := make(map[rune]int)
	for _, char := range str {
		if _, ok := charCount[char]; ok {
			charCount[char] += 1
		} else {
			charCount[char] = 1
		}
	}
	for _, v := range charCount {
		if v == numDups {
			return 1
		}
	}
	return 0
}

func main() {
	cwd, _ := os.Getwd()
	file, err := os.Open(cwd + "\\data\\2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numTwos, numThrees := 0, 0
	for scanner.Scan() {
		t := scanner.Text()
		numTwos += hasDupChars(t, 2)
		numThrees += hasDupChars(t, 3)

	}

	fmt.Println(numTwos * numThrees)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
