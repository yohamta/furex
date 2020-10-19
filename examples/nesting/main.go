package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/furex"
)

type Game struct {
	view *furex.View
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
	g.view.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.view.Draw(screen)
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
	// root container
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)
	rootFlex.Direction = furex.Column
	rootFlex.Justify = furex.JustifySpaceBetween
	rootFlex.AlignContent = furex.AlignContentCenter

	// top container
	top := furex.NewFlex(0, 0, screenWidth, screenHeight/2)
	top.Direction = furex.Row
	top.Justify = furex.JustifyCenter
	top.AlignItems = furex.AlignItemStart
	top.AddChild(furex.NewBox(50, 50, color.RGBA{0xff, 0, 0, 0xff}))
	top.AddChild(furex.NewBox(50, 50, color.RGBA{0, 0xff, 0, 0xff}))
	rootFlex.AddChild(top)

	// bottom container
	bottom := furex.NewFlex(0, 0, screenWidth, screenHeight/2)
	bottom.Direction = furex.Row
	bottom.Justify = furex.JustifyCenter
	bottom.AlignItems = furex.AlignItemEnd
	bottom.AddChild(furex.NewBox(50, 50, color.RGBA{0, 0xff, 0, 0xff}))
	bottom.AddChild(furex.NewBox(50, 50, color.RGBA{0xff, 0, 0, 0xff}))
	rootFlex.AddChild(bottom)

	// layer
	layer := furex.NewLayerWithContainer(rootFlex)

	// view
	g.view = furex.NewView()
	g.view.Layout(0, 0, screenWidth, screenHeight)
	g.view.AddLayer(layer)
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
