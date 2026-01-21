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
	return "Makes next move not end turn."
}

func (p *SpeedPotion) UseOnEntity(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	targetMovement, ok := target.(HasMovement)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}

	targetMovement.GetMovement().ExtraMoves += 1

	return fmt.Sprintf(
		"%v uses %v on %v. %v's next move won't end their turn!\n",
		user.GetName(),
		p.GetName(),
		target.GetName(),
		target.GetName(),
	), true, true
}
