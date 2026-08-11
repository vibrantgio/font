// Package notosansmono is the optional symbol face — Noto Sans Mono Regular,
// packaged as a Gio font face under the typeface name "Noto Sans Mono".
//
// # What it is for
//
// It is a fallback of last resort for glyphs Roboto and Roboto Mono do not
// carry: arrows, box drawing, block elements, geometric shapes, and the
// punctuation and mathematical operators outside Latin-1. Nothing in the
// design system names this typeface, and nothing needs to: the shaper reaches
// it only when the faces ahead of it have no glyph for a rune.
//
// # It is optional, and that is the design
//
// It is deliberately not part of theme's DefaultTypography.Faces. The
// default shaper leaves system fonts enabled, so on a normal desktop the
// platform already covers these characters — and covers far more of them than
// one embedded face ever could, emoji and CJK included. Paying 596 KB in every
// binary to duplicate what the OS already has would be a poor trade.
//
// Append it when the platform is not there to fall back to: a container, a
// kiosk, a scratch image, anything shipping its own world. One line does it,
// on the typography token rather than by rebuilding the collection:
//
//	typ := tokens.DefaultTypography.WithFaces(notosansmono.FontFace())
//
// The same line is how a golden test that legitimately draws an arrow keeps
// its faces pinned — see Typography.DeterministicShaper.
//
// # Coverage, measured
//
// Counted against the embedded TTF's cmap, glyphs present per Unicode block:
//
//	Box Drawing              128/128     Arrows                    23/112
//	Block Elements            32/32      Mathematical Operators   104/256
//	Geometric Shapes          96/96      Miscellaneous Technical  118/256
//	General Punctuation      111/112     Superscripts/Subscripts   42/48
//	Letterlike Symbols        80/80      Currency Symbols          32/32
//
// The arrow and operator blocks are partial by design of the face, not by
// subsetting: the 23 arrows are the cardinal, double and long-tailed forms
// (← ↑ → ↓ ↔ ↕ ⇐ ⇑ ⇒ ⇓ ⇔ and kin), and the 104 operators are the ones prose
// uses (− ≤ ≥ ≈ ≠ ≡ ∞ √ ∂ ∈ ⊕ …), with × ÷ ± · ° coming from Latin-1 rather
// than this block. What is *not* here: diagonal arrows (↗ ↘), the large
// operators (∑ ∏ ∫), dingbats (✓ ✗ ★), emoji and CJK. On a desktop the system
// fallback serves those, and a kiosk that needs them ships a face for them
// beside this one — WithFaces takes as many as you give it.
//
// # Provenance
//
// The TTF is the static Regular instance of Noto Sans Mono v2.014, from the
// Noto project's own release,
// https://github.com/notofonts/latin-greek-cyrillic/releases/tag/NotoSansMono-v2.014
// (NotoSansMono/googlefonts/ttf/NotoSansMono-Regular.ttf), retrieved
// 2026-08-06; SHA-256
// 74fd536351d0f30a73410e1bd223a0ebf763dd4808eb844aa2aed18c0d6e3c84. Google
// Fonts distributes the same family under ofl/notosansmono, as a variable font
// carrying nine weights and four widths — the static Regular is taken instead,
// because one weight is all a fallback needs. The font is licensed under the
// SIL Open Font License 1.1 — see OFL.txt beside this file.
//
// # Layout
//
// The roboto and robotomono packages nest a leaf package per face, so that a
// program can link one weight of twelve or of four. This family is one weight,
// so the family package is the leaf: it embeds the TTF, parses it lazily
// behind a sync.Once, and offers both the leaf shape (Font, FontFace) and the
// family shape (FontFaces, one entry) the other two packages have.
package notosansmono

import (
	_ "embed"
	"sync"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

//go:embed NotoSansMono-Regular.ttf
var ttf []byte

// Font is the one face this package carries.
var Font = font.Font{Typeface: "Noto Sans Mono", Style: font.Regular, Weight: font.Normal}

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
