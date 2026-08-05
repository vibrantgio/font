// Package bold is the Roboto Mono italic bold face.
//
// The TTF is the static instance from the Roboto Mono project,
// https://github.com/googlefonts/RobotoMono (fonts/ttf/RobotoMono-BoldItalic.ttf),
// licensed under the SIL Open Font License 1.1 — see OFL.txt in the
// robotomono package root.
package bold

import (
	_ "embed"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

//go:embed RobotoMono-BoldItalic.ttf
var ttf []byte

var Font = font.Font{Typeface: "Roboto Mono", Style: font.Italic, Weight: font.Bold}

var face struct {
	once  sync.Once
	value font.Face
}

func FontFace() font.FontFace {
	face.once.Do(func() {
		if value, err := opentype.Parse(ttf); err == nil {
			face.value = value
		} else {
			panic("failed to parse font")
		}
	})
	return font.FontFace{Font: Font, Face: face.value}
}
