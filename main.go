package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BladenWard/melee-api/server"
)

func main() {
	listenAddr := flag.String("listen", ":3030", "Listen address")
	seed := flag.Bool("seed", false, "Seed the database")
	flag.Parse()

	s := server.NewServer(*listenAddr)
	if *seed {
		// TODO: Seed the database
		// s.Seed()
	}

	fmt.Println("Listening on", s.ListenAddr)
	log.Fatal(s.Start())
}
