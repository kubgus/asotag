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
	return game.FmtItem(k.Prefix + " Healing Potion")
}

func (k *HealingPotion) GetDesc() string {
	return fmt.Sprintf(
		"Restores %v when used.",
		game.FmtHealth(strconv.Itoa(k.HealAmount) + " health"),
	)
}

func (k *HealingPotion) Use(user, target game.Entity) (string, bool, bool) {
	healAmount := k.HealAmount

	response, alive := target.AddHealth(healAmount)

	return fmt.Sprintf(
		"%v uses %v on %v, restoring %v! %v\n",
		user.GetName(),
		k.GetName(),
		target.GetName(),
		game.FmtHealth(strconv.Itoa(healAmount) + " health"),
		response,
	), alive, true
}
