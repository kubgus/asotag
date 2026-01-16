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
	return game.FmtItem(s.Name)
}

func (s *Sword) GetDesc() string {
	return fmt.Sprintf(
		"Deals %v to %v close-range damage.",
		game.FmtDamage(strconv.Itoa(s.MinDamage)),
		game.FmtDamage(strconv.Itoa(s.MaxDamage)),
		)
}

func (s *Sword) Use(user, target game.Entity) (string, bool, bool) {
	damage := rand.IntN(s.MaxDamage - s.MinDamage + 1) + s.MinDamage

	response, alive := target.AddHealth(-damage)

	if !alive {
		loot := target.Loot(user)
		if len(loot) > 0 {
			player, ok := user.(*Player)
			if ok {
				player.Inventory = append(player.Inventory, loot...)
			}

			response += fmt.Sprintf(
				"\nObtained items: %v",
				game.ListItems(loot),
				)
		}
	}

	return fmt.Sprintf(
		"%v slices %v with a %v for %v damage! %v\n",
		user.GetName(),
		target.GetName(),
		s.GetName(),
		game.FmtDamage(strconv.Itoa(damage)),
		response,
		), true, false
}
