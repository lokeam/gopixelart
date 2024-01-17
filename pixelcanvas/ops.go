package pixelcanvas

import "fyne.io/fyne/v2"

func (pixelCanvas *PixelCanvas) Pan(previousCoordinate, currentCoordinate fyne.PointEvent) {
	xCoordDiff := currentCoordinate.Position.X - previousCoordinate.Position.X
	yCoordDiff := currentCoordinate.Position.Y - previousCoordinate.Position.Y

	pixelCanvas.CanvasOffset.X += xCoordDiff
	pixelCanvas.CanvasOffset.Y += yCoordDiff
	pixelCanvas.Refresh()
}

func (pixelCanvas *PixelCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pixelCanvas.PixelSize += 1
	case direction < 0:
		if pixelCanvas.PixelSize > 2 {
			pixelCanvas.PixelSize -= 1
		}
	default:
		pixelCanvas.PixelSize = 10
	}
}
