package ui

import (
	"sixam/gopixel/apptype"
	"sixam/gopixel/pixelcanvas"
	"sixam/gopixel/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixelCanvas *pixelcanvas.PixelCanvas
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
