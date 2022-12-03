package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

func createPanels(img *ebiten.Image, r image.Rectangle, border int) map[string]*ganim8.Sprite {
	ret := map[string]*ganim8.Sprite{}

	// top left
	g := ganim8.NewGrid(border, border, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X, r.Min.Y)
	ret["top_left"] = ganim8.NewSprite(img, g.Frames())
	// top
	g = ganim8.NewGrid(r.Dx()-border*2, r.Dy()-border*2, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X+border, r.Min.Y)
	ret["top"] = ganim8.NewSprite(img, g.Frames())
	// top right
	g = ganim8.NewGrid(border, border, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X+r.Dx()-border, r.Min.Y)
	ret["top_right"] = ganim8.NewSprite(img, g.Frames())
	// left
	g = ganim8.NewGrid(border, r.Dy()-border*2, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X, r.Min.Y+border)
	ret["left"] = ganim8.NewSprite(img, g.Frames())
	// center
	g = ganim8.NewGrid(r.Dx()-border*2, r.Dy()-border*2, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X+border, r.Min.Y+border)
	ret["center"] = ganim8.NewSprite(img, g.Frames())
	// right
	g = ganim8.NewGrid(border, r.Dy()-border*2, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X+r.Dx()-border, r.Min.Y+border)
	ret["right"] = ganim8.NewSprite(img, g.Frames())
	// bottom left
	g = ganim8.NewGrid(border, border, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X, r.Min.Y+r.Dy()-border)
	ret["bottom_left"] = ganim8.NewSprite(img, g.Frames())
	// bottom
	g = ganim8.NewGrid(r.Dx()-border*2, border, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X+border, r.Min.Y+r.Dy()-border)
	ret["bottom"] = ganim8.NewSprite(img, g.Frames())
	// bottom right
	g = ganim8.NewGrid(border, border, img.Bounds().Dx(), img.Bounds().Dy(), r.Min.X+r.Dx()-border, r.Min.Y+r.Dy()-border)
	ret["bottom_right"] = ganim8.NewSprite(img, g.Frames())

	return ret
}
