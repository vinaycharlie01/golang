package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func chekAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s id DowM\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s ->status code %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			boybyte, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("writing responce body to %s\n", file)
			err = ioutil.WriteFile(file, boybyte, 0664)
			if err != nil {

				log.Fatal(err)
			}
		}

	}
	wg.Done()

}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go chekAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))

	}
	fmt.Println("No of goroutines", runtime.NumGoroutine())
	wg.Wait()
}
