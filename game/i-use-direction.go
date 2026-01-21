package game

type ItemUseDirection interface {
	UseInDirection(
		user Entity,
		dx, dy int,
		direction string,
		context *Context,
	) (response string, ok, consume bool)
}
