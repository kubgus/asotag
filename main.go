package main

import (
	"asotag/content"
	"asotag/game"
	"fmt"
	"math/rand/v2"
	"os/user"
	"strings"
	"time"
)

const (
	worldSize = 15

	enemyCountGoblin = 7

	locationCountChest        = 10
	locationCountWorkbench    = 5
	locationCountDepositWood  = 20
	locationCountDepositStone = 10
	locationCountDepositIron  = 7
	locationCountDepositGold  = 4
)

func main() {
	fmt.Print(game.ColSystem("Welcome to ASOTAG!\n"))
	fmt.Print(game.ColSystem("Defeat all enemies to win, but beware of your health!\n\n\n"))

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
		if player.GetHealth().CurrentHealth <= 0 {
			lose = true
		}

		// Win by all enemies defeated
		allEnemiesDefeated := true
		for _, enemy := range enemies {
			if enemyHealth, ok := enemy.(game.HasHealth); ok {
				if enemyHealth.GetHealth().CurrentHealth > 0 {
					allEnemiesDefeated = false
					break
				}
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
		fmt.Print(game.ColSystem(
			"\nWith the last enemy defeated, you stand victorious!\n",
		))
	} else if lose {
		fmt.Print(game.ColSystem(
			"\nYou have fallen in battle. Your adventure ends here.\n",
		))
	}
	time.Sleep(3 * time.Second)
	fmt.Print(game.ColSystem("Thank you for playing! Feel free to try again.\n"))
	time.Sleep(4 * time.Second)

	fmt.Scan()
}

func addPlayer(context *game.Context) *content.Player {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error retrieving current user:", err)
		return content.NewPlayer("Hero")
	}
	userName := currentUser.Username
	lastIdx := strings.LastIndexAny(userName, "/\\@")
	if lastIdx != -1 {
		userName = userName[lastIdx+1:]
	}

	worldX := worldSize / 2
	worldY := worldSize / 2

	player := content.NewPlayer(userName)
	context.World.Add(
		player,
		worldX,
		worldY,
		true,
	)

	player.GetInventory().AddItems([]game.Item{
		content.NewPickaxeHand(),
		content.NewSwordWooden(),
		content.NewSpearWooden(),
		content.NewHealingPotionMinor(),
	})

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
			true,
		)
		locations = append(locations, chest)
	}

	for range locationCountWorkbench {
		workbench := content.NewWorkbench()
		context.World.Add(
			workbench,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			true,
		)
		locations = append(locations, workbench)
	}

	for range locationCountDepositWood {
		woodDeposit := content.NewDepositTree()
		context.World.Add(
			woodDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			true,
		)
		locations = append(locations, woodDeposit)
	}

	for range locationCountDepositStone {
		stoneDeposit := content.NewDepositRock()
		context.World.Add(
			stoneDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			true,
		)
		locations = append(locations, stoneDeposit)
	}

	for range locationCountDepositIron {
		ironDeposit := content.NewDepositIronVein()
		context.World.Add(
			ironDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			true,
		)
		locations = append(locations, ironDeposit)
	}

	for range locationCountDepositGold {
		goldDeposit := content.NewDepositGoldVein()
		context.World.Add(
			goldDeposit,
			rand.IntN(worldSize),
			rand.IntN(worldSize),
			true,
		)
		locations = append(locations, goldDeposit)
	}

	return locations
}
