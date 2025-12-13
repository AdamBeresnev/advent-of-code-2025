package internal

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Day8Challenge2(file *os.File) {
	var boxes []*junctionBox
	var lStrings []*lightString
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		newBox := createBox(line)
		lStrings = append(lStrings, boxConnect(&boxes, &newBox)...)
		boxes = append(boxes, &newBox)
	}

	sort.Slice(lStrings, func (i, j int) bool {
		return lStrings[i].distance < lStrings[j].distance
	})

	ammountOfBoxes := len(boxes)
	connectedBoxes := 0

	for _, v := range lStrings {
		connectedBoxes += connectString(v)

		if ammountOfBoxes == connectedBoxes {
			fmt.Println(v.box1.x * v.box2.x)
			break
		}
	}
}

func connectString(lString *lightString) (result int) {
	if !lString.box1.inCircuit {
		lString.box1.inCircuit = true
		result++
	}

	if !lString.box2.inCircuit {
		lString.box2.inCircuit = true
		result++
	}

	return result
}