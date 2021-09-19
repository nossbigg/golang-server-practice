package http_server

import (
	"encoding/json"
	"log"
	"net/http"
)

type HandlerMapType = map[string]func(http.ResponseWriter, *http.Request)

var userEndpointHandlers = HandlerMapType{
	"GET": getHandler,
}

func StartHttpServer() {
	http.HandleFunc("/user", userHttpRequestHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func userHttpRequestHandler(w http.ResponseWriter, r *http.Request) {
	matchingHandler, ok := userEndpointHandlers[r.Method]
	if !ok {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	matchingHandler(w, r)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"name": "jeff",
	}
	result, _ := json.Marshal(data)
	w.Write(result)
}
