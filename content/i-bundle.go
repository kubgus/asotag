package content

import (
	"asotag/game"
	"fmt"
)

type Bundle struct {
	Items []game.Item
}

func NewBundle(items []game.Item) *Bundle {
	return &Bundle{
		Items: items,
	}
}

func (b *Bundle) GetName() string {
	return game.ColItem("Bundle")
}

func (b *Bundle) GetDesc() string {
	return fmt.Sprintf(
		"Contains: %v. Can be used for crafting or trading.",
		game.ListItems(b.Items),
	)
}

func (b *Bundle) UseOnEntity(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	craftingAllowed := false
	var craftedItem game.Item

	if workbench := target.(*Workbench); workbench != nil {
		craftingAllowed = true
		craftedItem = workbench.Craft(user, b.Items)
	}

	if craftedItem != nil {
		if player, ok := user.(*Player); ok {
			addResponse := player.GetInventory().AddItems([]game.Item{craftedItem})

			return fmt.Sprintf(
				"Crafting successful!\n%v%v",
				addResponse,
			), true, true
		} else {
			return game.SnipItemCannotBeUsedBy(user, b), false, false
		}
	}

	if craftingAllowed {
		return fmt.Sprintf(
			"%v tries to craft something out of %v, but cannot figure anything out.\n",
			user.GetName(),
			game.ListItems(b.Items),
		), false, false
	}

	return game.SnipCannotUseItemOn(user, target, b), false, false
}
