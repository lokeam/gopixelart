package pixelcanvas

import (
	"image"
	"image/color"
	"sixam/gopixel/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type PixelCanvasMouseState struct {
	previousCoordinate *fyne.PointEvent
}

type PixelCanvas struct {
	widget.BaseWidget
	apptype.PixelCanvasConfig
	renderer    *PixelCanvasRenderer
	PixelData   image.Image
	mouseState  PixelCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (pixelCanvas *PixelCanvas) Bounds() image.Rectangle {
	x0 := int(pixelCanvas.CanvasOffset.X)
	y0 := int(pixelCanvas.CanvasOffset.Y)
	x1 := int(pixelCanvas.PixelCols*pixelCanvas.PixelSize + int(pixelCanvas.CanvasOffset.X))
	y1 := int(pixelCanvas.PixelRows*pixelCanvas.PixelSize + int(pixelCanvas.CanvasOffset.Y))
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

func NewPixelCanvas(state *apptype.State, config apptype.PixelCanvasConfig) *PixelCanvas {
	pixelCanvas := &PixelCanvas{
		PixelCanvasConfig: config,
		appState:          state,
	}
	pixelCanvas.PixelData = NewBlankImage(pixelCanvas.PixelCols, pixelCanvas.PixelRows, color.NRGBA{128, 128, 128, 255})
	pixelCanvas.ExtendBaseWidget(pixelCanvas)
	return pixelCanvas
}

func (pixelCanvas *PixelCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pixelCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)

	// set defaults for pixel canvas borders
	for i := 0; i < len(canvasBorder); i++ {
		// dark gray canvas border
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PixelCanvasRenderer{
		pixelCanvas:  pixelCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}
	pixelCanvas.renderer = renderer
	return renderer
}

// Attempt to pan using the scroll wheel
// Left click: primary
// Right click: secondary
// Scroll wheel: tertiary
func (pixelCanvas *PixelCanvas) TryToPan(previousCoordinate *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoordinate != nil && ev.Button == desktop.MouseButtonTertiary {
		pixelCanvas.Pan(*previousCoordinate, ev.PointEvent)
	}
}

// Brushable interface
func (pixelCanvas *PixelCanvas) SetColor(c color.Color, x, y int) {

	// need to check file format before setting pixel color
	if nrgba, ok := pixelCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}
	if rgba, ok := pixelCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}
	pixelCanvas.Refresh()
}

// Check if moused over canvas
func (pixelCanvas *PixelCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := pixelCanvas.Bounds()
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}

	pixelSize := float32(pixelCanvas.PixelSize)
	xOffset := pixelCanvas.CanvasOffset.X
	yOffset := pixelCanvas.CanvasOffset.Y

	// Get pixel points within image
	x := int((ev.Position.X - xOffset) / pixelSize)
	y := int((ev.Position.Y - yOffset) / pixelSize)

	return &x, &y
}
