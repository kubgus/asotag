package game

type Item interface {
	GetName() string
	GetDesc() string
	Use(user, target Entity) string
}
