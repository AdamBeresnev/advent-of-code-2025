package internal

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Day2Challenge1 (file *os.File) {
	var result int
	reader := csv.NewReader(file)
	idRangeList, err := reader.ReadAll()	
	if err != nil {
		log.Fatal("Could not read input\n", err)
	}
	idRanges := idRangeList[0]

	for _, idRange := range idRanges {
		result += invalidIdSum(idRange)
	}

	fmt.Println(result)
}

func invalidIdSum (idRange string) int {
	var sum int
	
	rangeSeparatorPos := bytes.IndexByte([]byte(idRange), '-')
	
	rangeStart, err := strconv.Atoi(idRange[:rangeSeparatorPos])
	if err != nil {
		log.Fatal("Couldn't read starting position in range", idRange)
	}
	
	rangeEnd, err := strconv.Atoi(idRange[rangeSeparatorPos + 1:])
	if err != nil {
		log.Fatal("Couldn't read ending position in range", idRange)
	}

	lenRangeStart := intLength(rangeStart)
	lenRangeEnd := intLength(rangeEnd)

	var multiplier int
	var invalidId int
	for i := lenRangeStart; i <= lenRangeEnd; i++ {
		if i % 2 != 0 {
			continue
		}
		
		multiplier = int(math.Pow10(i/2))

		for o := multiplier/10; o < multiplier; o++ {
			invalidId = o * multiplier + o
			
			if invalidId > rangeEnd {
				break
			}
			
			if invalidId >= rangeStart {
				sum += invalidId
			}
		}
	}
	return sum
}

// func invalidIdsInRange (idRange string) int {
// 	var invalidIds int

// 	rangeSeparatorPos := bytes.IndexByte([]byte(idRange), '-')
	
// 	rangeStart, err := strconv.Atoi(idRange[:rangeSeparatorPos])
// 	if err != nil {
// 		log.Fatal("Couldn't read starting position in range", idRange)
// 	}
	
// 	rangeEnd, err := strconv.Atoi(idRange[rangeSeparatorPos + 1:])
// 	if err != nil {
// 		log.Fatal("Couldn't read ending position in range", idRange)
// 	}
	
// 	lenRangeStart := intLength(rangeStart)
// 	lenRangeEnd := intLength(rangeEnd)
	
// 	halfLengthStart := lenRangeStart / 2
// 	halfLengthEnd := lenRangeEnd / 2
	
// 	for i := lenRangeStart; i <= lenRangeEnd; i++ {
// 		if i % 2 == 0 {
// 			invalidIds += 9 * int(math.Pow10(i / 2 - 1))
// 		}
// 	}
	
// 	if lenRangeStart % 2 == 0 {
// 		invalidIds -= (rangeStart - int(math.Pow10(lenRangeStart - 1))) / int(math.Pow10(halfLengthStart)) 

// 		if rangeStart / int(math.Pow10(halfLengthStart)) < rangeStart % int(math.Pow10(halfLengthStart)) {
// 			invalidIds--
// 		}
// 	}
	
// 	if lenRangeEnd % 2 == 0 {
// 		invalidIds -= (int(math.Pow10(lenRangeEnd)) - rangeEnd) / int(math.Pow10(halfLengthEnd)) 

// 		if rangeEnd / int(math.Pow10(halfLengthEnd)) > rangeEnd % int(math.Pow10(halfLengthEnd)) {
// 			invalidIds--
// 		}
// 	}
// 	return invalidIds
// }

func intLength (number int) int {
	if number == 0 {
		return 1
	}
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}
