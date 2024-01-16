package ui

import (
	"image/color"
	"sixam/gopixel/swatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	// create buffer of canvas objects
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)

	for index := 0; index < cap(app.Swatches); index++ {
		initialColor := color.NRGBA{255, 255, 255, 255}

		sw := swatch.NewSwatch(app.State, initialColor, index, func(sw *swatch.Swatch) {
			for index2 := 0; index2 < len(app.Swatches); index2++ {
				// de-select swatches, remove any highlights around borders
				app.Swatches[index2].Selected = false
				canvasSwatches[index2].Refresh()
			}
			app.State.SwatchSelected = sw.SwatchIndex
			app.State.BrushColor = sw.Color
		})

		if index == 0 {
			sw.Selected = true
			app.State.SwatchSelected = 0
			// ensure that UI reflects internal state
			sw.Refresh()
		}

		app.Swatches = append(app.Swatches, sw)
		canvasSwatches = append(canvasSwatches, sw)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
