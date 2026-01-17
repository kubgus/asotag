package content

import (
	"fmt"
	"text-adventure-game/game"
)

type Key struct {}

func NewKey() *Key {
	return &Key{}
}

func (k *Key) GetName() string {
	return game.ColItem("Key")
}

func (k *Key) GetDesc() string {
	return "A small key that can unlock chests."
}

func (k *Key) Use(user, target game.Entity) (string, bool, bool) {
	chest, ok := target.(*Chest)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, k), false, false
	}
	if chest.IsUnlocked {
		return game.SnipCannotUseItemOn(user, target, k), false, false
	}

	loot := chest.GetLoot(user)

	if player, ok := user.(*Player); ok {
		player.CollectLoot(loot)
	} else {
		return game.SnipItemCannotBeUsedBy(user, k), false, false
	}

	return fmt.Sprintf(
		"%v uses %v to unlock %v and finds: %v.\n",
		user.GetName(),
		k.GetName(),
		target.GetName(),
		game.ListItems(loot),
	), true, true
}
