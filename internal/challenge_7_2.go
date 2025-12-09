package internal

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
)

type Beam struct {
	position int
	left *Beam
	right *Beam
}


func Day7Challenge2 (file *os.File) {
	var objectPositions []int
	var rootBeam *Beam

	beams := list.New()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	rootBeam = &Beam{
		position: bytes.IndexByte([]byte(line), 'S'),
	}
	beams.PushBack(rootBeam)

	for scanner.Scan() {
		line = scanner.Text()
		objectPositions = bytePositionArray([]byte(line), '^')
		beamQuantumTravel(objectPositions, beams)
	}
	fmt.Println("time to count")

	fmt.Println(timelineCounter(rootBeam))
}

func timelineCounter (beam *Beam) int {
	var result int

	if beam.left == nil && beam.right == nil {
		return 1
	}

	result += timelineCounter(beam.left)
	result += timelineCounter(beam.right)

	return result
}

func beamQuantumTravel(splitters []int, beams *list.List) {
	currentElement := beams.Front()
	var newBeam *Beam

	for _, v := range splitters {
		
		for element := currentElement; element != nil; element = element.Next() {
			beam := element.Value.(*Beam)
			
			if beam.position != v {
				continue
			}

			if element.Prev() == nil || element.Prev().Value.(*Beam).position != v - 1 {
				newBeam = &Beam{position: v - 1}
				beam.left = newBeam
				beams.InsertBefore(newBeam, element)
			} else {
				beam.left = element.Prev().Value.(*Beam)
			}

			if element.Next() == nil || element.Next().Value.(*Beam).position != v + 1 {
				newBeam = &Beam{position: v + 1}
				beam.right = newBeam
				beams.InsertAfter(newBeam, element)
			} else {
				beam.right = element.Next().Value.(*Beam)
			}

			currentElement = element.Next()
			beams.Remove(element)

			break
		}
	}
}
