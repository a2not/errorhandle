package main

import (
	"github.com/a2not/errorhandle"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(errorhandle.Analyzer) }
