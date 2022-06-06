// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func deposit(b *int, n int, wg *sync.WaitGroup, mt *sync.Mutex) {
// 	mt.Lock()
// 	*b += n
// 	mt.Unlock()

// 	wg.Done()
// }

// func withdraw(b *int, n int, wg *sync.WaitGroup, mt *sync.Mutex) {
// 	mt.Lock()
// 	*b -= n
// 	mt.Unlock()
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var mt sync.Mutex

// 	wg.Add(200)

// 	balance := 100

// 	for i := 0; i < 100; i++ {
// 		go deposit(&balance, i, &wg, &mt)
// 		go withdraw(&balance, i, &wg, &mt)
// 	}
// 	wg.Wait()

// 	fmt.Println("Final balance value:", balance)
// }
package main

import (
	"fmt"
	"sync"
)

// pass pointer to sync.Mutex (mutex is a value type)
func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b -= n
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	// the final balance value be always 100
	fmt.Println("Final balance value:", balance)
}
