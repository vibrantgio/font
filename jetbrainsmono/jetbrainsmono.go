// Package jetbrainsmono is JetBrains Mono — a monospace face designed for
// reading code — packaged as Gio font faces under the typeface name
// "JetBrains Mono".
//
// The family mirrors the robotomono package's layout: the four faces the
// markdown/highlight path shapes code with — normal and bold, in upright and
// italic.
//
// The TTFs are the static instances from the JetBrains Mono project,
// https://github.com/JetBrains/JetBrainsMono (fonts/ttf/), retrieved from the
// master branch on 2026-08-21. The fonts are licensed under the SIL Open
// Font License 1.1 — see OFL.txt beside this file.
package jetbrainsmono

import (
	"gioui.org/font"

	italicbold "github.com/vibrantgio/font/jetbrainsmono/italic/bold"
	italicnormal "github.com/vibrantgio/font/jetbrainsmono/italic/normal"
	regularbold "github.com/vibrantgio/font/jetbrainsmono/regular/bold"
	regularnormal "github.com/vibrantgio/font/jetbrainsmono/regular/normal"
)

var (
	RegularNormal = font.Font{Typeface: "JetBrains Mono", Style: font.Regular, Weight: font.Normal}
	RegularBold   = font.Font{Typeface: "JetBrains Mono", Style: font.Regular, Weight: font.Bold}

	ItalicNormal = font.Font{Typeface: "JetBrains Mono", Style: font.Italic, Weight: font.Normal}
	ItalicBold   = font.Font{Typeface: "JetBrains Mono", Style: font.Italic, Weight: font.Bold}
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
