package content

type Material int

// Ordered by increasing strength/durability
const (
	MaterialWood Material = iota
	MaterialStone
	MaterialIron
	MaterialGold
)

func (m Material) String() string {
	switch m {
	case MaterialWood:
		return "Wood"
	case MaterialStone:
		return "Stone"
	case MaterialIron:
		return "Iron"
	case MaterialGold:
		return "Gold"
	default:
		return "Mysterious Material"
	}
}

func (m Material) AdjectiveString() string {
	switch m {
	case MaterialWood:
		return "Wooden"
	case MaterialStone:
		return "Stone"
	case MaterialIron:
		return "Iron"
	case MaterialGold:
		return "Golden"
	default:
		return "Mysterious"
	}
}
