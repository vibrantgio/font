// Package regular is the two upright Roboto Mono faces.
package regular

import (
	"gioui.org/font"

	"github.com/vibrantgio/font/robotomono/regular/bold"
	"github.com/vibrantgio/font/robotomono/regular/normal"
)

var (
	Normal = font.Font{Typeface: "Roboto Mono", Style: font.Regular, Weight: font.Normal}
	Bold   = font.Font{Typeface: "Roboto Mono", Style: font.Regular, Weight: font.Bold}
)

// FontFaces returns the two upright faces, parsed lazily by their leaf
// packages.
func FontFaces() []font.FontFace {
	return []font.FontFace{normal.FontFace(), bold.FontFace()}
}
