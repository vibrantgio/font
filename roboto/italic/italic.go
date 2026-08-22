package italic

import (
	"gioui.org/font"

	"github.com/vibrantgio/font/roboto/italic/black"
	"github.com/vibrantgio/font/roboto/italic/bold"
	"github.com/vibrantgio/font/roboto/italic/light"
	"github.com/vibrantgio/font/roboto/italic/medium"
	"github.com/vibrantgio/font/roboto/italic/normal"
	"github.com/vibrantgio/font/roboto/italic/thin"
)

var (
	Thin   = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Thin}
	Light  = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Light}
	Normal = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Normal}
	Medium = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Medium}
	Bold   = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Bold}
	Black  = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Black}
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
