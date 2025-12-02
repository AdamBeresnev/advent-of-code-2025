package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1Challenge1 (file *os.File) {
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
		reached, position := zeroReached(scanner.Text(), position, target)
		if reached {
			result++
		}
	}

	fmt.Println(result)
}

func zeroReached (rotation string, position int, target int) (reached bool, position int) {
	move, err := strconv.Atoi(rotation[1:])
	if err != nil {
		log.Fatal("Could not read line: ", rotation, "\n", err)
	}

	switch rotation[0] {
	case 'R':
		position += move
	case 'L':
		position -= move
	default:
		log.Fatal("Incorrect prefix on line: ", rotation)
	}

	position = position % 100

	return position == target, position
}