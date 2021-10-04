package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miyahoyo/furex"
	"github.com/miyahoyo/furex/examples/shared"
)

type Game struct{}

const desktopScreenScale = 2

var (
	screenWidth   int
	screenHeight  int
	isInitialized = false
	rootFlex      *furex.Flex
)

func (g *Game) Update() error {
	if isInitialized == false {
		g.buildUI()
		isInitialized = true
	}
	rootFlex.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	rootFlex.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	screenWidth = outsideWidth / desktopScreenScale
	screenHeight = outsideHeight / desktopScreenScale
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) buildUI() {
	// root flex container
	rootFlex = furex.NewFlex(screenWidth, screenHeight)
	rootFlex.Direction = furex.Column
	rootFlex.Justify = furex.JustifySpaceBetween
	rootFlex.AlignItems = furex.AlignItemCenter

	// flex item: box0
	b0 := shared.NewBox(100, 100, color.RGBA{0xff, 0, 0, 0xff})
	rootFlex.AddChild(b0)

	// flex item: box1
	b1 := shared.NewBox(100, 100, color.RGBA{0, 0xff, 0, 0xff})
	rootFlex.AddChild(b1)
}

func main() {
	ebiten.SetWindowSize(480, 640)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
