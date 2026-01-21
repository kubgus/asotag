package content

import (
	"asotag/game"
	"fmt"
	"strconv"
)

type HealingPotion struct {
	Prefix    string
	Magnitude int
}

func NewHealingPotion(prefix string, magnitude int) *HealingPotion {
	return &HealingPotion{
		Prefix:    prefix,
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
		"Restores %v health when used.",
		game.ColHealth(strconv.Itoa(k.Magnitude)),
	)
}

func (k *HealingPotion) UseOnEntity(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	targetHealth, ok := target.(game.HasHealth)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, k), false, false
	}

	healAmount := k.Magnitude

	response := targetHealth.GetHealth().Change(healAmount)
	return response, true, true
}
