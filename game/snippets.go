package game

import "fmt"

func SnipCannotUseItemOn(user, target Entity, item Item) string {
	return fmt.Sprintf(
		"%v cannot use %v on %v.\n",
		user.GetName(),
		item.GetName(),
		target.GetName(),
		)
}

func SnipItemCannotBeUsedBy(user Entity, item Item) string {
	return fmt.Sprintf(
		"%v cannot use %v.\n",
		user.GetName(),
		item.GetName(),
		)
}
