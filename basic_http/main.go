package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	var helloRequest struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(req.Body).Decode(&helloRequest); err != nil {
		fmt.Fprintf(w, `{"message":"error decode json input"}`)
		return
	}

	var helloResponse struct {
		ServerStatus string `json:"server_status"`
		Message      string `json:"message"`
	}

	helloResponse.ServerStatus = "Normal"
	helloResponse.Message = fmt.Sprintf("Hello %s", helloRequest.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(helloResponse)

}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":7777", nil)
}
