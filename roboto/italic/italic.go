package italic

import (
	"fmt"
	"sync"

	"eliasnaur.com/font/roboto/robotoblackitalic"
	"eliasnaur.com/font/roboto/robotobolditalic"
	"eliasnaur.com/font/roboto/robotoitalic"
	"eliasnaur.com/font/roboto/robotolightitalic"
	"eliasnaur.com/font/roboto/robotomediumitalic"
	"eliasnaur.com/font/roboto/robotothinitalic"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	Thin   = font.Font{Typeface: "Roboto", Variant: "", Style: font.Italic, Weight: font.Thin}
	Light  = font.Font{Typeface: "Roboto", Variant: "", Style: font.Italic, Weight: font.Light}
	Normal = font.Font{Typeface: "Roboto", Variant: "", Style: font.Italic, Weight: font.Normal}
	Medium = font.Font{Typeface: "Roboto", Variant: "", Style: font.Italic, Weight: font.Medium}
	Bold   = font.Font{Typeface: "Roboto", Variant: "", Style: font.Italic, Weight: font.Bold}
	Black  = font.Font{Typeface: "Roboto", Variant: "", Style: font.Italic, Weight: font.Black}
)

var fontfaces struct {
	once       sync.Once
	collection []font.FontFace
}

func FontFaces() []font.FontFace {
	register := func(f font.Font, ttf []byte) {
		face, err := opentype.Parse(ttf)
		if err != nil {
			panic(fmt.Sprintf("failed to parse font: %v", err))
		}
		fontfaces.collection = append(fontfaces.collection, font.FontFace{Font: f, Face: face})
	}
	fontfaces.once.Do(func() {
		register(Normal, robotoitalic.TTF)
		register(Thin, robotothinitalic.TTF)
		register(Light, robotolightitalic.TTF)
		register(Medium, robotomediumitalic.TTF)
		register(Bold, robotobolditalic.TTF)
		register(Black, robotoblackitalic.TTF)
	})
	n := len(fontfaces.collection)
	return fontfaces.collection[0:n:n]
}
