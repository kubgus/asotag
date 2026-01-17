package content

import (
	"fmt"
	"text-adventure-game/game"
	"text-adventure-game/utils"
)

var (
	defaultWorkbenchRecipes = map[game.Item][]game.Item{
		NewSword("Wooden Sword", 5, 10): {
			NewResource(MaterialWood),
			NewResource(MaterialWood),
			NewResource(MaterialWood),
			NewResource(MaterialWood),
		},
		NewSword("Stone Sword", 9, 16): {
			NewResource(MaterialStone),
			NewResource(MaterialStone),
			NewResource(MaterialStone),
			NewResource(MaterialWood),
		},
		NewSword("Iron Sword", 12, 20): {
			NewResource(MaterialIron),
			NewResource(MaterialIron),
			NewResource(MaterialIron),
			NewResource(MaterialWood),
		},
		NewSword("Gold Sword", 17, 35): {
			NewResource(MaterialGold),
			NewResource(MaterialGold),
			NewResource(MaterialGold),
			NewResource(MaterialWood),
		},

		NewPickaxe(MaterialWood): {
			NewResource(MaterialWood),
			NewResource(MaterialWood),
			NewResource(MaterialWood),
		},
		NewPickaxe(MaterialStone): {
			NewResource(MaterialStone),
			NewResource(MaterialStone),
			NewResource(MaterialWood),
			NewResource(MaterialWood),
		},
		NewPickaxe(MaterialIron): {
			NewResource(MaterialIron),
			NewResource(MaterialIron),
			NewResource(MaterialWood),
			NewResource(MaterialWood),
		},
		NewPickaxe(MaterialGold): {
			NewResource(MaterialGold),
			NewResource(MaterialGold),
			NewResource(MaterialWood),
			NewResource(MaterialWood),
		},
	}
)

type Workbench struct {
	Recipes map[game.Item][]game.Item
}

func NewWorkbench() *Workbench {
	return &Workbench{
		Recipes: defaultWorkbenchRecipes,
	}
}

func (w *Workbench) GetName() string {
	return game.FmtLocation("Workbench")
}

func (w *Workbench) GetDesc() string {
	return "A sturdy workbench used for crafting items."
}

func (w *Workbench) GetHealth() int {
	return 0
}

func (w *Workbench) GetHealthString(includeWordHealth bool) string {
	return game.FmtHealth("Solid")
}

func (w *Workbench) AddHealth(amount int) (string, bool) {
	return fmt.Sprintf(
		"%v seems completely unaffected.",
		w.GetName(),
	), true
}

func (w *Workbench) Examine(user game.Entity) string {
	craftableItems := make([]game.Item, 0, len(w.Recipes))
	for item := range w.Recipes {
		craftableItems = append(craftableItems, item)
	}

	if len(craftableItems) == 0 {
		return fmt.Sprintf(
			"%v has no available recipes at the moment.",
			w.GetName(),
		)
	}

	recipesList := utils.JoinWithMapFunc(
		craftableItems,
		"\n",
		func(i int, item game.Item) string {
			return fmt.Sprintf(
				"%v <- %v",
				item.GetName(),
				game.ListItems(w.Recipes[item]),
				)
		},
		)

	return fmt.Sprintf(
		"%v\n%v\n",
		fmt.Sprintf(
			game.FmtTooltip("%v can be used to craft the following items:"),
			w.GetName(),
		),
		recipesList,
	)
}

func (w *Workbench) Loot(user game.Entity) []game.Item {
	return []game.Item{}
}

func (w *Workbench) Reset(context *game.Context) { }

func (w *Workbench) Move(context *game.Context) (string, bool) {
	return fmt.Sprintf(
		"%v remains firmly in place.",
		w.GetName(),
	), true
}

func (w *Workbench) Craft(user game.Entity, items []game.Item) game.Item {
	for item, recipe := range w.Recipes {
		if game.ItemsMatchUnordered(items, recipe) {
			return item
		}
	}

	return nil
}
