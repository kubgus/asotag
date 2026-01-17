package game

import (
	"asotag/utils"
	"fmt"
	"strconv"
)

type Entity interface {
	GetName() string
	GetStatus() string
	// Returns a descripton when examined by the user.
	GetDesc(user Entity) string
}

func ListEntities(entities []Entity) string {
	if len(entities) == 0 {
		return ColTooltip("None")
	}

	var keys []string
	for _, entity := range entities {
		keys = append(keys, fmt.Sprintf(
			"%v(%v)",
			entity.GetName(),
			entity.GetStatus(),
		))
	}
	return utils.JoinWithLast(keys, ", ", " and ")
}

func ListOrderedEntities(items []Entity) string {
	if len(items) == 0 {
		return ColTooltip("None")
	}

	return utils.JoinWithMapFunc(items, "\n", func(i int, entity Entity) string {
		return fmt.Sprintf(
			"%v: %v(%v)",
			ColAction(strconv.Itoa(i)),
			entity.GetName(),
			entity.GetStatus(),
		)
	})
}
