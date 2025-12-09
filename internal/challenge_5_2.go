package internal

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day5Challenge2 (file *os.File) {
	var result int
	freshList := list.New()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		
		if line == "" {
			break
		}
		
		rangeSeparatorPos := bytes.IndexByte([]byte(line), '-')
		
		start, err := strconv.Atoi(line[:rangeSeparatorPos])
		if err != nil {
			log.Fatal("Couldn't read starting position in range", line)
		}
		
		end, err := strconv.Atoi(line[rangeSeparatorPos + 1:])
		if err != nil {
			log.Fatal("Couldn't read ending position in range", line)
		}

		addFreshRange(freshList, start, end)
	}

	var currRange Range
	
	for element := freshList.Front(); element != nil; element = element.Next() {
		currRange = element.Value.(Range)

		result += currRange.end - currRange.start + 1
	}

	fmt.Println(result)
}

func addFreshRange(list *list.List, start int, end int) {
	var currRange Range
	
	for element := list.Front(); element != nil; element = element.Next() {
		currRange = element.Value.(Range)
		
		if currRange.end < start && end > currRange.end {
			continue
		} 
		
		if currRange.start > start && end < currRange.start {
			list.InsertBefore(Range{start: start, end: end}, element)
			return
		}
		
		if currRange.start < start && end < currRange.end {
			return
		}

		currRange.start = min(start, currRange.start)
		currRange.end = max(end, currRange.end)

		currRange.start = joinLeft(list, element, currRange.start)
		currRange.end = joinRight(list, element, currRange.end)

		element.Value = currRange

		return
	}

	list.PushBack(Range{start: start, end: end})
}

func joinLeft(freshList *list.List, startingElement *list.Element, start int) int {
	var currRange Range
	var garbage []*list.Element

	leftBoundry := start
	
	for element := startingElement.Prev(); element != nil; element = element.Prev() {
		currRange = element.Value.(Range)

		if start > currRange.end {
			break
		}

		leftBoundry = currRange.start
		garbage = append(garbage, element)
		
		if start > currRange.start {
			break
		}
	}

	for _, e := range garbage {
		freshList.Remove(e)
	}

	return leftBoundry
}

func joinRight(freshList *list.List, startingElement *list.Element, end int) int {
	var currRange Range
	var garbage []*list.Element

	rightBoundry := end
	
	for element := startingElement.Next(); element != nil; element = element.Next() {
		currRange = element.Value.(Range)

		if end < currRange.start {
			break
		}

		rightBoundry = currRange.end
		garbage = append(garbage, element)
		
		if end < currRange.end {
			break
		}
	}

	for _, e := range garbage {
		freshList.Remove(e)
	}

	return rightBoundry
}