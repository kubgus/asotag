package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
)

var workbenchRecipes = map[game.Item][]game.Item{
	NewSwordWooden(): {
		NewResource(MaterialWood),
		NewResource(MaterialWood),
		NewResource(MaterialWood),
		NewResource(MaterialWood),
	},
	NewSwordStone(): {
		NewResource(MaterialStone),
		NewResource(MaterialStone),
		NewResource(MaterialStone),
		NewResource(MaterialWood),
	},
	NewSwordIron(): {
		NewResource(MaterialIron),
		NewResource(MaterialIron),
		NewResource(MaterialIron),
		NewResource(MaterialWood),
	},
	NewSwordGolden(): {
		NewResource(MaterialGold),
		NewResource(MaterialGold),
		NewResource(MaterialGold),
		NewResource(MaterialWood),
	},

	NewSpearWooden(): {
		NewResource(MaterialWood),
		NewResource(MaterialWood),
	},
	NewSpearIron(): {
		NewResource(MaterialIron),
		NewResource(MaterialWood),
	},

	NewPickaxeWooden(): {
		NewResource(MaterialWood),
		NewResource(MaterialWood),
		NewResource(MaterialWood),
	},
	NewPickaxeStone(): {
		NewResource(MaterialStone),
		NewResource(MaterialStone),
		NewResource(MaterialWood),
		NewResource(MaterialWood),
	},
	NewPickaxeIron(): {
		NewResource(MaterialIron),
		NewResource(MaterialIron),
		NewResource(MaterialWood),
		NewResource(MaterialWood),
	},
	NewPickaxeGolden(): {
		NewResource(MaterialGold),
		NewResource(MaterialGold),
		NewResource(MaterialWood),
		NewResource(MaterialWood),
	},
}

type Workbench struct {
	Recipes map[game.Item][]game.Item
}

func NewWorkbench() *Workbench {
	return &Workbench{
		Recipes: workbenchRecipes,
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
			itemClone, ok := utils.CloneInterface(item)
			if !ok {
				panic("failed to clone workbench item")
			}
			return itemClone
		}
	}

	return nil
}
