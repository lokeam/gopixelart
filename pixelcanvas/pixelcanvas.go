package pixelcanvas

import (
	"image"
	"image/color"
	"sixam/gopixel/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/widget"
)

type PixelCanvasMouseState struct {
	previousCoordinate *fyne.PointEvent
}

type PixelCanvas struct {
	widget.BaseWidget
	apptype.PixelCanvasConfig
	renderer    *PixelCanvasRenderer
	mouseState  PixelCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (PixelCanvas *PixelCanvas) Bounds() image.Rectangle {
	x0 := int(pixelCanvas.CanvasOffset.X)
	y0 := int(pixelCanvas.CanvasOffset.Y)
	x1 := int(pixelCanvas.PixelCols*pixelCanvas.PixelSize + int(PixelCanvas.CanvasOffset.X))
	y1 := int(pixelCanvas.PixelRows*pixelCanvas.PixelSize + int(PixelCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
		pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) &&
		pos.X < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPixelCanvas(state *apptype.State, config apptype.PixelCanvasConfig) *PixcelCanvas {
	pixelCanvas := &PixelCanvas{
		PixelCanvasConfig: config,
		appState:          state,
	}
	pixelCanvas.PixelData = NewBlankImage(pixelCanvas.PixelCols, pixelCanvas.PixelRows, color.NRGBA{128, 128, 128, 255})
	pixelCanvas.ExtendBaseWidget(pixelCanvas)
	return pixelCanvas
}
