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

	cheatCommand = "UUDDLRLRBA"
	cheatHealthBoost = 200
)

type Player struct {
	Name string
	Health int

	Inventory []game.Item

	// Trackers
	looksThisTurn int
	movesThisTurn int

	// Potion effects
	extraMoves int
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
	return game.ColHero(p.Name)
}

func (p *Player) GetHealth() int {
	return p.Health
}

func (p *Player) GetStatus() string {
	return game.FormatHealth(p.Health, false)
}

func (p *Player) AddHealth(amount int) (string, bool) {
	p.Health += amount

	if p.Health <= 0 {
		return fmt.Sprintf(
			"%v has been knocked out.",
			p.GetName(),
			), false
	}

	return game.GetHealthStatusResponse(p), true
}

func (p *Player) GetDesc(user game.Entity) string {
	if user != p {
		return fmt.Sprintf(
			"%v looks at %v with a puzzled expression.\n",
			p.GetName(),
			user.GetName(),
			)
	}
	return fmt.Sprintf(
		"%v contemplates existence.\n",
		p.GetName(),
		)
}

func (p *Player) BeforeTurn(context *game.Context) {
	p.looksThisTurn = 0
	p.movesThisTurn = 0
}

func (p *Player) OnTurn(context *game.Context) (string, bool) {
	fmt.Printf(
		"%v %v\n",
		game.ColTooltip("Nearby:"),
		game.ListEntities(context.World.GetOccupantsSameTile(p)),
		)

	fmt.Printf(
		"%v %v\n",
		game.ColTooltip("Actions:"),
		listActions(actions),
		)

	if p.extraMoves > 0 {
		fmt.Printf(
			"%v %v\n",
			game.ColTooltip("Extra Moves:"),
			game.ColSystem(strconv.Itoa(p.extraMoves)),
		)
	}

	fmt.Println("\nChoose your action:")
	input := game.Input()

	if input == cheatCommand {
		return p.ApplyCheats(context), false
	}

	action, exists := actions[input]
	if !exists {
		return fmt.Sprintf(
			"Invalid action. (%v)\n",
			game.ColTooltip(input),
			), false
	}

	return action(p, context)
}

func (p *Player) ApplySpeedPotion(magnitude int) {
	p.extraMoves += magnitude
}

func (p *Player) CollectLoot(loot []game.Item) (string, bool) {
	if len(loot) == 0 {
		return "", false
	}

	p.Inventory = append(p.Inventory, loot...)

	return fmt.Sprintf(
		"%v collects %v.\n",
		p.GetName(),
		game.ListItems(loot),
		), true
}

func (p *Player) ApplyCheats(context *game.Context) string {
	context.CheatRevealMap = true
	if px, py, ok := context.World.GetEntityPos(p); ok {
		context.World.Add(NewWorkbench(), px, py, false)

		context.World.Add(NewDepositTree(50, 100), px, py, false)
		context.World.Add(NewDepositRock(50, 100), px, py, false)
		context.World.Add(NewDepositIronVein(50, 100), px, py, false)
		context.World.Add(NewDepositGoldVein(50, 100), px, py, false)
	}
	originalInventoryLen := len(p.Inventory)
	p.Inventory = append(
		p.Inventory,
		NewSwordGold(),
		NewPickaxe(MaterialGold),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		)
	return fmt.Sprintf(
		"%v activated a cheat code!\n" +
		"Map revealed!\n" +
		"A workbench appears nearby!\n" +
		"Gained %v!\n",
		p.GetName(),
		game.ListItems(p.Inventory[originalInventoryLen:]),
		)
}

type actionFunc func(player *Player, context *game.Context) (string, bool)

var actions = map[string]actionFunc{
	"bundle": func(player *Player, context *game.Context) (string, bool) {
		if len(player.Inventory) == 0 {
			return "No items in inventory to bundle.\n", false
		}

		bundleIdxs := []int{}

		for i := 0; ; i++ {
			fmt.Println(game.ColTooltip("Select an item to bundle or bundle to open:"))
			fmt.Println(game.ListOrderedItemsWithMapFunc(
				player.Inventory, func(idx int, item game.Item) string {
					if slices.Contains(bundleIdxs, idx) {
						return game.ColSystem("(selected for bundling)")
					}
					return item.GetDesc()
				},
				))
			input := game.Input()

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
				// Sort bundleIdxs in descending order
				// to safely remove items from inventory
				// without affecting the indexes of yet-to-be-removed items
				sort.Slice(bundleIdxs, func(i, j int) bool {
					return bundleIdxs[i] > bundleIdxs[j]
				})
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
				game.ColTooltip("(Enter 'x' to finish bundling)"),
				)
		}
	},
	"inventory": func(player *Player, context *game.Context) (string, bool) {
		if len(player.Inventory) == 0 {
			return game.ColTooltip("No inventory items.\n"), false
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

		fmt.Println(game.ColTooltip("Select an item to use:"))
		fmt.Println(game.ListOrderedItems(player.Inventory))
		input := game.Input()

		index, err := strconv.Atoi(input)
		if err != nil || index < 0 || index >= len(player.Inventory) {
			return fmt.Sprintf(
				"Invalid item selection. (%v)\n",
				game.ColTooltip(input),
				), false
		}
		item := player.Inventory[index]

		fmt.Println(game.ColTooltip("Select a target:"))
		fmt.Println(game.ListOrderedEntities(neighbors))

		var targetInput string
		fmt.Print(game.ColTooltip("> "))
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

		result, ok, consume := item.Use(player, target, context)

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

		fmt.Println(game.ColTooltip("Select what you want to examine:"))
		fmt.Println(game.ListOrderedEntities(neighbors))
		input := game.Input()

		index, err := strconv.Atoi(input)
		if err != nil || index < 0 || index >= len(neighbors) {
			return fmt.Sprintf(
				"Invalid selection. (%v)\n",
				game.ColTooltip(input),
				), false
		}

		entity := neighbors[index]
		response := entity.GetDesc(player)
		return response, false
	},
	"move": func(player *Player, context *game.Context) (string, bool) {
		fmt.Printf(
			"%v %v\n",
			game.ColTooltip("Choose a direction to move:"),
			game.ListDirections())
		input := game.Input()

		dx, dy, valid := game.DirToDelta(input)
		if !valid {
			return game.SnipInvalidDirection(input), false
		}

		response, ok := context.World.MoveInDirection(player, dx, dy)
		player.movesThisTurn++

		if player.movesThisTurn > 1 {
			player.extraMoves--
		}

		return response, ok && player.extraMoves == 0
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
			game.ColTooltip("Choose a direction to look:"),
			game.ListDirections())
		input := game.Input()

		dx, dy, valid := game.DirToDelta(input)
		if !valid {
			return fmt.Sprintf(
				"Invalid direction. (%v)\n",
				game.ColTooltip(input),
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
						game.ColItem(distanceStr)), false
				}
				return fmt.Sprintf("%v sees %v moving about %v.\n",
					player.GetName(),
					game.ColItem("something"),
					game.ColItem(distanceStr)), false
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
		return game.ColAction(action)
	})
}
