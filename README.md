# furex
A simple UI library with a subset of flexbox layout for [Ebiten](https://ebiten.org/)

## Example Usage

```go
import "github.com/yohamta/furex"

type Game struct {
	cont *furex.Controller
}

func (g *Game) initUI() {
	// root flex container
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	for i := 0; i < 20; i++ {
		rootFlex.AddChild(furex.NewBox(50, 50, colors[i%3]))
	}

	// layer
	layer := furex.NewLayerWithContainer(rootFlex)

	// controller
	g.cont = furex.NewController()
	g.cont.Layout(0, 0, screenWidth, screenHeight)
	g.cont.AddLayer(layer)
}

func (g *Game) Update() {
	// Update the UI 
	g.cont.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the UI 
	g.cont.Draw(screen)
}
```

### Result
![image](https://user-images.githubusercontent.com/1475839/95682206-0279fa80-0c1f-11eb-8dd5-03bec58325e8.png)