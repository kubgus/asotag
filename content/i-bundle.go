package content

import (
	"fmt"
	"text-adventure-game/game"
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
	return game.FmtItem("Bundle")
}

func (b *Bundle) GetDesc() string {
	return fmt.Sprintf(
		"Contains: %v. Can be used for crafting or trading.",
		game.ListItems(b.Items),
		)
}

func (b *Bundle) Use(user, target game.Entity) (string, bool, bool) {
	craftingAllowed := false
	var craftedItem game.Item

	if workbench := target.(*Workbench); workbench != nil {
		craftingAllowed = true
		craftedItem = workbench.Craft(user, b.Items)
	}

	if craftedItem != nil {
		player, ok := user.(*Player)
		if ok {
			player.Inventory = append(player.Inventory, craftedItem)
		}

		return fmt.Sprintf(
			"%v uses %v to craft %v!\n",
			user.GetName(),
			game.ListItems(b.Items),
			craftedItem.GetName(),
			), true, true
	}

	if craftingAllowed {
		return fmt.Sprintf(
			"%v tries to craft something from %v, but cannot figure anything out.\n",
			user.GetName(),
			game.ListItems(b.Items),
			), false, false
	}

	return fmt.Sprintf(
		"%v does not accept the gift from %v.\nMaybe %v should use it for crafting instead.\n",
		target.GetName(),
		user.GetName(),
		user.GetName(),
		), false, false
}
