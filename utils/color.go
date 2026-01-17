package utils

import (
	"regexp"
	"strings"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func StripANSI(str string) string {
	return re.ReplaceAllString(str, "")
}

const (
	ColorReset = "\033[0m"

	ColorFgBold   = "\033[1m"
	ColorFgDim    = "\033[2m"
	ColorFgItalic = "\033[3m"
	ColorFgUnder  = "\033[4m"
	ColorFgBlink  = "\033[5m"
	ColorFgRev    = "\033[7m"
	ColorFgHidden = "\033[8m"

	ColorFgBlack  = "\033[30m"
	ColorFgRed    = "\033[31m"
	ColorFgGreen  = "\033[32m"
	ColorFgYellow = "\033[33m"
	ColorFgBlue   = "\033[34m"
	ColorFgPurple = "\033[35m"
	ColorFgCyan   = "\033[36m"
	ColorFgWhite  = "\033[37m"

	ColorBgRed    = "\033[41m"
	ColorBgGreen  = "\033[42m"
	ColorBgYellow = "\033[43m"
	ColorBgBlue   = "\033[44m"
	ColorBgPurple = "\033[45m"
	ColorBgCyan   = "\033[46m"
	ColorBgWhite  = "\033[47m"

	ColorFgBrightBlack  = "\033[90m"
	ColorFgBrightRed    = "\033[91m"
	ColorFgBrightGreen  = "\033[92m"
	ColorFgBrightYellow = "\033[93m"
	ColorFgBrightBlue   = "\033[94m"
	ColorFgBrightPurple = "\033[95m"
	ColorFgBrightCyan   = "\033[96m"
	ColorFgBrightWhite  = "\033[97m"

	ColorBgBrightBlack  = "\033[100m"
	ColorBgBrightRed    = "\033[101m"
	ColorBgBrightGreen  = "\033[102m"
	ColorBgBrightYellow = "\033[103m"
	ColorBgBrightBlue   = "\033[104m"
	ColorBgBrightPurple = "\033[105m"
	ColorBgBrightCyan   = "\033[106m"
	ColorBgBrightWhite  = "\033[107m"
)

var GlobalStyleStack = &styleStack{}

func NewColor(colors ...string) func(string) string {
	code := Join(colors, "")

	return func(text string) string {
		GlobalStyleStack.Push(code)
		current := GlobalStyleStack.GetActive()

		// FIX: If the internal text has a Reset, replace it with
		// Reset + the current styles so the color "continues"
		// after the nested element finishes.
		sanitizedText := strings.ReplaceAll(text, ColorReset, current)

		GlobalStyleStack.Pop()
		previous := GlobalStyleStack.GetActive()

		return current + sanitizedText + previous
	}
}

type styleStack struct {
	stack []string
}

func (s *styleStack) GetActive() string {
	if len(s.stack) == 0 {
		return ColorReset
	}
	return ColorReset + Join(s.stack, "")
}

func (s *styleStack) Push(code string) {
	s.stack = append(s.stack, code)
}

func (s *styleStack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}
