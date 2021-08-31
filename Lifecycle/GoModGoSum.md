# The go.mod and go.sum files

## go.mod

Contains go version used and dependencies plus some extra options.

## go.sum

Contains checksums, should be committed, but not edited. 


## Flows

* add dependencies, e.g. by using "go get ..."
* run "go mod tidy" to keep "go.mod" tidy
* edit "go.mod" to change e.g. a required version of a dependency
* run "go mod vendor" to vendor dependencies into "vendor/" directory

The go tool will try to be informative, e.g. helping with commands to run:

```
$ go mod download github.com/pravj/geopattern
```

Note: "geopattern" is a bad example, as it contains a few backward compatibility
issues.

## Get a specific version of a dependency

```
$ go get example.com/theirmodule@v1.3.4
```

Or by default, get the latest:

```
$ go get example.com/theirmodule@latest
``` 

It's also possible to use a commit hash or branch name.

```
$ go get example.com/theirmodule@4cf76c2
$ go get example.com/theirmodule@bugfixes
```

## List available updates

```
$ go list -m -u all
```

## Cleanup and sync

```
$ go mod tidy
``` 

## Module Path

> The module path should be a path from which Go tools can download the module source. In practice, this is typically the module source’s repository domain and path to the module code within the repository. The go command relies on this form when downloading module versions to resolve dependencies on the module user’s behalf.

> Even if you’re not at first intending to make your module available for use from other code, using its repository path is a best practice that will help you avoid having to rename the module if you publish it later.

> If at first you don’t know the module’s eventual repository location, consider temporarily using a safe substitute, such as the name of a domain you own or example.com, along with a path following from the module’s name or source directory.

> For example, if you’re developing in a stringtools directory, your temporary
module path might be example.com/stringtools, as in the following example:


## Version numbers

> When you run a go command such as go get, Go inserts require directives for
> each module containing imported packages. When a module isn’t yet tagged in
> its repository, Go assigns a pseudo-version number it generates when you run
> the command.


## Other directives

You can "redirect" an import, e.g. to a local copy that may contain local fixes.

## Proxy

To use a different module proxy server, set `GOPROXY` 

```
GOPROXY="https://proxy.golang.org,direct"
```

> To specify another module proxy server for Go tools use, set the GOPROXY
> environment variable to the URL of one or more servers. Go tools will try each
> URL in the order you specify. By default, GOPROXY specifies a public
> Google-run module proxy first, then direct download from the module’s
> repository (as specified in its module path)