// Package italic is the two italic Roboto Mono faces.
package italic

import (
	"gioui.org/font"

	"github.com/vibrantgio/font/robotomono/italic/bold"
	"github.com/vibrantgio/font/robotomono/italic/normal"
)

var (
	Normal = font.Font{Typeface: "Roboto Mono", Style: font.Italic, Weight: font.Normal}
	Bold   = font.Font{Typeface: "Roboto Mono", Style: font.Italic, Weight: font.Bold}
)

// FontFaces returns the two italic faces, parsed lazily by their leaf
// packages.
func FontFaces() []font.FontFace {
	return []font.FontFace{normal.FontFace(), bold.FontFace()}
}
