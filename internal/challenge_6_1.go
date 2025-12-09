package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day6Challenge1 (file *os.File) {
	var lineContents []string
	var lineNumbers []int 
	var allNumbers [][]int
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineContents = lineToArray(line)
		
		if line[0] == '*' || line[0] == '+' {
			break
		}
		
		lineNumbers = contentsToIntArray(lineContents)
		allNumbers = append(allNumbers, lineNumbers)
	}

	fmt.Println(weirdMathResult(lineContents, allNumbers))
}

func weirdMathResult (operations []string, numbers [][]int) int {
	var result int
	var lineResult int
	
	lineCount := len(numbers)
	
	for i, operation := range operations {
		switch operation {
		case "*":
			lineResult = 1
		case "+":
			lineResult = 0
		default:
			log.Fatal("Unexpected operation", operation, "at position", i)
		}

		for j := 0; j < lineCount; j++ {
			switch operation {
			case "*":
				lineResult *= numbers[j][i]
			case "+":
				lineResult += numbers[j][i]
			}
		}

		result += lineResult
	}

	return result
}

func lineToArray (line string) []string {
	var result []string
	var currString string

	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			if currString != "" {
				result = append(result, currString)
				currString = ""
			}

			continue
		}

		currString += string(line[i])
	}

	if currString != "" {
		result = append(result, currString)
	}

	return result
}

func contentsToIntArray (contents []string) []int {
	var result []int

	for _, v := range contents {
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Couldn't read convert to int", v)
		}

		result = append(result, number)
	}

	return result
}