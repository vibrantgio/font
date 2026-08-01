# font

Roboto packaged as Gio font faces, for [Vibrant Gio](https://github.com/vibrantgio),
a design system for native desktop applications on macOS, Windows and Linux,
written in pure Go on [Gio](https://gioui.org). This repository is the
typeface, and nothing else: no theme, no scale, no widgets.

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

Three files anywhere in the organization import it today:
[style](https://github.com/vibrantgio/style), which pulls in five of the
leaves, and two `mvu/example` programs, which pull in one each. Everything else
that draws text — every component in prism, pulse, cadence and markdown —
compiles in gofont instead. That is the defect Phase C of the
[org plan](https://github.com/vibrantgio/.github) fixes; C1.2 puts
`roboto.FontFaces()` behind the typography token as the default face
collection, and makes spectrum depend on this module, which is why the tier
table carries a `font` row at all.

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

The root package `github.com/vibrantgio/font` is empty — see Status.

The counts are exact: `roboto.FontFaces()` returns 12, `roboto/regular` and
`roboto/italic` return 6 each, and a leaf's `FontFace()` returns one.

## Usage

One weight, one face, one shaper — from `mvu/example/edit`, which draws into a
`widget.Editor` and needs nothing but body text:

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

- **This module is not wired into the component stack.** No library source
  file in prism, pulse, cadence, markdown or spectrum imports it. Every one of
  those components builds
  `text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))`
  for itself when no shaper is passed, so an application that forgets a
  `Props.Shaper` renders in the Go typeface and nothing warns it. Phase C of
  the [org plan](https://github.com/vibrantgio/.github) fixes this: C1.2 makes
  `roboto.FontFaces()` the default face collection on the typography theme
  token, and C2 removes the gofont fallbacks component by component behind a
  CI lint that forbids the import in library source.
- **The module root package is empty.** `doc.go` declares `package font` and
  nothing else — no exports, and no package comment — so importing
  `github.com/vibrantgio/font` gets you nothing, and the module's landing page
  on pkg.go.dev is blank. Everything real is under `roboto/`.
- **A parse failure panics.** Neither `FontFaces()` nor a leaf's `FontFace()`
  returns an error; both `panic` if `opentype.Parse` rejects the TTF. The bytes
  are compiled in, so this cannot fail at run time for a build that linked, but
  there is no seam for a caller-supplied font file either.
- **Roboto is the only family.** There is no API to register another typeface,
  so an application that wants its own brand face cannot get one through this
  module — it builds the `[]font.FontFace` itself. Phase C's `Typography` token
  carries the face collection precisely so a caller can substitute it there.
- **Eleven of the sixteen packages have no consumer anywhere.** Only five leaves
  are imported by anything in the organization — `roboto/regular/thin`,
  `light`, `normal`, `medium` and `bold`, all five by `style`, and `normal`
  again by two `mvu` examples. The other eleven — `roboto/regular/black`, all
  six italic leaves, the three aggregate packages `roboto`, `roboto/regular`
  and `roboto/italic`, and the empty module root — are imported by nothing at
  all. `FontFaces()`, the twelve-face call this README opens with and the one
  C1.2 is written against, has no caller yet.
- **There are no tests.** `go test ./...` reports "no test files" for all
  sixteen packages. The face counts and weight assignments in this README were
  measured against the built module, not asserted by a test in it.

## License

MIT — see [LICENSE](./LICENSE).
