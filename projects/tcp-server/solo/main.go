package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	rw.Write([]byte("Hello, Welcome to Harish's webpage!\n"))
}

func schoolFriends(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	rw.Write([]byte("Sundar\nGanesh\nTarun\nZenith\n"))
}

func bTechFriends(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Uchit\nBeauty\nAlphin\n"))
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/schoolfriends/", schoolFriends)
	http.HandleFunc("/collegefriends", bTechFriends)

	fmt.Println("Starting a HTTP server on localhost and 6543 port")
	fmt.Println("To kill the server : Ctrl + C ")
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":6543", nil))
}
