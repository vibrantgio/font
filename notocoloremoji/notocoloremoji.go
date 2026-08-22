// Package notocoloremoji is the optional color-emoji face — Noto Color Emoji
// Regular, packaged as a Gio font face under the typeface name "Noto Color
// Emoji".
//
// # What it is for
//
// It is a fallback for emoji Roboto, Roboto Mono and Noto Sans Mono do not
// carry. Nothing in the design system names this typeface, and nothing needs
// to: the shaper reaches it only when the faces ahead of it have no glyph for
// a rune. It is not another JetBrains Mono.
//
// # It is optional, and that is the design
//
// It is deliberately not part of theme's DefaultTypography.Faces. Putting
// 9.9 MB in the default would parse that TTF on every golden and every pinned
// shaper in the organization, and no existing golden contains emoji. Append
// it when a document actually draws emoji:
//
//	typ := tokens.DefaultTypography.WithFaces(notocoloremoji.FontFace())
//
// # Coverage, measured
//
// The face is CBDT/PNG color emoji: one 109 ppem strike, format 17 PNG.
// opentype.Parse on Gio v0.10.2 succeeds. The glyphs have no outlines —
// shaper.Shape skips them; they paint only through shaper.Bitmaps.
//
// Single-codepoint probes present in the cmap: 😀 😂 😉 😍 🔥 🎉 ❤ ✨ 🚀.
// Latin letters are not: 'A' is absent, so a collection of only this face
// resolves a grin and does not resolve a letter.
//
// # ZWJ sequences
//
// go-text applies the face's GSUB. Shaping 👨‍👩‍👧‍👦 (U+1F468 ZWJ U+1F469
// ZWJ U+1F467 ZWJ U+1F466) on a NoSystemFonts shaper of this collection
// yields one glyph, not four; that glyph's GlyphData is a PNG. The same is
// true of flag, profession, skin-tone and regional-indicator sequences this
// face ligates. This package does not compose sequences itself.
//
// # Provenance
//
// The TTF is Noto Color Emoji Regular; SHA-256
// 2eeac855a08803c6d209f8eb74ed5f798af46e128bc93dd3913e04de57523a7c
// (10,353,636 bytes). The font is licensed under the SIL Open Font License
// 1.1 — see OFL.txt beside this file.
//
// # Layout
//
// The roboto and robotomono packages nest a leaf package per face, so that a
// program can link one weight of twelve or of four. This family is one weight,
// so the family package is the leaf: it embeds the TTF, parses it lazily
// behind a sync.Once, and offers both the leaf shape (Font, FontFace) and the
// family shape (FontFaces, one entry) the other packages have.
package notocoloremoji

import (
	_ "embed"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

//go:embed NotoColorEmoji.ttf
var ttf []byte

// Font is the one face this package carries.
var Font = font.Font{Typeface: "Noto Color Emoji", Style: font.Regular, Weight: font.Normal}

var face struct {
	once  sync.Once
	value font.Face
}

// FontFace returns the parsed face, parsing the embedded TTF on the first
// call. It panics if opentype.Parse rejects the bytes, which a build that
// linked cannot do at run time — the bytes are compiled in.
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

// FontFaces returns the family's faces — one, here — so that this package can
// stand in for roboto and robotomono wherever a collection is assembled.
func FontFaces() []font.FontFace {
	return []font.FontFace{FontFace()}
}
