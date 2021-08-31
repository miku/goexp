# Packages and Modules


## Packages

Packages were there from the beginning.

* package: A collection of source files in the same directory that are compiled together. 

For example, you cannot have two files with two different package names one directory.

```
$ cat *
package a
package b

$ go build .
found packages a (a.go) and b (b.go) in /home/tir/code/miku/goexp/Lifecycle/PackagesAndModules/SingleName
``` 

## Modules

Modules are newer.

> Go 1.11 and 1.12 include preliminary support for modules, Go’s new dependency
> management system that makes dependency version information explicit and
> easier to manage.

* module: A collection of packages that are released, versioned, and distributed together.

Other notions:

* module cache: A local directory storing downloaded modules, located in GOPATH/pkg/mod. See Module cache.
* module graph: The directed graph of module requirements, rooted at the main module. Each vertex in the graph is a module; each edge is a version from a require statement in a go.mod file (subject to replace and exclude statements in the main module’s go.mod file.
* module graph pruning: A change in Go 1.17 that reduces the size of the module graph by omitting transitive dependencies of modules that specify go 1.17 or higher. See Module graph pruning.
* module path: A path that identifies a module and acts as a prefix for package import paths within the module. For example, "golang.org/x/net".
* module proxy: A web server that implements the GOPROXY protocol. The go command downloads version information, go.mod files, and module zip files from module proxies.
* module root directory: The directory that contains the go.mod file that defines a module.


When choosing a module path, we often use "github.com/username/project" because that gives us:

* a unique name
* a location where we can go get the repo

> A module path should describe both what the module does and where to find it.
> Typically, a module path consists of a repository root path, a directory
> within the repository (usually empty), and a major version suffix (only for
> major version 2 or higher).

The go tool has a mechanism to go from "module path" to repository location.

* [https://golang.org/ref/mod#vcs-find](https://golang.org/ref/mod#vcs-find)



## What is the comment behind a package name

That is a canonical import path.

* https://golang.org/doc/go1.4#canonicalimports

```
package pdf // import "rsc.io/pdf"
```

It fixes the name to be used to import a package.
