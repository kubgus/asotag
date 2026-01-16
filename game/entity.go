package game

type Entity interface {
	GetName() string

	GetHealth() int
	GetHealthString(includeWordHealth bool) string
	AddHealth(amount int) string

	Reset(context *Context) // Helper to reset any per-turn state
	Move(context *Context) (string, bool)
}
