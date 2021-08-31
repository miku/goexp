# Compilation Options

Typically, you will build your project with `go build`, but the compilation step can be invoked with:

```go
$ go tool compile
```

* there is no `-O2`, but you can disable optimization with `-N`

A way to shrink a binary is to strip debug information (in linking step):

```
$ go build -ldflags="-w -s" ...
```

* `-s`, omit the symbol table and debug information
* `-w`, omit the DWARF symbol table

Further options: [https://pkg.go.dev/cmd/link](https://pkg.go.dev/cmd/link)