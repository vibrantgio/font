# AGENTS.md — font

Roboto packaged as Gio font faces: `font/roboto` names the twelve
weight-and-style combinations as `font.Font` values and returns them all,
lazily parsed, from `FontFaces()`; one leaf package per face —
`roboto/regular/normal`, `roboto/italic/bold` and the rest — carries a
single `Font` and its own `FontFace()`, so a program can link one weight
instead of twelve.

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
