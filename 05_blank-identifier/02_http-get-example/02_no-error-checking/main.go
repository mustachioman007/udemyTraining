package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	res, _ := http.Get("http://www.powerlinks.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("%s", page)
}