package content

import (
	"asotag/game"
	"asotag/utils"
)

func NewChest() *Chest {
	isUnlocked, _ := utils.RandChoice([]bool{true, false})

	return &Chest{
		IsUnlocked: isUnlocked,
		Contents:   game.GetRandomLoot(lootTableChest, utils.RandIntInRange(1, 4)),
	}
}
