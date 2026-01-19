package content

import (
	"asotag/game"
	"fmt"
)

var (
	lootTableChest = map[game.Item]int{
		NewHealingPotionMinor(): 50,
		NewKey():                30,
		NewSpeedPotion():        30,
		NewHealingPotionMajor(): 15,
		NewSwordIron():          5,
		NewSpearIron():          5,
	}
)

type Chest struct {
	IsUnlocked bool
	Contents   []game.Item
}

func (c *Chest) GetName() string {
	return game.ColLocation("Chest")
}

func (c *Chest) GetStatus() string {
	if c.IsUnlocked {
		if len(c.Contents) == 0 {
			return game.ColHealth("Empty")
		}
		return game.ColHealth("Unlocked")
	}
	return game.ColHealth("Locked")
}

func (c *Chest) GetDesc(user game.Entity) string {
	if c.IsUnlocked {
		if len(c.Contents) > 0 {
			return fmt.Sprintf(
				"%v is unlocked and contains: %v. Can be mined to loot.\n",
				c.GetName(),
				game.ListItems(c.Contents),
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

func (c *Chest) GetLoot(user game.Entity) []game.Item {
	c.IsUnlocked = true

	response := c.Contents
	c.Contents = []game.Item{}

	return response
}
