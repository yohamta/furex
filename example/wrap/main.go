package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/furex"
)

type Game struct {
	vc *furex.ViewController
}

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

var (
	windowRect         image.Rectangle
	screenWidth        = 240
	screenHeight       = 360
	desktopScreenScale = 2
	isWindowSizeSet    = false
	isInitialized      = false
)

func (g *Game) Update(screen *ebiten.Image) error {
	if isInitialized == false {
		g.buildUI()
		isInitialized = true
	}
	// update flex container and it's children
	g.vc.Update()
	return nil
}

func (g *Game) SetWindowSize(width, height int) {
	screenHeight = int(float64(screenWidth) / float64(width) * float64(height))
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw flex items
	g.vc.Draw(screen, windowRect)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) buildUI() {
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	for i := 0; i < 20; i++ {
		var c color.RGBA
		switch i % 3 {
		case 0:
			c = color.RGBA{0xff, 0, 0, 0xff}
		case 1:
			c = color.RGBA{0, 0xff, 0, 0xff}
		default:
			c = color.RGBA{0, 0, 0xff, 0xff}
		}
		rootFlex.AddChild(furex.NewBox(50, 50, c))
	}

	// view controller
	g.vc = furex.NewViewController()
	g.vc.SetRootView(rootFlex)
}

func main() {
	windowRect = image.Rect(0, 0, screenWidth*desktopScreenScale, screenHeight*desktopScreenScale)
	size := windowRect.Size()
	ebiten.SetWindowSize(size.X, size.Y)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	game.SetWindowSize(size.X, size.Y)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
