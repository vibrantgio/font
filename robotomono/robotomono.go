// Package robotomono is Roboto Mono — Roboto's companion monospace face —
// packaged as Gio font faces under the typeface name "Roboto Mono".
//
// The family mirrors the roboto package's per-weight layout, but only as far
// as real use: the markdown/highlight path shapes code at normal and bold, in
// upright and italic, so exactly those four faces are packaged — no thin,
// light, medium or black.
//
// The TTFs are the static instances from the Roboto Mono project,
// https://github.com/googlefonts/RobotoMono (fonts/ttf/), retrieved from the
// main branch on 2026-08-05. Google Fonts distributes the same family under
// ofl/robotomono. The fonts are licensed under the SIL Open Font License 1.1
// — see OFL.txt beside this file. (Roboto Mono was Apache-2.0-licensed
// historically; the project relicensed to the OFL, and google/fonts moved it
// from apache/ to ofl/ accordingly.)
package robotomono

import (
	"gioui.org/font"

	italicbold "github.com/vibrantgio/font/robotomono/italic/bold"
	italicnormal "github.com/vibrantgio/font/robotomono/italic/normal"
	regularbold "github.com/vibrantgio/font/robotomono/regular/bold"
	regularnormal "github.com/vibrantgio/font/robotomono/regular/normal"
)

var (
	RegularNormal = font.Font{Typeface: "Roboto Mono", Style: font.Regular, Weight: font.Normal}
	RegularBold   = font.Font{Typeface: "Roboto Mono", Style: font.Regular, Weight: font.Bold}

	ItalicNormal = font.Font{Typeface: "Roboto Mono", Style: font.Italic, Weight: font.Normal}
	ItalicBold   = font.Font{Typeface: "Roboto Mono", Style: font.Italic, Weight: font.Bold}
)

// FontFaces returns all four faces, parsed lazily by their leaf packages.
func FontFaces() []font.FontFace {
	return []font.FontFace{
		regularnormal.FontFace(),
		regularbold.FontFace(),
		italicnormal.FontFace(),
		italicbold.FontFace(),
	}
}
