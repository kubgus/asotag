package content

import (
	"fmt"
	"math"
	"math/rand/v2"
	"text-adventure-game/game"
	"text-adventure-game/utils"
)

const (
	defaultGoblinName = "Goblin"
	defaultGoblinHealth = 30
	defaultGoblinDamage = 5

	surnameChanceGoblin = 1

	dropChanceGoblin = 0.25

	moveChanceGoblinChase = 0.6
	moveChanceGoblinRandom = 0.2
	moveChanceGoblinIdle = 0.2
)

var (
	dropItemsGoblin = []game.Item{
		NewKey(),
		NewHealingPotion("Suspicious", 15),
	}
)

var (
	randomGoblinSurnames = []string{
		"Gerlach",
		"Grok",
		"Zig",
		"Steve",
		"Juliette",
		"Gnasher",
		"Skritch",
		"Zog",
		"Bilge",
		"Rattle",
		"Snaggle",
		"Kroak",
		"Fizzle",
		"Blight",
		"Gribbit",
		"Muck",
		"Scab",
		"Thistle",
		"Grime",
		"Pox",
		"Jinx",
		"Skitter",
		"Grub",
		"Vex",
		"Wrench",
		"Spite",
		"Krag",
		"Nibble",
		"Sludge",
		"Twitch",
		"Guzzle",
		"Blot",
		"Shank",
		"Murk",
		"Zap",
	}
)

var (
	fmtGoblin = utils.NewColor(utils.ColorFgBold, utils.ColorFgGreen)
)

type Goblin struct {
	Name string
	Health int
	Damage int
}

func NewGoblin() *Goblin {
	surname := ""
	if rand.Float32() < surnameChanceGoblin {
		surname = fmt.Sprintf(" %v",
			randomGoblinSurnames[rand.IntN(len(randomGoblinSurnames))],
			)
	}

	return &Goblin{
		Name: defaultGoblinName + surname,
		Health: defaultGoblinHealth,
		Damage: defaultGoblinDamage,
	}
}

func (g *Goblin) GetName() string {
	return fmtGoblin(g.Name)
}

func (g *Goblin) GetHealth() int { return g.Health }

func (g *Goblin) GetHealthString(includeWordHealth bool) string {
	result := fmt.Sprintf("%d", g.GetHealth())
	if includeWordHealth {
		result += " health"
	}
	return game.FmtHealth(result)
}

func (g *Goblin) AddHealth(amount int) (string, bool) {
	g.Health += amount
	if g.Health <= 0 {
		return fmt.Sprintf(
			"%v has perished.",
			g.GetName(),
			), false
	}
	return fmt.Sprintf(
		"%v is now at %v.",
		g.GetName(),
		g.GetHealthString(true),
		), true
}

func (g *Goblin) Examine(user game.Entity) string {
	return fmt.Sprintf(
		"%v tries to strike up a conversation with %v, but it seems uninterested in talking.\n",
		user.GetName(),
		g.GetName(),
		)
}

func (g *Goblin) Loot(user game.Entity) []game.Item {
	rgn := rand.Float32()
	switch {
	case rgn < dropChanceGoblin:
		item := dropItemsGoblin[rand.IntN(len(dropItemsGoblin))]
		return []game.Item{item}
	default:
		return []game.Item{}
	}
}

func (g *Goblin) Reset(context *game.Context) { }

func (g *Goblin) Move(context *game.Context) (string, bool) {
	occupants := context.World.GetOccupantsSameTile(g)
	for _, entity := range occupants {
		if _, isGoblin := entity.(*Goblin); !isGoblin && entity.GetHealth() > 0 {
			response, _ := entity.AddHealth(-g.Damage)
			return fmt.Sprintf("%v attacks %v for %d damage! %v\n",
				g.GetName(), entity.GetName(), g.Damage, response), true
		}
	}

	rgn := rand.Float32()
	switch {
	case rgn < moveChanceGoblinRandom:
		return g.moveRandomly(context), true
	case rgn < moveChanceGoblinRandom+moveChanceGoblinChase:
		return g.moveTowardsClosestNonGoblin(context), true
	default:
		return fmt.Sprintf("%v stays still, observing the surroundings.\n", g.GetName()), true
	}
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
		if _, isGoblin := entity.(*Goblin); isGoblin || entity.GetHealth() <= 0 {
			continue
		}

		// Calculate Euclidean distance
		dist := math.Sqrt(math.Pow(float64(pos.X-currentPos.X), 2) + math.Pow(float64(pos.Y-currentPos.Y), 2))

		if dist < minDistance {
			minDistance = dist
			target = entity
		}
	}

	if target == nil {
		return 0, 0, nil
	}

	targetPos := context.World.Positions[target]

	dx, dy := 0, 0
	if targetPos.X > currentPos.X { dx = 1 } else if targetPos.X < currentPos.X { dx = -1 }
	if targetPos.Y > currentPos.Y { dy = 1 } else if targetPos.Y < currentPos.Y { dy = -1 }

	return dx, dy, target
}
