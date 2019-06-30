package log

import "fmt"

// RedBold returns a red Bold string
func RedBold(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", Bold(message))
}

func Green(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
}

func Bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[21m", message)
}

func LightCyan(message string) string {
	return fmt.Sprintf("\x1b[96m%s\x1b[0m", message)
}

func LightMagenta(message string) string {
	return fmt.Sprintf("\x1b[95m%s\x1b[0m", message)
}
func Sad(message string) string {
	return fmt.Sprintf("\U0001f622 \t%s", message)
}

func Beer(message string) string {
	return fmt.Sprintf("\U0001f37a \t%s", message)
}
