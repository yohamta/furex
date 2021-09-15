package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
	"github.com/yohamta/furex/examples/shared"
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

func (g *Game) Update() error {
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
	rootFlex := furex.NewFlex(screenWidth, screenHeight)
	rootFlex.Direction = furex.Column
	rootFlex.Justify = furex.JustifySpaceBetween
	rootFlex.AlignContent = furex.AlignContentCenter

	// top container
	top := furex.NewFlex(screenWidth-20, 70)
	top.Direction = furex.Row
	top.Justify = furex.JustifySpaceBetween
	top.AlignItems = furex.AlignItemCenter
	top.AddChild(shared.NewBox(50, 50, color.RGBA{0xff, 0, 0, 0xff}))
	top.AddChild(shared.NewBox(100, 30, color.RGBA{0xff, 0xff, 0xff, 0xff}))
	top.AddChild(shared.NewBox(50, 50, color.RGBA{0, 0xff, 0, 0xff}))
	rootFlex.AddChild(top)

	// center
	rootFlex.AddChild(shared.NewButton(100, 50))

	// bottom container
	bottom := furex.NewFlex(screenWidth, 70)
	bottom.Direction = furex.Row
	bottom.Justify = furex.JustifySpaceAround
	bottom.AlignItems = furex.AlignItemEnd
	bottom.AddChild(shared.NewButton(50, 30))
	bottom.AddChild(shared.NewButton(50, 30))
	bottom.AddChild(shared.NewButton(50, 30))
	bottom.AddChild(shared.NewButton(50, 30))
	rootFlex.AddChild(bottom)

	// view
	g.view = furex.NewView(image.Rect(0, 0, screenWidth, screenHeight), rootFlex)
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
