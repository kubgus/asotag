package content

import (
	"asotag/game"
	"fmt"
	"strconv"
	"strings"
)

type HasDefense interface {
	GetDefense() *DefenseModule
}

type DefenseModule struct {
	entity   game.Entity

	// 0.0-1.0 = % of damage applied
	Defense 	int
	TempDefense int
}

func (d *DefenseModule) Init(e game.Entity) {
	d.entity = e
}

func (d *DefenseModule) EffectiveDefense() int {
	return d.Defense + d.TempDefense
}

func (d *DefenseModule) Get() string {
	return fmt.Sprintf(
		"%s now has %d%% damage reduction.\n",
		d.entity.GetName(),
		100-d.EffectiveDefense()*100,
	)
}

func (d *DefenseModule) Apply(amount int, persist bool) string {
	var response strings.Builder

	fmt.Fprintf(&response, "%s's defense", d.entity.GetName())

	if amount > 0 {
		response.WriteString(" increases")
	} else if amount < 0 {
		response.WriteString(" decreases")
	}

	fmt.Fprintf(&response, " by %s",
	game.ColItem(strconv.Itoa(abs(amount)*100))+"%")

	if persist {
		applyDefense(&d.Defense, amount)
		response.WriteString(".\n")
	} else {
		applyDefense(&d.TempDefense, amount)
		response.WriteString(" for the next attack.\n")
	}

	response.WriteString(d.Get())

	return response.String()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func applyDefense(v *int, amount int) {
	if amount > 0 {
		*v *= amount
	} else if amount < 0 {
		*v /= -amount
	}
}
