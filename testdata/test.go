package main

/*
// test case one
func resolvePaths(paths, skip []string) []string {
	if len(paths) == 0 {
		return []string{"."}
	}

	skipPath := newPathFilter(skip)
	dirs := newStringSet()
	for _, path := range paths {
		if strings.HasSuffix(path, "/...") {
			root := filepath.Dir(path)
			_ = filepath.Walk(root, func(p string, i os.FileInfo, err error) error {
				if err != nil {
					warning("invalid path %q: %s", p, err)
					return err
				}

				skip := skipPath(p)
				switch {
				case i.IsDir() && skip:
					return filepath.SkipDir
				case !i.IsDir() && !skip && strings.HasSuffix(p, ".go"):
					dirs.add(filepath.Clean(filepath.Dir(p)))
				}
				return nil
			})
		} else {
			dirs.add(filepath.Clean(path))
		}
	}
	out := make([]string, 0, dirs.size())
	for _, d := range dirs.asSlice() {
		out = append(out, relativePackagePath(d))
	}
	sort.Strings(out)
	for _, d := range out {
		debug("linting path %s", d)
	}
	return out
}*/

/*func cuapo() {
	//myList := make([]uint32, 0, 10)
	var myList []uint32
	existing := make([]uint32, 0, 5)
	for _, ex := range existing {
		myList = append(myList, ex)
	}

}*/

func slinky() ([]string, error) {
	var ofiles []string
	for _, sfile := range sfiles {
		ofile := obj + sfile[:len(sfile)-len(".s")] + ".o"
		ofiles = append(ofiles, ofile)
		a := append(args, "-o", ofile, mkAbs(p.Dir, sfile))
		if err := b.run(p.Dir, p.ImportPath, nil, a...); err != nil {
			return nil, err
		}
	}
	return ofiles, nil
}
