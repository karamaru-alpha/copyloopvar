package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/karamaru-alpha/copyloopvar"
)

func main() { singlechecker.Main(copyloopvar.Analyzer) }
