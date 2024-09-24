package types

import (
	"database/sql"
	"encoding/json"
)

func (c *Character) seedStats(charId int, infoMap map[string]interface{}) {
	c.ID = uint(charId + 1)
	c.Name = infoMap["name"].(string)
	c.Weight = uint(infoMap["weight"].(float64))
	c.FastfallSpeed = infoMap["fastfall_speed"].(float64)
	c.DashSpeed = infoMap["dash_speed"].(float64)
	if infoMap["run_speed"] != nil {
		c.RunSpeed = infoMap["run_speed"].(float64)
	}
	c.WavedashLength = uint(infoMap["wavedash_length_rank"].(float64))
	c.Galint = uint(infoMap["galint"].(float64))
	c.JumpSquat = uint(infoMap["jump_squat"].(float64))
	c.Walljump = infoMap["walljump"].(bool)
}

func (c *Character) seedGroundAttacks(groundMap []interface{}) {
	c.GroundAttacks = []GroundAttack{}
	groundAttackId := 0
	for _, ground := range groundMap {
		groundAttack := GroundAttack{}
		groundAttackMap := ground.(map[string]interface{})
		groundAttack.ID = uint(groundAttackId)
		groundAttack.Name = groundAttackMap["name"].(string)
		if groundAttackMap["start"] == nil {
			groundAttack.Start = sql.NullInt16{Int16: 0, Valid: false}
		} else {
			groundAttack.Start = sql.NullInt16{Int16: int16(groundAttackMap["start"].(float64)), Valid: true}
		}
		// groundAttack.Start = uint(groundAttackMap["start"].(float64))
		groundAttack.End = uint(groundAttackMap["end"].(float64))
		if groundAttackMap["frames"] != nil {
			var totalFrames uint = uint(groundAttackMap["frames"].(float64))
			groundAttack.TotalFrames = &totalFrames
		}
		if groundAttackMap["iasa_frames"] != nil {
			groundAttack.IasaFrames = groundAttackMap["iasa_frames"].(*uint)
		}
		// groundAttack.ShieldStun = uint(groundAttackMap["shield_stun"].(float64))
		if groundAttackMap["shield_stun"] == nil {
			groundAttack.ShieldStun = sql.NullInt16{Int16: 0, Valid: false}
		} else {
			groundAttack.ShieldStun = sql.NullInt16{Int16: int16(groundAttackMap["shield_stun"].(float64)), Valid: true}
		}
		groundAttack.BaseDamage = uint(groundAttackMap["base_damage"].(float64))
		if groundAttackMap["weak_damage"] != nil {
			var weak uint = uint(groundAttackMap["weak_damage"].(float64))
			groundAttack.WeakDamage = &weak
		}
		groundAttack.CharacterID = c.ID
		c.GroundAttacks = append(c.GroundAttacks, groundAttack)
		groundAttackId++
	}
}

func (c *Character) seedAerials(aerialMap []interface{}) {
	c.Aerials = []Aerial{}
	aerialAttackId := 0
	for _, aerial := range aerialMap {
		aerialAttack := Aerial{}
		aerialAttackMap := aerial.(map[string]interface{})
		aerialAttack.ID = uint(aerialAttackId)
		aerialAttack.Name = aerialAttackMap["name"].(string)
		aerialAttack.Start = uint(aerialAttackMap["start"].(float64))
		if aerialAttackMap["end"] != nil {
			var end uint = uint(aerialAttackMap["end"].(float64))
			aerialAttack.End = &end
		}
		if aerialAttackMap["frames"] != nil {
			var totalFrames uint = uint(aerialAttackMap["frames"].(float64))
			aerialAttack.TotalFrames = &totalFrames
		}
		aerialAttack.ShieldStun = uint(aerialAttackMap["shield_stun"].(float64))
		aerialAttack.BaseDamage = uint(aerialAttackMap["base_damage"].(float64))
		if aerialAttackMap["weak_damage"] != nil {
			var weak uint = uint(aerialAttackMap["weak_damage"].(float64))
			aerialAttack.WeakDamage = &weak
		}
		aerialAttack.LandingLag = uint(aerialAttackMap["landing_lag"].(float64))
		if aerialAttackMap["lcancel_lag"] != nil {
			aerialAttack.LCancelLag.Int16 = int16(aerialAttackMap["lcancel_lag"].(float64))
		}
		if aerialAttackMap["auto_cancel"] != nil {
			var autoCancel int = int(aerialAttackMap["auto_cancel"].(float64))
			aerialAttack.AutoCancel = &autoCancel
		}
		aerialAttack.CharacterID = c.ID
		c.Aerials = append(c.Aerials, aerialAttack)
		aerialAttackId++
	}
}

