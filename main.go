package main

import (
	"advent-of-code/internal"
	"log"
	"os"
)

func main() {
	challengeNumber := os.Args[1]
	inputFile := os.Args[2]
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatal("Could not open file ", inputFile, "\n", err)
	}
	defer file.Close()

	switch challengeNumber {
	case "11":
		internal.Day1Challenge1(file)
	case "12":
		internal.Day1Challenge2(file)
	case "21":
		internal.Day2Challenge1(file)
	case "22":
		internal.Day2Challenge2(file)
	case "31":
		internal.Day3Challenge1(file)
	case "32":
		internal.Day3Challenge2(file)
	case "41":
		internal.Day4Challenge1(file)
	case "42":
		internal.Day4Challenge2(file)
	case "51":
		internal.Day5Challenge1(file)
	case "52":
		internal.Day5Challenge2(file)
	case "61":
		internal.Day6Challenge1(file)
	case "62":
		internal.Day6Challenge2(file)
	case "71":
		internal.Day7Challenge1(file)
	case "72":
		internal.Day7Challenge2(file)
	default:
		log.Fatal("Incorrect challenge number specified")
	}
}
