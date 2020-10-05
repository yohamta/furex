package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/furex"
)

type Game struct {
	rootFlex *furex.Flex
}

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

var (
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
	g.rootFlex.Update()
	return nil
}

func (g *Game) SetWindowSize(width, height int) {
	screenHeight = int(float64(screenWidth) / float64(width) * float64(height))
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw flex items
	g.rootFlex.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) buildUI() {
	g.rootFlex = furex.NewFlex(0, 0, screenWidth, screenHeight)
	g.rootFlex.Direction = furex.Row
	g.rootFlex.Justify = furex.JustifyCenter
	g.rootFlex.AlignItems = furex.AlignItemCenter
	g.rootFlex.AlignContent = furex.AlignContentCenter
	g.rootFlex.Wrap = furex.Wrap

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
		g.rootFlex.AddChild(furex.NewBox(50, 50, c))
	}

	g.rootFlex.Layout()
}

func main() {
	windowSize := image.Point{screenWidth * desktopScreenScale, screenHeight * desktopScreenScale}
	ebiten.SetWindowSize(windowSize.X, windowSize.Y)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	game.SetWindowSize(windowSize.X, windowSize.Y)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
