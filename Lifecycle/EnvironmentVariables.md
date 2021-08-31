# Environment Variables

View them all with:

```
$ go env
GOPATH="/home/tir/go"
GOROOT="/usr/local/go"
GOARCH="amd64"
GOOS="linux"
GOENV="/home/tir/.config/go/env"
CGO_ENABLED="1"
...
```

* Prefix CGO is relevant for non-Go code (might be in non "cgo" projects as well).
* module handling: `GOMODCACHE`, `GOSUMDB`, `GONOSUMDB`, `GOPRIVATE`


----

## Complete List

```
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/tir/.cache/go-build"
GOENV="/home/tir/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/tir/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/tir/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.16.5"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/home/tir/code/miku/goexp/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build3399041964=/tmp/go-build -gno-record-gcc-switches"
```