package ui

import (
	"errors"
	"image/png"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, e error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PixelCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PixelWindow)
				return
			}
			app.State.SetFilePath(uri.URI().Path())
		}
		// associate with window
	}, app.PixelWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		// if we haven't created a file yet, open save dialog
		if app.State.FilePath == "" {
			saveFileDialog(app)
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()

				if err != nil {
					dialog.ShowError(err, app.PixelWindow)
				}
			}
			// create file path
			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PixelWindow)
			}
			err = png.Encode(fh, app.PixelCanvas.PixelData)

			if err != nil {
				dialog.ShowError(err, app.PixelWindow)
				return
			}
		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		// Ensure user creates an image with a size > 0px
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("please enter a positive Integer")
			}
			if width <= 0 {
				return errors.New("your file must be larger than 0 pixels in size")
			}
			return nil
		}
		widthInput := widget.NewEntry()
		widthInput.Validator = sizeValidator

		heightInput := widget.NewEntry()
		heightInput.Validator = sizeValidator

		widthForm := widget.NewFormItem("Width", widthInput)
		heightForm := widget.NewFormItem("Height", heightInput)

		formItems := []*widget.FormItem{widthForm, heightForm}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if widthInput.Validate() != nil {
					dialog.ShowError(errors.New("invalid width"), app.PixelWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthInput.Text)
				}
				if heightInput.Validate() != nil {
					dialog.ShowError(errors.New("invalid height"), app.PixelWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightInput.Text)
				}
				app.PixelCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PixelWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu(
		"File",
		BuildNewMenu(app),
		BuildSaveMenu(app),
		BuildSaveAsMenu(app),
	)
}

func SetupMenus(app *AppInit) {
	menus := BuildMenus(app)
	mainMenu := fyne.NewMainMenu(menus)
	app.PixelWindow.SetMainMenu(mainMenu)
}
