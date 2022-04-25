package main

import (
	"fmt"
	"math/rand"
	"time"
)

func waitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()
	d := <-ch
	fmt.Println("parent : recv'd signal :", d)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func main() {
	waitForResult()
}
