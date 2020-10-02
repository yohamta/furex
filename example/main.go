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
	// root flex container
	g.rootFlex = furex.NewFlex(0, 0, screenWidth, screenHeight)
	g.rootFlex.Direction = furex.Column
	g.rootFlex.Justify = furex.JustifySpaceBetween

	// flex item: box0
	b0 := furex.NewBox(100, 100, color.RGBA{0xff, 0, 0, 0xff})
	g.rootFlex.AddChild(b0)

	// flex item: box1
	b1 := furex.NewBox(100, 100, color.RGBA{0, 0xff, 0, 0xff})
	g.rootFlex.AddChild(b1)
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
