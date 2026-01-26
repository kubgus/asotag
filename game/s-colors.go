package game

import (
	"asotag/utils"
)

var (
	ColReset = utils.NewColor(utils.ColorReset)

	ColHealth = utils.NewColor(utils.ColorFgBold, utils.ColorFgRed)
	ColDamage = utils.NewColor(utils.ColorFgBrightRed)

	ColItem     = utils.NewColor(utils.ColorFgBold, utils.ColorFgBlue)
	ColLocation = utils.NewColor(utils.ColorFgBold, utils.ColorFgItalic, utils.ColorFgWhite)
	ColHero     = utils.NewColor(utils.ColorFgBold, utils.ColorFgCyan)

	ColAction  = utils.NewColor(utils.ColorFgBold, utils.ColorFgYellow)
	ColActionSec = utils.NewColor(utils.ColorFgBold, utils.ColorFgBrightPurple)
	ColTooltip = utils.NewColor(utils.ColorFgDim, utils.ColorFgWhite)
	ColSystem  = utils.NewColor(utils.ColorFgBold, utils.ColorFgBrightBlue)
)
