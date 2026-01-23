package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
)

type TeleportPotion struct{}

func NewTeleportPotion() *TeleportPotion {
	return &TeleportPotion{}
}

func (p *TeleportPotion) GetName() string {
	return game.ColItem("Teleportation Potion")
}

func (p *TeleportPotion) GetDesc() string {
	return "Teleports the user to a random location."
}

func (p *TeleportPotion) UseOnEntity(
	user, target game.Entity,
	context *game.Context,
) (string, bool, bool) {
	targetMovement, ok := target.(HasMovement)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}

	for range 100 {
		dx := utils.RandIntInRange(-10, 10)
		dy := utils.RandIntInRange(-10, 10)
		_, ok := targetMovement.GetMovement().Move(dx, dy, &context.World)
		if ok {
			return fmt.Sprintf(
				"%s uses %s on %s.\n%s is teleported to a new location!\n",
				user.GetName(),
				p.GetName(),
				target.GetName(),
				target.GetName(),
			), false, true
		}
	}

	return fmt.Sprintf(
		"%v tries to use %v on %v, but it doesn't work for some reason.\n",
		user.GetName(),
		p.GetName(),
		target.GetName(),
	), false, false
}
