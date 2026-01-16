package game

import "text-adventure-game/utils"

var (
	FmtReset = utils.NewColor(utils.ColorReset)
	FmtHealth = utils.NewColor(utils.ColorFgBold, utils.ColorFgRed)
	FmtDamage = utils.NewColor(utils.ColorFgRed)
	FmtItem = utils.NewColor(utils.ColorFgBold, utils.ColorFgBlue)
	FmtTooltip = utils.NewColor(utils.ColorFgDim, utils.ColorFgWhite)
	FmtHero = utils.NewColor(utils.ColorFgBold, utils.ColorFgCyan)
	FmtAction = utils.NewColor(utils.ColorFgBold, utils.ColorFgYellow)
	FmtSystem = utils.NewColor(utils.ColorFgBold, utils.ColorFgPurple)
	FmtLocation = utils.NewColor(utils.ColorFgBold, utils.ColorFgItalic, utils.ColorFgWhite)
)
