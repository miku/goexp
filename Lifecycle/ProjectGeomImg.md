# Project "geomimg"

Example project to try out modules. We want a command line tool to generate random images.

We will use
[https://github.com/pravj/geopattern](https://github.com/pravj/geopattern) as a
dependency. This library does the heavy lifting or generating SVG given some
options - for example the pattern type (like "plaid") and some random seed
string. This is a library, we want a command line tool.

Here are some nice features to add:

* the user should be able to override options for patterns and pattern types
* the tool should be default generate a random image as SVG and write it to stdout
    * use a random string for phrase
    * use a random color string for color
    * choose a random pattern type
* add an option "-o FILENAME" to write to an explicit file

Available patterns:

```
chevrons
concentric-circles
diamonds
hexagons
mosaic-squares
nested-squares
octagons
overlapping-circles
overlapping-rings
plaid
plus-signs
sine-waves
squares
tessellation
triangles
xes
```


## Steps

1. Create a new directory.

This will be the project directory.

2. Create a Go modules.

```
$ go mod init geomimg
```

A "go.mod" file will be created.

3. Implement your tool.

You can just use a single "main.go" for now.

4. A bit more structure.

* move the "main" into a "cmd/geomimg" subfolder
* create a file "randutil.go" in the root directory

The "randutil.go" will contain our helper functions.

5. Write a Test for a "randutil.go" function (e.g. for "RandomString")

* implement the test and run it
* run with coverage report

6. Write a Benchmark for e.g. "RandomString"

* implement a benchmark and run it

7. Write sub-benchmarks, if applicable, e.g. for "RandomString" with various length

* implement subbenchmarks and run them

8. Build a binary

Maybe write a short Makefile to automate the process

Use `go build -ldflags="-w -s" ...` to optimize binary size.
