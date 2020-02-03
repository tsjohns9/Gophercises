package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Problem defines a problem
type Problem struct {
	question string
	answer   int
	correct  bool
}

// AllProblems defines some problems
type AllProblems []Problem

func main() {
	file := flag.String("file", "", "CSV file to parse")
	time := flag.Int("time", 10, "Time limit for the quiz")
	flag.Parse()

	filePointer := *file

	if !(strings.Contains(filePointer, ".csv")) {
		fmt.Println("Must be a .csv file")
		return
	}

	c := make(chan string)

	go startTimer(*time, c)

	problemsArr := createProblems(readFile(filePointer))

	fmt.Println("GoLang Quiz!")

	askQuestions(problemsArr)
}

func readFile(file string) [][]string {
	first := file[:1]
	fmt.Println(first)
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Unable to read data from", file)
		os.Exit(1)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	lines, errors := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading csv file: ", errors)
		os.Exit(1)
	}

	return lines
}

func createProblems(problems [][]string) AllProblems {
	var arrProblems AllProblems
	for _, problem := range problems {
		arrProblems = append(arrProblems, Problem{
			question: problem[0],
			answer:   parseInt(problem[1]),
		})
	}
	return arrProblems
}

// TODO: add option to shuffle questions
func askQuestions(problemsArr AllProblems) {
	reader := bufio.NewReader(os.Stdin)
	for i := range problemsArr {
		p := &problemsArr[i]
		fmt.Printf("Problem %+v: %+v = ", i+1, p.question)

		text, _ := reader.ReadString('\n')

		s := strings.Trim(strings.ToLower(text[:len(text)-1]), " ")
		// TODO: allow string answeers
		p.correct = parseInt(s) == p.answer
	}
	problemsArr.getScore()
	fmt.Println()
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error parsing to integer: ", err)
		os.Exit(1)
	}
	return i
}

func (all AllProblems) getScore() {
	sum := 0
	for _, p := range all {
		if p.correct {
			sum++
		}
	}
	fmt.Printf("You answered %+v out of %+v correct!", sum, len(all))
}

func startTimer(limit int, c chan string) {
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	<-timer.C
	fmt.Println()
	fmt.Println("Time up!")
	os.Exit(1)
}
