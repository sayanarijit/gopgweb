package app

/*
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopgweb#gopgweb
*/

import (
	"log"
	"net/http"
	"strconv"
)

// ServeHTTP starts HTTP server
func ServeHTTP(host string, port int) {
	http.HandleFunc("/", Index)
	http.HandleFunc("/static/", Static)
	http.HandleFunc("/api/generate/", GenerateAPI)

	url := host + ":" + strconv.Itoa(port)

	log.Println("Serving HTTP at", url)
	log.Fatalln(http.ListenAndServe(url, nil))
}
