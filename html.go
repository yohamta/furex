package furex

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/vanng822/go-premailer/premailer"
	"golang.org/x/net/html"
)

// The Component can be either a handler instance (e.g., DrawHandler) or
// a factory function func() furex.Handler. This allows flexibility in usage:
// If you want to reuse the same handler instance for multiple HTML tags, pass the instance;
// otherwise, pass the factory function to create separate handler instances for each tag.
type Component interface{}

// ComponentsMap is a type alias for a dictionary that associates
// custom HTML tags with their respective components.
// This enables a convenient way to manage and reference components
// based on their corresponding HTML tags.
type ComponentsMap map[string]Component

// ParseOptions represents options for parsing HTML.
type ParseOptions struct {
	// Components is a map of component name and handler.
	// For example, if you have a component named "start-button", you can define a handler
	// for it like this:
	// 	opts.Components := map[string]Handler{
	// 		"start-button": <your handler>,
	//  }
	// The component name must be in kebab-case.
	// You can use the component in your HTML like this:
	// 	<start-button></start-button>
	// Note: self closing tag is not supported.
	Components ComponentsMap
	// Width and Height is the size of the root view.
	// This is useful when you want to specify the width and height
	// outside of the HTML.
	Width  int
	Height int
}

func Parse(input string, opts *ParseOptions) *View {
	if opts == nil {
		opts = &ParseOptions{}
	}

	inlinedHTML := inlineCSS(input)
	z := html.NewTokenizer(strings.NewReader(inlinedHTML))
	dummy := &View{}
	stack := &stack{stack: []*View{dummy}}
	depth := 0
	inBody := false
Loop:
	for {
		tt := z.Next()
		tn, _ := z.TagName()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				break Loop
			}
			panic(z.Err())
		case html.StartTagToken:
			if string(tn) == "body" {
				inBody = true
				continue
			}
			if !inBody {
				continue
			}
			view := processTag(z, string(tn), opts, depth)
			if view == nil {
				continue
			}
			stack.peek().AddChild(view)
			stack.push(view)

			depth++
		case html.SelfClosingTagToken:
			view := processTag(z, string(tn), opts, depth)
			if view == nil {
				continue
			}
			stack.peek().AddChild(view)
		case html.TextToken:
			if stack.len() > 0 {
				stack.peek().Text = strings.TrimSpace(string(z.Text()))
			}
		case html.EndTagToken:
			if string(tn) == "body" {
				inBody = false
				continue
			}
			if !inBody {
				continue
			}
			stack.pop()
			depth--
		}
	}
	if len(dummy.children) != 1 {
		panic(fmt.Sprintf("invalid html: %s", input))
	}
	return dummy.PopChild()
}

func inlineCSS(doc string) string {
	prem, err := premailer.NewPremailerFromString(doc, &premailer.Options{})
	if err != nil {
		println(fmt.Errorf("invalid css: %s", err))
		return doc
	}
	html, err := prem.Transform()
	if err != nil {
		println(fmt.Errorf("error transform html: %s", err))
		return doc
	}
	return html
}

type stack struct {
	stack []*View
}

func (s *stack) push(v *View) {
	s.stack = append(s.stack, v)
}

func (s *stack) len() int {
	return len(s.stack)
}

func (s *stack) peek() *View {
	return s.stack[len(s.stack)-1]
}

