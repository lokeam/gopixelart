package brush

import (
	"sixam/gopixel/apptype"

	"fyne.io/fyne/v2/driver/desktop"
)

const (
	Pixel = iota
)

// Separate out TryBrush and TryToPaint so that we may add new brushes later
func TryBrush(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryToPaintPixel(appState, canvas, ev)
	default:
		return false
	}
}

func TryToPaintPixel(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}
