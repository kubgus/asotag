package game

import "asotag/utils"

type EntityLoot interface {
	GetLoot(user Entity) []Item
}

func GetRandomLoot(lootTable map[Item]int, count int) []Item {
	var dropped []Item
	totalWeight := 0
	for _, weight := range lootTable {
		totalWeight += weight
	}

	for range count {
		randWeight := utils.RandIntInRange(1, totalWeight)
		currentWeight := 0

		for item, weight := range lootTable {
			currentWeight += weight
			if randWeight <= currentWeight {
				dropped = append(dropped, item)
				break
			}
		}
	}

	return dropped
}
