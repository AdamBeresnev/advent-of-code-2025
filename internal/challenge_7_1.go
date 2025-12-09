package internal

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
)

func Day7Challenge1 (file *os.File) {
	var result int
	var objectPositions []int

	beams := list.New()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	objectPositions = bytePositionArray([]byte(line), 'S')
	beams.PushBack(objectPositions[0])

	for scanner.Scan() {
		line = scanner.Text()
		objectPositions = bytePositionArray([]byte(line), '^')
		result += beamTravel(objectPositions, beams)
	}

	fmt.Println(result)
}

func beamTravel(splitters []int, beams *list.List) int {
	var result int
	
	currentElement := beams.Front()

	for _, v := range splitters {
		for element := currentElement; element != nil; element = element.Next() {
			beamPosition := element.Value.(int)
			
			if beamPosition != v {
				continue
			}

			result++

			if element.Prev() == nil || element.Prev().Value.(int) != v - 1 {
				beams.InsertBefore(v - 1, element)
			}

			if element.Next() == nil || element.Next().Value.(int) != v + 1 {
				beams.InsertAfter(v + 1, element)
			}

			currentElement = element.Next()
			beams.Remove(element)

			break
		}
	}	
	return result
}

func bytePositionArray (line []byte, target byte) []int {
	var result []int
	var fullIndex int

	for i := bytes.IndexByte(line, target); i != -1; i = bytes.IndexByte(line[fullIndex:], target) {
		fullIndex += i
		result = append(result, fullIndex)
		fullIndex++
	}

	return result
}
