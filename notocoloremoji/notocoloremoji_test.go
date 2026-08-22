package notocoloremoji_test

import (
	"testing"

	giofont "gioui.org/font"
	"gioui.org/text"
	gotext "github.com/go-text/typesetting/font"
	"golang.org/x/image/math/fixed"

	"github.com/vibrantgio/font/notocoloremoji"
)

// TestFontFace parses the embedded TTF (FontFace panics if opentype.Parse
// rejects it) and asserts the face carries the metadata its name claims.
func TestFontFace(t *testing.T) {
	face := notocoloremoji.FontFace()
	if face.Font != notocoloremoji.Font {
		t.Errorf("FontFace().Font = %+v, want %+v", face.Font, notocoloremoji.Font)
	}
	if face.Font.Typeface != "Noto Color Emoji" {
		t.Errorf("Typeface = %q, want %q", face.Font.Typeface, "Noto Color Emoji")
	}
	if face.Font.Style != giofont.Regular || face.Font.Weight != giofont.Normal {
		t.Errorf("style/weight = %v/%v, want Regular/Normal", face.Font.Style, face.Font.Weight)
	}
	if face.Face == nil {
		t.Fatal("FontFace().Face is nil")
	}
	if again := notocoloremoji.FontFace(); again.Face != face.Face {
		t.Error("FontFace() parsed a second face instead of caching the first")
	}
}

// TestFontFaces asserts the family aggregate is the one leaf face, so that
// this package can stand in for roboto and robotomono where a collection is
// assembled.
func TestFontFaces(t *testing.T) {
	faces := notocoloremoji.FontFaces()
	if len(faces) != 1 {
		t.Fatalf("FontFaces() returned %d faces, want 1", len(faces))
	}
	if faces[0] != notocoloremoji.FontFace() {
		t.Errorf("FontFaces()[0] = %+v, want FontFace()'s %+v", faces[0].Font, notocoloremoji.FontFace().Font)
	}
}

func parsedFace(t *testing.T) *gotext.Face {
	t.Helper()
	parsed, ok := notocoloremoji.FontFace().Face.(interface{ Face() *gotext.Face })
	if !ok {
		t.Fatalf("face %T does not expose the parsed font", notocoloremoji.FontFace().Face)
	}
	return parsed.Face()
}

// TestCoverage asserts the single-codepoint probes the package comment
// promises are in the face's cmap, and that a Latin letter is not. It reads
// the cmap directly rather than shaping: what a shaper resolves depends on
// the collection, and this assertion is about this file alone.
func TestCoverage(t *testing.T) {
	f := parsedFace(t)
	probes := []rune{'😀', '😂', '😉', '😍', '🔥', '🎉', '\u2764', '✨', '🚀'}
	for _, r := range probes {
		if _, ok := f.NominalGlyph(r); !ok {
			t.Errorf("U+%04X %q is not in the face's cmap", r, r)
		}
	}
	if _, ok := f.NominalGlyph('A'); ok {
		t.Error("U+0041 'A' is in the face's cmap; this face is emoji, not Latin")
	}
}

// TestGlyphDataPNG asserts a probe glyph's GlyphData is a PNG bitmap — the
// face is CBDT/CBLC, format 17, with no outlines.
func TestGlyphDataPNG(t *testing.T) {
	f := parsedFace(t)
	gid, ok := f.NominalGlyph('😀')
	if !ok {
		t.Fatal("😀 is not in the face's cmap")
	}
	data := f.GlyphData(gid)
	bm, ok := data.(gotext.GlyphBitmap)
	if !ok {
		t.Fatalf("GlyphData(%d) is %T, want GlyphBitmap", gid, data)
	}
	if bm.Format != gotext.PNG {
		t.Errorf("bitmap format = %d, want PNG (%d)", bm.Format, gotext.PNG)
	}
	if len(bm.Data) == 0 {
		t.Error("PNG bitmap has no bytes")
	}
}

// TestShaperResolution asserts a NoSystemFonts shaper on this collection
// alone resolves 😀 to a real glyph and does not resolve 'A'. Glyph ID 0 is
// .notdef: a shaper that found no face for a rune still returns a glyph, and
// this is how it says so.
func TestShaperResolution(t *testing.T) {
	shaper := text.NewShaper(text.NoSystemFonts(), text.WithCollection(notocoloremoji.FontFaces()))
	params := text.Parameters{
		Font:     notocoloremoji.Font,
		PxPerEm:  fixed.I(16),
		MaxWidth: 1000,
	}

	grin := resolvedGID(t, shaper, params, "😀")
	if grin == 0 {
		t.Error("😀 resolved to glyph ID 0 (.notdef) on this collection")
	}
	letter := resolvedGID(t, shaper, params, "A")
	if letter != 0 {
		t.Errorf("'A' resolved to glyph ID %d, want 0 (.notdef)", letter)
	}
}

func resolvedGID(t *testing.T, shaper *text.Shaper, params text.Parameters, s string) uint32 {
	t.Helper()
	shaper.LayoutString(params, s)
	g, ok := shaper.NextGlyph()
	if !ok {
		t.Fatalf("%q: shaper produced no glyph at all", s)
	}
	return uint32(g.ID)
}
