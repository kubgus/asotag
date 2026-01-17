package content

import (
	"fmt"
	"text-adventure-game/game"
	"text-adventure-game/utils"
)

type Spear struct {
	Name string
	MinDamage int
	MaxDamage int
}

func NewSpear(name string, minDamage, maxDamage int) *Spear {
	return &Spear{
		Name: name,
		MinDamage: minDamage,
		MaxDamage: maxDamage,
	}
}

func (s *Spear) GetName() string {
	return game.ColItem(s.Name)
}

func (s *Spear) GetDesc() string {
	return fmt.Sprintf(
		"Deals %v to %v damage when thrown at an adjacent square.",
		game.FormatDamage(s.MinDamage, false),
		game.FormatDamage(s.MaxDamage, false),
	)
}

func (s *Spear) Use(user, _ game.Entity, context *game.Context) (string, bool, bool) {
	fmt.Printf(
		game.ColTooltip("Select a direction to throw: %v\n"),
		game.ListDirections(),
		)
	input := game.Input()
	dx, dy, valid := game.DirToDelta(input)
	if !valid {
		return game.SnipInvalidDirection(input), false, false
	}

	px, py, ok := context.World.GetEntityPos(user)
	if !ok {
		return game.SnipItemCannotBeUsedBy(user, s), false, false
	}

	targets := context.World.GetEntitiesAt(
		px + dx,
		py + dy,
	)

	target, ok := utils.RandChoice(targets)
	if !ok {
		return fmt.Sprintf(
			"%v throws %v %v into the distance, hitting nothing.\n",
			user.GetName(),
			s.GetName(),
			game.ColAction(input),
		), true, true
	}

	targetHealth, ok := target.(game.EntityHealth)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, s), false, false
	}

	damage := utils.RandIntInRange(s.MinDamage, s.MaxDamage)

	response, alive := targetHealth.AddHealth(-damage)

	if !alive {
		if targetLoot, ok := target.(game.EntityLoot); ok {
			loot := targetLoot.GetLoot(user)
			if player, ok := user.(*Player); ok {
				if responseLoot, ok := player.CollectLoot(loot); ok {
					response += "\n" + responseLoot
				}
			} else {
				return game.SnipItemCannotBeUsedBy(user, s), false, false
			}
		}
	}

	return fmt.Sprintf(
		"%v throws %v %v at %v, dealing %v damage! %v\n",
		user.GetName(),
		s.GetName(),
		game.ColAction(input),
		target.GetName(),
		game.FormatDamage(damage, false),
		response,
	), true, true
}
