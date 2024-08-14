package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listen", ":3030", "Listen address")
	seed := flag.Bool("seed", false, "Seed the database")
	flag.Parse()

	s := NewServer(*listenAddr)
	if *seed {
		// TODO: Seed the database
		// s.Seed()
	}

	fmt.Println("Listening on", s.listenAddr)
	log.Fatal(s.Start())
}
