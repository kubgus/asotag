package content

import (
	"fmt"
	"math/rand/v2"
	"text-adventure-game/game"
)

const (
	defaultPickaxeMinDamage = 5
	defaultPickaxeMaxDamage = 10
)

type Pickaxe struct {
	Material Material
}

func NewPickaxe(material Material) *Pickaxe {
	return &Pickaxe{
		Material: material,
	}
}

func (p *Pickaxe) GetName() string {
	return game.FmtItem(p.Material.String() + " Pickaxe")
}

func (p *Pickaxe) GetDesc() string {
	return fmt.Sprintf(
		"Can mine %v.",
		game.FmtItem(p.MaxMineable().String()),
	)
}

func (p *Pickaxe) Use(user, target game.Entity) (string, bool, bool) {
	deposit, ok := target.(*Deposit)
	if !ok {
		damage := rand.IntN(defaultPickaxeMaxDamage-defaultPickaxeMinDamage+1) + defaultPickaxeMinDamage

		response, alive := target.AddHealth(-damage)

		if !alive {
			loot := target.Loot(user)
			if len(loot) > 0 {
				player, ok := user.(*Player)
				if ok {
					player.Inventory = append(player.Inventory, loot...)
				}

				response += fmt.Sprintf(
					"\nObtained items: %v",
					game.ListItems(loot),
					)
			}
		}

		return fmt.Sprintf(
			"%v swings %v at %v and deals %d damage.\n",
			user.GetName(),
			p.GetName(),
			target.GetName(),
			damage,
			), true, false
	}

	if deposit.Type > p.MaxMineable() {
		return fmt.Sprintf(
			"%v tries to mine %v with %v, but it's too hard to break!\n",
			user.GetName(),
			deposit.GetName(),
			p.GetName(),
			), false, false
	}

	loot := deposit.Loot(user)

	if len(loot) == 0 {
		return fmt.Sprintf(
			"%v tries to mine %v with %v, but there's nothing left to mine!\n",
			user.GetName(),
			deposit.GetName(),
			p.GetName(),
			), false, false
	}

	player, ok := user.(*Player)
	if ok {
		player.Inventory = append(player.Inventory, loot...)
	}

	return fmt.Sprintf(
		"%v uses %v to mine %v and obtains %v.\n",
		user.GetName(),
		p.GetName(),
		deposit.GetName(),
		game.ListItems(loot),
		), true, false
}


func (p *Pickaxe) MaxMineable() Material {
	return p.Material + 1
}
