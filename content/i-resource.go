package content

import "asotag/game"

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
	return "Can be used for crafting or trading."
}
