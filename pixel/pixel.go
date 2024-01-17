package main

import (
	"image/color"
	"sixam/gopixel/apptype"
	"sixam/gopixel/pixelcanvas"
	"sixam/gopixel/swatch"
	"sixam/gopixel/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixel")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixelCanvasConfig := apptype.PixelCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PixelRows:    10,
		PixelCols:    10,
		PixelSize:    30,
	}

	pixelCanvas := pixelcanvas.NewPixelCanvas(&state, pixelCanvasConfig)

	appInit := ui.AppInit{
		PixelCanvas: pixelCanvas,
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
