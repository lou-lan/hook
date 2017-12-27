package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		fmt.Println("err")
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
