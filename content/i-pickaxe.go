package content

import (
	"fmt"
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
	return game.ColItem(p.Material.String() + " Pickaxe")
}

func (p *Pickaxe) GetDesc() string {
	return fmt.Sprintf(
		"Can mine %v and softer materials.",
		game.ColItem(p.MaxMineable().String()),
	)
}

func (p *Pickaxe) Use(user, target game.Entity) (string, bool, bool) {
	deposit, ok := target.(*Deposit)
	if !ok {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}
	if deposit.Type > p.MaxMineable() || deposit.Amount == 0 {
		return game.SnipCannotUseItemOn(user, target, p), false, false
	}

	loot := deposit.GetLoot(user)

	if player, ok := user.(*Player); ok {
		player.CollectLoot(loot)
	} else {
		return game.SnipItemCannotBeUsedBy(user, p), false, false
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
