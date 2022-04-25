package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	var ch chan string
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %v\n", r)
		}
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()

	d := <-ch
	fmt.Println("parent : recv'd signal :", d)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
