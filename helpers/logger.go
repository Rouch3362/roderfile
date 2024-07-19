package helpers

import (
	"time"

	"github.com/fatih/color"
)

func GreenLog(text string) {
	time.Sleep(time.Second)
	color.Green(text)
}

func RedLog(text string) {
	time.Sleep(time.Second)
	color.Red(text)
}