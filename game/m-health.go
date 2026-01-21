package game

import (
	"fmt"
	"strconv"
	"strings"
)

type HasHealth interface {
	GetHealth() *HealthModule
}

type HealthModule struct {
	entity Entity

	CurrentHealth int
	MaxHealth     int
}

func (hm *HealthModule) Init(entity Entity) {
	hm.entity = entity
}

func (hm *HealthModule) Get() string {
	if hm.MaxHealth > 0 {
		return fmt.Sprintf(
			"%s is now at %s/%s health.\n",
			hm.entity.GetName(),
			ColHealth(strconv.Itoa(hm.CurrentHealth)),
			ColHealth(strconv.Itoa(hm.MaxHealth)),
		)
	}

	return fmt.Sprintf(
		"%s is now at %s health.\n",
		hm.entity.GetName(),
		ColHealth(strconv.Itoa(hm.CurrentHealth)),
	)
}

func (hm *HealthModule) Change(amount int) string {
	var response strings.Builder

	effectiveAmount := amount

	if amount > 0 {
		fmt.Fprintf(&response, "%s gains %s health.\n",
			hm.entity.GetName(),
			ColHealth(strconv.Itoa(effectiveAmount)))
	} else if amount < 0 {
		//if defense, ok := hm.entity.(HasDefense); ok {
		//	totalDefense := defense.GetDefense().EffectiveDefense()
		//	effectiveAmount = -calculateEffectiveDamage(-amount, totalDefense)
		//}

		fmt.Fprintf(&response, "%s loses %s health.\n",
			hm.entity.GetName(),
			ColHealth(strconv.Itoa(-effectiveAmount)))
	}

	hm.CurrentHealth += effectiveAmount

	if hm.MaxHealth > 0 && hm.CurrentHealth > hm.MaxHealth {
		hm.CurrentHealth = hm.MaxHealth
	}

	response.WriteString(hm.Get())

	if hm.CurrentHealth < 0 {
		hm.CurrentHealth = 0

		fmt.Fprintf(&response, "%s has been knocked out.\n",
			hm.entity.GetName())
	}

	return response.String()
}

func calculateEffectiveDamage(damage int, totalDefense int) int {
	effectiveDamage := float64(damage) * (1.0 - float64(totalDefense)/100.0)
	if effectiveDamage < 0 {
		return 0
	}
	return int(effectiveDamage)
}
