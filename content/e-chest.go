package content

import (
	"fmt"
	"text-adventure-game/game"
	"text-adventure-game/utils"
)

var (
	lootTableChest = map[game.Item]int{
		NewHealingPotion("Minor", 20): 50,
		NewHealingPotion("Major", 50): 20,
		NewSpeedPotion("Minor", 1): 30,
		NewSpeedPotion("Major", 3): 10,
		NewSword("Iron Sword", 12, 20): 10,
		NewSpear("Iron Spear", 11, 18): 10,
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
