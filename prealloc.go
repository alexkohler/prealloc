package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexkohler/prealloc/pkg"
)

// Support: (in order of priority)
//  * Full make suggestion with type?
//	* Test flag
//  * Embedded ifs?
//  * Use an import rather than the duplcated import.go

const (
	pwd = "./"
)

func init() {
	// Ignore build flags
	build.Default.UseAllFiles = true
}

func usage() {
	log.Printf("Usage of %s:\n", os.Args[0])
	log.Printf("\nprealloc [flags] # runs on package in current directory\n")
	log.Printf("\nprealloc [flags] [packages]\n")
	log.Printf("Flags:\n")
	flag.PrintDefaults()
}

func main() {

	// Remove log timestamp
	log.SetFlags(0)

	simple := flag.Bool("simple", true, "Report preallocation suggestions only on simple loops that have no returns/breaks/continues/gotos in them")
	includeRangeLoops := flag.Bool("rangeloops", true, "Report preallocation suggestions on range loops")
	includeForLoops := flag.Bool("forloops", false, "Report preallocation suggestions on for loops")
	setExitStatus := flag.Bool("set_exit_status", false, "Set exit status to 1 if any issues are found")
	flag.Usage = usage
	flag.Parse()

	fset := token.NewFileSet()

	hints, err := checkForPreallocations(
		flag.Args(),
		fset,
		*simple,
		*includeRangeLoops,
		*includeForLoops,
	)
	if err != nil {
		log.Println(err)
	}

	for _, hint := range hints {
		log.Println(hint.StringFromFS(fset))
	}
	if *setExitStatus && len(hints) > 0 {
		os.Exit(1)
	}
}

func checkForPreallocations(
	args []string,
	fset *token.FileSet,
	simple, includeRangeLoops, includeForLoops bool,
) ([]pkg.Hint, error) {

	files, err := parseInput(args, fset)
	if err != nil {
		return nil, fmt.Errorf("could not parse input %v", err)
	}

	hints := pkg.Check(files, simple, includeRangeLoops, includeForLoops)

	return hints, nil
}

func parseInput(args []string, fset *token.FileSet) ([]*ast.File, error) {
	var directoryList []string
	var fileMode bool
	files := make([]*ast.File, 0)

	if len(args) == 0 {
		directoryList = append(directoryList, pwd)
	} else {
		for _, arg := range args {
			if strings.HasSuffix(arg, "/...") && isDir(arg[:len(arg)-len("/...")]) {

				for _, dirname := range allPackagesInFS(arg) {
					directoryList = append(directoryList, dirname)
				}

			} else if isDir(arg) {
				directoryList = append(directoryList, arg)

			} else if exists(arg) {
				if strings.HasSuffix(arg, ".go") {
					fileMode = true
					f, err := parser.ParseFile(fset, arg, nil, 0)
					if err != nil {
						return nil, err
					}
					files = append(files, f)
				} else {
					return nil, fmt.Errorf("invalid file %v specified", arg)
				}
			} else {

				//TODO clean this up a bit
				imPaths := importPaths([]string{arg})
				for _, importPath := range imPaths {
					pkg, err := build.Import(importPath, ".", 0)
					if err != nil {
						return nil, err
					}
					var stringFiles []string
					stringFiles = append(stringFiles, pkg.GoFiles...)
					// files = append(files, pkg.CgoFiles...)
					stringFiles = append(stringFiles, pkg.TestGoFiles...)
					if pkg.Dir != "." {
						for i, f := range stringFiles {
							stringFiles[i] = filepath.Join(pkg.Dir, f)
						}
					}

					fileMode = true
					for _, stringFile := range stringFiles {
						f, err := parser.ParseFile(fset, stringFile, nil, 0)
						if err != nil {
							return nil, err
						}
						files = append(files, f)
					}

				}
			}
		}
	}

	// if we're not in file mode, then we need to grab each and every package in each directory
	// we can to grab all the files
	if !fileMode {
		for _, fpath := range directoryList {
			pkgs, err := parser.ParseDir(fset, fpath, nil, 0)
			if err != nil {
				return nil, err
			}

			for _, pkg := range pkgs {
				for _, f := range pkg.Files {
					files = append(files, f)
				}
			}
		}
	}

	return files, nil
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
