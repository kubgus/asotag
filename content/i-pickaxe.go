package content

import (
	"asotag/game"
	"fmt"
)

type Pickaxe struct {
	Material Material
}

func (p *Pickaxe) GetName() string {
	return game.ColItem(p.Material.String() + " Pickaxe")
}

func (p *Pickaxe) GetDesc() string {
	return fmt.Sprintf(
		"Can mine %v and softer materials.",
		game.ColItem(p.MaxMineable().String()),
	)
}

func (p *Pickaxe) UseOnEntity(user, target game.Entity, _ *game.Context) (string, bool, bool) {
	var loot []game.Item

	if deposit, ok := target.(*Deposit); ok {
		if deposit.Type > p.MaxMineable() || deposit.Amount == 0 {
			return game.SnipCannotUseItemOn(user, target, p), false, false
		}
		loot = deposit.GetLoot(user)
	} else if chest, ok := target.(*Chest); ok {
		if !chest.IsUnlocked {
			return game.SnipCannotUseItemOn(user, target, p), false, false
		}
		loot = chest.GetLoot(user)
	} else {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}

	if player, ok := user.(*Player); ok {
		player.CollectLoot(loot)
	} else {
		return game.SnipItemCannotBeUsedBy(user, p), false, false
	}

	return fmt.Sprintf(
		"%v uses %v on %v and obtains %v.\n",
		user.GetName(),
		p.GetName(),
		target.GetName(),
		game.ListItems(loot),
	), true, false
}

func (p *Pickaxe) MaxMineable() Material {
	return p.Material + 1
}
