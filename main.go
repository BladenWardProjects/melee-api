package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listen", ":3030", "Listen address")
	flag.Parse()

	s := NewServer(*listenAddr)

	fmt.Println("Listening on", s.listenAddr)
	log.Fatal(s.Start())
}
