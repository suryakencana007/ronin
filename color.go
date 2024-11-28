package ronin

import (
	"fmt"
	"os"
)

const (
	ColorBlack = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite

	ColorBold     = 1
	ColorDarkGray = 90
)

// Colorize function for send (colored or common) message to output.
func Colorize(msg string, c int) string {
	// Send common or colored caption.
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, msg)
}

// ColorizeLevel function for send (colored or common) message to output.
func ColorizeLevel(level string) string {
	var (
		icon string
		c    int
	)
	// Switch color.
	switch level {
	case "success":
		icon = "[OK]"
		c = ColorGreen
	case "error":
		icon = "[ERROR]"
		c = ColorRed
	case "info":
		icon = "[INFO]"
		c = ColorBlue
	}
	return Colorize(icon, c)
}

// ShowMessage function for showing output messages.
func ShowMessage(level, text string, startWithNewLine, endWithNewLine bool) {
	// Define variables.
	var startLine, endLine string

	if startWithNewLine {
		startLine = "\n" // set a new line
	}

	if endWithNewLine {
		endLine = "\n" // set a new line
	}

	// Formatting message.
	message := fmt.Sprintf("%s %s %s %s", startLine, ColorizeLevel(level), text, endLine)

	// Return output.
	_, err := fmt.Fprintln(os.Stdout, message)
	if err != nil {
		return
	}
}

// ShowError function for send error message to output.
func ShowError(text string) error {
	return fmt.Errorf("%s %s", ColorizeLevel("error"), text)
}
