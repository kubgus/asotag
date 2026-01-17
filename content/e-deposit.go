package content

import (
	"fmt"
	"text-adventure-game/game"
	"text-adventure-game/utils"
)

type Deposit struct {
	Name string
	Type Material
	Amount int
}

func NewDeposit(name string, depositType Material, amount int) *Deposit {
	return &Deposit{
		Name: name,
		Type: depositType,
		Amount: amount,
	}
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
	if d.Amount > 0 {
		return game.ColHealth("Full")
	}
	return game.ColHealth("Empty")
}

func (d *Deposit) GetDesc(user game.Entity) string {
	return fmt.Sprintf(
		"%v contains a deposit of %d units of %v.\n",
		d.GetName(),
		d.Amount,
		game.ColItem(d.Type.String()),
	)
}

func (d *Deposit) GetLoot(user game.Entity) []game.Item {
	result := make([]game.Item, 0, d.Amount)
	for i := 0; i < d.Amount; i++ {
		result = append(result, NewResource(d.Type))
	}
	d.Amount = 0
	return result
}
