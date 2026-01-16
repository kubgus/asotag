package main

import (
	"fmt"
	"math/rand/v2"
	"os/user"
	"text-adventure-game/entities"
	"text-adventure-game/game"
	"text-adventure-game/items"
	"text-adventure-game/utils"
	"time"
)

const (
	worldSize = 10

	enemyCountGoblin = 5
)

var (
	fmtWelcome = utils.NewColor(utils.ColorFgBold, utils.ColorFgPurple)
)

func main() {
	fmt.Print(fmtWelcome(
		"Welcome to the Text Adventure Game!\n" +
		"Defeat all enemies to win, but beware of your health!\n\n",
		))

	context := game.Context{
		World: *game.NewWorld(worldSize),
	}

	player := addPlayer(&context)
	enemies := addEnemies(&context)

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
		fmt.Print(fmtWelcome(
			"\nWith the last enemy defeated, you stand victorious!\n",
			))
	} else if lose {
		fmt.Print(fmtWelcome(
			"\nYou have fallen in battle. Your adventure ends here.\n",
			))
	}
	time.Sleep(3 * time.Second)
	fmt.Print(fmtWelcome("Thank you for playing! Feel free to try again.\n"))
	time.Sleep(2 * time.Second)
}

func addPlayer(context *game.Context) *entities.Player {
	currentUser, err := user.Current()
	playerName := ""
	if err == nil {
		playerName = currentUser.Username
	}

	player := entities.NewPlayer(playerName)
	context.World.Add(
		player,
		worldSize/2,
		worldSize/2,
		true,
		)

	player.Inventory = append(player.Inventory, &items.Sword{
		Name: "Iron Sword",
		Damage: 100,
	})

	return player
}

func addEnemies(context *game.Context) []game.Entity {
	enemies := []game.Entity{}

	for range enemyCountGoblin {
		goblin := entities.NewGoblin()
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
