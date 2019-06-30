package log

import "fmt"

func Error(msg string) {
	fmt.Println(Sad(RedBold(msg)))
}
func Success(msg string) {
	fmt.Println(Beer(Green(msg)))
}

func Warn(msg string) {
	fmt.Println(Beer(LightMagenta(msg)))
}

func ShowBanner(banner string) {
	fmt.Println(LightMagenta(banner))
}
