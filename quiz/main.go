package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var flagVal = flag.String("file", "problems.csv", "file name which contains the questions")
	flag.Parse()

	file, err := os.Open(*flagVal)
	if err != nil {
		fmt.Printf("Failed to open file with name: %s\n", *flagVal)
		os.Exit(1)
	}

	// close the file at the end of the program
	defer file.Close()

	r := csv.NewReader(file)

	questionCount := 1
	correctAnswerCount := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)

		}
		question := record[0]
		answer := record[1]
		var userAnswer string

		fmt.Printf("Problem # %d: %v = ?\n", questionCount, question)
		fmt.Scanln(&userAnswer)
		if userAnswer == answer {
			correctAnswerCount++
		}
		questionCount++
	}

	fmt.Printf("You got %v out of %v correct \n", correctAnswerCount, questionCount-1)
}
