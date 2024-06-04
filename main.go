package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type QuestionAnswer struct {
	Question string
	Answer   string
}

func main() {
	strPtr := flag.String("filename", "problems.csv", "CSV File Name")
	flag.Parse()
	var commandIs string = *strPtr
	if *strPtr != "problems.csv" {
		commandIs = *strPtr + ".csv"
	}

	f, err := os.Open(commandIs)
	checkNilErr(err)
	defer f.Close()

	stat, err := f.Stat()
	checkNilErr(err)
	if stat.Size() == 0 {
		log.Fatal("empty CSV file given")
	}

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	checkNilErr(err)

	if len(data) == 0 {
		log.Fatal("no questions found in the CSV file")
	}

	var questionAnswers []QuestionAnswer
	for _, row := range data {
		if len(row) != 2 {
			log.Fatal("invalid CSV format: each row must contain exactly two columns")
		}
		questionAnswers = append(questionAnswers, QuestionAnswer{
			Question: row[0],
			Answer:   row[1],
		})
	}

	var score int
	reader := bufio.NewReader(os.Stdin)
	for i, q := range questionAnswers {
		fmt.Printf("%d. %s ", i+1, q.Question)
		val, err := reader.ReadString('\n')
		checkNilErr(err)
		val = strings.TrimSpace(val)
		if val == q.Answer {
			score++
		}
	}

	fmt.Printf("You have scored %d out of %d\n", score, len(questionAnswers))
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
