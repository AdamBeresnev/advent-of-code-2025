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

func Day2Challenge2 (file *os.File) {
	var result int
	reader := csv.NewReader(file)
	idRangeList, err := reader.ReadAll()	
	if err != nil {
		log.Fatal("Could not read input\n", err)
	}
	idRanges := idRangeList[0]

	for _, idRange := range idRanges {
		rangeSeparatorPos := bytes.IndexByte([]byte(idRange), '-')
	
		start, err := strconv.Atoi(idRange[:rangeSeparatorPos])
		if err != nil {
			log.Fatal("Couldn't read starting position in range", idRange)
		}
		
		end, err := strconv.Atoi(idRange[rangeSeparatorPos + 1:])
		if err != nil {
			log.Fatal("Couldn't read ending position in range", idRange)
		}

		result += invalidIdSumByPattern(start, end)
	}

	fmt.Println(result)
}

func invalidIdSumByPattern (start int, end int) int {
	var sum int
	var sliceStart int
	var sliceEnd int

	lenStart := intLength(start)
	lenEnd := intLength(end)

	for i := lenStart; i <= lenEnd; i++ {
		for j := 2; j <= lenEnd; j++ {
			if i % j != 0 {
				continue
			}

			sliceStart = max(start, int(math.Pow10(i - 1)))
			sliceEnd = min(end, int(math.Pow10(i)) - 1)
			
			sum += invalidIdPatternSum(sliceStart, sliceEnd, j, i/j)
		}
	}
	return sum
}

func invalidIdPatternSum (start int, end int, splits int, patternLength int) int {
	var sum int
	var invalidId int

	patternStart := int(math.Pow10(patternLength - 1))
	patternEnd := int(math.Pow10(patternLength)) - 1

	for i := patternStart; i <= patternEnd; i++ {
		for j := splits; j <= 1; j++ {
			// if i > 
		}
	}

	return sum
}
