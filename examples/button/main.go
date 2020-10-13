package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/furex"
	"github.com/yotahamada/furex/examples/button/components"
)

type Game struct {
	cont *furex.Controller
}

const desktopScreenScale = 2

var (
	screenWidth   int
	screenHeight  int
	isInitialized = false
)

func (g *Game) Update(screen *ebiten.Image) error {
	if isInitialized == false {
		g.buildUI()
		isInitialized = true
	}
	g.cont.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.cont.Draw(screen)
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
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)
	rootFlex.Direction = furex.Column
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter

	// flex item: button
	button := components.NewSampleButton(100, 100)
	rootFlex.AddChild(button)

	// layer
	layer := furex.NewLayerWithContainer(rootFlex)

	// controller
	g.cont = furex.NewController()
	g.cont.Layout(0, 0, screenWidth, screenHeight)
	g.cont.AddLayer(layer)
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(480, 640)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
