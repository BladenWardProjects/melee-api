package types

type Move interface {
	GetMoveByName(string) interface{}
}

func (c *Character) GetMoveByName(name string) interface{} {
	for _, move := range c.GroundAttacks {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Aerials {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Specials {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Grabs {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Throws {
		if move.Name == name {
			return move
		}
	}
	for _, move := range c.Dodges {
		if move.Name == name {
			return move
		}
	}
	return nil
}