func (s *stack) pop() *View {
	v := s.peek()
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func processTag(z *html.Tokenizer, tagName string, opts *ParseOptions, depth int) *View {
	attrs := readAttrs(z)
	view := &View{}

	if depth == 0 {
		if opts.Width != 0 {
			view.Width = opts.Width
		}
		if opts.Height != 0 {
			view.Height = opts.Height
		}
	}

	parseStyle(view, attrs.style, opts.Components)
	view.ID = attrs.id
	view.Raw = string(z.Raw())
	view.Attrs = attrs.miscs

	if c, ok := opts.Components[tagName]; ok {
		if reflect.TypeOf(c).Kind() == reflect.Func {
			if c, ok := c.(func() Handler); ok {
				view.Handler = c()
			} else {
				panic(fmt.Sprintf("invalid component: %s", tagName))
			}
		} else {
			view.Handler = c
		}
	}

	return view
}

func parseStyle(view *View, style string, handlers ComponentsMap) {
	pairs := strings.Split(style, ";")
	errs := &ErrorList{}
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		if len(kv) != 2 {
			continue
		}
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])

		mapper, ok := styleMapper[k]
		if !ok {
			errs.Add(fmt.Errorf("unknown style: %s", k))
			continue
		}
		parsed, err := mapper.parseFunc(v)
		if err != nil {
			errs.Add(err)
			continue
		}
		mapper.setFunc(view, parsed)
	}
	if errs.HasErrors() {
		println(fmt.Sprintf("parse style errors: %v", errs))
	}
}

var styleMapper = map[string]mapper[View]{
	"left": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.Left = val }),
	},
	"top": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.Top = val }),
	},
	"width": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.Width = val }),
	},
	"height": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.Height = val }),
	},
	"margin-left": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.MarginLeft = val }),
	},
	"margin-top": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.MarginTop = val }),
	},
	"margin-right": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.MarginRight = val }),
	},
	"margin-bottom": {
		parseFunc: parseNumber,
		setFunc:   setFunc(func(v *View, val int) { v.MarginBottom = val }),
	},
	"position": {
		parseFunc: parsePosition,
		setFunc:   setFunc(func(v *View, val Position) { v.Position = val }),
	},
	"direction": {
		parseFunc: parseDirection,
		setFunc:   setFunc(func(v *View, val Direction) { v.Direction = val }),
	},
	"flex-direction": {
		parseFunc: parseDirection,
		setFunc:   setFunc(func(v *View, val Direction) { v.Direction = val }),
	},
	"flex-wrap": {
		parseFunc: parseWrap,
		setFunc:   setFunc(func(v *View, val FlexWrap) { v.Wrap = val }),
	},
	"wrap": {
		parseFunc: parseWrap,
		setFunc:   setFunc(func(v *View, val FlexWrap) { v.Wrap = val }),
	},
	"justify": {
		parseFunc: parseJustify,
		setFunc:   setFunc(func(v *View, val Justify) { v.Justify = val }),
	},
	"justify-content": {
		parseFunc: parseJustify,
		setFunc:   setFunc(func(v *View, val Justify) { v.Justify = val }),
	},
	"align-items": {
		parseFunc: parseAlignItem,
		setFunc:   setFunc(func(v *View, val AlignItem) { v.AlignItems = val }),
	},
	"align-content": {
		parseFunc: parseAlignContent,
		setFunc:   setFunc(func(v *View, val AlignContent) { v.AlignContent = val }),
	},
	"flex-grow": {
		parseFunc: parseFloat,
		setFunc:   setFunc(func(v *View, val float64) { v.Grow = val }),
	},
	"grow": {
		parseFunc: parseFloat,
		setFunc:   setFunc(func(v *View, val float64) { v.Grow = val }),
	},
	"flex-shrink": {
		parseFunc: parseFloat,
		setFunc:   setFunc(func(v *View, val float64) { v.Shrink = val }),
	},
	"shrink": {
		parseFunc: parseFloat,
		setFunc:   setFunc(func(v *View, val float64) { v.Shrink = val }),
	},
	"display": {
		parseFunc: parseString,
		setFunc:   setFunc(func(v *View, val string) { /* noop */ }),
	},
}

