package errorhandle_test

import (
	"testing"

	"github.com/a2not/errorhandle"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, errorhandle.Analyzer, "a")
}
