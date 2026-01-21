package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
	"strconv"
)

type Sword struct {
	Name      string
	MinDamage int
	MaxDamage int
}

func NewSword(name string, minDamage, maxDamage int) *Sword {
	return &Sword{
		Name:      name,
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

func NewSwordGolden() *Sword {
	return NewSword("Golden Sword", 17, 35)
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

func (s *Sword) UseOnEntity(
	user, target game.Entity,
	_ *game.World,
) (string, bool, bool) {
	targetHealth, ok := target.(game.HasHealth)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, s), false, false
	}

	damage := utils.RandIntInRange(s.MinDamage, s.MaxDamage)

	response := targetHealth.GetHealth().Change(-damage)

	// TODO: temporary solution
	if targetHealth.GetHealth().CurrentHealth <= 0 {
		if targetLoot, hasLoot := target.(HasLoot); hasLoot {
			if player, isPlayer := user.(*Player); isPlayer {
				loot := targetLoot.GetLoot().Drop()
				response += "\n" + player.GetInventory().AddItems(loot)
			} else {
				return game.SnipItemCannotBeUsedBy(user, s), false, false
			}
		}
	}

	return fmt.Sprintf(
		"%s slices %s with a %s!\n%s\n",
		user.GetName(),
		target.GetName(),
		s.GetName(),
		response,
	), true, false
}
