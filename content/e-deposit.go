package content

import (
	"fmt"
	"text-adventure-game/game"
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

func (d *Deposit) GetName() string {
	return game.ColLocation(d.Name)
}

func (d *Deposit) GetHealth() int {
	return 0
}

func (d *Deposit) GetStatus() string {
	if d.Amount > 0 {
		return game.ColHealth("Full")
	}
	return game.ColHealth("Empty")
}

func (d *Deposit) AddHealth(amount int) (string, bool) {
	return fmt.Sprintf(
		"%v seems completely unaffected.",
		d.GetName(),
	), true
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

func (d *Deposit) BeforeTurn(context *game.Context) { }

func (d *Deposit) OnTurn(context *game.Context) (string, bool) {
	return fmt.Sprintf(
		"%v can't be bothered.",
		d.GetName(),
	), false
}
