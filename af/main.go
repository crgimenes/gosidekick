package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", b)

}
