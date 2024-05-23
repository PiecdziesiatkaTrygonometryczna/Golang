package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type SharkAttack struct {
	Date     string `json:"date"`
	Country  string `json:"country"`
	Injury   string `json:"injury"`
	FatalYN  string `json:"fatal_y_n"`
	Type     string `json:"type"`
	Activity string `json:"activity"`
}

func main() {
	file, err := os.Open("global-shark-attack.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var attacks []SharkAttack
	err = json.Unmarshal(data, &attacks)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	selectedRecords := make([]SharkAttack, 10)
	for i := 0; i < 10; i++ {
		selectedRecords[i] = attacks[rand.Intn(len(attacks))]
	}

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

	outputJSON, err := json.MarshalIndent(outputData, "", " ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("data.json", outputJSON, 0644)
	if err != nil {
		panic(err)
	}
}
