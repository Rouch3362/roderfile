package helpers

import (
	"time"

	"github.com/fatih/color"
)

func GreenLog(text string) {
	time.Sleep(time.Millisecond*150)
	color.Green(text)
}

func RedLog(text string) {
	time.Sleep(time.Millisecond*150)
	color.Red(text)
}


func YellowLog(text string) {
	time.Sleep(time.Millisecond*150)
	color.Yellow(text)
}