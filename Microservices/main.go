package main

import (
	"fmt"
	"net/http"
)

func main() {
	startServer()

}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there!\n")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
		value, ok := query["abc"] //type Values map[string][]string -> returns map of <string, list<string>>

		if !ok || len(value[0]) < 1 {
			fmt.Fprintf(w, "Url Param 'abc' is missing")
			return
		}
		fmt.Fprintf(w, "Url Param 'abc' = %s\n", value[0]) // We pass w parameter to printF to write the response directly to the client
	})
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
