package internal

import (
	"bufio"
	"fmt"
	"os"
)

func Day4Challenge2 (file *os.File) {
	var moved int
	var result int
	var gridX int
	var gridY int
	var allRolls []byte

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		gridY++
		allRolls = append(allRolls, line...)
	}
	gridX = len(allRolls) / gridY

	moved = movableRolls(&allRolls, gridX, gridY)

	for moved > 0 {
		result += moved
		moved = movableRolls(&allRolls, gridX, gridY)
	}

	fmt.Println(result)
}
