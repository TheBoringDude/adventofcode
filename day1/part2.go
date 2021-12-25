package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertStrToNum(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data), "\n")

	increases := 0
	prev := 0

	for index := range input {
		if index+3 > len(input) {
			continue
		}

		s := convertStrToNum(input[index]) + convertStrToNum(input[index+1]) + convertStrToNum(input[index+2])

		if index != 0 && s > prev {
			increases += 1
		}

		prev = s
	}

	fmt.Println(increases)
}
