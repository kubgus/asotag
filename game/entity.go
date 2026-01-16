package game

import (
	"fmt"
	"strconv"
	"text-adventure-game/utils"
)

type Entity interface {
	GetName() string

	GetHealth() int
	GetHealthString(includeWordHealth bool) string
	AddHealth(amount int) (string, bool)

	Examine(user Entity) string
	Loot(user Entity) []Item

	Reset(context *Context) // Helper to reset any per-turn state
	Move(context *Context) (string, bool)
}

func ListEntities(entities []Entity) string {
	if len(entities) == 0 {
		return FmtTooltip("None")
	}

	var keys []string
	for _, entity := range entities {
		keys = append(keys, fmt.Sprintf(
			"%v(%v)",
			entity.GetName(),
			entity.GetHealthString(false),
			))
	}
	return utils.JoinWithLast(keys, ", ", " and ")
}

func ListOrderedEntities(items []Entity) string {
	if len(items) == 0 {
		return FmtTooltip("None")
	}

	return utils.JoinWithMapFunc(items, "\n", func(i int, entity Entity) string {
		return fmt.Sprintf(
			"%v: %v(%v)",
			FmtAction(strconv.Itoa(i)),
			entity.GetName(),
			entity.GetHealthString(false),
			)
	})
}
