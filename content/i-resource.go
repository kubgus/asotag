package content

import (
	"fmt"
	"text-adventure-game/game"
)

type Resource struct {
	Type Material
}

func NewResource(resourceType Material) *Resource {
	return &Resource{
		Type: resourceType,
	}
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

func (r *Resource) Use(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	return game.SnipCannotUseItemOn(user, target, r), false, false
}
