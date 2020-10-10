package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Vachira",
	}

	p2 := person{
		First: "Nut",
	}
	xp := []person{p1, p2}
	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		log.Println("Bad data encode", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {

	people := []person{}

	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("Bad data to decode", err)
	}

	fmt.Println(people)
}
