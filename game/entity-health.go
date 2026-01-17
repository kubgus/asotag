package game

import "fmt"

type EntityHealth interface {
	GetHealth() int
	AddHealth(amount int) (response string, alive bool)
}

func GetHealthStatusResponse[T interface{ Entity; EntityHealth }](entity T) string {
	return fmt.Sprintf(
		"%v is now at %v.",
		entity.GetName(),
		FormatHealth(entity.GetHealth(), true),
		)
}
