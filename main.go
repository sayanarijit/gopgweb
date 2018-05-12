package main

/*
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopgweb#gopgweb
*/

import (
	"flag"
	"os"
	"strconv"

	"github.com/sayanarijit/gopgweb/app"
)

func getenv(v string, d string) string {
	x := os.Getenv(v)
	if x == "" {
		return d
	}
	return x
}

func main() {
	host := flag.String("host", "0.0.0.0", "HTTP server url")
	defport, _ := strconv.Atoi(getenv("PORT", "8080"))
	port := flag.Int("port", defport, "HTTP server port")

	flag.Parse()

	app.ServeHTTP(*host, *port)
}
