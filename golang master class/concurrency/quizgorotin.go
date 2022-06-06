// package main

// import (
// 	"fmt"
// 	"math"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	for i := 100; i < 150; i++ {
// 		go func(f float64, wg *sync.WaitGroup) {
// 			x := math.Sqrt(f)
// 			fmt.Println(x)
// 			wg.Done()

// 		}(float64(i), &wg)

// 	}
// 	wg.Wait()

// }

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(50)

	// IMPORTANT: i starts from 100. not 100
	// i is float64, not int. math.Sqrt() takes in a float64.
	for i := 100.; i < 150.; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			x := math.Sqrt(f)
			fmt.Printf("Square root of %.2f is %.6f\n", f, x)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
