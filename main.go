package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("错误")
	}
	fmt.Println(string(b))
}

func main() {
	http.HandleFunc("/", webhook)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
