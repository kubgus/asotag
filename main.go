package main

import (
	"fmt"
	"math/rand/v2"
	"os/user"
	"strings"
	"text-adventure-game/content"
	"text-adventure-game/game"
	"time"
)

const (
	worldSize = 10

	enemyCountGoblin = 7

	locationCountChest = 3
	locationCountWorkbench = 2
	locationCountDepositWood = 15
	locationCountDepositStone = 7
	locationCountDepositIron = 5
	locationCountDepositGold = 3

	depositCountPerResourceMin = 3
	depositCountPerResourceMax = 6
)

func main() {
	fmt.Print(game.FmtSystem(
		"Welcome to the Text Adventure Game!\n" +
		"Defeat all enemies to win, but beware of your health!\n\n\n",
		))

	context := game.Context{
		World: *game.NewWorld(worldSize),
	}

	player := addPlayer(&context)
	enemies := addEnemies(&context)
	addLocations(&context)

	win := false
	lose := false
	for {
		context.ExecuteRound()

		// Lose by player death
		if player.Health <= 0 {
			lose = true
		}

		// Win by all enemies defeated
		allEnemiesDefeated := true
		for _, enemy := range enemies {
			if enemy.GetHealth() > 0 {
				allEnemiesDefeated = false
				break
			}
		}
		if allEnemiesDefeated {
			win = true
		}

		if win || lose {
			break
		}
	}

	time.Sleep(2 * time.Second)
	if win {
		fmt.Print(game.FmtSystem(
			"\nWith the last enemy defeated, you stand victorious!\n",
			))
	} else if lose {
		fmt.Print(game.FmtSystem(
			"\nYou have fallen in battle. Your adventure ends here.\n",
			))
	}
	time.Sleep(3 * time.Second)
	fmt.Print(game.FmtSystem("Thank you for playing! Feel free to try again.\n"))
	time.Sleep(2 * time.Second)
}

func addPlayer(context *game.Context) *content.Player {
	currentUser, err := user.Current()
	playerName := ""
	if err == nil {
		userName := currentUser.Username
		lastIdx := strings.LastIndexAny(userName, "/\\@")
		if lastIdx != -1 {
			userName = userName[lastIdx+1:]
		}
		playerName = userName
	}

	worldX := worldSize/2
	worldY := worldSize/2

	player := content.NewPlayer(playerName)
	context.World.Add(
		player,
		worldX,
		worldY,
		true,
		)

	player.Inventory = []game.Item{
		content.NewSword("Stick", 4, 8),
		content.NewPickaxe(content.MaterialWood),
		content.NewHealingPotion("Minor", 20),
	}

	return player
}

func addEnemies(context *game.Context) []game.Entity {
	enemies := []game.Entity{}

	for range enemyCountGoblin {
		goblin := content.NewGoblin()
		context.World.Add(
			goblin,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			true,
			)
		enemies = append(enemies, goblin)
	}

	return enemies
}

func addLocations(context *game.Context) []game.Entity {
	locations := []game.Entity{}

	for range locationCountChest {
		chest := content.NewChest()
		context.World.Add(
			chest,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			false,
			)
		locations = append(locations, chest)
	}

	for range locationCountWorkbench {
		workbench := content.NewWorkbench()
		context.World.Add(
			workbench,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			false,
			)
		locations = append(locations, workbench)
	}

	for range locationCountDepositWood {
		woodDeposit := content.NewDeposit(
			"Tree",
			content.MaterialWood,
			randomDepositAmount(),
			)
		context.World.Add(
			woodDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			false,
			)
		locations = append(locations, woodDeposit)
	}

	for range locationCountDepositStone {
		stoneDeposit := content.NewDeposit(
			"Rock",
			content.MaterialStone,
			randomDepositAmount(),
			)
		context.World.Add(
			stoneDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			false,
			)
		locations = append(locations, stoneDeposit)
	}

	for range locationCountDepositIron {
		ironDeposit := content.NewDeposit(
			"Iron Vein",
			content.MaterialIron,
			randomDepositAmount(),
			)
		context.World.Add(
			ironDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			false,
			)
		locations = append(locations, ironDeposit)
	}

	for range locationCountDepositGold {
		goldDeposit := content.NewDeposit(
			"Gold Vein",
			content.MaterialGold,
			randomDepositAmount(),
			)
		context.World.Add(
			goldDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			false,
			)
		locations = append(locations, goldDeposit)
	}

	return locations
}

func randomDepositAmount() int {
	return rand.IntN(depositCountPerResourceMax - depositCountPerResourceMin + 1) + depositCountPerResourceMin
}
