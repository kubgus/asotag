package content

import (
	"fmt"
	"text-adventure-game/game"
	"text-adventure-game/utils"
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

func NewSwordWooden() *Sword {
	return NewSword("Wooden Sword", 5, 10)
}

func NewSwordStone() *Sword {
	return NewSword("Stone Sword", 9, 16)
}

func NewSwordIron() *Sword {
	return NewSword("Iron Sword", 12, 20)
}

func NewSwordGold() *Sword {
	return NewSword("Gold Sword", 17, 35)
}

func (s *Sword) GetName() string {
	return game.ColItem(s.Name)
}

func (s *Sword) GetDesc() string {
	return fmt.Sprintf(
		"Deals %v to %v close-range damage.",
		game.FormatDamage(s.MinDamage, false),
		game.FormatDamage(s.MaxDamage, false),
		)
}

func (s *Sword) Use(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	targetHealth, ok := target.(game.EntityHealth)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, s), false, false
	}

	damage := utils.RandIntInRange(s.MinDamage, s.MaxDamage)

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
		game.FormatDamage(damage, false),
		response,
		), true, false
}
