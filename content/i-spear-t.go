package content

func NewSpear(name string, minDamage, maxDamage int) *Spear {
	return &Spear{
		Name:      name,
		MinDamage: minDamage,
		MaxDamage: maxDamage,
	}
}

func NewSpearWooden() *Spear {
	return NewSpear("Wooden Spear", 4, 8)
}

func NewSpearIron() *Spear {
	return NewSpear("Iron Spear", 11, 18)
}
