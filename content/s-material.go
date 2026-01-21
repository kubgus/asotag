package content

type Material int

// Ordered by increasing strength/durability
const (
	MaterialVoid Material = iota
	MaterialWood
	MaterialStone
	MaterialIron
	MaterialGold
)

func (m Material) String() string {
	switch m {
	case MaterialVoid:
		return "Void"
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
