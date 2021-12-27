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

	oxygenNums := input
	co2Nums := input

	// --------- oxygen
	for i := 0; i < 12; i++ {
		if len(co2Nums) == 1 {
			break
		}

		vars := make([]string, 12)

		for _, v := range oxygenNums {
			for j, k := range v {
				vars[j] = vars[j] + string(k)
			}
		}

		delim := ""

		ones := strings.Count(vars[i], "1")
		zeros := strings.Count(vars[i], "0")

		if ones > zeros || ones == zeros {
			delim = "1"
		} else {
			delim = "0"
		}

		newnums := []string{}

		for _, v := range oxygenNums {
			if string(v[i]) == delim {
				newnums = append(newnums, v)
			}
		}

		oxygenNums = newnums
	}

	// ---------- co2
	for i := 0; i < 12; i++ {
		if len(co2Nums) == 1 {
			break
		}

		vars := make([]string, 12)

		for _, v := range co2Nums {
			for j, k := range v {
				vars[j] = vars[j] + string(k)
			}
		}

		delim := ""

		ones := strings.Count(vars[i], "1")
		zeros := strings.Count(vars[i], "0")

		if ones > zeros || ones == zeros {
			delim = "0"
		} else {
			delim = "1"
		}

		newnums := []string{}

		for _, v := range co2Nums {
			if string(v[i]) == delim {
				newnums = append(newnums, v)
			}
		}

		co2Nums = newnums
	}

	ox := binToDecimal(oxygenNums[0])
	co := binToDecimal(co2Nums[0])

	fmt.Println(ox * co)
}
