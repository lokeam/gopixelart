package pixelcanvas

import (
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
	pixelCanvas.TryToPan(pixelCanvas.mouseState.previousCoordinate, ev)
	pixelCanvas.Refresh()
	pixelCanvas.mouseState.previousCoordinate = &ev.PointEvent

}
func (pixelCanvas *PixelCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pixelCanvas *PixelCanvas) MouseOut()                      {}
