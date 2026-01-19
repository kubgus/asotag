package content

import (
	"asotag/game"
	"fmt"
)

type Key struct{}

func (k *Key) GetName() string {
	return game.ColItem("Key")
}

func (k *Key) GetDesc() string {
	return "A small key that can unlock chests."
}

func (k *Key) UseOnEntity(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	chest, ok := target.(*Chest)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, k), false, false
	}
	if chest.IsUnlocked {
		return game.SnipCannotUseItemOn(user, target, k), false, false
	}

	chest.IsUnlocked = true

	return fmt.Sprintf(
		"%v uses %v to unlock %v.\n",
		user.GetName(),
		k.GetName(),
		target.GetName(),
	), false, true
}
