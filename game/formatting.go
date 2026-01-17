package game

import (
	"asotag/utils"
	"fmt"
)

var (
	ColReset    = utils.NewColor(utils.ColorReset)
	ColHealth   = utils.NewColor(utils.ColorFgBold, utils.ColorFgRed)
	ColDamage   = utils.NewColor(utils.ColorFgRed)
	ColItem     = utils.NewColor(utils.ColorFgBold, utils.ColorFgBlue)
	ColTooltip  = utils.NewColor(utils.ColorFgDim, utils.ColorFgWhite)
	ColHero     = utils.NewColor(utils.ColorFgBold, utils.ColorFgCyan)
	ColAction   = utils.NewColor(utils.ColorFgBold, utils.ColorFgYellow)
	ColSystem   = utils.NewColor(utils.ColorFgBold, utils.ColorFgPurple)
	ColLocation = utils.NewColor(utils.ColorFgBold, utils.ColorFgItalic, utils.ColorFgWhite)
)

func FormatHealth(health int, extended bool) string {
	response := fmt.Sprintf("%d", health)
	if extended {
		response += " health"
	}
	return ColHealth(response)
}

func FormatDamage(damage int, extended bool) string {
	response := fmt.Sprintf("%d", damage)
	if extended {
		response += " damage"
	}
	return ColDamage(response)
}
