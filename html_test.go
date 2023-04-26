package furex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected *View
		opts     *ParseOptions
		after    func(t *testing.T, v *View)
	}{
		{
			name: "simple",
			html: `
				<body>
					<view style="
						left: 50;
						top: 100;
						width: 200;
						height: 300;
						margin-left: 120;
						margin-top: 130;
						margin-right: 140;
						margin-bottom: 150;
						position: absolute;
						direction: row;
						flex-wrap: wrap;
						justify-content: flex-end;
						align-items: flex-end;
						align-content: center;
						flex-grow: 2;
						flex-shrink: 3;
					">
						<view style="width: 100; height: 200;" />
					</view>
				</body>`,
			expected: (&View{
				Left:         50,
				Top:          100,
				Width:        200,
				Height:       300,
				MarginLeft:   120,
				MarginTop:    130,
				MarginRight:  140,
				MarginBottom: 150,
				Position:     PositionAbsolute,
				Direction:    Row,
				Wrap:         Wrap,
				Justify:      JustifyEnd,
				AlignItems:   AlignItemEnd,
				AlignContent: AlignContentCenter,
				Grow:         2,
				Shrink:       3,
			}).AddChild(&View{
				Width:  100,
				Height: 200,
			}),
		},
		{
			name: "nested",
			html: `
				<view>
					<view>
						<view>
						</view>
					</view>
					<view>
						<view>
						</view>
					</view>
				</view>
						`,
			expected: (&View{}).
				AddChild(
					(&View{}).AddChild(
						(&View{}),
					),
					(&View{}).AddChild(
						(&View{}),
					),
				),
		},
		{
			name: "root width and height",
			html: `
				<view>
					<view>
						<view>
						</view>
					</view>
				</view>
						`,
			expected: (&View{
				Width:  200,
				Height: 300,
			}).
				AddChild(
					(&View{}).AddChild(
						(&View{}),
					),
				),
			opts: &ParseOptions{
				Width:  200,
				Height: 300,
			},
		},
		{
			name: "with-handlers",
			html: `
						<body>
							<mock-handler>
								<view />
							</mock-handler>
						</body>`,
			opts: &ParseOptions{
				Components: map[string]Component{
					"mock-handler": func() Handler {
						return &mockHandler{}
					},
				},
			},
			expected: (&View{}).AddChild((&View{})),
			after: func(t *testing.T, v *View) {
				v.Update()
				h := v.Handler.(*mockHandler)
				require.True(t, h.IsUpdated)
			},
		},
		{
			name: "hidden attribute",
			html: `
				<view>
					<view id="test" hidden>
					</view>
				</view>
						`,
			expected: (&View{
				Width:  200,
				Height: 300,
			}).AddChild((&View{})),
			opts: &ParseOptions{
				Width:  200,
				Height: 300,
			},
			after: func(t *testing.T, v *View) {
				elem, ok := v.GetByID("test")
				require.True(t, ok)
				require.Equal(t, true, elem.Hidden)
			},
		},
		{
			name: "complex",
			html: `
		<head>
		    <style>
		        game-ui {
		            flex-direction: column;
		            justify-content: space-between;
		            align-items: stretch;
		            align-content: stretch;
		        }

		        container {
		            justify-content: center;
		            align-items: center;
		            margin-top: 50px;
		            flex-grow: 1;
		        }

		        panel {
		            width: 300px;
		            height: 300px;
		            margin-top: 120px;
		            margin-left: 130px;
		            flex-direction: column;
		            align-items: center;
		            justify-content: center;
		        }

		        panel-inner {
		            flex-direction: column;
		            align-items: center;
		            justify-content: center;
		            margin-top: 20px;
		            width: 245px;
		            height: 200px;
		        }

		        gauge-container {
		            width: 180px;
		            height: 38px;
		            align-items: flex-start;
		            justify-content: flex-start;
		            flex-direction: column;
		        }

		        gauge-container2 {
		            width: 180px;
		            height: 38px;
		            align-items: flex-start;
		            justify-content: flex-start;
		            flex-direction: column;
					margin-top: 30px;
		        }

		        .gauge-text {
		            width: 180px;
		            height: 20px;
		            margin-bottom: 2px;
		        }

		        .gauge {
		            width: 180px;
		            height: 18px;
		        }

		        buttons {
		            flex-direction: row;
		            align-items: center;
		            justify-content: center;
		            margin-top: 20px;
		            margin-bottom: 20px;
		            flex-grow: 1;
		        }

		        button-inventory {
		            width: 190px;
		            height: 49px;
		        }

		        button-ok {
		            width: 45px;
		            height: 49px;
		            margin-left: 10;
		        }

		        bottom-buttons {
		            justify-content: center;
		            align-items: flex-end;
		            margin-bottom: 20px;
		        }

		        .bottom-button {
		            width: 45px;
		            height: 49px;
		            margin-top: 5px;
		            margin-bottom: 10px;
		            margin-left: 20px;
		            margin-right: 20px;
		        }

		        play-game-container {
		            position: absolute;
		            left: 20;
		            top: 52;
		        }

		        play-game-inner-panel {
		            width: 260px;
		            height: 140px;
		            flex-direction: column;
		            align-items: center;
		            justify-content: center;
		        }

		        play-game-text {
		            width: 100px;
		            height: 8px;
		            margin-bottom: 20px;
		            flex-direction: row;
		            align-items: center;
		            justify-content: center;
		        }

		        play-game-buttons {
		            width: 100;
		            height: 50;
		            flex-direction: row;
		            align-items: center;
		            justify-content: center;
		        }

		        .play-game-button {
		            width: 100px;
		            height: 50px;
		            margin-left: 20px;
		        }

		        close-button {
		            position: absolute;
		            left: 283px;
		            top: -15;
		            width: 35px;
		            height: 38px
		        }

		        close-button-sprite {
		            position: absolute;
		            left: 18;
		            top: 17
		        }
		    </style>
		</head>

		<body>
		    <game-ui>
		        <container>
		            <panel>
		                <panel-inner>
		                    <gauge-container>
		                        <health-text class="gauge-text"></health-text>
		                        <health-gauge class="gauge"></health-gauge>
		                    </gauge-container>
		                    <gauge-container2>
		                        <mana-text class="gauge-text"></mana-text>
		                        <mana-gauge class="gauge"></mana-gauge>
		                    </gauge-container2>
		                </panel-inner>
		                <buttons>
		                    <button-inventory></button-inventory>
		                    <button-ok></button-ok>
		                </buttons>
		                <close-button>
		                    <close-button-sprite></close-button-sprite>
		                </close-button>
		            </panel>
		        </container>
		    </game-ui>
		</body>

		</html>
					`,
			opts: &ParseOptions{
				Width:  640,
				Height: 800,
			},
			expected: (&View{
				Width:        640,
				Height:       800,
				Direction:    Column,
				Justify:      JustifySpaceBetween,
				AlignItems:   AlignItemStretch,
				AlignContent: AlignContentStretch,
			}).AddChild(
				(&View{
					MarginTop:  50,
					Grow:       1,
					AlignItems: AlignItemCenter,
					Justify:    JustifyCenter,
				}).AddChild(
					(&View{
						Width:      300,
						Height:     300,
						MarginTop:  120,
						MarginLeft: 130,
						Direction:  Column,
						AlignItems: AlignItemCenter,
						Justify:    JustifyCenter,
					}).AddChild(
						(&View{
							MarginTop:  20,
							Width:      245,
							Height:     200,
							Direction:  Column,
							AlignItems: AlignItemCenter,
							Justify:    JustifyCenter,
						}).AddChild(
							(&View{
								Width:      180,
								Height:     38,
								AlignItems: AlignItemStart,
								Justify:    JustifyStart,
								Direction:  Column,
							}).AddChild(
								&View{
									Height:       20,
									Width:        180,
									MarginBottom: 2,
								},
								&View{
									Width:  180,
									Height: 18,
								},
							),

							(&View{
								Width:      180,
								Height:     38,
								AlignItems: AlignItemStart,
								Justify:    JustifyStart,
								Direction:  Column,
								MarginTop:  30,
							}).AddChild(
								&View{
									Height:       20,
									Width:        180,
									MarginBottom: 2,
								},
								&View{
									Width:  180,
									Height: 18,
								},
							),
						),
						(&View{
							MarginTop:    20,
							MarginBottom: 20,
							Grow:         1,
							Direction:    Row,
							AlignItems:   AlignItemCenter,
							Justify:      JustifyCenter,
						}).AddChild(
							&View{
								Width:  190,
								Height: 49,
							},
							&View{
								Width:      45,
								Height:     49,
								MarginLeft: 10,
							},
						),
						(&View{
							Position: PositionAbsolute,
							Left:     300 - 35/2,
							Top:      4 - 38/2,
							Width:    35,
							Height:   38,
						}).AddChild(
							&View{
								Position: PositionAbsolute,
								Left:     18,
								Top:      17,
							},
						),
					),
				),
			)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Parse(tt.html, tt.opts)
			testViewStyle(t, v, tt.expected)
		})
	}
}

func testViewStyle(t *testing.T, v *View, expected *View) {
	t.Helper()
	require.Equal(t, expected.Config(), v.Config())
}
