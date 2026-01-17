package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
	"math"
	"math/rand/v2"
)

const (
	defaultGoblinName   = "Goblin"
	defaultGoblinHealth = 30
	defaultGoblinDamage = 5

	surnameChanceGoblin = 1

	dropChanceGoblin = 0.25

	moveChanceGoblinChase  = 0.6
	moveChanceGoblinRandom = 0.2
	moveChanceGoblinIdle   = 0.2
)

var (
	lootTableGoblin = map[game.Item]int{
		NewKey():                           30,
		NewHealingPotion("Suspicious", 15): 70,
	}

	takenGoblinSurnames  = map[string]bool{}
	randomGoblinSurnames = []string{
		"Archibald",
		"Blot",
		"Cook",
		"Drake",
		"Elph",
		"Fenrir",
		"Grok",
		"Hugo",
		"Jinx",
		"Kip",
		"Lug",
		"Muck",
		"Nibble",
		"Oscar",
		"Pax",
		"Quill",
		"Rattle",
		"Shank",
		"Till",
		"Urk",
		"Vorp",
		"Zig",
	}
)

var (
	colGoblin = utils.NewColor(utils.ColorFgBold, utils.ColorFgGreen)
)

// TODO: Implement speed potion
type Goblin struct {
	Name   string
	Health int
	Damage int

	// Potion effects
	speedPotion bool
}

func NewGoblin() *Goblin {
	surname := ""
	if rand.Float32() < surnameChanceGoblin {
		var randomSurname string
		for {
			randomSurname, _ = utils.RandChoice(randomGoblinSurnames)
			if !takenGoblinSurnames[randomSurname] && len(takenGoblinSurnames) < len(randomGoblinSurnames) {
				takenGoblinSurnames[randomSurname] = true
				break
			}
		}
		surname = fmt.Sprintf(" %v", randomSurname)
	}

	return &Goblin{
		Name:   defaultGoblinName + surname,
		Health: defaultGoblinHealth,
		Damage: defaultGoblinDamage,
	}
}

func (g *Goblin) GetName() string {
	return colGoblin(g.Name)
}

func (g *Goblin) GetStatus() string {
	return game.FormatHealth(g.Health, false)
}

func (g *Goblin) GetDesc(user game.Entity) string {
	return fmt.Sprintf(
		"%v tries to strike up a conversation with %v, but it seems uninterested in talking.\n",
		user.GetName(),
		g.GetName(),
	)
}

func (g *Goblin) GetHealth() int {
	return g.Health
}

func (g *Goblin) AddHealth(amount int) (string, bool) {
	g.Health += amount

	if g.Health <= 0 {
		return fmt.Sprintf(
			"%v has perished.",
			g.GetName(),
		), false
	}

	return game.GetHealthStatusResponse(g), true
}

func (g *Goblin) GetLoot(user game.Entity) []game.Item {
	return game.GetRandomLoot(lootTableGoblin, utils.RandIntInRange(0, 1))
}

func (g *Goblin) BeforeTurn(context *game.Context) {}

func (g *Goblin) OnTurn(context *game.Context) (string, bool) {
	occupants := context.World.GetOccupantsSameTile(g)
	for _, entity := range occupants {
		if _, isGoblin := entity.(*Goblin); isGoblin {
			continue
		}

		if entityHealth, hasHealth := entity.(game.EntityHealth); hasHealth && entityHealth.GetHealth() > 0 {
			response, _ := entityHealth.AddHealth(-g.Damage)
			return fmt.Sprintf(
				"%v attacks %v for %v damage! %v\n",
				g.GetName(),
				entity.GetName(),
				game.FormatDamage(g.Damage, false),
				response,
			), true
		}
	}

	speedPotionActive := g.speedPotion
	rgn := rand.Float32()
	switch {
	case rgn < moveChanceGoblinRandom:
		g.speedPotion = false
		return g.moveRandomly(context), true && !speedPotionActive
	case rgn < moveChanceGoblinRandom+moveChanceGoblinChase:
		g.speedPotion = false
		return g.moveTowardsClosestNonGoblin(context), true && !speedPotionActive
	default:
		return fmt.Sprintf("%v stays still, observing the surroundings.\n", g.GetName()), true
	}
}

func (g *Goblin) ApplySpeedPotion() {
	g.speedPotion = true
}

func (g *Goblin) moveRandomly(context *game.Context) string {
	dx, dy := rand.IntN(3)-1, rand.IntN(3)-1
	result, _ := context.World.MoveInDirection(g, dx, dy)
	return result
}

func (g *Goblin) moveTowardsClosestNonGoblin(context *game.Context) string {
	dx, dy, target := g.getDirectionToClosestNonGoblin(context)

	if target != nil && (dx != 0 || dy != 0) {
		result, _ := context.World.MoveInDirection(g, dx, dy)
		return fmt.Sprintf("%sMoving towards %v.\n", result, target.GetName())
	}

	return fmt.Sprintf("%v looks around, but finds no one to hunt.\n", g.GetName())
}

func (g *Goblin) getDirectionToClosestNonGoblin(context *game.Context) (int, int, game.Entity) {
	currentPos, ok := context.World.Positions[g]
	if !ok {
		return 0, 0, nil
	}

	var target game.Entity
	minDistance := math.MaxFloat64

	for entity, pos := range context.World.Positions {
		if _, isGoblin := entity.(*Goblin); isGoblin {
			continue
		}
		if entityHealth, hasHealth := entity.(game.EntityHealth); hasHealth && entityHealth.GetHealth() > 0 {
			// Calculate Euclidean distance
			dist := math.Sqrt(math.Pow(float64(pos.X-currentPos.X), 2) + math.Pow(float64(pos.Y-currentPos.Y), 2))

			if dist < minDistance {
				minDistance = dist
				target = entity
			}
		}

	}

	if target == nil {
		return 0, 0, nil
	}

	targetPos := context.World.Positions[target]

	dx, dy := 0, 0
	if targetPos.X > currentPos.X {
		dx = 1
	} else if targetPos.X < currentPos.X {
		dx = -1
	}
	if targetPos.Y > currentPos.Y {
		dy = 1
	} else if targetPos.Y < currentPos.Y {
		dy = -1
	}

	return dx, dy, target
}
