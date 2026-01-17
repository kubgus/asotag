package content

import (
	"fmt"
	"text-adventure-game/game"
)

type HealingPotion struct {
	Prefix string
	Magnitude int
}

func NewHealingPotion(prefix string, magnitude int) *HealingPotion {
	return &HealingPotion{
		Prefix: prefix,
		Magnitude: magnitude,
	}
}


func NewHealingPotionMinor() *HealingPotion {
	return NewHealingPotion("Minor", 20)
}

func NewHealingPotionMajor() *HealingPotion {
	return NewHealingPotion("Major", 50)
}

func NewHealingPotionSuperior() *HealingPotion {
	return NewHealingPotion("Superior", 100)
}

func (k *HealingPotion) GetName() string {
	return game.ColItem(k.Prefix + " Healing Potion")
}

func (k *HealingPotion) GetDesc() string {
	return fmt.Sprintf(
		"Restores %v when used.",
		game.FormatHealth(k.Magnitude, true),
	)
}

func (k *HealingPotion) Use(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	targetHealth, ok := target.(game.EntityHealth)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, k), false, false
	}

	healAmount := k.Magnitude

	response, alive := targetHealth.AddHealth(healAmount)

	return fmt.Sprintf(
		"%v uses %v on %v, restoring %v! %v\n",
		user.GetName(),
		k.GetName(),
		target.GetName(),
		game.FormatHealth(healAmount, true),
		response,
	), alive, true
}
