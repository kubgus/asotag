package game

import (
	"asotag/utils"
	"fmt"
	"time"
)

type Context struct {
	World World

	CheatRevealMap bool
}

func (c *Context) ExecuteRound() {
	for _, entity := range c.World.EntityOrder {
		if entityHealth, ok := entity.(HasHealth); ok && entityHealth.GetHealth().CurrentHealth <= 0 {
			c.World.Remove(entity, true)
			continue
		}

		if entityActive, ok := entity.(EntityActive); ok {
			entityActive.BeforeTurn(c)

			time.Sleep(time.Duration(500+utils.RandIntInRange(0, 500)) * time.Millisecond)

			fmt.Printf("%v %v's turn %v\n\n",
				ColTooltip("======>"),
				entity.GetName(),
				ColTooltip("<======"),
			)

			for {
				response, endTurn := entityActive.OnTurn(c)

				fmt.Println(response)

				if entityHealth, ok := entity.(HasHealth); ok && entityHealth.GetHealth().CurrentHealth <= 0 {
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
