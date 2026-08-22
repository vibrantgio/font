// Package normal is the JetBrains Mono upright regular face.
//
// The TTF is the static instance from the JetBrains Mono project,
// https://github.com/JetBrains/JetBrainsMono (fonts/ttf/JetBrainsMono-Regular.ttf),
// licensed under the SIL Open Font License 1.1 — see OFL.txt in the
// jetbrainsmono package root.
package normal

import (
	_ "embed"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

//go:embed JetBrainsMono-Regular.ttf
var ttf []byte

var Font = font.Font{Typeface: "JetBrains Mono", Style: font.Regular, Weight: font.Normal}

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
