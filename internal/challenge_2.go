package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Challenge2 (file *os.File) {
	var position int
	var target int
	var result int
	fmt.Print("Starting position: ")
	fmt.Scanln(&position)

	fmt.Print("Target position: ")
	fmt.Scanln(&target)

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		move, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Could not read line: ", line, "\n", err)
		}

		result += move / 100
		move = move % 100
		
		switch line[0] {
		case 'R':
			if position + move > 100 {
				result++
			}
			position += move
		case 'L':
			if position - move < 0 && position != 0 {
				result++
			}
			position -= move
		default:
			log.Fatal("Incorrect prefix on line: ", line)
		}

		position = (position + 100) % 100 
		
		if position == target {
			result++
		}
	}

	fmt.Println(result)
}