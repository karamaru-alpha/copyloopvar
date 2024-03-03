package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/karamaru-alpha/copyloopvar"
)

func main() {
	var skipRename bool
	flag.BoolVar(&skipRename, "skip-rename", false,
		"Skip renamed copy")
	flag.Parse()

	unitchecker.Main(copyloopvar.NewAnalyzer(&copyloopvar.Setting{
		SkipRename: skipRename,
	}))
}
