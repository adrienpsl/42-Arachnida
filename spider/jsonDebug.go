package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveJsonDebug(data Data) error {
	logger.Debug("save data to file")
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return nil
	}
	// Write the JSON data to a file
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return nil
	}

	return err
}
