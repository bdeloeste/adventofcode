package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	cwd, _ := os.Getwd()
	file, err := os.Open(cwd + "\\data\\1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := []int{}
	for scanner.Scan() {
		t := scanner.Text()
		if s, err := strconv.Atoi(t); err == nil {
			nums = append(nums, s)
		}
	}

	i := 0
	m := make(map[int]int)
	for k := 0; ; k++ {
		if k == len(nums) {
			k = 0
		}

		val := nums[k]
		if _, ok := m[i]; ok {
			fmt.Println(i, "exists")
			break
		} else {
			m[i] = 1
		}
		i += val
	}

	fmt.Println(i)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
