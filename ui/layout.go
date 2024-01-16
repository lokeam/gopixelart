package ui

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.PixelWindow.SetContent(swatchesContainer)
}
