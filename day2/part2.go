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

	depth := 0
	horizontalPosition := 0
	aim := 0

	for _, v := range input {
		x := strings.Split(v, " ")
		command, value := x[0], convertStrToNum(x[1])

		switch command {
		case "forward":
			{
				horizontalPosition += value
				depth += (value * aim)
			}
		case "down":
			{
				aim += value
			}
		case "up":
			{
				aim -= value
			}
		}
	}

	fmt.Println(horizontalPosition * depth)
}
