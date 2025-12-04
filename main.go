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
	case "1":
		internal.Day1Challenge1(file)
	case "2":
		internal.Day1Challenge2(file)
	case "3":
		internal.Day2Challenge1(file)
	case "4":
		internal.Day2Challenge2(file)
	case "5":
		internal.Day3Challenge1(file)
	case "6":
		internal.Day3Challenge2(file)
	default:
		log.Fatal("Incorrect challenge number specified")
	}
}
