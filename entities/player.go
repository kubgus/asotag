package entities

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strconv"
	"text-adventure-game/game"
	"text-adventure-game/utils"
)

const (
	defaultPlayerName = "Hero"
	defaultPlayerHealth = 100

	maxLooksPerTurnPlayer = 2
	lookActionDiscoveryChancePlayer = 0.8
)

type Player struct {
	Name string
	Health int

	Inventory []game.Item

	looksThisTurn int
}

func NewPlayer(name string) *Player {
	nameToUse := name
	if nameToUse == "" {
		nameToUse = defaultPlayerName
	}

	return &Player{
		Name: nameToUse,
		Health: defaultPlayerHealth,

		Inventory: []game.Item{},
	}
}

func (p *Player) GetName() string {
	return game.FmtHero(p.Name)
}

func (p *Player) GetHealth() int {
	return p.Health
}

func (p *Player) GetHealthString(includeWordHealth bool) string {
	var result string
	if includeWordHealth {
		result = fmt.Sprintf("%d health", p.GetHealth())
	} else {
		result = fmt.Sprintf("%d", p.GetHealth())
	}
	return game.FmtHealth(result)
}

func (p *Player) AddHealth(amount int) string {
	p.Health += amount

	if p.Health <= 0 {
		return fmt.Sprintf(
			game.FmtHealth("%v has perished."),
			p.GetName(),
			)
	}

	return fmt.Sprintf(
		"%v is now at %v.",
		p.GetName(),
		p.GetHealthString(true),
		)
}

func (p *Player) Reset(context *game.Context) {
	p.looksThisTurn = 0
}

func (p *Player) Move(context *game.Context) (string, bool) {
	fmt.Printf(
		"%v %v\n",
		game.FmtTooltip("Nearby:"),
		listEntities(context.World.GetOccupantsSameTile(p)),
		)

	fmt.Printf(
		"%v %v\n",
		game.FmtTooltip("Actions:"),
		listActions(actions),
		)

	fmt.Println("\nChoose your action:")

	var input string
	fmt.Print(game.FmtTooltip("> "))
	fmt.Scan(&input)
	fmt.Println()

	action, exists := actions[input]
	if !exists {
		return fmt.Sprintf(
			"Invalid action. (%v)\n",
			game.FmtTooltip(input),
			), false
	}

	return action(p, context)
}

type actionFunc func(player *Player, context *game.Context) (string, bool)

