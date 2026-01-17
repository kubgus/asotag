package game

import (
	"fmt"
	"text-adventure-game/utils"
)

type Point struct {
	X, Y int
}

type world struct {
	Positions map[Entity]Point
	Occupants map[Point][]Entity

	EntityOrder []Entity

	Size      int
}

func NewWorld(size int) *world {
	return &world{
		Positions: make(map[Entity]Point),
		Occupants: make(map[Point][]Entity),
		EntityOrder: make([]Entity, 0),
		Size:      size,
	}
}

func (w *world) Add(entity Entity, x, y int, addToOrder bool) bool {
	if x < 0 || x >= w.Size || y < 0 || y >= w.Size {
		return false
	}

	p := Point{x, y}
	w.Positions[entity] = p
	w.Occupants[p] = append(w.Occupants[p], entity)

	if addToOrder {
		w.EntityOrder = append(w.EntityOrder, entity)
	}

	return true
}

func (w *world) Remove(entity Entity, removeFromOrder bool) bool {
	pos, ok := w.Positions[entity]
	if !ok {
		return false
	}

	tileEntities := w.Occupants[pos]
	for i, e := range tileEntities {
		if e == entity {
			w.Occupants[pos] = append(tileEntities[:i], tileEntities[i+1:]...)
			break
		}
	}

	// Clean up empty keys to save memory
	if len(w.Occupants[pos]) == 0 {
		delete(w.Occupants, pos)
	}

	delete(w.Positions, entity)

	if removeFromOrder {
		for i, e := range w.EntityOrder {
			if e == entity {
				w.EntityOrder = append(w.EntityOrder[:i], w.EntityOrder[i+1:]...)
				break
			}
		}
	}

	return true
}

func (w *world) Move(entity Entity, x, y int) {
	w.Remove(entity, false)
	w.Add(entity, x, y, false)
}


func (w *world) MoveInDirection(entity Entity, dx, dy int) (string, bool) {
	pos, ok := w.Positions[entity]
	if !ok {
		return fmt.Sprintf("%v is not in the world.\n", entity.GetName()), false
	}

	newPos := Point{pos.X + dx, pos.Y + dy}

	if newPos.X < 0 || newPos.X >= w.Size || newPos.Y < 0 || newPos.Y >= w.Size {
		return fmt.Sprintf("%v tries to move, but hits a wall.\n", entity.GetName()), false
	}

	w.Remove(entity, false)
	w.Add(entity, newPos.X, newPos.Y, false)

	direction, _ := DeltaToDir(dx, dy)
	return fmt.Sprintf("%v moves %v.\n", entity.GetName(), ColAction(direction)), true
}

func (w *world) GetEntityPos(entity Entity) (int, int, bool) {
	p, ok := w.Positions[entity]
	return p.X, p.Y, ok
}

func (w *world) GetEntitiesAt(x, y int) []Entity {
	return w.Occupants[Point{x, y}]
}

func (w *world) GetOccupantsSameTile(entity Entity) []Entity {
	pos, ok := w.Positions[entity]
	if !ok {
		return nil
	}
	return w.Occupants[pos]
}

// Generated using Gemini
func (w *world) debugPrint(highlightEntity Entity) {
	// Header with X-axis coordinates
	fmt.Print("   ")
	for x := 0; x < w.Size; x++ {
		fmt.Printf(" %d ", x%10) // Print last digit of X
	}
	fmt.Println()

	// Top Border
	fmt.Print("  " + ColTooltip("┌"))
	for x := 0; x < w.Size; x++ {
		fmt.Print(ColTooltip("───"))
	}
	fmt.Println(ColTooltip("┐"))

	for y := 0; y < w.Size; y++ {
		// Y-axis coordinate
		fmt.Printf("%d " + ColTooltip("│"), y%10)

		for x := 0; x < w.Size; x++ {
			occupants := w.Occupants[Point{x, y}]
			if len(occupants) == 0 {
				fmt.Print(" ◻︎ ") // Empty space
			} else {
				first := occupants[0]
				if highlightEntity != nil {
					for _, e := range occupants {
						if e == highlightEntity {
							first = e
							break
						}
					}
				}

				color := utils.ColorFgRed
				if first == highlightEntity {
					color = utils.ColorBgYellow + utils.ColorFgBlack
				}

				symbol := color +
				utils.ColorFgBold +
				utils.StripANSI(first.GetName())[:1] +
				utils.ColorReset

				// If multiple entities are here, highlight it
				if len(occupants) >= 5 {
					fmt.Printf("{%s}", symbol)
				} else if len(occupants) == 4 {
					fmt.Printf("[%s]", symbol)
				} else if len(occupants) == 3 {
					fmt.Printf("<%s>", symbol)
				} else if len(occupants) == 2 {
					fmt.Printf("(%s)", symbol)
				} else {
					fmt.Printf(" %s ", symbol)
				}
			}
		}
		fmt.Println(ColTooltip("│"))
	}

	// Bottom Border
	fmt.Print("  " + ColTooltip("└"))
	for x := 0; x < w.Size; x++ {
		fmt.Print(ColTooltip("───"))
	}
	fmt.Println(ColTooltip("┘"))
}
