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
		NewSpeedPotion("Major", 2): 10,
		NewSword("Iron", 12, 20): 10,
		NewKey(): 40,
	}
)

type Chest struct {
	IsUnlocked bool
	Contents []game.Item
}

func NewChest() *Chest {
	return &Chest{
		IsUnlocked: utils.RandChoice([]bool{true, false}),
		Contents: game.GetRandomLoot(lootTableChest, utils.RandIntInRange(1, 4)),
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
		if len(c.Contents) == 0 {
			return game.ColHealth("Empty")
		}
		return game.ColHealth("Unlocked")
	}
	return game.ColHealth("Locked")
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
