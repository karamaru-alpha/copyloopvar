package copyloopvar

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		analysistest.Run(t, analysistest.TestData(), NewAnalyzer(), "basic")
	})

	t.Run("check-alias", func(t *testing.T) {
		analyzer := NewAnalyzer()
		if err := analyzer.Flags.Set("check-alias", "true"); err != nil {
			t.Error(err)
		}

		analysistest.Run(t, analysistest.TestData(), analyzer, "checkalias")
	})
}
