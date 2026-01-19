package content

import (
	"asotag/utils"
	"fmt"
)

const (
	defaultGoblinName   = "Goblin"
	defaultGoblinHealth = 30
	defaultGoblinDamage = 5

	surnameChanceGoblin = 1

)

var (
	takenGoblinSurnames  = map[string]bool{}
	randomGoblinSurnames = []string{
		"Archibald",
		"Blot",
		"Cook",
		"Drake",
		"Elph",
		"Fenrir",
		"Grok",
		"Hugo",
		"Jinx",
		"Kip",
		"Lug",
		"Muck",
		"Nibble",
		"Oscar",
		"Pax",
		"Quill",
		"Rattle",
		"Shank",
		"Till",
		"Urk",
		"Vorp",
		"Zig",
	}
)

func NewGoblin() *Goblin {
	surname := ""
	if utils.RandProbability(surnameChanceGoblin) {
		var randomSurname string
		for {
			randomSurname, _ = utils.RandChoice(randomGoblinSurnames)
			if !takenGoblinSurnames[randomSurname] && len(takenGoblinSurnames) < len(randomGoblinSurnames) {
				takenGoblinSurnames[randomSurname] = true
				break
			}
		}
		surname = fmt.Sprintf(" %v", randomSurname)
	}

	return &Goblin{
		Name:   defaultGoblinName + surname,
		Health: defaultGoblinHealth,
		Damage: defaultGoblinDamage,
	}
}
