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
				NewKey():                15,
				NewHealingPotionMinor(): 30,
				NewHealingPotionMajor(): 15,
				NewSpeedPotion():        15,
				NewTeleportPotion():     15,
				NewSwordIron():          5,
				NewSpearIron():          5,
				NewSwordGolden():        1,
			},
			AmountTable: map[int]int{
				1: 10,
				2: 40,
				3: 40,
				4: 10,
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
				"%v is unlocked and contains loot.\n",
				c.GetName(),
			)
		}
		return fmt.Sprintf(
			"%v is empty.\n",
			c.GetName(),
		)
	} else {
		return fmt.Sprintf(
			"%v is locked.\n",
			c.GetName(),
		)
	}
}
