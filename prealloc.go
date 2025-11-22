package main

import (
	"flag"
	"go/build"
	"log"

	"github.com/alexkohler/prealloc/pkg"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// Support: (in order of priority)
//  * Full make suggestion with type?
//	* Test flag
//  * Embedded ifs?
//  * Use an import rather than the duplicated import.go

func main() {
	singlechecker.Main(NewAnalyzer())
}

type prealloc struct {
	simple            bool
	includeRangeLoops bool
	includeForLoops   bool
}

func NewAnalyzer() *analysis.Analyzer {
	// Ignore build flags
	build.Default.UseAllFiles = true

	// Remove log timestamp
	log.SetFlags(0)

	p := &prealloc{}

	a := &analysis.Analyzer{
		Name: "prealloc",
		Doc:  "Find slice declarations that could potentially be preallocated",
		Run:  p.run,
	}
	a.Flags.Init("prealloc", flag.ExitOnError)
	a.Flags.BoolVar(&p.simple, "simple", true, "Report preallocation suggestions only on simple loops that have no returns/breaks/continues/gotos in them")
	a.Flags.BoolVar(&p.includeRangeLoops, "rangeloops", true, "Report preallocation suggestions on range loops")
	a.Flags.BoolVar(&p.includeForLoops, "forloops", false, "Report preallocation suggestions on for loops")
	return a
}

func (p *prealloc) run(pass *analysis.Pass) (any, error) {
	hints := pkg.Check(pass.Files, p.simple, p.includeRangeLoops, p.includeForLoops)

	for _, hint := range hints {
		pass.Report(hint)
	}

	return nil, nil
}
