package content

import (
	"asotag/game"
	"fmt"
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
		if deposit.Material > p.MaxMineable() {
			return game.SnipCannotUseItemOn(user, target, p), false, false
		}
		loot = deposit.GetLoot().Drop()
	} else if chest, ok := target.(*Chest); ok {
		if !chest.IsUnlocked {
			return game.SnipCannotUseItemOn(user, target, p), false, false
		}
		loot = chest.GetLoot().Drop()
	} else {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}

	if player, ok := user.(*Player); ok {
		response := player.GetInventory().AddItems(loot)
		return response, len(loot) > 0, false
	}

	return game.SnipItemCannotBeUsedBy(user, p), false, false
}

func (p *Pickaxe) MaxMineable() Material {
	return p.Material + 1
}
