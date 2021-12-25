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

	for i, v := range input {
		x := convertStrToNum(v)
		if i != 0 && x > prev {
			increases += 1
		}

		prev = x
	}

	fmt.Println(increases)
}
