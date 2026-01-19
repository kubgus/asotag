package content

func NewHealingPotion(prefix string, magnitude int) *HealingPotion {
	return &HealingPotion{
		Prefix:    prefix,
		Magnitude: magnitude,
	}
}

func NewHealingPotionMinor() *HealingPotion {
	return NewHealingPotion("Minor", 20)
}

func NewHealingPotionMajor() *HealingPotion {
	return NewHealingPotion("Major", 50)
}

func NewHealingPotionSuperior() *HealingPotion {
	return NewHealingPotion("Superior", 100)
}
