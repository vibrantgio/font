# AGENTS.md — font

Roboto packaged as Gio font faces: `font/roboto` names the twelve
weight-and-style combinations as `font.Font` values and returns them all,
lazily parsed, from `FontFaces()`; one leaf package per face —
`roboto/regular/normal`, `roboto/italic/bold` and the rest — carries a
single `Font` and its own `FontFace()`, so a program can link one weight
instead of twelve. `robotomono` packages the four mono faces the markdown
code path shapes, the same way.

**Layer.** Tier 0 of ADR-001's table — a leaf that imports nothing else in
the organization, only `eliasnaur.com/font` and Gio. Inside the org only
`style` and `mvu/example` require it directly today; C1.2 makes spectrum
depend on it for the default Roboto faces, which is why the tier table
carries a `font` row at all.

**Read the canonical guide before you write code against this module.** It is
the organization's one agent guide — the module inventory with current tags,
the application skeleton, the MVU loop and rx semantics, typography, and the
pitfalls that are not guessable. It lives exactly once, in `vibrantgio/.github`,
and this file links it rather than copying it:

    https://raw.githubusercontent.com/vibrantgio/.github/master/llms.txt

**Module.** `github.com/vibrantgio/font`, one module at the repository root.

**Build and test.** From the repository root:

    go build ./... && go test ./...

**`notosansmono` is optional, and keeping it optional is the design.** It is
Noto Sans Mono Regular — one weight — carrying the arrows, box drawing, block
elements, geometric shapes, punctuation and operators Roboto and Roboto Mono
lack. Do **not** add it to `tokens.DefaultTypography.Faces`: spectrum's default
shaper leaves system fonts on, so a desktop already covers those characters and
more, and putting this face in the default would link 596 KB into every binary
in the organization to duplicate the platform. It is for the case with no
platform to fall back to — a container, a kiosk — and for a test that
legitimately draws a symbol while keeping its faces pinned. Both append it the
same way, in one line:

    tokens.DefaultTypography.WithFaces(notosansmono.FontFace())

The package comment carries the measured coverage table and the file's
provenance and SHA-256; `notosansmono_test.go` asserts that table block by
block, so change the TTF and the test tells you what moved.