var actions = map[string]actionFunc{
	"inventory": func(player *Player, context *game.Context) (string, bool) {
		if len(player.Inventory) == 0 {
			return game.FmtTooltip("Your inventory is empty.\n"), false
		}
		return fmt.Sprintf("%s\n", listOrderedItems(player.Inventory)), false
	},
	"use": func(player *Player, context *game.Context) (string, bool) {
		if len(player.Inventory) == 0 {
			return "No item in inventory to use.", false
		}

		neighbors := context.World.GetOccupantsSameTile(player)

		if len(neighbors) == 0 {
			return "No targets available to use the item on.\n", false
		}

		fmt.Println(game.FmtTooltip("Select an item to use:"))
		fmt.Println(listOrderedItems(player.Inventory))

		var input string
		fmt.Print(game.FmtTooltip("> "))
		fmt.Scan(&input)
		fmt.Println()
		index, err := strconv.Atoi(input)
		if err != nil || index < 0 || index >= len(player.Inventory) {
			return fmt.Sprintf(
				"Invalid item selection. (%v)\n",
				game.FmtTooltip(input),
				), false
		}
		item := player.Inventory[index]

		fmt.Println(game.FmtTooltip("Select a target:"))
		fmt.Println(listOrderedEntities(neighbors))

		var targetInput string
		fmt.Print(game.FmtTooltip("> "))
		fmt.Scan(&targetInput)
		fmt.Println()
		targetIndex, err := strconv.Atoi(targetInput)
		if err != nil || targetIndex < 0 || targetIndex >= len(neighbors) {
			return fmt.Sprintf(
				"Invalid target selection. (%v)\n",
				targetInput,
				), false
		}
		target := neighbors[targetIndex]

		result := item.Use(player, target)
		return result, true
	},
	"move": func(player *Player, context *game.Context) (string, bool) {
		fmt.Printf("Choose a direction to move: %v\n", game.ListDirections())
		var input string
		fmt.Print(game.FmtTooltip("> "))
		fmt.Scan(&input)
		fmt.Println()

		dx, dy, valid := game.DirToDelta(input)
		if !valid {
			return fmt.Sprintf(
				"Invalid direction. (%v)\n",
				game.FmtTooltip(input),
				), false
		}

		return context.World.MoveInDirection(player, dx, dy)
	},
	"look": func(player *Player, context *game.Context) (string, bool) {
		if player.looksThisTurn >= maxLooksPerTurnPlayer {
			return "Your eyes are already strained from looking around this turn.\n", false
		}

		fmt.Printf("Choose a direction to look: %v\n", game.ListDirections())
		var input string
		fmt.Print(game.FmtTooltip("> "))
		fmt.Scan(&input)
		fmt.Println()

		dx, dy, valid := game.DirToDelta(input)
		if !valid {
			return fmt.Sprintf(
				"Invalid direction. (%v)\n",
				input,
			), false
		}

		currX, currY, ok := context.World.GetEntityPos(player)
		if !ok {
			return "You are lost.\n", false
		}

		player.looksThisTurn++

		for dist := 1; dist < context.World.Size; dist++ {
			tx, ty := currX+(dx*dist), currY+(dy*dist)

			if tx < 0 || tx >= context.World.Size || ty < 0 || ty >= context.World.Size {
				break
			}

			tileEntities := context.World.GetEntitiesAt(tx, ty)
			if len(tileEntities) > 0 {
				var discovered []string
				for _, e := range tileEntities {
					if rand.Float64() <= lookActionDiscoveryChancePlayer {
						discovered = append(discovered, e.GetName())
					}
				}

				distanceStr := fmt.Sprintf("%d tile(s) away", dist)
				if len(discovered) > 0 {
					return fmt.Sprintf("You spot %v about %v.\n",
						utils.JoinWithLast(discovered, ", ", " and "),
						game.FmtItem(distanceStr)), false
				}
				return fmt.Sprintf("You see %v moving about %v.\n",
					game.FmtItem("something"),
					game.FmtItem(distanceStr)), false
			}
		}

		return "You see nothing but the horizon.\n", false
	},
	"wait": func(player *Player, context *game.Context) (string, bool) {
		return fmt.Sprintf("%v waits.\n", player.GetName()), true
	},
}

func listActions(actions map[string]actionFunc) string {
	var keys []string
	for k := range actions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return utils.JoinWithMapFunc(keys, ", ", func(i int, action string) string {
		return game.FmtAction(action)
	})
}

func listEntities(entities []game.Entity) string {
	if len(entities) == 0 {
		return "None"
	}

	var keys []string
	for _, entity := range entities {
		keys = append(keys, fmt.Sprintf(
			"%v(%v)",
			entity.GetName(),
			entity.GetHealthString(false),
			))
	}
	return utils.JoinWithLast(keys, ", ", " and ")
}

func listOrderedEntities(items []game.Entity) string {
	if len(items) == 0 {
		return "Empty"
	}

	return utils.JoinWithMapFunc(items, "\n", func(i int, entity game.Entity) string {
		return fmt.Sprintf(
			"%v: %v(%v)",
			game.FmtAction(strconv.Itoa(i)),
			entity.GetName(),
			entity.GetHealthString(false),
			)
	})
}

func listOrderedItems(items []game.Item) string {
	if len(items) == 0 {
		return "Empty"
	}

	return utils.JoinWithMapFunc(items, "\n", func(i int, item game.Item) string {
		return fmt.Sprintf(
			"%v: %v %v %v",
			game.FmtAction(strconv.Itoa(i)),
			item.GetName(),
			game.FmtTooltip("-"),
			item.GetDesc(),
			)
	})
}
