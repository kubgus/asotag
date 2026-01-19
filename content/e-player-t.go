package content

import "asotag/game"

const (
	defaultPlayerName   = "Hero"
	defaultPlayerHealth = 100
)

func NewPlayer(name string) *Player {
	nameToUse := name
	if nameToUse == "" {
		nameToUse = defaultPlayerName
	}

	return &Player{
		Name:   nameToUse,
		Health: defaultPlayerHealth,

		Inventory: []game.Item{},
	}
}

