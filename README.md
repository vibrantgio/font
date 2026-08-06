# font

Roboto — and its companion monospace face, Roboto Mono — packaged as Gio font
faces, for [Vibrant Gio](https://github.com/vibrantgio),
a design system for native desktop applications on macOS, Windows and Linux,
written in pure Go on [Gio](https://gioui.org). This repository is the
typefaces, and nothing else: no theme, no scale, no widgets.

Gio does not ship Roboto. It ships `gioui.org/font/gofont`, the Go typeface,
and that is the collection every example in the Gio world reaches for. Getting
Roboto instead means finding the TTF bytes — `eliasnaur.com/font/roboto/*`
publishes one package per weight — calling `opentype.Parse` on them, handling
the error, and assembling a `[]font.FontFace` before you can build a
`*text.Shaper`. Every program that wanted Roboto wrote that loop. This module
writes it once, parses lazily behind a `sync.Once`, and offers the result at
three granularities: all twelve faces, the six of one style, or exactly one.

The granularity is the point. A `FontFace` is a parsed font, and parsing twelve
of them costs both time at first use and the TTF bytes linked into the binary.
An application that only ever draws body text imports
`roboto/regular/normal` and links one face; an application that wants the whole
family imports `roboto` and gets twelve. Both are one import.

## Where it sits

Tier 0 of the stack — `mvu → spectrum → prism → pulse → cadence → markdown` —
a leaf that imports nothing else in the organization, only
`eliasnaur.com/font` and Gio. The
[organization page](https://github.com/vibrantgio) has the full tier table.

Two modules in the organization import it today, and one of them is the one
that matters: [spectrum](https://github.com/vibrantgio/spectrum)'s `tokens`
package builds the default `Typography`'s face collection from
`roboto.FontFaces()` and `robotomono.FontFaces()` (C1.2), so every component
that shapes text through the theme renders these faces — which is why the
tier table carries a `font` row at all. The other is the frozen
[style](https://github.com/vibrantgio/style) module, whose deprecated scale
still pulls in five of the upright leaves. The old defect — every component
compiling in gofont — is gone the hard way: a no-gofont lint in prism, pulse,
cadence and markdown fails `go test` on the import.

```sh
go get github.com/vibrantgio/font
```

Every module in the organization is on gioui.org v0.10.1 and Go 1.25.1.

## Packages

| Package | |
| --- | --- |
| `roboto` | The whole family. Twelve `font.Font` values named `RegularThin` … `RegularBlack` and `ItalicThin` … `ItalicBlack`, and `FontFaces()`, which parses and returns all twelve. |
| `roboto/regular` | The six upright weights as `Thin`, `Light`, `Normal`, `Medium`, `Bold`, `Black`, and a six-face `FontFaces()`. |
| `roboto/italic` | The same six weights, italic, under the same six names. |
| `roboto/{regular,italic}/{thin,light,normal,medium,bold,black}` | Twelve leaf packages, one face each: a `Font` value and a `FontFace()` that parses exactly that TTF. Import one to link one weight. |
| `robotomono` | Roboto Mono, typeface name `"Roboto Mono"`. Four `font.Font` values — `RegularNormal`, `RegularBold`, `ItalicNormal`, `ItalicBold` — and `FontFaces()`, which parses and returns all four. Only the weights the markdown code path shapes are packaged (G-F0). |
| `robotomono/{regular,italic}` | The two weights of one style as `Normal` and `Bold`, and a two-face `FontFaces()`. |
| `robotomono/{regular,italic}/{normal,bold}` | Four leaf packages, one face each, embedding exactly that TTF. |

The root package `github.com/vibrantgio/font` is empty — see Status.

The counts are exact: `roboto.FontFaces()` returns 12, `roboto/regular` and
`roboto/italic` return 6 each, `robotomono.FontFaces()` returns 4,
`robotomono/regular` and `robotomono/italic` return 2 each, and a leaf's
`FontFace()` returns one.

## Usage

First, the case where you do nothing: a Vibrant Gio application gets these
faces through the theme — `tokens.DefaultTypography.Faces` is built from this
module, and `Typography.Shaper()` is the one shaper — so an app never imports
font directly. The calls below are for programs outside the theme system.

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
families — markdown's code spans resolve through the generic `monospace`
family, which only exists if system fonts are on.

Once the collection is in the shaper, its typefaces become the shaper's default
families, so a widget that lays text out with a zero `font.Font{}` — empty
typeface, Normal weight — resolves to Roboto without naming it. That is how
[style](https://github.com/vibrantgio/style)'s type scale gets Roboto while
naming only weights.

## For coding assistants

Read the canonical guide before writing code against this module — the module
inventory with current tags, the application skeleton, MVU and rx semantics,
typography, and the pitfalls that are not guessable:

<https://raw.githubusercontent.com/vibrantgio/.github/master/llms.txt>

[`AGENTS.md`](./AGENTS.md) in this repository has the build and test commands.

## Status

Honest about what does not work yet. Every count below is measured.

- **The module root package is empty.** `doc.go` declares `package font` and
  nothing else — no exports, and no package comment — so importing
  `github.com/vibrantgio/font` gets you nothing, and the module's landing page
  on pkg.go.dev is blank. Everything real is under `roboto/`.
- **A parse failure panics.** Neither `FontFaces()` nor a leaf's `FontFace()`
  returns an error; both `panic` if `opentype.Parse` rejects the TTF. The bytes
  are compiled in, so this cannot fail at run time for a build that linked, but
  there is no seam for a caller-supplied font file either.
- **Roboto and Roboto Mono are the only families.** There is no API to
  register another typeface, so an application that wants its own brand face
  cannot get one through this module — it builds the `[]font.FontFace` itself.
  The `Typography` theme token carries the face collection precisely so a
  caller can substitute it there.
- **The fine granularity is API, not practice.** Outside this module, only
  the two whole-family aggregates and five leaves are imported: `roboto` and
  `robotomono` by `spectrum/tokens` for the default face collection, and the
  five upright leaves `roboto/regular/{thin,light,normal,medium,bold}` by the
  frozen `style` module. The one-leaf-one-face granularity this README opens
  with has no consumer left since the `mvu` examples moved onto the theme's
  typography; it survives as API surface, exercised by nothing.
- **The Roboto packages have no tests.** `go test ./...` reports "no test
  files" for all sixteen Roboto-side packages; their face counts and weight
  assignments in this README were measured against the built module, not
  asserted by a test in it. The `robotomono` package does carry tests, which
  parse all four embedded TTFs and assert their metadata.

## License

MIT — see [LICENSE](./LICENSE). The font data carries its own licences: the
Roboto TTFs come from `eliasnaur.com/font` (Apache 2.0), and the embedded
Roboto Mono TTFs are the static instances from
[googlefonts/RobotoMono](https://github.com/googlefonts/RobotoMono) under the
SIL Open Font License 1.1 — see [robotomono/OFL.txt](./robotomono/OFL.txt).
(Roboto Mono was historically Apache 2.0; the project relicensed to the OFL,
and Google Fonts now distributes it under `ofl/robotomono`.)
