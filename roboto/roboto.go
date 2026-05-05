package roboto

import (
	"fmt"
	"sync"

	"eliasnaur.com/font/roboto/robotoblack"
	"eliasnaur.com/font/roboto/robotobold"
	"eliasnaur.com/font/roboto/robotolight"
	"eliasnaur.com/font/roboto/robotomedium"
	"eliasnaur.com/font/roboto/robotoregular"
	"eliasnaur.com/font/roboto/robotothin"

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
	RegularThin   = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Thin}
	RegularLight  = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Light}
	RegularNormal = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Normal}
	RegularMedium = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Medium}
	RegularBold   = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Bold}
	RegularBlack  = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Black}

	ItalicThin   = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Thin}
	ItalicLight  = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Light}
	ItalicNormal = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Normal}
	ItalicMedium = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Medium}
	ItalicBold   = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Bold}
	ItalicBlack  = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Black}
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
		register(RegularNormal, robotoregular.TTF)
		register(RegularThin, robotothin.TTF)
		register(RegularLight, robotolight.TTF)
		register(RegularMedium, robotomedium.TTF)
		register(RegularBold, robotobold.TTF)
		register(RegularBlack, robotoblack.TTF)

		register(ItalicNormal, robotoitalic.TTF)
		register(ItalicThin, robotothinitalic.TTF)
		register(ItalicLight, robotolightitalic.TTF)
		register(ItalicMedium, robotomediumitalic.TTF)
		register(ItalicBold, robotobolditalic.TTF)
		register(ItalicBlack, robotoblackitalic.TTF)
	})
	n := len(fontfaces.collection)
	return fontfaces.collection[0:n:n]
}
