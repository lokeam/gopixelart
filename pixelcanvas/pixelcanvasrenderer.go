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

	// First four border objects on LayoutBorder func
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}

	objects = append(objects, renderer.canvasImage)
	return objects
}

func (renderer *PixelCanvasRenderer) Destroy() {}

// Base Layout, layout canvas needs to be called first
func (renderer *PixelCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)
}

// Refresh - checks if we need to reload the image
func (renderer *PixelCanvasRenderer) Refresh() {
	if renderer.pixelCanvas.reloadImage {
		renderer.canvasImage = renderer.NewImageFromImage(renderer.pixelCanvas.PixelData)
		// change out image is scaled, smooth or pixel
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.pixelCanvas.reloadImage = false
	}
	renderer.Layout(renderer.pixelCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

// Layout img, px width and height
func (renderer *PixelCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imagePixelWidth := renderer.pixelCanvas.PixelCols
	imagePixelHeight := renderer.pixelCanvas.PixelRows
	pixelSize := renderer.pixelCanvas.PixelSize

	// Allow panning and zoom into pixels for ease of editing
	renderer.canvasImage.Move(fyne.NewPos(renderer.PixelCanvas.CanvasOffset.X, renderer.pixelCanvas.CanvasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imagePixelWidth*pixelSize), float32(imagePixelHeight*pixelSize)))
}

// Create borders of actual onscreen size of image
func (renderer *PixelCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pixelCanvas.CanvasOffset
	imageHeight := renderer.canvasImage.Size().Height
	imageWidth := renderer.canvasImage.Size().Width

	// Positions represent start and ending line positions
	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imageHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imageWidth, offset.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imageWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imageWidth, offset.Y+imageHeight)

	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X+imageWidth, offset.Y)
	bottom.Position2 = fyne.NewPos(offset.X+imageWidth, offset.Y+imageHeight)
}

func (renderer *PixelCanvasRenderer) Layout(size fyne.Size) {}
