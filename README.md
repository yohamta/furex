# furex
A simple and extensible UI library with small subset of flexbox-layout which is made to create a game with [Ebiten](https://ebiten.org/)

## Simple usage

```go
import "github.com/yohamta/furex"

func (g *Game) initUI() {
	// root flex container
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)
	rootFlex.Direction = furex.Column
	rootFlex.Justify = furex.JustifySpaceBetween
	rootFlex.AlignItems = furex.AlignItemCenter

	// flex item: box0
	b0 := furex.NewBox(100, 100, color.RGBA{0xff, 0, 0, 0xff})
	rootFlex.AddChild(b0)

	// flex item: box1
	b1 := furex.NewBox(100, 100, color.RGBA{0, 0xff, 0, 0xff})
	rootFlex.AddChild(b1)

	// view controller
	g.vc = furex.NewViewController()
	g.vc.SetRootView(rootFlex)
}

func (g *Game) Update() {
	// Update the container's children
	g.vc.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw children in the flex containers
	g.vc.Draw(screen, image.Rect(0, 0, screenWidth, screenHeight))
}
```
