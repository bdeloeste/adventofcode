package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Tuple struct {
	pair     string
	distance int
	diffIdx  int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Levenshtein Distance
func editDistance(s1, s2 []rune) int {
	s1len := len(s1)
	s2len := len(s2)
	column := make([]int, len(s1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastKey := x - 1
		for y := 1; y <= s1len; y++ {
			oldKey := column[y]
			incr := 0
			if s1[y-1] != s2[x-1] {
				incr = 1
			}

			column[y] = min(column[y]+1, min(column[y-1]+1, lastKey+incr))
			lastKey = oldKey
		}
	}
	return column[s1len]
}

func hashRunes(r []rune) int {
	result := 0
	for _, val := range r {
		result += int(val)
	}
	return result
}

func getDiffIndex(s1, s2 string) int {
	for i := range s1 {
		if s1[i] != s2[i] {
			return i
		}
	}
	return -1
}

func main() {
	cwd, _ := os.Getwd()
	file, err := os.Open(cwd + "\\data\\2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	strings := []string{}
	for scanner.Scan() {
		t := scanner.Text()
		strings = append(strings, t)
	}

	pair := make(map[int]Tuple)
	for _, i := range strings {
		for _, j := range strings {
			if i == j {
				continue
			}
			rs1 := []rune(i)
			rs2 := []rune(j)
			d := editDistance(rs1, rs2)
			if d == 1 {
				key := hashRunes(rs1) + hashRunes(rs2)
				if _, ok := pair[key]; ok == false {
					pair[key] = Tuple{pair: i + "," + j, distance: d, diffIdx: getDiffIndex(i, j)}
				}
			}
		}
	}

	fmt.Println(pair)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
