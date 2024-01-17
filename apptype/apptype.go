package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
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

// Separate package for brushes so that we may add more later
type Brushable interface {
	SetColor(c color.Color, x, y int)
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}
