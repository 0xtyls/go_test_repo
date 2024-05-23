package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Parsing query parameters
	query := r.URL.Query()
	num1Str := query.Get("num1")
	num2Str := query.Get("num2")
	operation := query.Get("operation")

	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		http.Error(w, "Invalid num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		http.Error(w, "Invalid num2", http.StatusBadRequest)
		return
	}

	// Performing the operation
	var result int
	switch operation {
	case "add":
		result = num1 + num2
	case "subtract":
		result = num1 - num2
	case "multiply":
		result = num1 * num2
	case "divide":
		if num2 == 0 {
			http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
			return
		}
		result = num1 / num2
	default:
		http.Error(w, "Unsupported operation", http.StatusBadRequest)
		return
	}

	// Sending the response
	_, err = fmt.Fprintf(w, "Result: %d", result)
	if err != nil {
		return
	}
}
