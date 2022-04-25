package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	const grs = 2
	var wg sync.WaitGroup
	wg.Add(grs)
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {
				value := counter
				value++
				counter = value
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
