package main

import (
	"encoding/json"
	"fmt"
	"microservice/model"
	"microservice/model/operation"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	startServer()

}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there!\n")
	})
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/operation", operationHandler)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func operationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		performOperation(w, r)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		executeGetMethod(r, w)
	case http.MethodPost:
		executePostMethod(r, w)
	default:
	}
}

// ================================ API ============================

func executeGetMethod(r *http.Request, w http.ResponseWriter) {
	query := r.URL.Query()
	value, ok := query["abc"]

	if !ok || len(value[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'abc' is missing")
		return
	}
	fmt.Fprintf(w, "Url Param 'abc' = %s\n", value[0])
}

func executePostMethod(r *http.Request, w http.ResponseWriter) {
	var requestData model.RequestData2

	// json decoding
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	requestData.PrintMe()
	found := false
	for _, color := range requestData.Colors {
		if color.Band == requestData.BandToFind {
			found = true
			break
		}
	}
	response := model.ResponseData{
		Status: found,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
// ================================ Operation ============================

func performOperation(w http.ResponseWriter, r *http.Request) {
	var operation operation.OperationRequest
	err := json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	// Validate the request data
	validate := validator.New()
	err = validate.Struct(operation)
	if err != nil {
		http.Error(w, fmt.Sprintf("Validation failed: %s", err.Error()), http.StatusBadRequest)
		return
	}

	str := operation.String()
	fmt.Printf("Received operation: %s", str)
	switch operation.Operation {
	case "add":
		result := add(operation.Operands)
		json.NewEncoder(w).Encode(result)
	case "subtract":
		result := subtract(operation.Operands)
		json.NewEncoder(w).Encode(result)
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
	}
}

func add(s []int) int {
	result := s[0] + s[1]
	return result
}

func subtract(s []int) int {
	result := s[0] - s[1]
	return result
}

