# Furex

Furex is a minimal GUI framework for [Ebitengine](https://ebiten.org/) that supports a subset of [Flex Layout Algorithm](https://www.w3.org/TR/css-flexbox-1/#layout-algorithm).

For now, Furex is not a component library but a framework for positioning and stacking virtual widgets, handling button and touch events. How they are rendered is completely up to the user.

[GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2)

## Motivation

[Flexbox](https://www.w3.org/TR/css-flexbox/) is a good mechanism for laying out items of different sizes. I wanted to use the same concept for game UI because I have experience in Web and ReactNative projects. I hope the library helps other people with the same thoughts.

## Features

| Feature            | How to use                                                                                                                                                                                                                                                                                                                                              | Example                                                                       |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| Flexbox layout     | The flex layout can be configured via the properties of a [View](https://pkg.go.dev/github.com/yohamta/furex/v2#View) instance. We can think of a `View` as a `div` element in HTML, which can be stacked or nested.
| Custom UI          | A `View` can receive a `Handler` instance. We can create any type of UI components by implementing the [DrawHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#DrawHandler) and/or the [UpdateHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#UpdateHandler) interface. |
| Buttons            | To create a button, you can implement the [ButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#ButtonHandler). It supports both touch and mouse.                                                                                                                                                                                                  | [Example](https://github.com/yohamta/furex/blob/master/widgets/button.go) |
| Touch events       | To handle touch events and positions, you can implement [TouchHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#TouchHandler).                                                                                                                                                                                                                    |                                                                               |
| Mouse click events | To handle mouse click events, you can implement [MouseLeftButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseLeftButtonHandler).                                                                                                                                                                                                        |                                                                               |
| Mouse move events  | To detect mouse positions, you can implement [MouseHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseHandler).                                                                                                                                                                                                                               |                                                                               |
| Mouse enter/leave events  | To detect mouse enter/leave events, you can implement [MouseEnterLeaveHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseEnterLeaveHandler).                                                                                                                                                                                                                               |                                                                               |
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
      Handler: &widgets.Box{
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
			Handler: &widgets.Button{
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
				Handler: &widgets.Box{Color: color.RGBA{0xff, 0, 0, 0xff}},
			},
			&furex.View{
				Width:   200,
				Height:  60,
				Handler: &widgets.Box{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}},
			},
			&furex.View{
				Width:   100,
				Height:  100,
				Handler: &widgets.Box{Color: color.RGBA{0, 0xff, 0, 0xff}},
			},
		),
	).AddChild(&furex.View{
		Width:  200,
		Height: 50,
		Handler: &widgets.Button{
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

## Debug

You can turn on the Debug Mode by setting the variable below.
```go
furex.Debug = true
```

## How to contribute?

Feel free to contribute in any way you want. Share ideas, submit issues, create pull requests, adding examples, or adding widgets. Thank you!
