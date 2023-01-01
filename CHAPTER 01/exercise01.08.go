package main

import "fmt"

func main() {
	//declaring
	query, limit, offset := "bat", 10, 0
	//changing values
	query, limit, offset = "ball", 35, 9

	fmt.Println(query, limit, offset)
}