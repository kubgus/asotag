package content

import (
	"asotag/game"
	"fmt"
	"sort"
)

type HasIntentory interface {
	GetInventory() *InventoryModule
}

type InventoryModule struct {
	entity  game.Entity
	context *game.Context

	Items []game.Item
}

func (im *InventoryModule) Init(entity game.Entity) {
	im.entity = entity
}

func (im *InventoryModule) HasIndex(index int) bool {
	return index >= 0 && index < len(im.Items)
}

func (im *InventoryModule) AddItems(items []game.Item) string {
	im.Items = append(im.Items, items...)
	return fmt.Sprintf(
		"%s picks up %s.\n",
		im.entity.GetName(),
		game.ListItems(items),
	)
}

func (im *InventoryModule) RemoveItems(indexes []int) string {
	items := make([]game.Item, 0, len(indexes))

	for _, index := range indexes {
		if !im.HasIndex(index) {
			return game.SnipInvalidItemIndex(index)
		}
		items = append(items, im.Items[index])
	}
	// Sort bundleIdxs in descending order
	// to safely remove items from inventory
	// without affecting the indexes of yet-to-be-removed items
	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i] > indexes[j]
	})
	// Remove bundled items from inventory
	// after collecting them to bundle
	// to avoid index shifting issues
	for _, index := range indexes {
		im.Items = append(
			im.Items[:index],
			im.Items[index+1:]...,
		)
	}

	for _, index := range indexes {
		if !im.HasIndex(index) {
			return game.SnipInvalidItemIndex(index)
		}
		items = append(items, im.Items[index])
	}

	return fmt.Sprintf(
		"%s removes %s from their inventory.\n",
		im.entity.GetName(),
		game.ListItems(items),
	)
}

func (im *InventoryModule) UseItemOnEntity(index int, target game.Entity) (string, bool) {
	if !im.HasIndex(index) {
		return game.SnipInvalidItemIndex(index), false
	}

	item := im.Items[index]
	if correctItem, ok := item.(game.ItemUseEntity); ok {
		response, ok, consume := correctItem.UseOnEntity(
			im.entity,
			target,
			im.context,
		)

		var removeResponse string
		if consume {
			removeResponse = im.RemoveItems([]int{index})
		}

		return fmt.Sprintf(
			"%s%s",
			response,
			removeResponse,
		), ok
	}

	return fmt.Sprintf(
		"%s has to use %s in a direction, not on %s.\n",
		im.entity.GetName(),
		item.GetName(),
		target.GetName(),
	), false
}

func (im *InventoryModule) UseItemInDirection(index int, dx, dy int, direction string) (string, bool) {
	if !im.HasIndex(index) {
		return game.SnipInvalidItemIndex(index), false
	}

	item := im.Items[index]
	if correctItem, ok := item.(game.ItemUseDirection); ok {
		response, ok, consume := correctItem.UseInDirection(
			im.entity,
			dx,
			dy,
			direction,
			im.context,
		)

		var removeResponse string
		if consume {
			removeResponse = im.RemoveItems([]int{index})
		}

		return fmt.Sprintf(
			"%s%s",
			response,
			removeResponse,
		), ok
	}

	return fmt.Sprintf(
		"%s has to use %s on an entity, not in a direction.\n",
		im.entity.GetName(),
		item.GetName(),
	), false
}

func (im *InventoryModule) FindItem(item game.Item) int {
	for i, invItem := range im.Items {
		if invItem == item {
			return i
		}
	}
	return -1
}
