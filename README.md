# furex

A simple and flexible UI framework for [Ebiten](https://ebiten.org/) with a subset of flexbox layout specifications.
[GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2)

## Motivation

When I was developing React Native apps, I thought the Flexbox layout was a very intuitive UI, so I thought it would be great if I could use the same concept for the UI of the games I make with Ebiten. I hope this library will help others with the same idea. Any types of Issue/PRs are welcomed :)

## Features

- Flexbox layout
  - Supports a subset of flexbox layout spec.
- Custom component
  - Supports any component that implements `DrawHandler` or `UpdateHandler`. See the [example](https://github.com/yohamta/furex/blob/master/components/box.go).
- Button component
  - Able to handle touch ID by implementing `ButtonHandler` interface. See the [example](https://github.com/yohamta/furex/blob/master/components/button.go).
- Touch event handling
  - Able to handle touch ID by implementing `TouchHandler` interface. See the [GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2#TouchHandler).
- Mouse click / Mouse move handling
  - Able to handle left click by implementing `MouseHandler` interface. See the [GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseHandler).
  - Able to handle mouse move by implementing `MouseLeftButtonHandler` interface. See the [GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseLeftButtonHandler).
- Margin
- Nesting
- Wrap
- Flex grow
- Align stretch
- Absolute-position with Left and Top
- Method chaning

## Examples
To check all examples, visit [this page](https://github.com/yohamta/furex/tree/main/examples).

## Simple Usage

```sh
go get github.com/yohamta/furex/v2
```

[Full source code of simple usage example](https://github.com/yohamta/furex/blob/master/examples/wrap/main.go)

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
  g.gameUI.Update()
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

### Result

<image src="https://user-images.githubusercontent.com/1475839/133445715-b94b8c7f-bcd3-4aef-b7a4-b58bbb29d556.png" width="500px" />

## Method chaining example

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

### Result

<image src="https://user-images.githubusercontent.com/1475839/165524288-53827304-731e-4f33-81cd-26bb6a42e0d4.png" width="500px" />


## How to contribute?

Feel free to contribute in any way you want. Share ideas, submit issues, create pull requests. 
Thank you!

