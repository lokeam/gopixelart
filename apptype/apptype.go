package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PixelCanvasConfig struct {
	DrawingArea          fyne.Size
	CanvasOffset         fyne.Position
	PixelRows, PixelCols int
	PixelSize            int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
