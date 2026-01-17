package content

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"text-adventure-game/game"
)

type Sword struct {
	Name string
	MinDamage int
	MaxDamage int
}

func NewSword(name string, minDamage, maxDamage int) *Sword {
	return &Sword{
		Name: name,
		MinDamage: minDamage,
		MaxDamage: maxDamage,
	}
}

func (s *Sword) GetName() string {
	return game.ColItem(s.Name)
}

func (s *Sword) GetDesc() string {
	return fmt.Sprintf(
		"Deals %v to %v close-range damage.",
		game.ColDamage(strconv.Itoa(s.MinDamage)),
		game.ColDamage(strconv.Itoa(s.MaxDamage)),
		)
}

func (s *Sword) Use(user, target game.Entity) (string, bool, bool) {
	targetHealth, ok := target.(game.EntityHealth)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, s), false, false
	}

	damage := rand.IntN(s.MaxDamage - s.MinDamage + 1) + s.MinDamage

	response, alive := targetHealth.AddHealth(-damage)

	if !alive {
		if targetLoot, ok := target.(game.EntityLoot); ok {
			loot := targetLoot.GetLoot(user)
			if player, ok := user.(*Player); ok {
				if responseLoot, ok := player.CollectLoot(loot); ok {
					response += "\n" + responseLoot
				}
			} else {
				return game.SnipItemCannotBeUsedBy(user, s), false, false
			}
		}
	}

	return fmt.Sprintf(
		"%v slices %v with a %v for %v damage! %v\n",
		user.GetName(),
		target.GetName(),
		s.GetName(),
		game.ColDamage(strconv.Itoa(damage)),
		response,
		), true, false
}
