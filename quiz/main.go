package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var fileInput = flag.String("file", "problems.csv", "file name which contains the questions")
	var timeDuration = flag.Int("time", 30, "time for how long the timer should run")
	flag.Parse()
	fmt.Printf("timeDuration: %v\n", *timeDuration)
	file, err := os.Open(*fileInput)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file with name: %s\n", *fileInput))
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

	go func() {
		delaySecond(time.Duration(*timeDuration))
		exit(fmt.Sprint("Time has run out"))
	}()

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

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}
