package pixelcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PixelCanvasRenderer struct {
	PixelCanvas  *PixelCanvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
}

// Widget Renderer Interface
func (renderer *PixelCanvasRenderer) MinSize() fyne.Size {
	return renderer.pixelCanvas.DrawingArea
}

// Draw out main app border and append image
func (renderer *PixelCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5) // 4 out of 5 objects will be border lines

	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}

	objects = append(objects, renderer.canvasImage)
	return objects
}

func (renderer *PixelCanvasRenderer) Destroy() {}

// Todo:
// Layout fn
// Layout img, px width and height
// Allow panning and zoom into pixels for ease of editing
// Actual onscreen size of image
