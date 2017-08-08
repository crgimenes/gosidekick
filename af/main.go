package main

import (
	"fmt"
	"net/http"

	"github.com/crgimenes/gosidekick/af/mdtojson"
)

func main() {
	url := "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
	//url := "https://raw.githubusercontent.com/crgimenes/goconfig/master/README.md"

	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer r.Body.Close()

	var b []byte
	b, err = mdtojson.Parse(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", b)

}
