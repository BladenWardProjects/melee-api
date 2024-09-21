package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BladenWard/melee-api/types"
	// "gopkg.in/yaml.v3"
	// "github.com/BladenWard/melee-api/types"
)

func Seed() {
	fmt.Println("Seeding database...")
	characterFile, _ := os.ReadFile("seed/characters/fox.json")

	characterJson := string(characterFile)

	// fmt.Println(string(characterJson))

	character := types.Character{}
	moveMap := map[string]interface{}{}
	json.Unmarshal([]byte(characterJson), &moveMap)

	character.Name = moveMap["name"].(string)
	character.Weight = uint(moveMap["weight"].(float64))
	character.FastfallSpeed = moveMap["fastfall_speed"].(float64)
	character.DashSpeed = moveMap["dash_speed"].(float64)
	character.RunSpeed = moveMap["run_speed"].(float64)
	character.WavedashLength = uint(moveMap["wavedash_length_rank"].(float64))
	character.Galint = uint(moveMap["galint"].(float64))
	character.JumpSquat = uint(moveMap["jump_squat"].(float64))
	character.Walljump = moveMap["walljump"].(bool)

	groundMap := moveMap["ground"].([]interface{})
	for _, ground := range groundMap {
		groundAttack := types.GroundAttack{}
		groundAttackMap := ground.(map[string]interface{})
		groundAttack.Name = groundAttackMap["name"].(string)
		groundAttack.Start = uint(groundAttackMap["start"].(float64))
		groundAttack.End = uint(groundAttackMap["end"].(float64))
		if groundAttackMap["frames"] != nil {
			groundAttack.TotalFrames = uint(groundAttackMap["frames"].(float64))
		}
		if groundAttackMap["iasa_frames"] != nil {
			groundAttack.IasaFrames = groundAttackMap["iasa_frames"].(*uint)
		}
		groundAttack.ShieldStun = uint(groundAttackMap["shield_stun"].(float64))
		groundAttack.BaseDamage = uint(groundAttackMap["base_damage"].(float64))
		if groundAttackMap["weak_damage"] != nil {
			var weak uint = uint(groundAttackMap["weak_damage"].(float64))
			groundAttack.WeakDamage = &weak
		}
		character.GroundAttacks = append(character.GroundAttacks, groundAttack)
	}
	//
	aerialMap := moveMap["aerial"].([]interface{})
	for _, aerial := range aerialMap {
		aerialAttack := types.Aerial{}
		aerialAttackMap := aerial.(map[string]interface{})
		aerialAttack.Name = aerialAttackMap["name"].(string)
		aerialAttack.Start = uint(aerialAttackMap["start"].(float64))
		aerialAttack.End = uint(aerialAttackMap["end"].(float64))
		aerialAttack.TotalFrames = uint(aerialAttackMap["frames"].(float64))
		aerialAttack.ShieldStun = uint(aerialAttackMap["shield_stun"].(float64))
		aerialAttack.BaseDamage = uint(aerialAttackMap["base_damage"].(float64))
		if aerialAttackMap["weak_damage"] != nil {
			var weak uint = uint(aerialAttackMap["weak_damage"].(float64))
			aerialAttack.WeakDamage = &weak
		}
		aerialAttack.LandingLag = uint(aerialAttackMap["landing_lag"].(float64))
		aerialAttack.LCancelLag = uint(aerialAttackMap["lcancel_lag"].(float64))
		if aerialAttackMap["auto_cancel"] != nil {
			var autoCancel int = int(aerialAttackMap["auto_cancel"].(float64))
			aerialAttack.AutoCancel = &autoCancel
		}
		character.Aerials = append(character.Aerials, aerialAttack)
	}
	//
	// specialMap := moveMap["special"].([]interface{})
	// for _, special := range specialMap {
	// 	specialAttack := types.Special{}
	// 	specialAttackMap := special.(map[string]interface{})
	// 	specialAttack.Name = specialAttackMap["name"].(string)
	// 	specialAttack.Start = specialAttackMap["start"].(uint)
	// 	specialAttack.End = specialAttackMap["end"].(uint)
	// 	specialAttack.TotalFrames = specialAttackMap["total_frames"].(uint)
	// 	specialAttack.ShieldStun = specialAttackMap["shield_stun"].(*uint)
	// 	specialAttack.BaseDamage = specialAttackMap["base_damage"].(uint32)
	// 	specialAttack.WeakDamage = specialAttackMap["weak_damage"].(*uint32)
	// 	specialAttack.LandingLag = specialAttackMap["landing_lag"].(*uint)
	// 	specialAttack.LandingFallSpecial = specialAttackMap["landing_fall_special"].(*uint)
	// 	character.Specials = append(character.Specials, specialAttack)
	// }
	//
	// grabMap := moveMap["grab"].([]interface{})
	// for _, grab := range grabMap {
	// 	grabAttack := types.Grab{}
	// 	grabAttackMap := grab.(map[string]interface{})
	// 	grabAttack.Name = grabAttackMap["name"].(string)
	// 	grabAttack.Start = grabAttackMap["start"].(uint)
	// 	grabAttack.TotalFrames = grabAttackMap["total_frames"].(uint)
	// 	character.Grabs = append(character.Grabs, grabAttack)
	// }
	//
	// throwMap := moveMap["throw"].([]interface{})
	// for _, throw := range throwMap {
	// 	throwAttack := types.Throw{}
	// 	throwAttackMap := throw.(map[string]interface{})
	// 	throwAttack.Name = throwAttackMap["name"].(string)
	// 	throwAttack.Start = throwAttackMap["start"].(uint)
	// 	throwAttack.End = throwAttackMap["end"].(uint)
	// 	throwAttack.TotalFrames = throwAttackMap["total_frames"].(uint)
	// 	throwAttack.BaseDamage = throwAttackMap["base_damage"].(uint32)
	// 	throwAttack.WeakDamage = throwAttackMap["weak_damage"].(*uint32)
	// 	character.Throws = append(character.Throws, throwAttack)
	// }
	//
	// dodgeMap := moveMap["dodge"].([]interface{})
	// for _, dodge := range dodgeMap {
	// 	dodgeAttack := types.Dodge{}
	// 	dodgeAttackMap := dodge.(map[string]interface{})
	// 	dodgeAttack.Name = dodgeAttackMap["name"].(string)
	// 	dodgeAttack.Start = dodgeAttackMap["start"].(uint)
	// 	dodgeAttack.End = dodgeAttackMap["end"].(uint)
	// 	dodgeAttack.TotalFrames = dodgeAttackMap["total_frames"].(uint)
	// 	dodgeAttack.LandingFallSpecial = dodgeAttackMap["landing_fall_special"].(*uint)
	// 	character.Dodges = append(character.Dodges, dodgeAttack)
	// }

	fmt.Println(character.Aerials)

	// char := types.Character{}
	// json.Unmarshal([]byte(characterYaml), &char)
	// fmt.Println(char)
	//
	// moveMap := map[string]map[string]interface{}{}
	// json.Unmarshal([]byte(characterYaml), &moveMap)
	// // fmt.Println(moveMap)
	//
	// aerialMap := moveMap["aerial"]
	// fmt.Println(aerialMap["neutral_air"]["base_damage"])

	// yaml.Unmarshal([]byte(characterYaml), &m)

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
