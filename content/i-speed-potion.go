package content

import (
	"fmt"
	"strconv"
	"text-adventure-game/game"
)

type SpeedPotion struct {
	Prefix string
	Magnitude int
}

func NewSpeedPotion(prefix string, magnitude int) *SpeedPotion {
	return &SpeedPotion{
		Prefix: prefix,
		Magnitude: magnitude,
	}
}

func (p *SpeedPotion) GetName() string {
	return game.ColItem(fmt.Sprintf("%v Speed Potion", p.Prefix))
}

func (p *SpeedPotion) GetDesc() string {
	return fmt.Sprintf(
		"Allows the use to move %v extra spaces in a single turn when used.",
		game.ColItem(strconv.Itoa(p.Magnitude)),
		)
}

func (p *SpeedPotion) Use(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	targetSpeedPotionable, ok := target.(EntitySpeedPotionable)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}

	targetSpeedPotionable.ApplySpeedPotion(p.Magnitude)

	return fmt.Sprintf(
		"%v uses %v on %v, granting the ability to move %v extra spaces in a single turn!\n",
		user.GetName(),
		p.GetName(),
		target.GetName(),
		game.ColItem(strconv.Itoa(p.Magnitude)),
	), true, true
}
