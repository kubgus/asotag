package game

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Context struct {
	World world
}

func (c *Context) ExecuteRound() {
	for _, entity := range c.World.MoveOrder {
		if entity.GetHealth() > 0 {
			entity.Reset(c)

			time.Sleep(time.Duration(1000 + rand.IntN(700)) * time.Millisecond)

			fmt.Printf("%v %v's turn %v\n\n",
				FmtTooltip("======>"),
				entity.GetName(),
				FmtTooltip("<======"),
				)
		} else {
			c.World.Remove(entity, true)
			continue
		}

		for {
			result, endTurn := entity.Move(c)

			fmt.Println(result)

			if entity.GetHealth() <= 0 {
				c.World.Remove(entity, true)
				break
			}

			if endTurn {
				break
			}
		}

		//c.World.debugPrint(entity)
	}
}
