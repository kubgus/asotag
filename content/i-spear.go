package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
)

type Spear struct {
	Name      string
	MinDamage int
	MaxDamage int
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

func (s *Spear) UseInDirection(
	user game.Entity,
	dx, dy int,
	direction string,
	context *game.Context,
) (string, bool, bool) {
	px, py, ok := context.World.GetEntityPos(user)
	if !ok {
		return game.SnipItemCannotBeUsedBy(user, s), false, false
	}

	targets := context.World.GetEntitiesAt(
		px+dx,
		py+dy,
	)

	var target game.Entity
	var targetHealth game.EntityHealth
	for _, entity := range utils.Shuffled(targets) {
		if entityHealth, ok := entity.(game.EntityHealth); ok {
			target = entity
			targetHealth = entityHealth
			break
		}
	}
	if target == nil {
		return fmt.Sprintf(
			"%v throws %v %v into the distance, hitting nothing.\n",
			user.GetName(),
			s.GetName(),
			game.ColAction(direction),
			), true, true
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
		game.ColAction(direction),
		target.GetName(),
		game.FormatDamage(damage, false),
		response,
	), true, true
}
