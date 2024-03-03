package copyloopvar

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, NewAnalyzer(&Setting{
		SkipRename: false,
	}), "basic")
	analysistest.Run(t, testdata, NewAnalyzer(&Setting{
		SkipRename: true,
	}), "skiprename")
}
