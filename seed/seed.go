package seed

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BladenWard/melee-api/types"
)

func Seed() {
	fmt.Println("Seeding database...")
	characterFile, _ := os.ReadFile("seed/characters/fox.json")

	characterJson := string(characterFile)

	character := types.Character{}
	infoMap := map[string]interface{}{}
	json.Unmarshal([]byte(characterJson), &infoMap)

	character.Name = infoMap["name"].(string)
	character.Weight = uint(infoMap["weight"].(float64))
	character.FastfallSpeed = infoMap["fastfall_speed"].(float64)
	character.DashSpeed = infoMap["dash_speed"].(float64)
	character.RunSpeed = infoMap["run_speed"].(float64)
	character.WavedashLength = uint(infoMap["wavedash_length_rank"].(float64))
	character.Galint = uint(infoMap["galint"].(float64))
	character.JumpSquat = uint(infoMap["jump_squat"].(float64))
	character.Walljump = infoMap["walljump"].(bool)

	groundMap := infoMap["ground"].([]interface{})
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

	aerialMap := infoMap["aerial"].([]interface{})
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

	specialMap := infoMap["special"].([]interface{})
	for _, special := range specialMap {
		specialAttack := types.Special{}
		specialAttackMap := special.(map[string]interface{})
		specialAttack.Name = specialAttackMap["name"].(string)
		specialAttack.Start = uint(specialAttackMap["start"].(float64))
		specialAttack.End = uint(specialAttackMap["end"].(float64))
		specialAttack.TotalFrames = uint(specialAttackMap["frames"].(float64))
		if specialAttackMap["shield_stun"] != nil {
			var shield uint = uint(specialAttackMap["shield_stun"].(float64))
			specialAttack.ShieldStun = &shield
		}
		specialAttack.BaseDamage = uint(specialAttackMap["base_damage"].(float64))
		if specialAttackMap["weak_damage"] != nil {
			var weak uint = uint(specialAttackMap["weak_damage"].(float64))
			specialAttack.WeakDamage = &weak
		}
		if specialAttackMap["landing_lag"] != nil {
			var landingLag uint = uint(specialAttackMap["landing_lag"].(float64))
			specialAttack.LandingLag = &landingLag
		}
		if specialAttackMap["landing_fall_special"] != nil {
			var landingFallSpecial uint = uint(specialAttackMap["landing_fall_special"].(float64))
			specialAttack.LandingFallSpecial = &landingFallSpecial
		}
		character.Specials = append(character.Specials, specialAttack)
	}

	grabMap := infoMap["grab"].([]interface{})
	for _, grab := range grabMap {
		grabAttack := types.Grab{}
		grabAttackMap := grab.(map[string]interface{})
		grabAttack.Name = grabAttackMap["name"].(string)
		grabAttack.Start = uint(grabAttackMap["start"].(float64))
		grabAttack.TotalFrames = uint(grabAttackMap["frames"].(float64))
		character.Grabs = append(character.Grabs, grabAttack)
	}

	throwMap := infoMap["throw"].([]interface{})
	for _, throw := range throwMap {
		throwAttack := types.Throw{}
		throwAttackMap := throw.(map[string]interface{})
		throwAttack.Name = throwAttackMap["name"].(string)
		throwAttack.Start = uint(throwAttackMap["start"].(float64))
		throwAttack.End = uint(throwAttackMap["end"].(float64))
		throwAttack.TotalFrames = uint(throwAttackMap["frames"].(float64))
		throwAttack.BaseDamage = uint(throwAttackMap["base_damage"].(float64))
		if throwAttackMap["weak_damage"] != nil {
			var weak uint = uint(throwAttackMap["weak_damage"].(float64))
			throwAttack.WeakDamage = &weak
		}
		character.Throws = append(character.Throws, throwAttack)
	}

	dodgeMap := infoMap["dodge"].([]interface{})
	for _, dodge := range dodgeMap {
		dodgeAttack := types.Dodge{}
		dodgeAttackMap := dodge.(map[string]interface{})
		dodgeAttack.Name = dodgeAttackMap["name"].(string)
		dodgeAttack.Start = uint(dodgeAttackMap["start"].(float64))
		dodgeAttack.End = uint(dodgeAttackMap["end"].(float64))
		dodgeAttack.TotalFrames = uint(dodgeAttackMap["frames"].(float64))
		if dodgeAttackMap["landing_fall_special"] != nil {
			var landingFallSpecial uint = uint(dodgeAttackMap["landing_fall_special"].(float64))
			dodgeAttack.LandingFallSpecial = &landingFallSpecial
		}
		character.Dodges = append(character.Dodges, dodgeAttack)
	}

	fmt.Println(character.GetMoveByName("up_air"))

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
