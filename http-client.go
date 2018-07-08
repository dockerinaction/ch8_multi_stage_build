package main

import (
	"net/http"
)
import "io/ioutil"
import "fmt"

func main() {
	resp, err := http.Get("https://google.com/search?q=docker")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response:\n", string(body))
}
