package internal

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Day3Challenge2 (file *os.File) {
	var sum int
	var bank []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		
		for _, v := range line {
			num, err := strconv.Atoi(string(v))
			
			if err != nil {
				fmt.Println("Could not convert", v, "to int")
			}

			bank = append(bank, num)
		}

		sum += maxJoltageOverride(bank)
		bank = []int{}
	}

	fmt.Println(sum)
}

func maxJoltageOverride (bank []int) int {
	var joltage int = 0
	var batteriesLeft int = 0

	lenBank := len(bank)
	maxPositions := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i := 0; i <= len(bank) - 12; i++ {
		if maxPositions[0] < 0 || bank[maxPositions[0]] < bank[i] {
			maxPositions[0] = i
		}
	}
	
	for i := 11; i >= 1; i-- {
		batteriesLeft = 12 - i
		
		for j := maxPositions[batteriesLeft - 1] + 1; j <= lenBank - i; j++ {
			if maxPositions[batteriesLeft] < 0 || bank[maxPositions[batteriesLeft]] < bank[j] {
				maxPositions[batteriesLeft] = j
			}
		}
	}

	for i, v := range maxPositions {
		joltage += bank[v] * int(math.Pow10(11 - i))
	}
	
	return joltage
}
