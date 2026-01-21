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

func NewSpear(name string, minDamage, maxDamage int) *Spear {
	return &Spear{
		Name:      name,
		MinDamage: minDamage,
		MaxDamage: maxDamage,
	}
}

func NewSpearWooden() *Spear {
	return NewSpear("Wooden Spear", 4, 8)
}

func NewSpearIron() *Spear {
	return NewSpear("Iron Spear", 11, 18)
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
	var targetHealth game.HasHealth
	for _, entity := range utils.Shuffled(targets) {
		if entityHealth, ok := entity.(game.HasHealth); ok {
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

	response := targetHealth.GetHealth().Change(-damage)

	// TODO: temporary solution
	if targetHealth.GetHealth().CurrentHealth <= 0 {
		if targetLoot, hasLoot := target.(HasLoot); hasLoot {
			if player, isPlayer := user.(*Player); isPlayer {
				loot := targetLoot.GetLoot().Drop()
				response += "\n" + player.GetInventory().AddItems(loot)
			} else {
				return game.SnipItemCannotBeUsedBy(user, s), false, false
			}
		}
	}

	return fmt.Sprintf(
		"%s throws %s %s at %s! %v\n",
		user.GetName(),
		s.GetName(),
		game.ColAction(direction),
		target.GetName(),
		response,
	), true, true
}
