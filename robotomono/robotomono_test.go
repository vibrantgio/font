package robotomono_test

import (
	"testing"

	"gioui.org/font"

	"github.com/vibrantgio/font/robotomono"
	"github.com/vibrantgio/font/robotomono/italic"
	"github.com/vibrantgio/font/robotomono/regular"
)

// TestFontFaces parses all four embedded TTFs (FontFaces panics if
// opentype.Parse rejects one) and asserts each face carries the metadata its
// name claims: typeface "Roboto Mono", the right style and weight, distinct
// parsed faces.
func TestFontFaces(t *testing.T) {
	faces := robotomono.FontFaces()
	if len(faces) != 4 {
		t.Fatalf("FontFaces() returned %d faces, want 4", len(faces))
	}
	want := []font.Font{
		robotomono.RegularNormal,
		robotomono.RegularBold,
		robotomono.ItalicNormal,
		robotomono.ItalicBold,
	}
	seen := map[font.Face]bool{}
	for i, face := range faces {
		if face.Font != want[i] {
			t.Errorf("faces[%d].Font = %+v, want %+v", i, face.Font, want[i])
		}
		if face.Font.Typeface != "Roboto Mono" {
			t.Errorf("faces[%d].Typeface = %q, want %q", i, face.Font.Typeface, "Roboto Mono")
		}
		if face.Face == nil {
			t.Fatalf("faces[%d].Face is nil", i)
		}
		if seen[face.Face] {
			t.Errorf("faces[%d] shares a parsed face with an earlier entry", i)
		}
		seen[face.Face] = true
	}
}

// TestAggregatesAgree asserts the style aggregates return the same faces the
// family aggregate does, two per style.
func TestAggregatesAgree(t *testing.T) {
	all := robotomono.FontFaces()
	upright, slanted := regular.FontFaces(), italic.FontFaces()
	if len(upright) != 2 || len(slanted) != 2 {
		t.Fatalf("regular returned %d faces and italic %d, want 2 and 2", len(upright), len(slanted))
	}
	for i, face := range append(upright, slanted...) {
		if face != all[i] {
			t.Errorf("aggregate face %d = %+v, want the family's %+v", i, face.Font, all[i].Font)
		}
	}
}
