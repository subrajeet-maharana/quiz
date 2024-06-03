package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type questionAnswer struct {
	question string
	answer   string
}

func main() {
	strPtr := flag.String("filename", "problems.csv", "CSV File Name")
	// intPtr := flag.Int("time", 30, "Time Limit")
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

	var score int = 0
	for i, row := range data {
		fmt.Printf("%d. %s = ", i+1, row[0])
		var val string
		fmt.Scan(&val)
		if val == row[1] {
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
