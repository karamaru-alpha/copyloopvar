package copyloopvar

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		testdata := testutil.WithModules(t, analysistest.TestData(), nil)
		analysistest.Run(t, testdata, NewAnalyzer(), "basic")
	})

	t.Run("ignore-alias", func(t *testing.T) {
		analyzer := NewAnalyzer()
		analyzer.Flags.Set("ignore-alias", "true")
		testdata := testutil.WithModules(t, analysistest.TestData(), nil)
		analysistest.Run(t, testdata, analyzer, "ignorealias")
	})
}