// setFunc creates a function that takes an entity and a value as an interface{}.
// The created function type asserts the value to the correct type U and then calls
// the given function f with the entity and the value of type U.
// If the type U is a pointer, the function will handle it accordingly:
// - If the input value is nil, it will pass a nil value of type U to the given function f.
// - If the input value is non-nil, it will create a pointer to the value, type-assert it to U, and pass it to f.
func setFunc[T, U any](f func(entity T, value U)) func(T, any) {
	return func(e T, v any) {
		argType := reflect.TypeOf(f).In(1)

		// If the value v is nil
		if v == nil {
			if argType.Kind() == reflect.Ptr {
				// pass nil value if the type is pointer
				nilValue := reflect.Zero(argType).Interface().(U)
				f(e, nilValue)
			} else {
				// pass deafult value if the type is not pointer
				var u U
				defaultValue := reflect.Zero(reflect.TypeOf(u))
				f(e, defaultValue.Interface().(U))
			}
			return
		}

		// If the second input parameter type of f is a pointer
		if argType.Kind() == reflect.Ptr {
			// Create a pointer to the value and set its value to v
			valuePtr := reflect.New(reflect.TypeOf(v)).Elem()
			valuePtr.Set(reflect.ValueOf(v))

			// If the pointer created can be type asserted to U, call f with e and the pointer
			if ptr, ok := valuePtr.Addr().Interface().(U); ok {
				f(e, ptr)
				return
			}
		}

		// If the value v is of type U, call f directly with e and v
		if v, ok := v.(U); ok {
			f(e, v)
			return
		}

		// If the type of the value is incorrect, panic with an error message
		var u U
		panic(fmt.Sprintf("type of the value is incorrect: %v, %T vs %T", v, v, u))
	}
}

type mapper[T any] struct {
	parseFunc func(string) (any, error)
	setFunc   func(*T, any)
}

func parseNumber(val string) (any, error) {
	val = strings.TrimSuffix(val, "px")
	return strconv.Atoi(val)
}

func parseFloat(val string) (any, error) {
	return strconv.ParseFloat(val, 64)
}

func parsePosition(val string) (any, error) {
	switch val {
	case "absolute":
		return PositionAbsolute, nil
	case "static", "relative":
		return PositionStatic, nil
	}
	return PositionStatic, fmt.Errorf("unknown position: %s", val)
}

func parseDirection(val string) (any, error) {
	switch val {
	case "row":
		return Row, nil
	case "column":
		return Column, nil
	}
	return Column, fmt.Errorf("unknown direction: %s", val)
}

func parseWrap(val string) (any, error) {
	switch val {
	case "wrap":
		return Wrap, nil
	case "nowrap":
		return NoWrap, nil
	}
	return NoWrap, fmt.Errorf("unknown wrap: %s", val)
}

func parseJustify(val string) (any, error) {
	switch val {
	case "flex-start", "start":
		return JustifyStart, nil
	case "flex-end", "end":
		return JustifyEnd, nil
	case "space-between":
		return JustifySpaceBetween, nil
	case "space-around":
		return JustifySpaceAround, nil
	case "center":
		return JustifyCenter, nil
	}
	return JustifyStart, fmt.Errorf("unknown justify: %s", val)
}

func parseAlignItem(val string) (any, error) {
	switch val {
	case "flex-start", "start":
		return AlignItemStart, nil
	case "flex-end", "end":
		return AlignItemEnd, nil
	case "center":
		return AlignItemCenter, nil
	case "stretch":
		return AlignItemStretch, nil
	}
	return AlignItemStretch, fmt.Errorf("unknown align-items: %s", val)
}

func parseAlignContent(val string) (any, error) {
	switch val {
	case "flex-start", "start":
		return AlignContentStart, nil
	case "flex-end", "end":
		return AlignContentEnd, nil
	case "center":
		return AlignContentCenter, nil
	case "stretch":
		return AlignContentStretch, nil
	case "space-between":
		return AlignContentSpaceBetween, nil
	case "space-around":
		return AlignContentSpaceAround, nil
	}
	return AlignContentStart, fmt.Errorf("unknown align-content: %s", val)
}

func parseString(val string) (any, error) {
	return val, nil
}

type attrs struct {
	id    string
	style string
	miscs map[string]string
}

func readAttrs(z *html.Tokenizer) attrs {
	attr := attrs{
		miscs: make(map[string]string),
	}
	for {
		key, val, more := z.TagAttr()
		attr.miscs[string(key)] = string(val)
		switch string(key) {
		case "id":
			attr.id = string(val)
		case "style":
			attr.style = string(val)
		}
		if !more {
			break
		}
	}
	return attr
}
