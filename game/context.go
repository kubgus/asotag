package game

import (
	"fmt"
	"math"
	"time"
)

type Context struct {
	World world

	CheatRevealMap bool
}

func (c *Context) ExecuteRound() {
	for _, entity := range c.World.EntityOrder {
		if entityHealth, ok := entity.(EntityHealth); ok && entityHealth.GetHealth() <= 0 {
			c.World.Remove(entity, true)
			continue
		}

		if entityActive, ok := entity.(EntityActive); ok {
			entityActive.BeforeTurn(c)

			time.Sleep(time.Duration((math.Round(6000 / float64(len(c.World.EntityOrder))))) * time.Millisecond)

			fmt.Printf("%v %v's turn %v\n\n",
				ColTooltip("======>"),
				entity.GetName(),
				ColTooltip("<======"),
				)

			for {
				response, endTurn := entityActive.OnTurn(c)

				fmt.Println(response)

				if entityHealth, ok := entity.(EntityHealth); ok && entityHealth.GetHealth() <= 0 {
					c.World.Remove(entity, true)
					break
				}

				if endTurn {
					break
				}
			}

			if c.CheatRevealMap {
				c.World.debugPrint(entity)
			}
		}
	}
}
