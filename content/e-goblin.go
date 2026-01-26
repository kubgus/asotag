package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
)

var (
	colGoblin = utils.NewColor(utils.ColorFgBold, utils.ColorFgGreen)
)

var (
	takenGoblinNames  = map[string]bool{}
	randomGoblinNames = []string{
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

type GoblinAction int

const (
	GoblinActionIdle GoblinAction = iota
	GoblinActionMoveRandom
	GoblinActionChaseEnemy
)

type Goblin struct {
	Name string

	health    game.HealthModule
	inventory InventoryModule
	loot      LootModule
	movement  MovementModule
}

func NewGoblin() *Goblin {
	goblin := Goblin{
		Name: "",

		health: game.HealthModule{
			CurrentHealth: 30,
		},
		inventory: InventoryModule{
			Items: []game.Item{
				NewSword("Rusty Dagger", 3, 7),
			},
		},
		loot: LootModule{
			LootTable: map[game.Item]int{
				NewHealingPotion("Suspicious", 30): 50,
				NewKey():                           40,
				NewResource(MaterialGold):          10,
			},
			AmountTable: map[int]int{
				0: 5,
				1: 70,
				2: 20,
				3: 5,
			},
			DropInventoryChance: 0.1,
			LootLimit:           1,
		},
		movement: MovementModule{},
	}

	var randomName string
	for {
		randomName, _ = utils.RandChoice(randomGoblinNames)
		if !takenGoblinNames[randomName] && len(takenGoblinNames) < len(randomGoblinNames) {
			takenGoblinNames[randomName] = true
			break
		}
	}
	goblin.Name = randomName

	return &goblin
}

func (g *Goblin) GetHealth() *game.HealthModule {
	g.health.Init(g)
	return &g.health
}

func (g *Goblin) GetInventory() *InventoryModule {
	g.inventory.Init(g)
	return &g.inventory
}

func (g *Goblin) GetLoot() *LootModule {
	g.loot.Init(g)
	return &g.loot
}

func (g *Goblin) GetMovement() *MovementModule {
	g.movement.Init(g)
	return &g.movement
}

func (g *Goblin) GetName() string {
	return colGoblin("Goblin " + g.Name)
}

func (g *Goblin) GetStatus() string {
	return game.ColHealth(strconv.Itoa(g.GetHealth().CurrentHealth))
}

func (g *Goblin) GetDesc(user game.Entity) string {
	return fmt.Sprintf(
		"%v tries to strike up a conversation with %v, but it seems uninterested in talking.\n",
		user.GetName(),
		g.GetName(),
	)
}

func (g *Goblin) BeforeTurn(context *game.Context) {}

func (g *Goblin) OnTurn(context *game.Context) (string, bool) {
	occupants := context.World.GetOccupantsSameTile(g)
	for _, entity := range occupants {
		if !g.isEnemy(entity) {
			continue
		}

		response, _ := g.GetInventory().UseItemOnEntity(
			0,
			entity,
			context,
		)
		return response, true
	}

	motivation := map[GoblinAction]int{
		GoblinActionChaseEnemy: g.GetHealth().CurrentHealth,
		GoblinActionMoveRandom: 30,
		GoblinActionIdle:       10,
	}

	action := utils.RandWeightedChoice(motivation)
	switch action {
	case GoblinActionMoveRandom:
		dx, dy := rand.IntN(3)-1, rand.IntN(3)-1
		response, endTurn := g.GetMovement().Move(dx, dy, &context.World)
		return response, endTurn
	case GoblinActionChaseEnemy:
		dx, dy, target := g.getDirectionToClosestEnemy(context)

		if target != nil {
			result, endTurn := g.GetMovement().Move(dx, dy, &context.World)
			return fmt.Sprintf(
				"%sMoving towards %s.\n",
				result,
				target.GetName(),
			), endTurn
		}

		return fmt.Sprintf("%v looks around, but finds no one to hunt.\n", g.GetName()), true
	default:
		// Idle
		return fmt.Sprintf("%v stays still, observing the surroundings.\n", g.GetName()), true
	}
}

func (g *Goblin) isEnemy(entity game.Entity) bool {
	if _, isActive := entity.(game.EntityActive); !isActive {
		return false
	}
	if entityHealth, hasHealth := entity.(game.HasHealth); !hasHealth || entityHealth.GetHealth().CurrentHealth <= 0 {
		return false
	}
	if _, isGoblin := entity.(*Goblin); isGoblin {
		return false
	}
	return true
}

func (g *Goblin) getDirectionToClosestEnemy(context *game.Context) (int, int, game.Entity) {
	currentPos, ok := context.World.Positions[g]
	if !ok {
		return 0, 0, nil
	}

	var target game.Entity
	minDistance := math.MaxFloat64

	for entity, pos := range context.World.Positions {
		if !g.isEnemy(entity) {
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
