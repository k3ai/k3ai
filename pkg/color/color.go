package color

import (
	"github.com/fatih/color"
)

func Alert() {
	color.Set(color.FgHiRed)
}

func Disable() {
	color.Unset()
}

func Done() {
	color.Set(color.FgGreen)
}

func InProgress() {
	color.Set(color.FgYellow)
}

func White() {
	color.Set(color.FgHiWhite)
}
