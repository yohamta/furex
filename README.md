# furex

A simple UI framework for [Ebiten](https://ebiten.org/) with a subset of flexbox layout specification.
[GoDoc](https://pkg.go.dev/github.com/miyahoyo/furex)

## Motivation

When I was developing some React Native app, I thought the Flexbox layout is very intuitive for UI and it would be great if I could use the same concept when building a UI for a game developed in Ebiten. I hope this library will help others with the same thoughts.

## Features

It provides minimum functionalities to implement user interactions such as buttons, as well as functionality to implement UI using Flexbox layout.

| Feature                 | Supported | Note                                                                                                                   |
|-------------------------|------------------|------------------------------------------------------------------------------------------------------------------------|
| Flexbox layout          | o                | Supports a subset of flexbox layout spec.                                                                              |
| Custom component   | o                | Supports any component that implements `Drawable` (and `Updatable`) interface. See the [example](https://github.com/miyahoyo/furex/blob/master/examples/shared/box.go). |
| Button event handling   | o                | Supports both of touch and mouse click on components that implements `Button` interface. See the [example](https://github.com/miyahoyo/furex/blob/master/examples/shared/button.go). |
| Touch handler interface | o                | Able to handle touch ID on components that implements the `TouchHandler` interface.                                                                             |
| Mouse handler           | o                | Able to handle left click on components that implements the `MouseHandler` interface.                                                                                |
| Margin           | o                | Margin is supported for components that implement `MarginedItem` interface.
| Padding           | -                | To be implemented when needed.                                                     |



## Layout Example

[Full source code of the example](https://github.com/miyahoyo/furex/blob/master/examples/nesting/main.go)

<image src="https://user-images.githubusercontent.com/1475839/133440846-dae6cc3e-22d4-4e13-965c-7989b50ed58a.png" width="500px" />


## Simple Usage

[Full source code of simple usage example](https://github.com/miyahoyo/furex/blob/master/examples/wrap/main.go)

```go
import "github.com/miyahoyo/furex"

var (
	colors = []color.Color{
		color.RGBA{0xaa, 0, 0, 0xff},
		color.RGBA{0, 0xaa, 0, 0xff},
		color.RGBA{0, 0, 0xaa, 0xff},
	}
)

var {
	rootFlex *furex.Flex
}

// Initialize the UI
func (g *Game) initUI() {
	// Make a instance of root flexbox container
	rootFlex = furex.NewFlex(screenWidth, screenHeight)

	// Set options for flexbox layout
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	// Add items to flexbox container
	for i := 0; i < 20; i++ {
		rootFlex.AddChild(NewBox(50, 50, colors[i%3]))
	}
}

func (g *Game) Update() {
	// Update the UI tree
	rootFlex.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the UI tree
	rootFlex.Draw(screen)
}
```

### Result
<image src="https://user-images.githubusercontent.com/1475839/133445715-b94b8c7f-bcd3-4aef-b7a4-b58bbb29d556.png" width="500px" />

