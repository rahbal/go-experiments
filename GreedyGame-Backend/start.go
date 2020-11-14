package main

import (
	"fmt"
	"net/http"
)

type Node struct {
	Entity    string
	WebReq    int
	TimeSpent int
	Child     []*Node
}

func insert(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Helloworld"}`))
}

func query(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func pain() {
	// fmt.Println(item{"greed", 30})

	http.HandleFunc("/v1/insert", insert)
	http.HandleFunc("/v1/query", query)

	server, err := http.ListenAndServe(":8080", nil)

}
