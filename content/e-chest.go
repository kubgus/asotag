package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
)

type Chest struct {
	IsUnlocked bool

	loot LootModule
}

func NewChest() *Chest {
	isUnlocked, _ := utils.RandChoice([]bool{true, false})

	chest := Chest{
		IsUnlocked: isUnlocked,

		loot: LootModule{
			LootTable: map[game.Item]int{
				NewHealingPotionMinor(): 50,
				NewKey():                30,
				NewSpeedPotion():        30,
				NewHealingPotionMajor(): 15,
				NewSwordIron():          5,
				NewSpearIron():          5,
			},
			AmountTable: map[int]int{
				1: 50,
				2: 75,
				3: 100,
				4: 50,
			},
			LootLimit: 1,
		},
	}

	return &chest
}

func (c *Chest) GetLoot() *LootModule {
	c.loot.Init(c)
	return &c.loot
}

func (c *Chest) GetName() string {
	return game.ColLocation("Chest")
}

func (c *Chest) GetStatus() string {
	if c.IsUnlocked {
		if c.GetLoot().LootLimit < 0 {
			return game.ColHealth("Empty")
		}
		return game.ColHealth("Unlocked")
	}
	return game.ColHealth("Locked")
}

func (c *Chest) GetDesc(user game.Entity) string {
	if c.IsUnlocked {
		if c.GetLoot().LootLimit > 0 {
			return fmt.Sprintf(
				"%v is unlocked. Mine it to gain loot.\n",
				c.GetName(),
			)
		}
		return fmt.Sprintf(
			"%v is empty.\n",
			c.GetName(),
		)
	} else {
		return fmt.Sprintf(
			"%v is locked. Unlock it to gain loot.\n",
			c.GetName(),
		)
	}
}
