package internal

import {
	"fmt"
	"bytes"
}

func Day2Challenge1 (file *os.File) {
	var result int
	scanner := bufio.NewReader(file)
	
	// File read logic

	fmt.Println(result)
}

func invalidIdsInRange (idRange string) int {
	var invalidIds int

	if idRange[len(idRange) - 1:] == ',' {
		idRange = idRange[:len(idRange) - 1]
	}

	rangeSeparatorPos := bytes.IndexByte(idRange, ':')
	
	rangeStart, err := strconv.Atoi(rotation[:rangeSeparatorPos])
	if err != nil {
		log.Fatal("Couldn't read starting position in range", idRange)
	}

	rangeEnd, err := strconv.Atoi(rotation[rangeSeparatorPos + 1:])
	if err != nil {
		log.Fatal("Couldn't read ending position in range", idRange)
	}

	lenRangeStart := intLength(rangeStart)
	lenRangeEnd := intLength(rangeEnd)

	halfLengthStart := lenRangeStart / 2
	halfLengthEnd := lenRangeEnd / 2

	for i := lenRangeStart; i <= lenRangeEnd; i++ {
		if i % 2 == 0 {
			invalidIds += 9 * 10 ^ (i / 2 - 1)
		}
	}
	
	if lenRangeStart % 2 == 0 {
		invalidIds -= (rangeStart - 10 ^ (lenRangeStart - 1)) / 10 ^ halfLengthStart 

		if rangeStart / 10 ^ halfLengthStart < rangeStart % 10 ^ halfLengthStart {
			invalidIds--
		}
	}
	
	if lenRangeEnd % 2 == 0 {
		invalidIds -= (10 ^ lenRangeEnd - rangeEnd) / 10 ^ halfLengthEnd 

		if rangeEnd / 10 ^ halfLengthEnd > rangeEnd % 10 ^ halfLengthEnd {
			invalidIds--
		}
	}

	return invalidIds
}

func intLength (number int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
