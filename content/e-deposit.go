package content

import (
	"asotag/game"
	"asotag/utils"
	"fmt"
)

type Deposit struct {
	Name     string
	Material Material // helper

	loot LootModule
}

func NewDeposit(name string, material Material, amount int) *Deposit {
	deposit := Deposit{
		Name:     name,
		Material: material,

		loot: LootModule{
			LootTable: map[game.Item]int{
				NewResource(material): 100,
			},
			AmountTable: map[int]int{
				amount: 100,
			},
			LootLimit: 1,
		},
	}

	return &deposit
}

func (d *Deposit) GetLoot() *LootModule {
	d.loot.Init(d)
	return &d.loot
}

func NewDepositTree(min, max int) *Deposit {
	return NewDeposit("Tree", MaterialWood, utils.RandIntInRange(min, max))
}

func NewDepositRock(min, max int) *Deposit {
	return NewDeposit("Rock", MaterialStone, utils.RandIntInRange(min, max))
}

func NewDepositIronVein(min, max int) *Deposit {
	return NewDeposit("Iron Vein", MaterialIron, utils.RandIntInRange(min, max))
}

func NewDepositGoldVein(min, max int) *Deposit {
	return NewDeposit("Gold Vein", MaterialGold, utils.RandIntInRange(min, max))
}

func (d *Deposit) GetName() string {
	return game.ColLocation(d.Name)
}

func (d *Deposit) GetStatus() string {
	if d.GetLoot().LootLimit < 0 {
		return game.ColHealth("Empty")
	}
	return game.ColHealth("Full")
}

func (d *Deposit) GetDesc(user game.Entity) string {
	return fmt.Sprintf(
		"%v provides %v when mined.\n",
		d.GetName(),
		game.ColItem(d.Material.String()),
	)
}
