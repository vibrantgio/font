package notosansmono_test

import (
	"testing"

	giofont "gioui.org/font"
	"github.com/go-text/typesetting/font"

	"github.com/vibrantgio/font/notosansmono"
)

// TestFontFace parses the embedded TTF (FontFace panics if opentype.Parse
// rejects it) and asserts the face carries the metadata its name claims.
func TestFontFace(t *testing.T) {
	face := notosansmono.FontFace()
	if face.Font != notosansmono.Font {
		t.Errorf("FontFace().Font = %+v, want %+v", face.Font, notosansmono.Font)
	}
	if face.Font.Typeface != "Noto Sans Mono" {
		t.Errorf("Typeface = %q, want %q", face.Font.Typeface, "Noto Sans Mono")
	}
	if face.Font.Style != giofont.Regular || face.Font.Weight != giofont.Normal {
		t.Errorf("style/weight = %v/%v, want Regular/Normal", face.Font.Style, face.Font.Weight)
	}
	if face.Face == nil {
		t.Fatal("FontFace().Face is nil")
	}
	if again := notosansmono.FontFace(); again.Face != face.Face {
		t.Error("FontFace() parsed a second face instead of caching the first")
	}
}

// TestFontFaces asserts the family aggregate is the one leaf face, so that
// this package can stand in for roboto and robotomono where a collection is
// assembled.
func TestFontFaces(t *testing.T) {
	faces := notosansmono.FontFaces()
	if len(faces) != 1 {
		t.Fatalf("FontFaces() returned %d faces, want 1", len(faces))
	}
	if faces[0] != notosansmono.FontFace() {
		t.Errorf("FontFaces()[0] = %+v, want FontFace()'s %+v", faces[0].Font, notosansmono.FontFace().Font)
	}
}

// TestCoverage asserts the characters the package comment promises are in the
// face's cmap — the whole reason this face is packaged at all. It reads the
// cmap directly rather than shaping: what a shaper resolves depends on the
// collection and on the machine's fonts, and this assertion is about this file
// alone. The resolution half — that a shaper carrying this face returns real
// glyphs for these runes rather than the missing-glyph glyph — is asserted in
// theme/tokens, where the shaper lives.
func TestCoverage(t *testing.T) {
	parsed, ok := notosansmono.FontFace().Face.(interface{ Face() *font.Face })
	if !ok {
		t.Fatalf("face %T does not expose the parsed font", notosansmono.FontFace().Face)
	}
	f := parsed.Face()

	groups := []struct {
		name  string
		runes []rune
	}{
		{"arrows", []rune{'←', '↑', '→', '↓', '↔', '↕', '⇐', '⇑', '⇒', '⇓', '⇔'}},
		{"box drawing", []rune{'─', '│', '┌', '┐', '└', '┘', '├', '┤', '┬', '┴', '┼', '═', '║', '╭', '╮', '╯', '╰'}},
		{"block elements", []rune{'█', '▀', '▄', '░', '▒', '▓'}},
		{"geometric shapes", []rune{'■', '□', '▪', '●', '○', '▲', '▼', '◆'}},
		{"mathematical operators", []rune{'−', '≤', '≥', '≈', '≠', '≡', '∞', '√', '∂', '∅', '∆', '∇', '∈', '⊕', '⊗'}},
		{"latin-1 operators", []rune{'×', '÷', '±', '¬', '°', '·', 'µ'}},
		{"punctuation", []rune{'–', '—', '‘', '’', '“', '”', '•', '…', '‰', '†', '‡', '′', '″', '⁄'}},
	}
	for _, g := range groups {
		for _, r := range g.runes {
			if _, ok := f.NominalGlyph(r); !ok {
				t.Errorf("%s: U+%04X %q is not in the face's cmap", g.name, r, r)
			}
		}
	}

	// Counted per Unicode block, the numbers the package comment publishes.
	// A face revision that quietly narrowed one of them would leave the
	// documentation lying; this fails instead.
	blocks := []struct {
		name   string
		lo, hi rune
		want   int
	}{
		{"General Punctuation", 0x2000, 0x206F, 111},
		{"Superscripts and Subscripts", 0x2070, 0x209F, 42},
		{"Currency Symbols", 0x20A0, 0x20BF, 32},
		{"Letterlike Symbols", 0x2100, 0x214F, 80},
		{"Arrows", 0x2190, 0x21FF, 23},
		{"Mathematical Operators", 0x2200, 0x22FF, 104},
		{"Miscellaneous Technical", 0x2300, 0x23FF, 118},
		{"Box Drawing", 0x2500, 0x257F, 128},
		{"Block Elements", 0x2580, 0x259F, 32},
		{"Geometric Shapes", 0x25A0, 0x25FF, 96},
	}
	for _, b := range blocks {
		got := 0
		for r := b.lo; r <= b.hi; r++ {
			if _, ok := f.NominalGlyph(r); ok {
				got++
			}
		}
		if got != b.want {
			t.Errorf("%s: %d glyphs of %d, want %d — the package comment's table is stale",
				b.name, got, b.hi-b.lo+1, b.want)
		}
	}
}
