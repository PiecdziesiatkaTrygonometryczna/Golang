package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// SharkAttack represents the structure of a shark attack record
type SharkAttack struct {
	Date     string `json:"date"`
	Country  string `json:"country"`
	Injury   string `json:"injury"`
	FatalYN  string `json:"fatal_y_n"`
	Type     string `json:"type"`
	Activity string `json:"activity"`
}

func main() {
	// Open and read the JSON file
	file, err := os.Open("global-shark-attack.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON data into a slice of SharkAttack structs
	var attacks []SharkAttack
	err = json.Unmarshal(data, &attacks)
	if err != nil {
		panic(err)
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Select 10 random records
	selectedRecords := make([]SharkAttack, 10)
	for i := 0; i < 10; i++ {
		selectedRecords[i] = attacks[rand.Intn(len(attacks))]
	}

	// Prepare data for writing to JSON file
	outputData := make([]map[string]string, 10)
	for i, record := range selectedRecords {
		outputRecord := map[string]string{
			"date":      record.Date,
			"country":   record.Country,
			"injury":    record.Injury,
			"fatal_y_n": record.FatalYN,
			"type":      record.Type,
			"activity":  record.Activity,
		}
		outputData[i] = outputRecord
	}

	// Marshal selected data to JSON
	outputJSON, err := json.MarshalIndent(outputData, "", " ")
	if err != nil {
		panic(err)
	}

	// Write selected data to a new JSON file
	err = ioutil.WriteFile("data.json", outputJSON, 0644)
	if err != nil {
		panic(err)
	}
}
