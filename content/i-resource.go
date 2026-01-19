package content

import (
	"asotag/game"
	"fmt"
)

type Resource struct {
	Type Material
}

func (r *Resource) GetName() string {
	return game.ColItem(r.Type.String())
}

func (r *Resource) GetDesc() string {
	return fmt.Sprintf(
		"A chunk of %v. Can be used for crafting or trading.",
		game.ColItem(r.Type.String()),
	)
}
