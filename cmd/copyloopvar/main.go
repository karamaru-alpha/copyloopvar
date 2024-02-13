package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/karamaru-alpha/copyloopvar"
)

func main() { unitchecker.Main(copyloopvar.Analyzer) }
