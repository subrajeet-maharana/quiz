package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//opening file in golang
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

	// reading from csv file
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll() //data will give [][]string
	checkNilErr(err)

	// printing data from csv
	// number of rows --> len(data) and columns --> len(data[0])
	var val string
	var score int = 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < 1; j++ {
			fmt.Printf("%d. %s = ", i+1, data[i][j])
		}
		for j := 0; j < 1; j++ {
			fmt.Scan(&val)
			if val == data[i][1] {
				score++
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("You have scored %d out of %d", score, len(data))
}
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
