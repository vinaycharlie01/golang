package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f)
		fmt.Println("", x)
		wg.Done()

	}(10.00, &wg)
	wg.Wait()

}
