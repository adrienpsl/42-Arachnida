package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const URL = "https://www.medical.fr/equipements-hospitaliers/chariot-hospitalier?page=2"
const DEEP = 2

type Data struct {
	Images []string
	Links  []string
}

func main() {
	// check deep = 0

	var data = make(map[int]Data, DEEP)
	data = LoopOnLinks(data, URL, DEEP)

	// Convert the data map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Write the JSON data to a file
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("JSON data written to data.json")
}
