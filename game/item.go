package game

import (
	"fmt"
	"strconv"
	"text-adventure-game/utils"
)

type Item interface {
	GetName() string
	GetDesc() string

	Use(user, target Entity) (response string, ok bool, consume bool)
}

func ListItems(items []Item) string {
	if len(items) == 0 {
		return FmtTooltip("None")
	}

	var keys []string
	for _, item := range items {
		keys = append(keys, item.GetName())
	}
	return utils.JoinWithLast(keys, ", ", " and ")
}

func ListOrderedItemsWithMapFunc(items []Item, f func(int, Item) string) string {
	if len(items) == 0 {
		return FmtTooltip("None")
	}

	return utils.JoinWithMapFunc(items, "\n", func(i int, item Item) string {
		return fmt.Sprintf(
			"%v: %v %v %v",
			FmtAction(strconv.Itoa(i)),
			item.GetName(),
			FmtTooltip("-"),
			f(i, item),
			)
	})
}

func ListOrderedItems(items []Item) string {
	return ListOrderedItemsWithMapFunc(items, func(i int, item Item) string {
		return item.GetDesc()
	})
}

func ItemsMatchUnordered(a, b []Item) bool {
	if len(a) != len(b) {
		return false
	}

	used := make([]bool, len(b))
	for _, itemA := range a {
		found := false
		for j, itemB := range b {
			if used[j] {
				continue
			}
			if itemA.GetName() == itemB.GetName() {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
