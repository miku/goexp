# Linting

* go vet
* gometalinter
* golangci-lint

```
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.0
```

Example output:

```
cmd/skate-wikipedia-doi/main.go:21:2: `bytesNewline` is unused (deadcode)
        bytesNewline = []byte("\n")
        ^
xio/util.go:71:2: unreachable: unreachable code (govet)
        return m, nil
        ^
reduce.go:742:6: `parseBiblioRef` is unused (deadcode)
func parseBiblioRef(s string) (r *BiblioRef, err error) {
     ^
unstructured_test.go:63:20: Error return value is not checked (errcheck)
                ParseUnstructured(&ref)
                                 ^
doi.go:31:19: S1010: should omit second index in slice, s[a:len(s)] is identical to s[a:] (gosimple)
                raw = raw[start:len(raw)]
                                ^
schema_test.go:271:5: S1004: should use !bytes.Equal(csl_json, csl_encoded) instead (gosimple)
        if bytes.Compare(csl_json, csl_encoded) != 0 {
           ^
...
```

A bunch of utilities combined.

```
$ golangci-lint help linters
Enabled by default linters:
deadcode: Finds unused code [fast: true, auto-fix: false] ...
errcheck: Errcheck is a program for checking for unchecked errors in go ...
gosimple: Linter for Go source code that specializes in simplifying a code ...
govet (vet, vetshadow): Vet examines Go source code and reports suspicious ...
ineffassign: Detects when assignments to existing variables are not used ...
staticcheck: Staticcheck is a go vet on steroids, applying a ton of  ...
structcheck: Finds unused struct fields [fast: true, auto-fix: false ...
typecheck: Like the front-end of a Go compiler, parses and type-chec ...
unused: Checks Go code for unused constants, variables, functions an ...
varcheck: Finds unused global variables and constants [fast: true, a ...
...
```
