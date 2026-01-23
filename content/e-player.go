package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strconv"
)

const (
	defaultNamePlayer               = "Hero"
	maxLooksPerTurnPlayer           = 2
	lookActionDiscoveryChancePlayer = 0.8

	cheatCommand = "UUDDLRLRBA"
)

type Player struct {
	Name string

	health    game.HealthModule
	inventory InventoryModule
	movement  MovementModule

	// Custom trackers
	looksThisTurn int
}

func NewPlayer(name string) *Player {
	player := Player{
		Name: name,

		health: game.HealthModule{
			CurrentHealth: 100,
			MaxHealth:     0,
		},
		inventory: InventoryModule{},
		movement:  MovementModule{},
	}

	if player.Name == "" {
		player.Name = defaultNamePlayer
	}

	return &player
}

func (p *Player) GetHealth() *game.HealthModule {
	p.health.Init(p)
	return &p.health
}

func (p *Player) GetInventory() *InventoryModule {
	p.inventory.Init(p)
	return &p.inventory
}

func (p *Player) GetMovement() *MovementModule {
	p.movement.Init(p)
	return &p.movement
}

func (p *Player) GetName() string {
	return game.ColHero(p.Name)
}

func (p *Player) GetStatus() string {
	return game.ColHealth(strconv.Itoa(p.GetHealth().CurrentHealth))
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

func (p *Player) ApplyCheats(context *game.Context) string {
	if px, py, ok := context.World.GetEntityPos(p); ok {
		context.World.Add(NewWorkbench(), px, py, false)

		context.World.Add(NewDepositTree(), px, py, false)
		context.World.Add(NewDepositRock(), px, py, false)
		context.World.Add(NewDepositIronVein(), px, py, false)
		context.World.Add(NewDepositGoldVein(), px, py, false)

		for range 5 {
			unlockedChest := NewChest()
			unlockedChest.IsUnlocked = true
			context.World.Add(unlockedChest, px, py, false)
		}
	}
	response := p.GetInventory().AddItems([]game.Item{
		NewSwordGolden(),
		NewPickaxeGolden(),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		NewHealingPotionSuperior(),
		NewSpeedPotion(),
		NewTeleportPotion(),
	})
	context.CheatRevealMap = true
	return fmt.Sprintf(
		"%v activates a cheat code!\n%v",
		p.GetName(),
		response,
	)
}

type actionFunc func(player *Player, context *game.Context) (string, bool)

var actions = map[string]actionFunc{
	"bundle": func(player *Player, context *game.Context) (string, bool) {
		if len(player.GetInventory().Items) == 0 {
			return "No items in inventory to bundle.\n", false
		}

		itemIndexes := []int{}
		for i := 0; ; {
			fmt.Println(game.ColTooltip("Select item to bundle or bundle to open:"))
			fmt.Println(game.ListOrderedItemsWithMapFunc(
				player.GetInventory().Items, func(idx int, item game.Item) string {
					if slices.Contains(itemIndexes, idx) {
						return game.ColSystem("(selected for bundling)")
					}
					return item.GetDesc()
				},
			))
			input := game.Input()

			itemIndex, indexErr := strconv.Atoi(input)

			if indexErr == nil && slices.Contains(itemIndexes, itemIndex) {
				itemIndexes = slices.Delete(
					itemIndexes,
					slices.Index(itemIndexes, itemIndex),
					slices.Index(itemIndexes, itemIndex)+1,
				)

				fmt.Printf(
					"%v deselected for bundling.\n\n",
					player.GetInventory().Items[itemIndex].GetName(),
				)

				i--
				continue
			}

			bundle, isBundleSelected := player.GetInventory().Items[itemIndex].(*Bundle)

			if i == 0 && isBundleSelected {
				removeResponse := player.GetInventory().RemoveItems([]int{itemIndex})
				addResponse := player.GetInventory().AddItems(bundle.Items)

				return addResponse + removeResponse, false
			}

			if !player.GetInventory().HasIndex(itemIndex) || indexErr != nil || isBundleSelected {
				if len(itemIndexes) < 1 {
					return "Nothing to bundle.\n", false
				}

				var response string
				if !isBundleSelected {
					bundle = NewBundle([]game.Item{})
					response = player.GetInventory().AddItems([]game.Item{bundle})
				} else {
					response = fmt.Sprintf(
						"%s adds items to %s.\n",
						player.GetName(),
						bundle.GetName(),
					)
				}

				for _, i := range itemIndexes {
					bundle.Items = append(
						bundle.Items,
						player.GetInventory().Items[i],
					)
				}

				removeResponse := player.GetInventory().RemoveItems(itemIndexes)

				return response + removeResponse, false
			}

			itemIndexes = append(itemIndexes, itemIndex)

			fmt.Printf(
				"%v selected for bundling.\n%v\n\n",
				player.GetInventory().Items[itemIndex].GetName(),
				game.ColTooltip("(Enter 'x' to finish bundling)"),
			)

			i++
		}
	},
	"inventory": func(player *Player, context *game.Context) (string, bool) {
		if len(player.GetInventory().Items) == 0 {
			return game.ColTooltip("No inventory items.\n"), false
		}
		return game.ListOrderedItems(player.GetInventory().Items) + "\n", false
	},
	"use": func(player *Player, context *game.Context) (string, bool) {
		if len(player.GetInventory().Items) == 0 {
			return "No item in inventory to use.", false
		}

		fmt.Println(game.ColTooltip("Select an item to use:"))
		fmt.Println(game.ListOrderedItems(player.GetInventory().Items))
		input := game.Input()

		itemIndex, indexErr := strconv.Atoi(input)

		if !player.GetInventory().HasIndex(itemIndex) || indexErr != nil {
			return game.SnipInvalidItemIndex(itemIndex), false
		}

		item := player.GetInventory().Items[itemIndex]

		if _, ok := item.(game.ItemUseEntity); ok {
			nearby := context.World.GetOccupantsSameTile(player)

			fmt.Printf(
				game.ColTooltip("Select a target:\n%v\n"),
				game.ListOrderedEntities(nearby),
			)
			var targetInput = game.Input()

			targetIndex, targetErr := strconv.Atoi(targetInput)

			if targetErr != nil || targetIndex < 0 || targetIndex >= len(nearby) {
				return fmt.Sprintf(
					"Invalid target selection. (%v)\n",
					game.ColTooltip(targetInput),
				), false
			}

			target := nearby[targetIndex]

			response, ok := player.GetInventory().UseItemOnEntity(
				itemIndex,
				target,
				context,
			)
			return response, ok
		} else if _, ok := item.(game.ItemUseDirection); ok {
			fmt.Printf(
				game.ColTooltip("Choose a direction to use: %v\n"),
				game.ListDirections(),
			)
			var directionInput = game.Input()

			var dx, dy, ok = game.DirToDelta(directionInput)
			if !ok {
				return game.SnipInvalidDirection(directionInput), false
			}

			response, ok := player.GetInventory().UseItemInDirection(
				itemIndex,
				dx,
				dy,
				directionInput,
				context,
			)
			return response, ok
		} else {
			return game.SnipItemCannotBeUsedBy(player, item), false
		}
	},
	"examine": func(player *Player, context *game.Context) (string, bool) {
		neighbors := context.World.GetOccupantsSameTile(player)
		if len(neighbors) == 0 {
			return "Nothing nearby to examine.\n", false
		}

		fmt.Printf(
			game.ColTooltip("Select what you want to examine:\n%v\n"),
			game.ListOrderedEntities(neighbors),
		)
		input := game.Input()

		targetIndex, indexErr := strconv.Atoi(input)

		if indexErr != nil || targetIndex < 0 || targetIndex >= len(neighbors) {
			return fmt.Sprintf(
				"Invalid selection. (%v)\n",
				game.ColTooltip(input),
			), false
		}

		target := neighbors[targetIndex]

		response := target.GetDesc(player)
		return response, false
	},
	"move": func(player *Player, context *game.Context) (string, bool) {
		fmt.Printf(
			game.ColTooltip("Choose a direction to move: %v\n"),
			game.ListDirections(),
		)
		input := game.Input()

		dx, dy, valid := game.DirToDelta(input)
		if !valid {
			return game.SnipInvalidDirection(input), false
		}

		response, endTurn := player.GetMovement().Move(dx, dy, &context.World)

		return response, endTurn
	},
	"look": func(player *Player, context *game.Context) (string, bool) {
		if player.looksThisTurn >= maxLooksPerTurnPlayer {
			return fmt.Sprintf(
				"%v's eyes are already strained from looking around this turn.\n",
				player.GetName(),
			), false
		}

		fmt.Printf(
			game.ColTooltip("Choose a direction to look: %v\n"),
			game.ListDirections(),
		)
		input := game.Input()

		dx, dy, valid := game.DirToDelta(input)
		if !valid {
			return game.SnipInvalidDirection(input), false
		}

		px, py, ok := context.World.GetEntityPos(player)
		if !ok {
			return fmt.Sprintf(
				"%v is lost.",
				player.GetName(),
			), false
		}

		player.looksThisTurn++

		for dist := 1; dist < context.World.Size; dist++ {
			tx, ty := px+(dx*dist), py+(dy*dist)

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
