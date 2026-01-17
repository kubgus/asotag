package content

import (
	"fmt"
	"text-adventure-game/game"
	"text-adventure-game/utils"
)

var (
	lootTableChest = map[game.Item]int{
		NewHealingPotionMinor(): 50,
		NewHealingPotionMajor(): 20,
		NewSpeedPotionMinor(): 30,
		NewSpeedPotionMajor(): 10,
		NewSwordIron(): 10,
		NewSpearIron(): 10,
		NewKey(): 40,
	}
)

type Chest struct {
	IsUnlocked bool
	Contents []game.Item
}

func NewChest() *Chest {
	isUnlocked, _ := utils.RandChoice([]bool{true, false})

	return &Chest{
		IsUnlocked: isUnlocked,
		Contents: game.GetRandomLoot(lootTableChest, utils.RandIntInRange(1, 4)),
	}
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
				"%v is unlocked and contains: %v.\n",
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
