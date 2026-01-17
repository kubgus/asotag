package content

import (
	"fmt"
	"math/rand/v2"
	"text-adventure-game/game"
)

var (
	lootItemsChest = []game.Item{
		NewHealingPotion("Minor", 20),
		NewHealingPotion("Major", 50),
		NewSword("Iron", 12, 20),
	}
)

type Chest struct {
	IsUnlocked bool
	Contents []game.Item
}

func NewChest() *Chest {
	return &Chest{
		IsUnlocked: false,
		Contents: []game.Item{
			lootItemsChest[rand.IntN(len(lootItemsChest))],
		},
	}
}

func (c *Chest) GetName() string {
	return game.ColLocation("Chest")
}

func (c *Chest) GetHealth() int {
	return 0
}

func (c *Chest) GetStatus() string {
	if c.IsUnlocked {
		return game.ColHealth("Unlocked")
	} else {
		return game.ColHealth("Locked")
	}
}

func (c *Chest) AddHealth(amount int) (string, bool) {
	return fmt.Sprintf(
		"%v seems completely unaffected.",
		c.GetName(),
	), true
}

func (c *Chest) GetDesc(user game.Entity) string {
	if c.IsUnlocked {
		return fmt.Sprintf(
			"%v is unlocked.",
			c.GetName(),
		)
	} else {
		return fmt.Sprintf(
			"%v is locked. Unlock it to gain loot.",
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

func (c *Chest) BeforeTurn(context *game.Context) { }

func (c *Chest) OnTurn(context *game.Context) (string, bool) {
	return fmt.Sprintf(
		"%v gains conciousness for a second.",
		c.GetName(),
	), true
}
