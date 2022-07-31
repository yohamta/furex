# Furex

Furex is a minimal GUI framework for [Ebitengine](https://ebiten.org/) that supports Flexbox layout specifications. [GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2)

For now, Furex is not a component library but a framework for positioning and stacking virtual components, handling button and touch events. How they are rendered is completely up to the user.

## Motivation

[Flexbox](https://www.w3.org/TR/css-flexbox/) is a good mechanism for laying out items of different sizes. I wanted to use the same concept for game UI because I have experience in Web and ReactNative projects. I hope the library helps other people with the same thoughts.

## Features

| Feature            | How to use                                                                                                                                                                                                                                                                                                                                              | Example                                                                       |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| Flexbox layout     | The layout can be adjusted by settings the properties of a [View](https://pkg.go.dev/github.com/yohamta/furex/v2#View). You can think of a View as a `div` in HTML. Views can be tiled or nested.                                                                                                                                                 | [Example](https://github.com/yohamta/furex/blob/master/examples/wrap/main.go) |
| Custom UI          | It supports any type of UI component. To create one, you can create a handler implements [DrawHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#DrawHandler) or [UpdateHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#UpdateHandler) interface. Each [View](https://pkg.go.dev/github.com/yohamta/furex/v2#View) can have one handler. | [Example](https://github.com/yohamta/furex/blob/master/components/button.go)  |
| Buttons            | To create a button, you can implement [ButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#ButtonHandler). It supports both touch and mouse.                                                                                                                                                                                                  | [Example](https://github.com/yohamta/furex/blob/master/components/button.go) |
| Touch events       | To handle touch events and positions, you can implement [TouchHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#TouchHandler).                                                                                                                                                                                                                    |                                                                               |
| Mouse click events | To handle mouse click events, you can implement [MouseLeftButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseLeftButtonHandler).                                                                                                                                                                                                        |                                                                               |
| Mouse move events  | To detect mouse positions, you can implement [MouseHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseHandler).                                                                                                                                                                                                                               |                                                                               |
| Swipe gestures     | It supports swipe gestures in four directions. To handle swipe events, you can implement [SwipeHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#SwipeHandler).                                                                                                                                                                                  |                                                                               |
| Margins            | A [View](https://pkg.go.dev/github.com/yohamta/furex/v2#View) can have margins by setting `MarginLeft`, `MarginTop`, `MarginRight`, `MarginBottom`.                                                                                                                                                                                                                                                            |                                                                               |
| Absolute Positions | A [View](https://pkg.go.dev/github.com/yohamta/furex/v2#View) position can be fixed by setting `PositionAbsolute` to the `Position` field, `Left`, and `Top` positions.                                                                                                                                                                                                                                      |                                                                               |

## Install

```sh
go get github.com/yohamta/furex/v2
```

## Examples
To check all examples, visit [here](examples).

### Simple example

[Full source code](examples/wrap/main.go)

```go
import "github.com/yohamta/furex/v2"

type Game struct {
  init   bool
  screen screen
  gameUI *furex.View
}

func (g *Game) Update() error {
  if !g.init {
    g.init = true
    g.setupUI()
  }
  g.gameUI.UpdateWithSize(ebiten.WindowSize())
  // g.gameUI.Update() // Update() is an alternate method for updating the UI and handling events.
  return nil
}

func (g *Game) setupUI() {
  // create a root view
  g.gameUI = &furex.View{
    Width:        g.screen.Width,
    Height:       g.screen.Height,
    Direction:    furex.Row,
    Justify:      furex.JustifyCenter,
    AlignItems:   furex.AlignItemCenter,
    AlignContent: furex.AlignContentCenter,
    Wrap:         furex.Wrap,
  }

  // create a child view
  for i := 0; i < 20; i++ {
    g.gameUI.AddChild(&furex.View{
      Width:  100,
      Height: 100,
      Handler: &components.Box{
        Color: colors[i%len(colors)],
      },
    })
  }
}

func (g *Game) Draw(screen *ebiten.Image) {
  // Draw the UI tree
  g.gameUI.Draw(screen)
}

var colors = []color.Color{
  color.RGBA{0xaa, 0, 0, 0xff},
  color.RGBA{0, 0xaa, 0, 0xff},
  color.RGBA{0, 0, 0xaa, 0xff},
}
```

<image src="https://user-images.githubusercontent.com/1475839/133445715-b94b8c7f-bcd3-4aef-b7a4-b58bbb29d556.png" width="500px" />

### Method chaining

View's `AddChild()` method returns itself, so it can be chained.

[Full source code of the example](https://github.com/yohamta/furex/blob/master/examples/nesting/main.go)

```go
func (g *Game) setupUI() {
	newButton := func() *furex.View {
		return (&furex.View{
			Width:        100,
			Height:       100,
			MarginTop:    5,
			MarginBottom: 10,
			MarginLeft:   5,
			MarginRight:  5,
			Handler: &components.Button{
				Text:    "Button",
				OnClick: func() { println("button clicked") },
			},
		})
	}

	g.gameUI = (&furex.View{
		Width:      g.screen.Width,
		Height:     g.screen.Height,
		Direction:  furex.Column,
		Justify:    furex.JustifySpaceBetween,
		AlignItems: furex.AlignItemCenter,
	}).AddChild(
		(&furex.View{
			Width:      g.screen.Width - 20,
			Height:     70,
			Justify:    furex.JustifySpaceBetween,
			AlignItems: furex.AlignItemCenter,
		}).AddChild(
			&furex.View{
				Width:   100,
				Height:  100,
				Handler: &components.Box{Color: color.RGBA{0xff, 0, 0, 0xff}},
			},
			&furex.View{
				Width:   200,
				Height:  60,
				Handler: &components.Box{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}},
			},
			&furex.View{
				Width:   100,
				Height:  100,
				Handler: &components.Box{Color: color.RGBA{0, 0xff, 0, 0xff}},
			},
		),
	).AddChild(&furex.View{
		Width:  200,
		Height: 50,
		Handler: &components.Button{
			Text:    "Button",
			OnClick: func() { println("button clicked") },
		},
	}).AddChild((&furex.View{
		Width:      g.screen.Width,
		Height:     140,
		Justify:    furex.JustifyCenter,
		AlignItems: furex.AlignItemEnd,
	}).AddChild(
		newButton(),
		newButton(),
		newButton(),
		newButton(),
	))
}
```

<image src="https://user-images.githubusercontent.com/1475839/165524288-53827304-731e-4f33-81cd-26bb6a42e0d4.png" width="500px" />

## How to contribute?

Feel free to contribute in any way you want. Share ideas, submit issues, create pull requests, adding examples, or adding components. Thank you!
