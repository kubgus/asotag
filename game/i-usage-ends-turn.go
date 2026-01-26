package game

type ItemUsageEndsTurn interface {
	EndTurnOnUse() bool
}
