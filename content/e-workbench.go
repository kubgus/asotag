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
	return game.ColLocation("Workbench")
}

func (w *Workbench) GetStatus() string {
	return game.ColHealth("Solid")
}

func (w *Workbench) GetDesc(user game.Entity) string {
	craftableItems := make([]game.Item, 0, len(w.Recipes))
	for item := range w.Recipes {
		craftableItems = append(craftableItems, item)
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
			game.ColTooltip("%v can be used to craft the following items:"),
			w.GetName(),
		),
		recipesList,
	)
}

func (w *Workbench) Craft(user game.Entity, items []game.Item) game.Item {
	for item, recipe := range w.Recipes {
		if game.ItemsMatchUnordered(items, recipe) {
			return item
		}
	}

	return nil
}
