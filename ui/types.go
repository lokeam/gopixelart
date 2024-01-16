package ui

import (
	"sixam/gopixel/apptype"
	"sixam/gopixel/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
