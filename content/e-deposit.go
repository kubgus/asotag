package content

import (
	"asotag/game"
	"fmt"
)

type Deposit struct {
	Name     string
	Material Material // helper

	loot LootModule
}

func NewDeposit(name string, material Material) *Deposit {
	deposit := Deposit{
		Name:     name,
		Material: material,

		loot: LootModule{
			LootTable: map[game.Item]int{
				NewResource(material): 100,
			},
			AmountTable: map[int]int{
				2: 5,
				3: 15,
				4: 35,
				5: 35,
				6: 10,
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

func NewDepositTree() *Deposit {
	return NewDeposit("Tree", MaterialWood)
}

func NewDepositRock() *Deposit {
	return NewDeposit("Rock", MaterialStone)
}

func NewDepositIronVein() *Deposit {
	return NewDeposit("Iron Vein", MaterialIron)
}

func NewDepositGoldVein() *Deposit {
	return NewDeposit("Gold Vein", MaterialGold)
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
