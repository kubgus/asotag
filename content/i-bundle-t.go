package content

import "asotag/game"

func NewBundle(items []game.Item) *Bundle {
	return &Bundle{
		Items: items,
	}
}

