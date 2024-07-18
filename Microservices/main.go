package main

import (
	"encoding/json"
	"fmt"
	"microservice/model"
	"net/http"
)

func main() {
	startServer()

}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there!\n")
	})
	http.HandleFunc("/api", handler)

	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		executeGetMethod(r, w)
	case http.MethodPost:
		executePostMethod(r, w)
	default:
	}
}

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
