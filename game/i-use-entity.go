package game

type ItemUseEntity interface {
	UseOnEntity(
		user, target Entity,
		context *Context,
	) (response string, ok, consume bool)
}
