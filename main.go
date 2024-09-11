package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BladenWard/melee-api/db"
	"github.com/BladenWard/melee-api/seed"
	"github.com/BladenWard/melee-api/server"
)

func main() {
	listenAddr := flag.String("listen", ":3030", "Listen address")
	seedFlag := flag.Bool("seed", false, "Seed the database")
	flag.Parse()

	server := server.NewServer(*listenAddr)

	db.Init()
	if *seedFlag {
		// TODO: Seed the database
		seed.Seed()
	}

	fmt.Println("Listening on", server.ListenAddr)
	log.Fatal(server.Start())
}
