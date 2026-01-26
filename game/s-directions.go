package game

import "asotag/utils"

var Directions = map[string][2]int{
	"north":     {0, -1},
	"northeast": {1, -1},
	"east":      {1, 0},
	"southeast": {1, 1},
	"south":     {0, 1},
	"southwest": {-1, 1},
	"west":      {-1, 0},
	"northwest": {-1, -1},
}

var DirectionShortcuts = map[string][2]int{
	"n": {0, -1},
	"ne": {1, -1},
	"e": {1, 0},
	"se": {1, 1},
	"s": {0, 1},
	"sw": {-1, 1},
	"w": {-1, 0},
	"nw": {-1, -1},
}

func ListDirections() string {
	var keys []string
	for k := range Directions {
		keys = append(keys, k)
	}
	return utils.JoinAdvanced(keys, ", ", " or ", func(_ int, direction string) string {
		return ColAction(direction)
	})
}

func DirToDelta(direction string) (int, int, bool) {
	var delta [2]int
	var ok bool
	for dir, dlt := range Directions {
		if dir == direction {
			delta = dlt
			ok = true
			break
		}
	}
	for dir, dlt := range DirectionShortcuts {
		if dir == direction {
			delta = dlt
			ok = true
			break
		}
	}

	if !ok {
		return 0, 0, false
	}

	return delta[0], delta[1], true
}

func DeltaToDir(dx, dy int) (string, bool) {
	for dir, delta := range Directions {
		if delta[0] == dx && delta[1] == dy {
			return dir, true
		}
	}

	return "in an unknown direction", false
}
