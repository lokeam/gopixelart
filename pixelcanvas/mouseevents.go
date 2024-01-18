package pixelcanvas

import (
	"sixam/gopixel/pixelcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// Scrollable Interface
func (pixelCanvas *PixelCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pixelCanvas.scale(int(ev.Scrolled.DY))
	pixelCanvas.Refresh()
}

// Hoverable interface
func (pixelCanvas *PixelCanvas) MouseMoved(ev *desktop.MouseEvent) {
	// if the mouse is somewhere over the canvas, draw a cursor
	if x, y := pixelCanvas.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(pixelCanvas.appState, pixelCanvas, ev)
		cursor := brush.Cursor(pixelCanvas.PixelCanvasConfig, pixelCanvas.appState.BrushType, ev, *x, *y)
		pixelCanvas.renderer.SetCursor(cursor)
	} else {
		pixelCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}

	pixelCanvas.TryToPan(pixelCanvas.mouseState.previousCoordinate, ev)
	pixelCanvas.Refresh()
	pixelCanvas.mouseState.previousCoordinate = &ev.PointEvent

}
func (pixelCanvas *PixelCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pixelCanvas *PixelCanvas) MouseOut()                      {}

// Mousable Interface
func (pixelCanvas *PixelCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pixelCanvas.appState, pixelCanvas, ev)
}

func (PixelCanvas *PixelCanvas) MouseUp(ev *desktop.MouseEvent) {}
