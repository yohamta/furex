package text

import "github.com/tinne26/etxt"

var (
	R    *etxt.Renderer
	Font = "x14y20pxScoreDozer"
)

func LoadFonts() {
	fontLib := etxt.NewFontLibrary()
	_, _, err := fontLib.ParseDirFonts("./assets/fonts")
	if err != nil {
		panic(err)
	}

	// create a new text renderer and configure it
	R = etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024) // 10MB
	R.SetCacheHandler(glyphsCache.NewHandler())
	R.SetFont(fontLib.GetFont(Font))
	R.SetAlign(etxt.YCenter, etxt.XCenter)
	R.SetSizePx(14)
}
