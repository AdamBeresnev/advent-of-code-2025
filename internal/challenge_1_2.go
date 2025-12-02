package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1Challenge2 (file *os.File) {
	var position int
	var target int
	var result int
	fmt.Print("Starting position: ")
	fmt.Scanln(&position)

	fmt.Print("Target position: ")
	fmt.Scanln(&target)

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		clicks := clicksFromRotation(scanner.Text(), &position, &target)
		result += clicks
	}

	fmt.Println(result)
}

func clicksFromRotation (rotation string, position *int, target *int) int {
	var clicks int

	move, err := strconv.Atoi(rotation[1:])
	if err != nil {
		log.Fatal("Could not read line: ", rotation, "\n", err)
	}

	clicks += move / 100
	move = move % 100
	
	switch rotation[0] {
	case 'R':
		if *position + move > 100 {
			clicks++
		}
		*position += move
	case 'L':
		if *position - move < 0 && *position != 0 {
			clicks++
		}
		*position -= move
	default:
		log.Fatal("Incorrect prefix on line: ", rotation)
	}

	*position = (*position + 100) % 100
	
	if position == target {
		clicks++
	}

	return clicks
}