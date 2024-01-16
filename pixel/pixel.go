package main

import (
	"image/color"
	"sixam/gopixel/apptype"
	"sixam/gopixel/swatch"
	"sixam/gopixel/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixel")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
