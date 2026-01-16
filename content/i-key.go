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
	return game.FmtItem("Key")
}

func (k *Key) GetDesc() string {
	return "A small key that can unlock chests."
}

func (k *Key) Use(user, target game.Entity) (string, bool, bool) {
	chest, ok := target.(*Chest)
	if !ok {
		return fmt.Sprintf(
			"%v pokes %v with %v. It's not very effective.\n",
			user.GetName(),
			target.GetName(),
			k.GetName(),
		), false, false
	}

	if chest.IsUnlocked {
		return fmt.Sprintf(
			"%v tries to use %v on %v, but it's already unlocked.\n",
			user.GetName(),
			k.GetName(),
			target.GetName(),
		), false, false
	}

	loot := target.Loot(user)

	player, ok := user.(*Player)
	if ok {
		player.Inventory = append(player.Inventory, loot...)
	}

	return fmt.Sprintf(
		"%v uses %v to unlock %v and finds: %v.\n",
		user.GetName(),
		k.GetName(),
		target.GetName(),
		game.ListItems(loot),
	), true, true
}
