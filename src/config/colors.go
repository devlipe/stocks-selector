package config

import "runtime"

var (
	Reset = "\033[0m"

	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	BackBlack  = "\033[40m"
	BackRed    = "\033[41m"
	BackGreen  = "\033[42m"
	BackYellow = "\033[43m"
	BackBlue   = "\033[44m"
	BackPurple = "\033[45m"
	BackCyan   = "\033[46m"
	BackGray   = "\033[47m"
	BackWhite  = "\033[107m"
)

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""

		Black = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""

		BackBlack = ""
		BackRed = ""
		BackGreen = ""
		BackYellow = ""
		BackBlue = ""
		BackPurple = ""
		BackCyan = ""
		BackGray = ""
		BackWhite = ""
	}
}
