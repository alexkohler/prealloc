# prealloc (work in progress)

prealloc is a Go static analysis tool to find slice declarations that could potentially be preallocated.

## Installation

    go get -u github.com/alexkohler/prealloc

## Usage

Similar to other Go static anaylsis tools (such as golint, go vet), prealloc can be invoked with one or more filenames, directories, or packages named by its import path. Prealloc also supports the `...` wildcard. 

    nakedret [flags] files/directories/packages

Currently, the only flag supported is -tests, which is an optional boolean flag to specify whether or not tests should be included in the analysis.

## Purpose

If the size of a slice is known at its creation time, it should be specified, as it will be much more performant than declaring a slice whose underlying array will potentially have to be reallocated multiple times.

//TODO - Add in benchmarks 

## Example

//TODO - include examples from Go source.

```Bash
$ prealloc
//TODO
```


## Contributing

Pull requests welcome!


## TODO

- Unit tests (may require some refactoring to do correctly)
- supporting toggling of `build.Context.UseAllFiles` may be useful for some. 
- Configuration on whether or not to run on test files
- Vim quickfix format?
- Globbing support (e.g. nakedret *.go)

## Other static analysis tools

If you've enjoyed prealloc, take a look at my other static anaylsis tools!
- [nakedret](https://github.com/alexkohler/nakedret) - Finds naked returns.
- [unimport](https://github.com/alexkohler/unimport) - Finds unnecessary import aliases
