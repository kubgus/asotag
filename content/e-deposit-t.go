package content

import "asotag/utils"

func NewDeposit(name string, depositType Material, amount int) *Deposit {
	return &Deposit{
		Name:   name,
		Type:   depositType,
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

