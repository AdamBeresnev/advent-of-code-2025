package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Range struct {
	start int
	end int
}

func Day5Challenge1 (file *os.File) {
	var freshList []Range
	var result int

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

		freshList = append(freshList, Range{start: start, end: end})
	}

	for scanner.Scan() {
		line := scanner.Text()

		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Couldn't read starting position in range", line)
		}

		if isFresh(freshList, id){
			result++
		}
	}

	fmt.Println(result)
}

func isFresh (list []Range, id int) bool {
	for _, v := range list {
		if v.start <= id && id <= v.end {
			return true
		}
	}

	return false
}
