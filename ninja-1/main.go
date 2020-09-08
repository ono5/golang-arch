package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	persons := []Person{{First: "Jenny"}, {First: "James"}}
	err := json.NewEncoder(w).Encode(persons)
	if err != nil {
		log.Println("Encoded bad data!", persons)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var persons []Person
	err := json.NewDecoder(r.Body).Decode(&persons)
	if err != nil {
		log.Println("Decoded bad data", persons)
	}
	log.Println(persons)
}
