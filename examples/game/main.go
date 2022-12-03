package main

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/sprites"
	"github.com/yohamta/furex/v2/examples/game/text"
	"github.com/yohamta/furex/v2/examples/game/widgets"
)

type Game struct {
	initOnce sync.Once
	screen   screen
	gameUI   *furex.View
}

type screen struct {
	Width  int
	Height int
}

func (g *Game) Update() error {
	g.initOnce.Do(func() {
		g.setupUI()
	})
	g.gameUI.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{63, 124, 182, 255})
	g.gameUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.screen.Width = outsideWidth
	g.screen.Height = outsideHeight
	return g.screen.Width, g.screen.Height
}

func NewGame() (*Game, error) {
	text.LoadFonts()
	sprites.LoadSprites(
		"assets/images/uipack_rpg_sheet.xml",
		"assets/images/uipack_rpg_sheet.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{
				"panelInset_beige.png": {
					Border: 32,
					Center: 36,
				},
				"panel_brown.png": {
					Border: 32,
					Center: 36,
				},
			},
		})
	sprites.LoadSprites(
		"assets/images/uipackSpace_sheet.xml",
		"assets/images/uipackSpace_sheet.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{
				"glassPanel_corners.png": {
					Border: 40,
					Center: 20,
				},
				"glassPanel_projection.png": {
					Border: 20,
					Center: 10,
				},
			},
		})
	game := &Game{}
	return game, nil
}

