// Package thin is the Roboto italic thin face.
//
// The TTF is embedded from this repository — see LICENSE in the roboto
// package root (Apache License 2.0).
package thin

import (
	_ "embed"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

// TTF is the embedded TrueType font file.
//
//go:embed Roboto-ThinItalic.ttf
var TTF []byte

var Font = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Thin}

var face struct {
	once  sync.Once
	value font.Face
}

func FontFace() font.FontFace {
	face.once.Do(func() {
		if value, err := opentype.Parse(TTF); err == nil {
			face.value = value
		} else {
			panic("failed to parse font")
		}
	})
	return font.FontFace{Font: Font, Face: face.value}
}
