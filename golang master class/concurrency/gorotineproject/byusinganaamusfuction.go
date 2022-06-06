package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func chekurl(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s id DowM\n", url)
		s += fmt.Sprintf("Errro: %v\n", err)
		fmt.Println("", s)
		c <- url
	} else {
		s := fmt.Sprintf("%s ->status code %d\n", url, resp.StatusCode)

		s += fmt.Sprintf("%s is up\n ", url)
		c <- s
	}

}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}
	c := make(chan string)
	for _, url := range urls {
		go chekurl(url, c)
	}
	fmt.Println("No of goroutines", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)

	// }
  //this is by uisng infimite
	// for {
	// 	go chekurl(<-c, c)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	//or this is by using range 

	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go chekurl(url, c)

	// }

	//or this is by using ananumynfuction

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			go chekurl(u, c)

		}(url)

	}

}
