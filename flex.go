// Referenced code: https://github.com/golang/exp/blob/master/shiny/widget/flex/flex.go
package furex

import (
	"fmt"
	"image"
	"math"

	"github.com/yohamta/furex/internal/container"
)

// Direction is the direction in which flex items are laid out
type Direction uint8

const (
	Row Direction = iota
	Column
)

// Justify aligns items along the main axis.
type Justify uint8

const (
	JustifyStart        Justify = iota // pack to start of line
	JustifyEnd                         // pack to end of line
	JustifyCenter                      // pack to center of line
	JustifySpaceBetween                // even spacing
	JustifySpaceAround                 // even spacing, half-size on each end
)

// FlexAlign represents align of flex children
type FlexAlign int

const (
	FlexCenter FlexAlign = iota
	FlexStart
	FlexEnd
	FlexSpaceBetween
)

// AlignItem aligns items along the cross axis.
type AlignItem uint8

const (
	AlignItemStart AlignItem = iota
	AlignItemEnd
	AlignItemCenter
)

// FlexWrap controls whether the container is single- or multi-line,
// and the direction in which the lines are laid out.
type FlexWrap uint8

const (
	NoWrap FlexWrap = iota
	Wrap
	WrapReverse
)

// AlignContent is the 'align-content' property.
// It aligns container lines when there is extra space on the cross-axis.
type AlignContent uint8

const (
	AlignContentStart AlignContent = iota
	AlignContentEnd
	AlignContentCenter
	AlignContentSpaceBetween
	AlignContentSpaceAround
)

// Flex is a container widget that lays out its children following the flexbox algorithm.
type Flex struct {
	containerEmbed

	Direction    Direction
	Wrap         FlexWrap
	Justify      Justify
	AlignItems   AlignItem
	AlignContent AlignContent
}

// NewFlex creates NewFlexContaienr
func NewFlex(width, height int) *Flex {
	f := new(Flex)

	f.Direction = Row
	f.Wrap = NoWrap
	f.Justify = JustifyStart
	f.AlignItems = AlignItemCenter
	f.AlignContent = AlignContentStart
	f.frame = image.Rect(0, 0, width, height)
	f.isDirty = true
	f.parent = nil

	return f
}

func (f *Flex) Update() {
	if f.isDirty {
		f.layout()
		f.isDirty = false
	}
	for c := range f.children {
		updatable, ok := f.children[c].Item.(Updatable)
		if ok && updatable != nil {
			updatable.Update()
		}
	}
	f.processEvent()
}

