package main

/*
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopgweb#gopgweb
*/

import (
	"flag"

	"github.com/sayanarijit/gopgweb/app"
)

func main() {
	host := flag.String("host", "0.0.0.0", "HTTP server url")
	port := flag.Int("port", 8080, "HTTP server port")

	flag.Parse()

	app.ServeHTTP(*host, *port)
}