func (c *Character) seedSpecials(specialMap []interface{}) {
	c.Specials = []Special{}
	specialAttackId := 0
	for _, special := range specialMap {
		specialAttack := Special{}
		specialAttackMap := special.(map[string]interface{})
		specialAttack.ID = uint(specialAttackId)
		specialAttack.Name = specialAttackMap["name"].(string)
		specialAttack.Start = uint(specialAttackMap["start"].(float64))
		if specialAttackMap["end"] != nil {
			var end uint = uint(specialAttackMap["end"].(float64))
			specialAttack.End = &end
		}
		if specialAttackMap["frames"] != nil {
			var totalFrames uint = uint(specialAttackMap["frames"].(float64))
			specialAttack.TotalFrames = &totalFrames
		}
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
		specialAttack.CharacterID = c.ID
		c.Specials = append(c.Specials, specialAttack)
		specialAttackId++
	}
}

func (c *Character) seedGrabs(grabMap []interface{}) {
	c.Grabs = []Grab{}
	grabId := 0
	for _, grab := range grabMap {
		grabAttack := Grab{}
		grabAttackMap := grab.(map[string]interface{})
		grabAttack.ID = uint(grabId)
		grabAttack.Name = grabAttackMap["name"].(string)
		if grabAttackMap["start"] != nil {
			var start uint = uint(grabAttackMap["start"].(float64))
			grabAttack.Start = &start
		}
		grabAttack.TotalFrames = uint(grabAttackMap["frames"].(float64))
		grabAttack.CharacterID = c.ID
		c.Grabs = append(c.Grabs, grabAttack)
		grabId++
	}
}

func (c *Character) seedThrows(throwMap []interface{}) {
	c.Throws = []Throw{}
	throwId := 0
	for _, throw := range throwMap {
		throwAttack := Throw{}
		throwAttackMap := throw.(map[string]interface{})
		throwAttack.ID = uint(throwId)
		throwAttack.Name = throwAttackMap["name"].(string)
		if throwAttackMap["start"] != nil {
			var start uint = uint(throwAttackMap["start"].(float64))
			throwAttack.Start = &start
		}
		if throwAttackMap["end"] != nil {
			var end uint = uint(throwAttackMap["end"].(float64))
			throwAttack.End = &end
		}
		if throwAttackMap["frames"] != nil {
			var totalFrames uint = uint(throwAttackMap["frames"].(float64))
			throwAttack.TotalFrames = totalFrames
		}
		throwAttack.BaseDamage = uint(throwAttackMap["base_damage"].(float64))
		if throwAttackMap["weak_damage"] != nil {
			var weak uint = uint(throwAttackMap["weak_damage"].(float64))
			throwAttack.WeakDamage = &weak
		}
		throwAttack.CharacterID = c.ID
		c.Throws = append(c.Throws, throwAttack)
		throwId++
	}
}

func (c *Character) seedDodges(dodgeMap []interface{}) {
	c.Dodges = []Dodge{}
	dodgeId := 0
	for _, dodge := range dodgeMap {
		dodgeAttack := Dodge{}
		dodgeAttackMap := dodge.(map[string]interface{})
		dodgeAttack.ID = uint(dodgeId)
		dodgeAttack.Name = dodgeAttackMap["name"].(string)
		dodgeAttack.Start = uint(dodgeAttackMap["start"].(float64))
		dodgeAttack.End = uint(dodgeAttackMap["end"].(float64))
		dodgeAttack.TotalFrames = uint(dodgeAttackMap["frames"].(float64))
		if dodgeAttackMap["landing_fall_special"] != nil {
			var landingFallSpecial uint = uint(dodgeAttackMap["landing_fall_special"].(float64))
			dodgeAttack.LandingFallSpecial = &landingFallSpecial
		}
		dodgeAttack.CharacterID = c.ID
		c.Dodges = append(c.Dodges, dodgeAttack)
		dodgeId++
	}
}

func (c *Character) SerializeStructure(charId int, characterJson *string) {
	infoMap := map[string]interface{}{}
	json.Unmarshal([]byte(*characterJson), &infoMap)

	c.seedStats(charId, infoMap)
	c.seedGroundAttacks(infoMap["ground"].([]interface{}))
	c.seedAerials(infoMap["aerial"].([]interface{}))
	c.seedSpecials(infoMap["special"].([]interface{}))
	c.seedGrabs(infoMap["grab"].([]interface{}))
	c.seedThrows(infoMap["throw"].([]interface{}))
	c.seedDodges(infoMap["dodge"].([]interface{}))
}
