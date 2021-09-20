package paint

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	imgOfAPixel *ebiten.Image
)

func createRectImg() *ebiten.Image {
	if imgOfAPixel != nil {
		return imgOfAPixel
	}
	imgOfAPixel := ebiten.NewImage(1, 1)
	return imgOfAPixel
}

func FillRect(target *ebiten.Image, r image.Rectangle, clr color.Color) {
	img := createRectImg()
	img.Fill(clr)

	op := &ebiten.DrawImageOptions{}

	size := r.Size()
	op.GeoM.Translate(float64(r.Min.X)*(1/float64(size.X)),
		float64(r.Min.Y)*(1/float64(size.Y)))
	op.GeoM.Scale(float64(size.X), float64(size.Y))

	target.DrawImage(img, op)
}

func DrawRect(target *ebiten.Image, r image.Rectangle, clr color.Color, width int) {
	FillRect(target, image.Rect(r.Min.X, r.Min.Y, r.Min.X+width, r.Max.Y), clr)
	FillRect(target, image.Rect(r.Max.X-width, r.Min.Y, r.Max.X, r.Max.Y), clr)
	FillRect(target, image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Min.Y+width), clr)
	FillRect(target, image.Rect(r.Min.X, r.Max.Y-width, r.Max.X, r.Max.Y), clr)
}
