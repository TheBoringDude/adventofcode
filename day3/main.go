package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func binToDecimal(v string) int64 {
	x, _ := strconv.ParseInt(v, 2, 64)
	return x
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data), "\n")

	gammaRate := ""
	epsilonRate := ""

	vars := make([]string, 12)

	for _, v := range input {
		for j, k := range v {
			vars[j] = vars[j] + string(k)
		}
	}

	for _, v := range vars {
		ones := strings.Count(v, "1")
		zeros := strings.Count(v, "0")

		if ones > zeros {
			gammaRate += "1"
			epsilonRate += "0"
			continue
		}

		gammaRate += "0"
		epsilonRate += "1"
	}

	decimalGammaRate := binToDecimal(gammaRate)
	decimalEpsilonRate := binToDecimal(epsilonRate)

	fmt.Println(decimalGammaRate * decimalEpsilonRate)
}
