package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{1, 2, 3}

	var wg sync.WaitGroup
	wg.Add(len(s))
	for i, v := range s {
		go func() {
			fmt.Printf("[%d] %d\n", i, v)
			wg.Done()
		}()
	}

	wg.Wait()

	for i, v := range s {
		fmt.Printf("[%d] %d -> %x\n", i, v, &v)
	}
}
