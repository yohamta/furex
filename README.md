# furex

A simple UI framework for [Ebiten](https://ebiten.org/) with a subset of flexbox layout specification.
[GoDoc](https://pkg.go.dev/github.com/yohamta/furex)

## Motivation

When I was developing the React Native app, I found the Flexbox layout to be very intuitive and useful, so I wanted to be able to use the same concept when building a UI for a game developed in EBITEN. I hope this library will help others with the same thoughts.

## Features

It provides minimum functionalities to implement user interactions such as buttons, as well as functionality to implement UI using Flexbox layout.

| Feature                 | Supported | Note                                                                                                                   |
|-------------------------|------------------|------------------------------------------------------------------------------------------------------------------------|
| Flexbox layout          | o                | Supports a subset of flexbox layout spec.                                                                              |
| Custom component   | o                | Supports any component that implements `Component` interface. See the [example](https://github.com/yohamta/furex/blob/master/examples/shared/box.go). |
| Button event handling   | o                | Supports both of touch and mouse click on components that implements `ButtonComponent` interface. See the [example](https://github.com/yohamta/furex/blob/master/examples/shared/button.go). |
| Touch handler interface | o                | Able to handle touch ID on components that implements the `TouchHandler` interface.                                                                             |
| Mouse handler           | o                | Able to handle left click on components that implements the `MouseHandler` interface.                                                                                |
| Padding           | -                | To be implemented when needed.                                                     |
| Margin           | -                | To be implemented when needed.                                                      |



## Layout Example

<image src="https://user-images.githubusercontent.com/1475839/133440846-dae6cc3e-22d4-4e13-965c-7989b50ed58a.png" width="500px" />

[Flex Layout Example using button and nesting flex layout](https://github.com/yohamta/furex/blob/master/examples/nesting/main.go)

## Button Component Example

[Button component example](https://github.com/yohamta/furex/blob/master/examples/shared/button.go)

## Simple Usage Example

[Source code of simple example](https://github.com/yohamta/furex/blob/master/examples/wrap/main.go)

```go
import "github.com/yohamta/furex"

var (
	colors = []color.Color{
		color.RGBA{0xaa, 0, 0, 0xff},
		color.RGBA{0, 0xaa, 0, 0xff},
		color.RGBA{0, 0, 0xaa, 0xff},
	}
)

var {
	view *furex.View
}

// Initialize the UI
func (g *Game) initUI() {
	// flexbox container
	rootFlex := furex.NewFlex(screenWidth, screenHeight)

	// Set the options for flexbox
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	// Add items to flexbox container
	for i := 0; i < 20; i++ {
		// Each flexbox item must have fixed size.
		// In this case, the width is 50, height is 50.
		// Box component is only an example custom component.
		rootFlex.AddChild(NewBox(50, 50, colors[i%3]))
	}

	// View: A View has a flex container as a root of component tree
	view = furex.NewView(image.Rect(0, 0, screenWidth, screenHeight), rootFlex)
}

func (g *Game) Update() {
	// Update the UI 
	view.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the UI 
	view.Draw(screen)
}
```

### Result
<image src="https://user-images.githubusercontent.com/1475839/133445715-b94b8c7f-bcd3-4aef-b7a4-b58bbb29d556.png" width="500px" />

## Use case
This [simple game](https://github.com/yohamta/godanmaku) is using furex for UI layout and interaction.

