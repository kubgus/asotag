package content

import "asotag/game"

type HasMovement interface {
	GetMovement() *MovementModule
}

type MovementModule struct {
	entity game.Entity

	ExtraMoves int
}

func (mm *MovementModule) Init(e game.Entity) {
	mm.entity = e
}

func (mm *MovementModule) Move(dx, dy int, world *game.World) (string, bool) {
	response, moved := world.MoveInDirection(mm.entity, dx, dy)
	if moved {
		mm.ExtraMoves--
	}
	return response, moved && mm.ExtraMoves <= 0
}
