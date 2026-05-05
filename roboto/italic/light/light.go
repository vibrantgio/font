package light

import (
	"sync"

	roboto "eliasnaur.com/font/roboto/robotolightitalic"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var Font = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Light}

var face struct {
	once  sync.Once
	value font.Face
}

func FontFace() font.FontFace {
	face.once.Do(func() {
		if value, err := opentype.Parse(roboto.TTF); err == nil {
			face.value = value
		} else {
			panic("failed to parse font")
		}
	})
	return font.FontFace{Font: Font, Face: face.value}
}
