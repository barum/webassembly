// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":9999", "listen address")
	dir    = flag.String("dir", "client", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	fs := http.FileServer(http.Dir("../client"))
	http.Handle("/", fs)
	err := http.ListenAndServe(*listen, nil)
	log.Fatalln(err)
}
