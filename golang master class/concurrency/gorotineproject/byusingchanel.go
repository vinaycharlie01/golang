package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func chekAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s id DowM\n", url)
		s += fmt.Sprintf("Errro: %v\n", err)
		c <- s
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s ->status code %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			boybyte, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			s += fmt.Sprintf("writing responce body to %s\n", file)
			err = ioutil.WriteFile(file, boybyte, 0664)
			if err != nil {

				log.Fatal(err)
				s += "Error writing file\n"
				c <- s
			}
			s += fmt.Sprintf("%s is up\n ", url)
			c <- s
		}

	}

}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}
	c := make(chan string)
	for _, url := range urls {
		go chekAndSaveBody(url, c)
		fmt.Println(strings.Repeat("#", 20))

	}
	fmt.Println("No of goroutines", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)

	}

}
