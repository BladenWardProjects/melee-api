package seed

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/BladenWard/melee-api/types"
)

func Seed() {
	fmt.Println("Seeding database...")
	characterFile, _ := os.ReadFile("seed/characters/fox.yaml")

	characterYaml := string(characterFile)

	// TODO: Import the yaml file
	fmt.Println(string(characterYaml))

	character := types.Character{}
	yaml.Unmarshal([]byte(characterYaml), &character)
	// yaml.Unmarshal([]byte(characterYaml), &m)
	fmt.Println(character)

	// t := &types.Character{}
	// err := yaml.Unmarshal(characterYaml[:count], &t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// d, err := yaml.Marshal(t)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(d))
	// fmt.Println(string(yamlFile))
}
