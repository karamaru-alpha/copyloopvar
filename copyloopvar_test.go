package copyloopvar

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		testdata := testutil.WithModules(t, analysistest.TestData(), nil)
		analysistest.Run(t, testdata, Analyzer, "basic")
	})

	t.Run("ignore-alias", func(t *testing.T) {
		err := Analyzer.Flags.Set("ignore-alias", "true")
		if err != nil {
			t.Error(err)
		}
		testdata := testutil.WithModules(t, analysistest.TestData(), nil)
		analysistest.Run(t, testdata, Analyzer, "ignorealias")
	})
}
