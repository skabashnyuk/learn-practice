package main

import (
	"fmt"
	"net/http"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s -> %s\n", url, err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}
	ch := make(chan string)

	go func() {
		for _, url2 := range urls {
			ch <- url2
		}
		close(ch)
	}()
	for i := range ch {
		returnType(i)
	}

}
