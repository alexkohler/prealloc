# prealloc (work in progress)

prealloc is a Go static analysis tool to find slice declarations that could potentially be preallocated.

## Installation

    go get -u github.com/alexkohler/prealloc

## Usage

Similar to other Go static anaylsis tools (such as golint, go vet), prealloc can be invoked with one or more filenames, directories, or packages named by its import path. Prealloc also supports the `...` wildcard. 

    prealloc [flags] files/directories/packages

Currently, the only flag supported is -tests, which is an optional boolean flag to specify whether or not tests should be included in the analysis.

## Purpose

If the size of a slice is known at the time of its creation, it should be specified, as it will be much more performant than declaring a slice whose underlying array will potentially have to be reallocated multiple times.

Consider the following benchmark: (Found in prealloc_test.go)

```Go
import "testing"

func BenchmarkNoPreallocate(b *testing.B) {
	existing := make([]int64, 1000, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		var init []int64
		for _, element := range existing {
			init = append(init, element)
		}
	}
}

func BenchmarkPreallocate(b *testing.B) {
	existing := make([]int64, 1000, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, 0, 1000)
		for _, element := range existing {
			init = append(init, element)
		}
	}
}
```

```Bash
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
BenchmarkNoPreallocate-4   	  200000	     10424 ns/op	   16376 B/op	      11 allocs/op
BenchmarkPreallocate-4     	 1000000	      1048 ns/op	       0 B/op	       0 allocs/op
```

As you can see, not preallocating causes a large performance hit, primarily due to Go having to reallocate the underlying array.   

## Example

//TODO - include examples from Go source.

```Bash
$ prealloc
//TODO
```

```Go
    // cmd/api/goapi.go:301
	// In next file, but not in API.
	var missing []string
	for feature := range optionalSet {
		missing = append(missing, feature)
	}

```

## TODO

- Unit tests (may require some refactoring to do correctly)
- supporting toggling of `build.Context.UseAllFiles` may be useful for some. 
- Configuration on whether or not to run on test files
- Globbing support (e.g. prealloc *.go)


## Contributing

Pull requests welcome!


## Other static analysis tools

If you've enjoyed prealloc, take a look at my other static anaylsis tools!
- [nakedret](https://github.com/alexkohler/nakedret) - Finds naked returns.
- [unimport](https://github.com/alexkohler/unimport) - Finds unnecessary import aliases
