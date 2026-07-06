package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	result := Result{
		Success: true,
		Message: "Go parser is running",
	}

	data, _ := json.Marshal(result)

	fmt.Println(string(data))
}