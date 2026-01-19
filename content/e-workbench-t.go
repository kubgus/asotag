package content

import "asotag/game"

var (
	defaultWorkbenchRecipes = map[game.Item][]game.Item{
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
		NewSwordGold(): {
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

func NewWorkbench() *Workbench {
	return &Workbench{
		Recipes: defaultWorkbenchRecipes,
	}
}
