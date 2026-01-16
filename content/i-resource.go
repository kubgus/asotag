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
	return game.FmtItem(r.Type.String())
}

func (r *Resource) GetDesc() string {
	return fmt.Sprintf(
		"A chunk of %v. Can be used for crafting or trading.",
		game.FmtItem(r.Type.String()),
		)
}

func (r *Resource) Use(user, target game.Entity) (string, bool, bool) {
	return fmt.Sprintf(
		"%v does not accept the gift from %v.\nMaybe %v should bundle it instead.\n",
		target.GetName(),
		user.GetName(),
		user.GetName(),
		), false, false
}
