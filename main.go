package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/gocarina/gocsv"
)

type QuestionAnswer struct {
	question string
	answer   string
}

func main() {
	strPtr := flag.String("filename", "problems.csv", "CSV File Name")
	flag.Parse()
	var commandIs string
	if *strPtr != "problems.csv" {
		commandIs = *strPtr + ".csv"
	} else {
		commandIs = *strPtr
	}
	f, err := os.Open(commandIs)
	checkNilErr(err)
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	checkNilErr(err)

	var questions []QuestionAnswer

	for _, d := range data[1:] {
		var q QuestionAnswer
		populateStructFromCSV(d, &q)
		questions = append(questions, q)
	}

	questionAnswers := []*QuestionAnswer{}
	if err := gocsv.UnmarshalFile(f, &questionAnswers); err != nil {
		panic(err)
	}
	var score int = 0
	for i, q := range questionAnswers {
		fmt.Printf("%d. %s ", i+1, q.question)
		reader := bufio.NewReader(os.Stdin)
		val, err := reader.ReadString('\n')
		checkNilErr(err)
		val = val[:len(val)-1]
		fmt.Printf("%s %s\n", val, q.answer)
		if val == q.answer {
			score++
		}
	}
	fmt.Printf("You have scored %d out of %d", score, len(data))
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func populateStructFromCSV(data []string, result interface{}) {
	val := reflect.ValueOf(result).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.CanSet() {
			field.SetString(data[i])
		}
	}
}
