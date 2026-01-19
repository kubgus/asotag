package content

func NewPickaxe(material Material) *Pickaxe {
	return &Pickaxe{
		Material: material,
	}
}
