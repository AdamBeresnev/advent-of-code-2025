package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day6Challenge2 (file *os.File) {
	var lines []string
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Println(weirderMathResult(lines))
}

func weirderMathResult (lines []string) int {
	var result int
	var numbers []int
	
	lineCount := len(lines)
	lineLength := len(lines[1])
	
	for i := lineLength - 1; i >= 0; i-- {
		var column string
		var columnResult int
		
		for j := 0; j < lineCount - 1; j++ {
			column += string(lines[j][i])
		}

		column = strings.TrimSpace(column)

		number, err := strconv.Atoi(column)
		if err != nil {
			log.Fatal("Could not convert to number ", column, " at column ", i)
		}

		numbers = append(numbers, number)

		if lines[lineCount - 1][i] == ' ' {
			continue
		}

		switch lines[lineCount - 1][i] {
		case '*':
			columnResult = 1
		case '+':
			columnResult = 0
		default:
			log.Fatal("Unexpected operation ", lines[lineCount - 1][i], " at column ", i)
		}

		for _, v := range numbers{
			switch lines[lineCount - 1][i] {
			case '*':
				columnResult *= v
			case '+':
				columnResult += v
			}
		}

		result += columnResult
		i--
		numbers = []int{}
	}

	return result
}
