package content

func NewSword(name string, minDamage, maxDamage int) *Sword {
	return &Sword{
		Name:      name,
		MinDamage: minDamage,
		MaxDamage: maxDamage,
	}
}

func NewSwordWooden() *Sword {
	return NewSword("Wooden Sword", 5, 10)
}

func NewSwordStone() *Sword {
	return NewSword("Stone Sword", 9, 16)
}

func NewSwordIron() *Sword {
	return NewSword("Iron Sword", 12, 20)
}

func NewSwordGold() *Sword {
	return NewSword("Gold Sword", 17, 35)
}
