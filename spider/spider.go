package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const URL = "https://www.medical.fr/equipements-hospitaliers/chariot-hospitalier?page=2"

// const URL = "https://dzone.com/articles/batch-processing-in-go"
const DEEP = 3

type Data struct {
	Images []string
	Links  []string
}

func main() {
	// !!test check deep = 0

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

//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//
//func mainFibo() {
//	c := make(chan int, 10)
//	go fibonacci(cap(c), c)
//}
//
//// ///////////////////////////////////////////////////////////////
//func calculateAll() {
//	s := []int{7, 6, 8, 1, 8, 8, 19}
//	c := make(chan int)
//	go sum(s[len(s)/2:], c)
//	go sum(s[:len(s)/2], c)
//	y := <-c + <-c
//	fmt.Println(y)
//}
//
//func sum(s []int, c chan int) {
//	sum := 0
//	for _, value := range s {
//		sum += value
//	}
//	c <- sum
//}
