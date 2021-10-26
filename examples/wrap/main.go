package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
	"github.com/yohamta/furex/examples/shared"
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

var (
	colors = []color.Color{
		color.RGBA{0xaa, 0, 0, 0xff},
		color.RGBA{0, 0xaa, 0, 0xff},
		color.RGBA{0, 0, 0xaa, 0xff},
	}
)

func (g *Game) buildUI() {
	rootFlex = furex.NewFlex(screenWidth, screenHeight)
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	for i := 0; i < 20; i++ {
		rootFlex.AddChild(shared.NewBox(50, 50, colors[i%3]))
	}
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
