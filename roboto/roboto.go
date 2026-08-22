// Package roboto is the Roboto family packaged as Gio font faces under the
// typeface name "Roboto".
//
// The TTFs are embedded in this repository and are licensed under the Apache
// License 2.0 — see LICENSE beside this file.
package roboto

import (
	"gioui.org/font"

	"github.com/vibrantgio/font/roboto/italic"
	"github.com/vibrantgio/font/roboto/regular"
)

var (
	RegularThin   = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Thin}
	RegularLight  = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Light}
	RegularNormal = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Normal}
	RegularMedium = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Medium}
	RegularBold   = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Bold}
	RegularBlack  = font.Font{Typeface: "Roboto", Style: font.Regular, Weight: font.Black}

	ItalicThin   = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Thin}
	ItalicLight  = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Light}
	ItalicNormal = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Normal}
	ItalicMedium = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Medium}
	ItalicBold   = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Bold}
	ItalicBlack  = font.Font{Typeface: "Roboto", Style: font.Italic, Weight: font.Black}
)

// FontFaces returns all twelve faces, parsed lazily by their leaf packages.
func FontFaces() []font.FontFace {
	return append(regular.FontFaces(), italic.FontFaces()...)
}
