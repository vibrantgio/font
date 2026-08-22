package regular

import (
	"gioui.org/font"

	"github.com/vibrantgio/font/roboto/regular/black"
	"github.com/vibrantgio/font/roboto/regular/bold"
	"github.com/vibrantgio/font/roboto/regular/light"
	"github.com/vibrantgio/font/roboto/regular/medium"
	"github.com/vibrantgio/font/roboto/regular/normal"
	"github.com/vibrantgio/font/roboto/regular/thin"
)

var (
	Thin   = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Thin}
	Light  = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Light}
	Normal = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Normal}
	Medium = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Medium}
	Bold   = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Bold}
	Black  = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Black}
)

// FontFaces returns all six faces, parsed lazily by their leaf packages.
func FontFaces() []font.FontFace {
	return []font.FontFace{
		normal.FontFace(),
		thin.FontFace(),
		light.FontFace(),
		medium.FontFace(),
		bold.FontFace(),
		black.FontFace(),
	}
}
