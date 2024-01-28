package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(ctype)
	}

}

func contentType(url string) (string, error) {

	//FIXME
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", errors.New("can't find Content-Type header")
	}
	return resp.Header.Get("Content-Type"), nil
}
