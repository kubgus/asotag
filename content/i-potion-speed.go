package content

import (
	"asotag/game"
	"fmt"
)

type SpeedPotion struct{}

func NewSpeedPotion() *SpeedPotion {
	return &SpeedPotion{}
}

func (p *SpeedPotion) GetName() string {
	return game.ColItem("Speed Potion")
}

func (p *SpeedPotion) GetDesc() string {
	return "Makes the user's next move not end their turn."
}

func (p *SpeedPotion) UseOnEntity(
	user, target game.Entity,
	_ *game.Context,
) (string, bool, bool) {
	targetMovement, ok := target.(HasMovement)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}

	targetMovement.GetMovement().ExtraMoves += 1

	return fmt.Sprintf(
		"%s uses %s on %s.\n%s's next move won't end their turn!\n",
		user.GetName(),
		p.GetName(),
		target.GetName(),
		target.GetName(),
	), false, true
}
