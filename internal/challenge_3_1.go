package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day3Challenge1 (file *os.File) {
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
		
		sum += maxJoltage(bank)
		bank = []int{}
	}

	fmt.Println(sum)
}

func maxJoltage (bank []int) int {
	firstPos := -1
	secondPos := -1

	for i := 0; i < len(bank) - 1; i++ {
		if firstPos < 0 || bank[firstPos] < bank[i] {
			firstPos = i
		}
	}
	
	for i := firstPos + 1; i < len(bank); i++ {
		if secondPos < 0 || bank[secondPos] < bank[i] {
			secondPos = i
		}
	}
	
	return bank[firstPos] * 10 + bank[secondPos]
}
