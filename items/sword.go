package items

import (
	"fmt"
	"strconv"
	"text-adventure-game/game"
)

type Sword struct {
	Name string
	Damage int
}

func (s *Sword) GetName() string {
	return game.FmtItem(s.Name)
}

func (s *Sword) GetDesc() string {
	return fmt.Sprintf(
		"Deals %v close-range damage.",
		game.FmtDamage(strconv.Itoa(s.Damage)),
		)
}

func (s *Sword) Use(user, target game.Entity) string {
	damage := s.Damage

	targetStatus := target.AddHealth(-damage)

	return fmt.Sprintf(
		"%v slices %v with a %v for %v damage! %v\n",
		user.GetName(),
		target.GetName(),
		s.GetName(),
		game.FmtDamage(strconv.Itoa(damage)),
		targetStatus,
		)
}
