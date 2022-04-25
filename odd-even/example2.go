// You can edit this code!
// Click here and start typing.
package main

import "fmt"

import "runtime"
import "time"
import "sync"

func printOdd(n int, wg *sync.WaitGroup) {
	for i := 1; i <= n; i += 2 {
		fmt.Println("Inside print odd before:", i, time.Now())
		time.Sleep(1 * time.Second)
		fmt.Println("Inside print odd after:", i, time.Now())
	}

	wg.Done()
}

func printEven(n int, wg *sync.WaitGroup) {
	for i := 2; i <= n; i += 2 {
		fmt.Println("Inside print even before:", i, time.Now())
		time.Sleep(1 * time.Second)
		fmt.Println("Inside print even after:", i, time.Now())
	}
	wg.Done()
}

func test1() {
	fmt.Println("Inside test1")
}

func test2() {
	fmt.Println("Inside test2")
}

func main() {
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup

	wg.Add(2)
	go printEven(10, &wg)
	go printOdd(10, &wg)

	wg.Wait()

	time.Sleep(4 * time.Second)
}
