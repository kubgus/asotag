package game

import "text-adventure-game/utils"

var Directions = map[string][2]int{
	"north": {0, -1},
	"northeast": {1, -1},
	"east": {1, 0},
	"southeast": {1, 1},
	"south": {0, 1},
	"southwest": {-1, 1},
	"west": {-1, 0},
	"northwest": {-1, -1},
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

func DirToDelta(dir string) (int, int, bool) {
	delta, ok := Directions[dir]
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
