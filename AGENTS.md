# AGENTS.md — font

The design system's embedded typefaces, packaged as Gio font faces.
`roboto` names the twelve weight-and-style combinations as `font.Font`
values and returns them all, lazily parsed, from `FontFaces()`; one leaf
package per face — `roboto/regular/normal`, `roboto/italic/bold` and the
rest — carries a single `Font` and its own `FontFace()`, so a program can
link one weight instead of twelve. `robotomono` packages the four monospace
faces the same way. `notosansmono` is a third family and a single face,
carrying the symbols the other two lack; it is opt-in rather than part of
the default typography's collection, and the note below says why that is
the design.

**Layer.** Tier 0 of ADR-001's table — a leaf, needing only Gio; every
TTF is embedded in this repository. It is in the tier table rather than
the support row because the faces it packages are the design system's own,
not a general-purpose library that happens to be usable from it. Its root
module imports nothing else in the organization. That direction is measured
rather than typed — `scripts/check-layers.sh --edges` reports the graph and
`scripts/sync-agents.sh` renders these sentences from it — so correcting
them here changes nothing. The other direction is measured too and
deliberately not written down: the gate checks the graph both ways, but a
public API's consumers are unknowable, so this file says what its module
needs and never who needs it.

**Read the canonical guide before you write code against this module.** It is
the organization's one agent guide — the module inventory with current tags,
the application skeleton, the MVU loop and rx semantics, typography, and the
pitfalls that are not guessable. It lives exactly once, in `vibrantgio/workbench` —
the repository that showcases building applications with Vibrant Gio —
and this file links it rather than copying it:

    https://raw.githubusercontent.com/vibrantgio/workbench/master/llms.txt

**Module.** `github.com/vibrantgio/font`, one module at the repository root.

**Build and test.** From the repository root:

    go build ./... && go test ./...

**`notosansmono` is optional, and keeping it optional is the design.** It is
Noto Sans Mono Regular — one weight — carrying the arrows, box drawing, block
elements, geometric shapes, punctuation and operators Roboto and Roboto Mono
lack. Do **not** add it to `tokens.DefaultTypography.Faces`: the default
typography's shaper leaves system fonts on, so a desktop already covers those
characters and more, and putting this face in the default would link 596 KB
into every binary that draws text to duplicate the platform. It is for the case with no
platform to fall back to — a container, a kiosk — and for a test that
legitimately draws a symbol while keeping its faces pinned. Both append it the
same way, in one line:

    tokens.DefaultTypography.WithFaces(notosansmono.FontFace())

The package comment carries the measured coverage table and the file's
provenance and SHA-256; `notosansmono_test.go` asserts that table block by
block, so change the TTF and the test tells you what moved.
