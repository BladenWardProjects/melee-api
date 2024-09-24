package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BladenWard/melee-api/db"
	"github.com/BladenWard/melee-api/seed"
	"github.com/BladenWard/melee-api/server"
	"github.com/BladenWard/melee-api/types"
)

func main() {
	listenAddr := flag.String("listen", ":3030", "Listen address")
	seedFlag := flag.Bool("seed", false, "Seed the database")
	flag.Parse()

	store := db.Init()
	if *seedFlag {
		seed.Seed(store)
	}

	char := types.Character{}
	store.DB.First(&char, 3)
	fmt.Println(char.Name)
	fmt.Println(char.Grabs)
	server := server.NewServer(*listenAddr, store)
	fmt.Println("Server running on port:", server.ListenAddr)
	log.Fatal(server.Start())
}
