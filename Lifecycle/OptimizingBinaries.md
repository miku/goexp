# Optimizing binaries

* use `go build -ldflags="-s -w" main.go` to omit symbol tables
* use [upx](https://upx.github.io/), to compress binaries 
