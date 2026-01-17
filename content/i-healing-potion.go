package content

import (
	"fmt"
	"strconv"
	"text-adventure-game/game"
)

type HealingPotion struct {
	Prefix string
	HealAmount int
}

func NewHealingPotion(prefix string, healAmount int) *HealingPotion {
	return &HealingPotion{
		Prefix: prefix,
		HealAmount: healAmount,
	}
}

func (k *HealingPotion) GetName() string {
	return game.ColItem(k.Prefix + " Healing Potion")
}

func (k *HealingPotion) GetDesc() string {
	return fmt.Sprintf(
		"Restores %v when used.",
		game.FormatHealth(k.HealAmount, true),
	)
}

func (k *HealingPotion) Use(user, target game.Entity) (string, bool, bool) {
	targetHealth, ok := target.(game.EntityHealth)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, k), false, false
	}

	healAmount := k.HealAmount

	response, alive := targetHealth.AddHealth(healAmount)

	return fmt.Sprintf(
		"%v uses %v on %v, restoring %v! %v\n",
		user.GetName(),
		k.GetName(),
		target.GetName(),
		game.ColHealth(strconv.Itoa(healAmount) + " health"),
		response,
	), alive, true
}
