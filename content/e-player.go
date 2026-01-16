package content

import (
	"fmt"
	"math/rand/v2"
	"slices"
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

func (p *Player) AddHealth(amount int) (string, bool) {
	p.Health += amount

	if p.Health <= 0 {
		return fmt.Sprintf(
			game.FmtHealth("%v has perished."),
			p.GetName(),
			), false
	}

	return fmt.Sprintf(
		"%v is now at %v.",
		p.GetName(),
		p.GetHealthString(true),
		), true
}

func (p *Player) Examine(user game.Entity) string {
	if user != p {
		return fmt.Sprintf(
			"%v looks at %v with a puzzled expression.\n",
			p.GetName(),
			user.GetName(),
			)
	}
	return fmt.Sprintf(
		"%v contemplates exsistence.\n",
		p.GetName(),
		)
}

func (p *Player) Loot(user game.Entity) []game.Item {
	return []game.Item{}
}

func (p *Player) Reset(context *game.Context) {
	p.looksThisTurn = 0
}

func (p *Player) Move(context *game.Context) (string, bool) {
	fmt.Printf(
		"%v %v\n",
		game.FmtTooltip("Nearby:"),
		game.ListEntities(context.World.GetOccupantsSameTile(p)),
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
	"bundle": func(player *Player, context *game.Context) (string, bool) {
		if len(player.Inventory) == 0 {
			return "No items in inventory to bundle.\n", false
		}

		bundleIdxs := []int{}

		for i := 0; ; i++ {
			fmt.Println(game.FmtTooltip("Select an item to bundle or bundle to open:"))
			fmt.Println(game.ListOrderedItemsWithMapFunc(
				player.Inventory, func(idx int, item game.Item) string {
					if slices.Contains(bundleIdxs, idx) {
						return game.FmtSystem("(selected for bundling)")
					}
					return item.GetDesc()
				},
				))

			var input string
			fmt.Print(game.FmtTooltip("> "))
			fmt.Scan(&input)
			fmt.Println()
			index, err := strconv.Atoi(input)

			if err == nil && slices.Contains(bundleIdxs, index) {
				bundleIdxs = append(
					bundleIdxs[:slices.Index(bundleIdxs, index)],
					bundleIdxs[slices.Index(bundleIdxs, index)+1:]...,
				)

				fmt.Printf(
					"%v deselected for bundling.\n\n",
					player.Inventory[index].GetName(),
					)

				continue
			}

			newBundle, isBundleSelected := player.Inventory[index].(*Bundle)

			if i == 0 && isBundleSelected {
				// Unpack existing bundle
				for _, item := range newBundle.Items {
					player.Inventory = append(player.Inventory, item)
				}
				player.Inventory = append(
					player.Inventory[:index],
					player.Inventory[index+1:]...,
				)

				return fmt.Sprintf(
					"%v unbundles %v.\n",
					player.GetName(),
					game.ListItems(newBundle.Items),
					), false
			}

			isInvalidInput := err != nil || index < 0 || index >= len(player.Inventory)
			if isInvalidInput || isBundleSelected {
				if len(bundleIdxs) < 1 {
					return "Nothing to bundle.\n", false
				}

				items := []game.Item{}
				for _, idx := range bundleIdxs {
					items = append(items, player.Inventory[idx])
				}
				// Remove bundled items from inventory
				// after collecting them to bundle
				// to avoid index shifting issues
				for _, idx := range bundleIdxs {
					player.Inventory = append(
						player.Inventory[:idx],
						player.Inventory[idx+1:]...,
					)
				}

				var wording string
				if isInvalidInput {
					wording = "creates a"
					newBundle = NewBundle(
						items,
						)
					player.Inventory = append(player.Inventory, newBundle)
				} else {
					wording = "fills the"
					newBundle.Items = append(newBundle.Items, items...)
				}

				return fmt.Sprintf(
					"%v %v %v with %v.\n",
					player.GetName(),
					wording,
					newBundle.GetName(),
					game.ListItems(newBundle.Items),
					), false
			}

			bundleIdxs = append(bundleIdxs, index)

			fmt.Printf(
				"%v selected for bundling.\n%v\n\n",
				player.Inventory[index].GetName(),
				game.FmtTooltip("(Enter 'x' to finish bundling)"),
				)
		}
	},
	"inventory": func(player *Player, context *game.Context) (string, bool) {
		if len(player.Inventory) == 0 {
			return game.FmtTooltip("Your inventory is empty.\n"), false
		}
		return fmt.Sprintf("%s\n", game.ListOrderedItems(player.Inventory)), false
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
		fmt.Println(game.ListOrderedItems(player.Inventory))

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
		fmt.Println(game.ListOrderedEntities(neighbors))

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

		result, ok, consume := item.Use(player, target)

		if consume {
			result += fmt.Sprintf(
				"%v removed from inventory.\n",
				item.GetName(),
				)

			player.Inventory = append(
				player.Inventory[:index],
				player.Inventory[index+1:]...,
			)
		}

		return result, ok
	},
	"examine": func(player *Player, context *game.Context) (string, bool) {
		neighbors := context.World.GetOccupantsSameTile(player)

		if len(neighbors) == 0 {
			return "Nothing nearby to examine.\n", false
		}

		fmt.Println(game.FmtTooltip("Select what you want to examine:"))
		fmt.Println(game.ListOrderedEntities(neighbors))

		var input string
		fmt.Print(game.FmtTooltip("> "))
		fmt.Scan(&input)
		fmt.Println()
		index, err := strconv.Atoi(input)
		if err != nil || index < 0 || index >= len(neighbors) {
			return fmt.Sprintf(
				"Invalid selection. (%v)\n",
				game.FmtTooltip(input),
				), false
		}

		entity := neighbors[index]
		response := entity.Examine(player)
		return response, false
	},
	"move": func(player *Player, context *game.Context) (string, bool) {
		fmt.Printf(
			"%v %v\n",
			game.FmtTooltip("Choose a direction to move:"),
			game.ListDirections())
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
			return fmt.Sprintf(
				"%v's eyes are already strained from looking around this turn.\n",
				player.GetName(),
				), false
		}

		fmt.Printf(
			"%v %v\n",
			game.FmtTooltip("Choose a direction to look:"),
			game.ListDirections())
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

		currX, currY, ok := context.World.GetEntityPos(player)
		if !ok {
			return fmt.Sprintf(
				"%v is lost.",
				player.GetName(),
			), false
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
					return fmt.Sprintf("%v spots %v about %v.\n",
						player.GetName(),
						utils.JoinWithLast(discovered, ", ", " and "),
						game.FmtItem(distanceStr)), false
				}
				return fmt.Sprintf("%v sees %v moving about %v.\n",
					player.GetName(),
					game.FmtItem("something"),
					game.FmtItem(distanceStr)), false
			}
		}

		return fmt.Sprintf(
			"%v sees nothing of interest in that direction.\n",
			player.GetName(),
			), false
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
