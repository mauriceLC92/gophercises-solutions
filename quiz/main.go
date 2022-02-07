package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var flagVal = flag.String("file", "problems.csv", "file name which contains the questions")
	flag.Parse()

	file, err := os.Open(*flagVal)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file with name: %s\n", *flagVal))
		os.Exit(1)
	}

	// close the file at the end of the program
	defer file.Close()

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprint("Failed parsing the provided CSV file"))
	}
	problems := parseLines(lines)

	correctAnswerCount := 0
	for i, problem := range problems {
		fmt.Printf("Problem # %d: %v = ?\n", i+1, problem.question)
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if userAnswer == problem.answer {
			correctAnswerCount++
		}
	}

	fmt.Printf("You got %v out of %v correct \n", correctAnswerCount, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
