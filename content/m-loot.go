package content

import (
	"asotag/game"
	"asotag/utils"
)

type HasLoot interface {
	GetLoot() *LootModule
}

type LootModule struct {
	entity game.Entity

	LootTable           map[game.Item]int
	AmountTable         map[int]int
	DropInventoryChance float64
	LootLimit           int
}

func (lm *LootModule) Init(e game.Entity) {
	lm.entity = e
}

func (lm *LootModule) Drop() []game.Item {
	if lm.LootLimit > 1 {
		lm.LootLimit -= 1
	} else if lm.LootLimit == 1 {
		lm.LootLimit -= 2
	} else if lm.LootLimit < 0 {
		return []game.Item{}
	}
	// lm.LootLimit == 0 -> unlimited loot

	amount := randomWeightTableChoice(lm.AmountTable)

	loot := make([]game.Item, 0, amount)
	for range amount {
		item := randomWeightTableChoice(lm.LootTable)
		if itemClone, ok := utils.CloneInterface(item); ok {
			loot = append(loot, itemClone)
		} else {
			panic("failed to clone loot item")
		}
	}

	if utils.RandProbability(lm.DropInventoryChance) {
		if entityInv, ok := lm.entity.(HasIntentory); ok {
			loot = append(loot, entityInv.GetInventory().Items...)
			entityInv.GetInventory().Items = []game.Item{}
		}
	}

	return loot
}

func randomWeightTableChoice[T comparable](table map[T]int) T {
	totalWeight := 0
	for _, weight := range table {
		totalWeight += weight
	}

	randWeight := utils.RandIntInRange(1, totalWeight)
	currentWeight := 0

	for item, weight := range table {
		currentWeight += weight
		if randWeight <= currentWeight {
			return item
		}
	}

	var zero T
	return zero
}
