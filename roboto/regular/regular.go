package regular

import (
	"fmt"
	"sync"

	"eliasnaur.com/font/roboto/robotoblack"
	"eliasnaur.com/font/roboto/robotobold"
	"eliasnaur.com/font/roboto/robotolight"
	"eliasnaur.com/font/roboto/robotomedium"
	"eliasnaur.com/font/roboto/robotoregular"
	"eliasnaur.com/font/roboto/robotothin"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	Thin   = font.Font{Typeface: "Roboto", Variant: "", Style: font.Regular, Weight: font.Thin}
	Light  = font.Font{Typeface: "Roboto", Variant: "", Style: font.Regular, Weight: font.Light}
	Normal = font.Font{Typeface: "Roboto", Variant: "", Style: font.Regular, Weight: font.Normal}
	Medium = font.Font{Typeface: "Roboto", Variant: "", Style: font.Regular, Weight: font.Medium}
	Bold   = font.Font{Typeface: "Roboto", Variant: "", Style: font.Regular, Weight: font.Bold}
	Black  = font.Font{Typeface: "Roboto", Variant: "", Style: font.Regular, Weight: font.Black}
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
		register(Normal, robotoregular.TTF)
		register(Thin, robotothin.TTF)
		register(Light, robotolight.TTF)
		register(Medium, robotomedium.TTF)
		register(Bold, robotobold.TTF)
		register(Black, robotoblack.TTF)
	})
	n := len(fontfaces.collection)
	return fontfaces.collection[0:n:n]
}
