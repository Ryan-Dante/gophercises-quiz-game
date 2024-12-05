package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Define the shuffle flag
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz questions")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")

	// Parse the flags
	flag.Parse()

	ticker := time.NewTicker(time.Duration(*timeLimit) * time.Second)
	defer ticker.Stop()

	// Open the csv file
	quizcsv, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new reader
	reader := csv.NewReader(quizcsv)
	// Created a slice of slices to store the records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the function
	defer quizcsv.Close()

	// Welcome messages
	fmt.Println("Welcome to the quiz game!")
	fmt.Println("Press enter to start the quiz")
	fmt.Scanln()

	// Start the timer
	go func() {
		<-ticker.C
		fmt.Println("\nTime's up!")
		os.Exit(0)
	}()

	// Randomize the order of the questions if shuffle flag is set
	if *shuffle {
		for i := range records {
			j := rand.Intn(i + 1)
			records[i], records[j] = records[j], records[i]
		}
	}

	// Start the quiz
	score := 0

	for i, record := range records {

		// Assign the question and answer to variables based on the record slice
		question := record[0]
		answer := TrimAndLower(record[1])

		fmt.Printf("Question %d: %s\n", i+1, question)
		fmt.Print("Answer: ")
		var response string
		fmt.Scanln(&response)
		// Clean and trim the response string, accept both upper and lower case answers
		response = TrimAndLower(response)

		if response == answer {
			score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}

		if i == len(records)-1 {
			fmt.Println("Quiz complete!")
			fmt.Printf("You scored %d out of %d\n", score, len(records))
		}
	}
}

// TrimAndLower trims and lowercases a string
func TrimAndLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
