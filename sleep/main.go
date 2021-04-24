package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkIfWorking() {
	fmt.Println("inside checkIfWorking")
	sleepGoRoutineTest()
	return
}

func sleepGoRoutineTest() {
	fmt.Println("inside sleepGoRoutineTest")
	time.Sleep(30 * time.Second)
	go checkIfWorking()
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my world %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Hello, playground")
	go checkIfWorking()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
