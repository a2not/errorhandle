package main

import (
	"github.com/Khdbble/errorhandle"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(errorhandle.Analyzer) }
