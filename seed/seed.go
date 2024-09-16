package seed

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/BladenWard/melee-api/types"
)

func Seed() {
	fmt.Println("Seeding database...")
	characterFile, _ := os.Open("seed/characters/fox.yaml")
	defer characterFile.Close()

	characterYaml := make([]byte, 10000)
	count, error := characterFile.Read(characterYaml)
	characterYaml = characterYaml[:count]
	if error != nil {
		log.Fatalf("error: %v", error)
	}

	// TODO: Import the yaml file
	// fmt.Println(string(characterYaml))

	t := &types.Character{}
	// err := yaml.Unmarshal(characterYaml[:count], &t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	d, err := yaml.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
	// fmt.Println(string(yamlFile))
}