func (g *Game) setupUI() {
	newButton := func(txt string) *furex.View {
		return (&furex.View{
			Width:        45,
			Height:       49,
			MarginTop:    5,
			MarginBottom: 10,
			MarginLeft:   20,
			MarginRight:  20,
			Handler: &widgets.Button{
				Text:          txt,
				Color:         color.RGBA{210, 178, 144, 255},
				Sprite:        "buttonSquare_brown.png",
				SpritePressed: "buttonSquare_brown_pressed.png",
				OnClick:       func() { println("button clicked") },
			},
		})
	}

	g.gameUI = (&furex.View{
		Width:        g.screen.Width,
		Height:       g.screen.Height,
		Direction:    furex.Column,
		Justify:      furex.JustifySpaceBetween,
		AlignItems:   furex.AlignItemStretch,
		AlignContent: furex.AlignContentStretch,
	}).AddChild(
		(&furex.View{
			MarginTop:  50,
			Grow:       1,
			AlignItems: furex.AlignItemCenter,
			Justify:    furex.JustifyCenter,
		}).AddChild(
			// panel
			(&furex.View{
				Width:      300,
				Height:     300,
				MarginTop:  120,
				MarginLeft: 130,
				Handler:    &widgets.Panel{Sprite: "panel_brown.png"},
				Direction:  furex.Column,
				AlignItems: furex.AlignItemCenter,
				Justify:    furex.JustifyCenter,
			}).AddChild(
				// panel inside panel
				(&furex.View{
					MarginTop:  20,
					Width:      245,
					Height:     200,
					Handler:    &widgets.Panel{Sprite: "panelInset_beige.png"},
					Direction:  furex.Column,
					AlignItems: furex.AlignItemCenter,
					Justify:    furex.JustifyCenter,
				}).AddChild(
					// gauges
					(&furex.View{
						Width:      180,
						Height:     38,
						AlignItems: furex.AlignItemStart,
						Justify:    furex.JustifyStart,
						Direction:  furex.Column,
					}).AddChild(
						&furex.View{
							Height:       20,
							Width:        180,
							MarginBottom: 2,
							Handler: &widgets.Text{
								Color: color.RGBA{50, 48, 41, 255},
								Value: "Health",
							},
						},
						&furex.View{
							Width:  180,
							Height: 18,
							Handler: &widgets.Bar{
								Color: "Green",
								Value: .8,
							},
						},
					),
					(&furex.View{
						Width:      180,
						Height:     38,
						AlignItems: furex.AlignItemStart,
						Justify:    furex.JustifyStart,
						Direction:  furex.Column,
						MarginTop:  20,
					}).AddChild(
						&furex.View{
							Height:       20,
							Width:        180,
							MarginBottom: 2,
							Handler: &widgets.Text{
								Color: color.RGBA{50, 48, 41, 255},
								Value: "Mana",
							},
						},
						&furex.View{
							Width:  180,
							Height: 18,
							Handler: &widgets.Bar{
								Color: "Blue",
								Value: .5,
							},
						},
					),
				),
				// buttons inside panel
				(&furex.View{
					MarginTop:    20,
					MarginBottom: 20,
					Grow:         1,
					Direction:    furex.Row,
					AlignItems:   furex.AlignItemCenter,
					Justify:      furex.JustifyCenter,
				}).AddChild(
					// button 1
					&furex.View{
						Width:  190,
						Height: 49,
						Handler: &widgets.Button{
							Text:          "Inventory",
							Sprite:        "buttonLong_blue.png",
							SpritePressed: "buttonLong_blue_pressed.png",
							OnClick:       func() { println("button clicked") },
						},
					},
					// button 2
					&furex.View{
						Width:      45,
						Height:     49,
						MarginLeft: 10,
						Handler: &widgets.Button{
							Text:          "OK",
							Sprite:        "buttonSquare_blue.png",
							SpritePressed: "buttonSquare_blue_pressed.png",
							OnClick:       func() { println("button clicked") },
						},
					},
				),
				// close button
				(&furex.View{
					Position: furex.PositionAbsolute,
					Left:     300 - 35/2,
					Top:      4 - 38/2,
					Width:    35,
					Height:   38,
					Handler: &widgets.Button{
						Sprite:  "buttonRound_blue.png",
						OnClick: func() { println("button clicked") },
					},
				}).AddChild(
					&furex.View{
						Position: furex.PositionAbsolute,
						Left:     18,
						Top:      17,
						Handler: &widgets.Sprite{
							Sprite: "iconCross_beige.png",
						},
					},
				),
			),
		),
	).AddChild(
		// buttons at the bottom
		(&furex.View{
			Width:        g.screen.Width,
			Height:       140,
			Justify:      furex.JustifyCenter,
			AlignItems:   furex.AlignItemEnd,
			MarginBottom: 20,
		}).AddChild(
			newButton("A"),
			newButton("B"),
			newButton("C"),
			newButton("D"),
		),
	).AddChild(
		(&furex.View{
			Position: furex.PositionAbsolute,
			Left:     20,
			Top:      52,
		}).AddChild(
			// panel
			(&furex.View{
				Width:      260,
				Height:     140,
				Handler:    &widgets.Panel{Sprite: "glassPanel_corners.png"},
				Direction:  furex.Column,
				AlignItems: furex.AlignItemCenter,
				Justify:    furex.JustifyCenter,
			}).AddChild(
				&furex.View{
					Width:        100,
					Height:       8,
					MarginBottom: 20,
					Direction:    furex.Row,
					AlignItems:   furex.AlignItemCenter,
					Justify:      furex.JustifyCenter,
					Handler: &widgets.Text{
						Color:     color.RGBA{45, 73, 94, 255},
						Value:     "PLAY THE GAME?",
						HorzAlign: etxt.XCenter,
						VertAlign: etxt.YCenter,
					},
				},
				(&furex.View{
					Width:      100,
					Height:     50,
					Direction:  furex.Row,
					AlignItems: furex.AlignItemCenter,
					Justify:    furex.JustifyCenter,
				}).AddChild(
					&furex.View{
						Width:  100,
						Height: 50,
						Handler: &widgets.Panel{
							Sprite:  "glassPanel_projection.png",
							Text:    "YES",
							OnClick: func() { println("button clicked") },
						},
					},
					&furex.View{
						Width:      100,
						Height:     50,
						MarginLeft: 20,
						Handler: &widgets.Panel{
							Sprite:  "glassPanel_projection.png",
							Text:    "NO",
							OnClick: func() { println("button clicked") },
						},
					},
				),
			),
		),
	).AddChild(
		// panels that draws mouse cursor
		&furex.View{
			Width:    g.screen.Width,
			Height:   g.screen.Height,
			Position: furex.PositionAbsolute,
			Left:     0,
			Top:      0,
			Handler:  &widgets.Mouse{},
		},
	)
}

func main() {
	ebiten.SetWindowSize(480, 640)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
