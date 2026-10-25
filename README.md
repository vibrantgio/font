# font

Roboto — its companion monospace face, Roboto Mono, a second monospace
family, JetBrains Mono, an optional symbol face, and an optional
color-emoji face — packaged as Gio font faces, for
[Vibrant Gio](https://github.com/vibrantgio),
a design system for native desktop applications on macOS, Windows and Linux,
written in pure Go on [Gio](https://gioui.org). This repository is the
typefaces, and nothing else: no theme, no scale, no components.

Gio does not ship Roboto. It ships `gioui.org/font/gofont`, the Go typeface,
and that is the collection every example in the Gio world reaches for. Getting
Roboto instead means finding the TTF bytes, calling `opentype.Parse` on them,
handling the error, and assembling a `[]font.FontFace` before you can build a
`*text.Shaper`. Every program that wants Roboto writes that loop. This module
embeds the twelve faces, writes that loop once, parses lazily behind a
`sync.Once`, and offers the result at three granularities: all twelve faces,
the six of one style, or exactly one. A caller that is not a Gio shaper
takes the same file from a leaf's `TTF` bytes.

The granularity is the point. A `FontFace` is a parsed font, and parsing twelve
of them costs both time at first use and the TTF bytes linked into the binary.
An application that only ever draws body text imports
`roboto/regular/normal` and links one face; an application that wants the whole
family imports `roboto` and gets twelve. Both are one import.

## Where it sits

Tier 0 of the stack — `mvu → theme → components → effects → patterns → markdown` —
a leaf that imports nothing else in the organization, only
Gio; every TTF is embedded in this repository. The
[organization page](https://github.com/vibrantgio) has the full tier table.

```sh
go get github.com/vibrantgio/font
```

The Gio and Go versions are the ones `go.mod` names.

## Packages

| Package | |
| --- | --- |
| `roboto` | The whole family. Twelve `font.Font` values named `RegularThin` … `RegularBlack` and `ItalicThin` … `ItalicBlack`, and `FontFaces()`, which parses and returns all twelve. |
| `roboto/regular` | The six upright weights as `Thin`, `Light`, `Normal`, `Medium`, `Bold`, `Black`, and a six-face `FontFaces()`. |
| `roboto/italic` | The same six weights, italic, under the same six names. |
| `roboto/{regular,italic}/{thin,light,normal,medium,bold,black}` | Twelve leaf packages, one face each: a `Font` value, a `FontFace()` that parses exactly that TTF, and the raw bytes as `TTF`. Import one to link one weight. |
| `robotomono` | Roboto Mono, typeface name `"Roboto Mono"`. Four `font.Font` values — `RegularNormal`, `RegularBold`, `ItalicNormal`, `ItalicBold` — and `FontFaces()`, which parses and returns all four. Only the regular and bold weights are packaged. |
| `robotomono/{regular,italic}` | The two weights of one style as `Normal` and `Bold`, and a two-face `FontFaces()`. |
| `robotomono/{regular,italic}/{normal,bold}` | Four leaf packages, one face each, embedding exactly that TTF. |
| `jetbrainsmono` | JetBrains Mono, typeface name `"JetBrains Mono"`. Four `font.Font` values — `RegularNormal`, `RegularBold`, `ItalicNormal`, `ItalicBold` — and `FontFaces()`, which parses and returns all four. The same four-face layout as `robotomono`. |
| `jetbrainsmono/{regular,italic}` | The two weights of one style as `Normal` and `Bold`, and a two-face `FontFaces()`. |
| `jetbrainsmono/{regular,italic}/{normal,bold}` | Four leaf packages, one face each, embedding exactly that TTF. |
| `notosansmono` | The optional symbol face, typeface name `"Noto Sans Mono"`. One weight, so the family package is the leaf: `Font`, `FontFace()`, and a one-entry `FontFaces()`. Arrows, box drawing, block elements, geometric shapes and the punctuation and operators Roboto lacks. |
| `notocoloremoji` | The optional color-emoji face, typeface name `"Noto Color Emoji"`. One weight, so the family package is the leaf: `Font`, `FontFace()`, and a one-entry `FontFaces()`. CBDT/PNG color emoji the rest of the collection cannot resolve. A shaper reaches it as a fallback. |

The root package `github.com/vibrantgio/font` is empty — see Status.

The counts are exact: `roboto.FontFaces()` returns 12, `roboto/regular` and
`roboto/italic` return 6 each, `robotomono.FontFaces()` returns 4,
`robotomono/regular` and `robotomono/italic` return 2 each,
`jetbrainsmono.FontFaces()` returns 4,
`jetbrainsmono/regular` and `jetbrainsmono/italic` return 2 each,
`notosansmono.FontFaces()` returns 1, `notocoloremoji.FontFaces()`
returns 1, and a leaf's `FontFace()` returns one.

## Usage

A Vibrant Gio application normally takes its faces and its shaper from the
[theme](https://github.com/vibrantgio/theme)'s `Typography`. The calls below
build a shaper straight from this module.

One weight, one face, one shaper — for a program that draws nothing but body
text:

```go
import "github.com/vibrantgio/font/roboto/regular/normal"

shaper := text.NewShaper(text.WithCollection([]font.FontFace{normal.FontFace()}))
```

The shaper is built once, outside the render closure. A `*text.Shaper` owns
glyph caches and is not safe for concurrent use, so build one per window at
layer-building scope and hand it to everything that draws — never one per
frame, and never one shared between windows.

The whole family is the same call without the slice literal:

```go
import "github.com/vibrantgio/font/roboto"

shaper := text.NewShaper(text.WithCollection(roboto.FontFaces()))
```

Note what is *not* passed: `text.NoSystemFonts()`. Leaving the system fonts
loaded keeps a fallback for glyphs Roboto lacks and for explicitly named
families — the generic `monospace` family only exists if system fonts are on.

### The optional symbol face

Roboto and Roboto Mono carry no arrow, no box-drawing character and no
dingbat. On a desktop that costs nothing — the system fallback above serves
them, and serves far more besides. Where there is no system to fall back to —
a container, a kiosk, a scratch image — append `notosansmono`:

```go
import "github.com/vibrantgio/font/notosansmono"

typ := tokens.DefaultTypography.WithFaces(notosansmono.FontFace())
```

It is deliberately **not** in `tokens.DefaultTypography.Faces`. Adding it there
would link 596 KB into every binary that uses the default typography to
duplicate what the OS already has. It is a fallback for the case where there
is none.

Measured coverage, per Unicode block: Box Drawing 128/128, Block Elements
32/32, Geometric Shapes 96/96, General Punctuation 111/112, Letterlike Symbols
80/80, Currency Symbols 32/32, Miscellaneous Technical 118/256, Mathematical
Operators 104/256, Arrows 23/112. The last two are partial: the arrows are the
cardinal, double and long-tailed forms, and the operators are the ones prose
uses. Diagonal arrows, the large operators (∑ ∏ ∫), dingbats (✓ ✗ ★), emoji and
CJK are not covered — `WithFaces` takes as many faces as you give it.

### The optional color-emoji face

Gio's system fallback does not supply a color-emoji face. Apple Color Emoji
can be installed on the machine and still lose: both `Typography.Shaper()`
and `DeterministicShaper()` resolve `😀` to Roboto's `.notdef` unless this
face is in the collection. Append it the same way as the symbol face:

```go
import "github.com/vibrantgio/font/notocoloremoji"

typ := tokens.DefaultTypography.WithFaces(notocoloremoji.FontFace())
```

It is deliberately **not** in `tokens.DefaultTypography.Faces`. Adding it
there would parse 9.9 MB for every shaper built on the default typography. A
shaper reaches it only as a fallback.

The face is one 109 ppem CBDT strike of format-17 PNGs. go-text applies
the face's GSUB, so ZWJ sequences this face ligates (`👨‍👩‍👧‍👦`, flags,
professions, skin tones) shape to one glyph. This package does not
compose sequences itself.

Once the collection is in the shaper, its typefaces become the shaper's default
families, so a component that lays text out with a zero `font.Font{}` — empty
typeface, Normal weight — resolves to Roboto without naming it, and a
`font.Font` that names only a weight gets that weight of Roboto.

## For coding assistants

Read this repository's [AGENTS.md](./AGENTS.md) before you write code against this module.

Build and test the root module with `go build ./...` and `go test ./...`.

## Status

Honest about what does not work yet.

- **The module root package is empty.** `doc.go` declares `package font` and
  nothing else — no exports, and no package comment — so importing
  `github.com/vibrantgio/font` gets you nothing, and the module's landing page
  on pkg.go.dev is blank. Everything real is under `roboto/`.
- **A parse failure panics.** Neither `FontFaces()` nor a leaf's `FontFace()`
  returns an error; both `panic` if `opentype.Parse` rejects the TTF. The bytes
  are compiled in, so this cannot fail at run time for a build that linked, but
  there is no seam for a caller-supplied font file either.
- **These five families are the only ones.** There is no API to register a
  typeface beyond these five, so an application that wants its own brand face cannot get
  one through this module — it builds the `font.FontFace` itself. Adding it to
  the theme is one line, though: `tokens.DefaultTypography.WithFaces(face)`.
- **The symbol face is one weight and one style.** `notosansmono` is Noto Sans
  Mono Regular and nothing else. Bold or italic symbols are not available, and
  text that asks for them gets the regular face.
- **The color-emoji face is one weight and one style.** `notocoloremoji` is
  Noto Color Emoji Regular and nothing else. It has no Latin, and text that
  asks it for `'A'` gets `.notdef`.
- **The Roboto packages have no tests.** `go test ./...` reports "no test
  files" for all fifteen Roboto packages (`roboto`, `roboto/regular`,
  `roboto/italic` and the twelve leaves); their face counts and weight
  assignments are not asserted by a test. The `robotomono`, `jetbrainsmono`,
  `notosansmono` and `notocoloremoji` packages do carry tests, which parse
  the embedded TTFs and assert their metadata —
  `notosansmono`'s also asserts the coverage table above, block by block, so
  the documentation cannot drift from the file, and `notocoloremoji`'s
  asserts the emoji probes, that `GlyphData` is a PNG, and that a pinned
  shaper on this collection resolves 😀 and does not resolve `'A'`.

## License

MIT — see [LICENSE](./LICENSE). The font data has its own licences: the
embedded Roboto TTFs are under the Apache License 2.0 — see
[roboto/LICENSE](./roboto/LICENSE) — and the embedded
Roboto Mono TTFs are the static instances from
[googlefonts/RobotoMono](https://github.com/googlefonts/RobotoMono) under the
SIL Open Font License 1.1 — see [robotomono/OFL.txt](./robotomono/OFL.txt). The JetBrains
Mono TTFs are the static instances from
[JetBrains/JetBrainsMono](https://github.com/JetBrains/JetBrainsMono) under
the SIL Open Font License 1.1 — see
[jetbrainsmono/OFL.txt](./jetbrainsmono/OFL.txt). The symbol face is
the static Regular instance of Noto Sans Mono v2.014, from the Noto project's
own [release](https://github.com/notofonts/latin-greek-cyrillic/releases/tag/NotoSansMono-v2.014),
also under the SIL Open Font License 1.1 — see
[notosansmono/OFL.txt](./notosansmono/OFL.txt). The color-emoji face is
Noto Color Emoji Regular, also under the SIL Open Font License 1.1 — see
[notocoloremoji/OFL.txt](./notocoloremoji/OFL.txt).
