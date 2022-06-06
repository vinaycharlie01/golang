package main

import (
	"fmt"
	"time"
)

func main() {
	s := time.Now().UnixNano()
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "hello!"

	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "salut!"

	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("reciverd", msg1)
		case msg2 := <-c2:
			fmt.Println("reciverd", msg2)

		}

	}
	e := time.Now().UnixNano()
	fmt.Println(e - s)

}
