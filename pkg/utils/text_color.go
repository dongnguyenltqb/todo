package utils

func WithColor(color string, s string) string {
	switch color {
	case "white":
		return "\x1B[37m" + s + "\x1B[39m"
	case "red":
		return "\x1B[31m" + s + "\x1B[39m"
	case "blue":
		return "\x1B[34m" + s + "\x1B[39m"
	case "green":
		return "\x1B[32m" + s + "\x1B[39m"
	case "mangeta":
		return "\x1B[35m" + s + "\x1B[39m"
	case "yellow":
		return "\x1B[33m" + s + "\x1B[39m"
	default:
		return s
	}
}

func WithStyle(style string, s string) string {
	switch style {
	case "bold":
		return "\x1B[1m" + s + "\x1B[22m"
	case "italic":
		return "\x1B[3m" + s + "\x1B[23m"
	case "underline":
		return "\x1B[4m" + s + "\x1B[24m"
	case "inverse":
		return "\x1B[7m" + s + "\x1B[27m"
	case "strikethrough":
		return "\x1B[9m" + s + "\x1B[29m"
	default:
		return s
	}
}

func ApplyStyle(style, color, s string) string {
	return WithStyle(style, WithColor(color, s))
}
