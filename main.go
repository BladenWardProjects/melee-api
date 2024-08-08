package main

import (
	"flag"
	"fmt"
)

func main() {
	listenAddr := flag.String("listen", ":6666", "Listen address")
	flag.Parse()

	fmt.Println("Listening on", *listenAddr)
}