// layout is the main routine that implements a subset of flexbox layout
// https://www.w3.org/TR/css-flexbox-1/#layout-algorithm
func (f *Flex) layout() {
	// 9.2. Line Length Determination
	// Determine the available main and cross space for the flex items.
	containerMainSize := float64(f.mainSize(f.Size()))
	containerCrossSize := float64(f.crossSize(f.Size()))

	// Determine the flex base size and hypothetical main size of each item:
	var children []element
	for i := 0; i < len(f.children); i++ {
		c := f.children[i]
		absolute, ok := c.Item.(AbsolutePositionItem)
		if ok {
			x, y := absolute.Position()
			w, h := absolute.Size()
			c.Bounds = image.Rect(x, y, x+w, y+h)
			continue
		}
		children = append(children, element{
			flexBaseSize: float64(f.flexBaseSize(c)),
			node:         c,
		})
	}

	// §9.3. Main Size Determination
	// Collect flex items into flex lines
	var lines []flexLine
	if f.Wrap == NoWrap {
		// Single line
		line := flexLine{child: make([]*element, len(children))}
		for i := range children {
			child := &children[i]
			line.child[i] = child
			line.mainSize += child.flexBaseSize
		}
		lines = []flexLine{line}
	} else {
		// Multi line
		var line flexLine
		for i := range children {
			child := &children[i]

			// Use flexBaseSize as hypotheticalMainSize for now
			hypotheticalMainSize := child.flexBaseSize

			if line.mainSize > 0 && line.mainSize+hypotheticalMainSize > containerMainSize {
				lines = append(lines, line)
				line = flexLine{}
			}
			line.child = append(line.child, child)
			line.mainSize += hypotheticalMainSize
		}

		if len(line.child) > 0 || len(children) == 0 {
			lines = append(lines, line)
		}
	}

	for l := range lines {
		line := &lines[l]

		// Calculate free space
		freeSpace := float64(f.mainSize(f.Size()))
		for _, child := range line.child {
			freeSpace -= float64(f.flexBaseSize(child.node))
		}

		// Distribute free space
		for _, child := range line.child {
			child.mainSize = float64(f.flexBaseSize(child.node))
		}
	}

	// §9.4. Cross Size Determination
	// Determine the hypothetical cross size of each item
	for l := range lines {
		for _, child := range lines[l].child {
			fixedSizeC, _ := child.node.Item.(FixedSizeItem)
			if fixedSizeC != nil {
				child.crossSize = float64(f.crossSize(fixedSizeC.Size()))
			} else {
				panic("flexible size is not available for now")
			}
		}
	}

	// §9.4.8 Calculate the cross size of each flex line.
	if len(lines) == 1 {
		// Single line
		lines[0].crossSize = containerCrossSize
	} else {
		// Multi line
		for l := range lines {
			line := &lines[l]
			max := 0.0
			for _, child := range line.child {
				if child.crossSize > max {
					max = child.crossSize
				}
			}
			line.crossSize = max
		}
	}

	off := 0.0
	for l := range lines {
		line := &lines[l]
		line.crossOffset = off
		off += line.crossSize
	}

	// §9.5. Main-Axis Alignment
	for l := range lines {
		line := &lines[l]
		total := 0.0
		for _, child := range line.child {
			total += child.mainSize
		}
		remFree := containerMainSize - total
		off, spacing := 0.0, 0.0
		switch f.Justify {
		case JustifyStart:
		case JustifyEnd:
			off = remFree
		case JustifyCenter:
			off = remFree / 2
		case JustifySpaceBetween:
			spacing = remFree / float64(len(line.child)-1)
		case JustifySpaceAround:
			spacing = remFree / float64(len(line.child))
			off = spacing / 2
		}
		for _, child := range line.child {
			child.mainOffset = off
			off += spacing + child.mainSize
		}
	}

	// §9.6. Cross axis alignment
	for l := range lines {
		line := &lines[l]
		for _, child := range line.child {
			child.crossOffset = line.crossOffset
			if child.crossSize == line.crossSize {
				continue
			}
			diff := line.crossSize - child.crossSize
			switch f.AlignItems {
			case AlignItemStart:
				// already laid out correctly
			case AlignItemEnd:
				child.crossOffset = line.crossOffset + diff
			case AlignItemCenter:
				child.crossOffset = line.crossOffset + diff/2
			}
		}
	}

	// §9.6.15 determine container cross size used
	crossSize := lines[len(lines)-1].crossOffset + lines[len(lines)-1].crossSize
	remFree := containerCrossSize - crossSize

	// §9.6.16 align flex lines, 'align-content'.
	if remFree > 0 {
		spacing, off := 0.0, 0.0
		switch f.AlignContent {
		case AlignContentStart:
			// already laid out correctly
		case AlignContentEnd:
			off = remFree
		case AlignContentCenter:
			off = remFree / 2
		case AlignContentSpaceBetween:
			spacing = remFree / float64(len(lines)-1)
		case AlignContentSpaceAround:
			spacing = remFree / float64(len(lines))
			off = spacing / 2
		}
		if f.AlignContent != AlignContentStart {
			for l := range lines {
				line := &lines[l]
				line.crossOffset += off
				for _, child := range line.child {
					child.crossOffset += off
				}
				off += spacing
			}
		}
	}

	// Layout complete. Update children position
	for l := range lines {
		line := &lines[l]
		for _, child := range line.child {
			switch f.Direction {
			case Row:
				child.node.Bounds = image.Rect(round(child.mainOffset),
					round(child.crossOffset),
					round(child.mainOffset+child.mainSize),
					round(child.crossOffset+child.crossSize))
				container, ok := child.node.Item.(Container)
				if ok && container != nil {
					container.SetFrame(child.node.Bounds.Add(f.frame.Min))
				}
			case Column:
				child.node.Bounds = image.Rect(round(child.crossOffset),
					round(child.mainOffset),
					round(child.crossOffset+child.crossSize),
					round(child.mainOffset+child.mainSize))
				container, ok := child.node.Item.(Container)
				if ok && container != nil {
					container.SetFrame(child.node.Bounds.Add(f.frame.Min))
				}
			default:
				panic(fmt.Sprint("flex: bad direction ", f.Direction))
			}
		}
	}
}

type element struct {
	node         *container.Child
	flexBaseSize float64
	mainSize     float64
	mainOffset   float64
	crossSize    float64
	crossOffset  float64
}

type flexLine struct {
	mainSize    float64
	crossSize   float64
	crossOffset float64
	child       []*element
}

func (f *Flex) mainSize(x, y int) int {
	switch f.Direction {
	case Row:
		return x
	case Column:
		return y
	default:
		panic(fmt.Sprint("flex: bad direction ", f.Direction))
	}
}

func (f *Flex) crossSize(x, y int) int {
	switch f.Direction {
	case Row:
		return y
	case Column:
		return x
	default:
		panic(fmt.Sprint("flex: bad direction ", f.Direction))
	}
}

func (f *Flex) flexBaseSize(c *container.Child) int {
	fixedSizeC, _ := c.Item.(FixedSizeItem)
	if fixedSizeC != nil {
		return f.mainSize(fixedSizeC.Size())
	}
	panic("flexible size is not available for now")
}

func round(f float64) int {
	return int(math.Floor(f + .5))
}
