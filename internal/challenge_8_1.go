package internal

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type junctionBox struct {
	x int
	y int
	z int
	lStrings []*lightString
	inCircuit bool
}

type lightString struct {
	distance float64
	box1 *junctionBox
	box2 *junctionBox
	needsToBeConnected bool
	connected bool
}

func Day8Challenge1(file *os.File, howMuchToConnect string) {
	var boxes []*junctionBox
	var lStrings []*lightString

	numOfCircuits, err := strconv.Atoi(howMuchToConnect)
	if err != nil {
		log.Fatal("Incorrect amount of circuits specified to connect", howMuchToConnect)
	}
	
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

	lStrings = lStrings[:numOfCircuits]
	
	for _, v := range lStrings {
		v.needsToBeConnected = true
	}
	
	var circuitsLengths []int

	for _, ls := range lStrings {
		newCircuit := circuitMap(ls)
		circuitsLengths = append(circuitsLengths, newCircuit)
	}

	sort.Ints(circuitsLengths)

	result := circuitsLengths[numOfCircuits - 1]
	result *= circuitsLengths[numOfCircuits - 2]
	result *= circuitsLengths[numOfCircuits - 3]

	fmt.Println(result)
}

func createBox(line string) (newBox junctionBox) {
	var err error

	coordinates := strings.Split(line, ",")

	newBox.x, err = strconv.Atoi(coordinates[0])
	if err != nil {
		log.Fatal("Error getting X from line ", line)
	}
	
	newBox.y, err = strconv.Atoi(coordinates[1])
	if err != nil {
		log.Fatal("Error getting Y from line ", line)
	}
	
	newBox.z, err = strconv.Atoi(coordinates[2])
	if err != nil {
		log.Fatal("Error getting Z from line ", line)
	}

	return newBox
}

func circuitMap(startingLString *lightString) (circuitLen int) {
	if startingLString.connected {
		return 0
	} 

	startingLString.connected = true

	if !startingLString.box1.inCircuit {
		startingLString.box1.inCircuit = true

		for _, lString := range startingLString.box1.lStrings {
			if lString == startingLString || lString.connected || !lString.needsToBeConnected {
				continue
			}
	
			circuitLen += circuitMap(lString)
		}

		circuitLen++
	}
	
	if !startingLString.box2.inCircuit {
		startingLString.box2.inCircuit = true

		for _, lString := range startingLString.box2.lStrings {
			if lString == startingLString || lString.connected || !lString.needsToBeConnected {
				continue
			}
			
			circuitLen += circuitMap(lString)
		}

		circuitLen++
	}
	
	return circuitLen
}

func boxConnect(boxes *[]*junctionBox, newBox *junctionBox) (newLStrings []*lightString){
	var conLength float64

	for _, box := range *boxes {
		conLength = pointDistance(box, newBox)
		newLString := lightString{
			distance: conLength,
			box1: box,
			box2: newBox,
		}

		box.lStrings = append(box.lStrings, &newLString)
		newBox.lStrings = append(newBox.lStrings, &newLString)
		newLStrings = append(newLStrings, &newLString)
	}

	return newLStrings
}

func pointDistance(point1 *junctionBox, point2 *junctionBox) float64 {
	return math.Sqrt(
		float64(point1.x - point2.x) * float64(point1.x - point2.x) +
		float64(point1.y - point2.y) * float64(point1.y - point2.y) +
		float64(point1.z - point2.z) * float64(point1.z - point2.z),
	)
}
