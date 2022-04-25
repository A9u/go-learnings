// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func printOdd(outputCh, waitCh chan int, n int) {
	fmt.Println("Inside printOdd")
	for i := 1; i <= n; i += 2 {
		<-waitCh
		outputCh <- i
		waitCh <- 1
	}
}

func printEven(outputCh, waitCh chan int, n int) {
	fmt.Println("Inside printEven")
	for i := 2; i < n+1; i += 2 {
		<-waitCh
		outputCh <- i
		waitCh <- 1
	}
}

func main() {
	outputCh := make(chan int)
	waitCh := make(chan int)

	go printEven(outputCh, waitCh, 5)
	go printOdd(outputCh, waitCh, 5)

	waitCh <- 1
	for i := 0; i < 5; i++ {
		fmt.Println(<-outputCh)
	}
}
