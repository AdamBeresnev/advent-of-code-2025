package internal

import (
	"bufio"
	"fmt"
	"os"
)

func Day4Challenge1 (file *os.File) {
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


	result := movableRolls(&allRolls, gridX, gridY)
	fmt.Println(result)
}

func movableRolls (allRolls *[]byte, gridX int, gridY int) int {
	var result int
	var adjacentSpaces []int

	maxIndex := gridX * gridY - 1

	for i := 0; i <= maxIndex; i++ {
		adjacentSpaces = surroundingPositions(i, gridX)

		if (*allRolls)[i] <= '.' {
			for _, v := range adjacentSpaces {
				if v < 0 || v > maxIndex {
					continue
				}
				
				(*allRolls)[v]--
			}
			continue
		}

		(*allRolls)[i] -= byte(8 - len(adjacentSpaces))

		for _, v := range surroundingPositions(i, gridX) {
			if v < 0 || v > maxIndex {
				(*allRolls)[i]--
			}
			
		}		
	}
	
	for i := 0; i <= maxIndex; i++ {
		if 55 < (*allRolls)[i] && (*allRolls)[i] < 60 {
			result++
		}

		if (*allRolls)[i] < 60 {
			(*allRolls)[i] = '.'
		} else {
			(*allRolls)[i] = '@'
		}
	}

	return result
}

func surroundingPositions (i int, gridX int) []int {
	positions := []int{
		i + gridX,
		i - gridX,
	}
	left := []int{
		i - 1,
		i + gridX - 1,
		i - gridX - 1,
	}
	right := []int{
		i + 1,
		i + gridX + 1,
		i - gridX + 1,
	}

	if i % gridX != 0 {
		positions = append(positions, left...)
	}

	if (i + 1) % gridX != 0 {
		positions = append(positions, right...)
	}

	return positions
} 