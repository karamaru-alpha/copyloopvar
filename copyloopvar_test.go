package copyloopvar

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc    string
		dir     string
		options map[string]string
	}{
		{
			desc: "basic",
			dir:  "basic",
		},
		{
			desc: "check-alias",
			dir:  "checkalias",
			options: map[string]string{
				"check-alias": "true",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			analyzer := NewAnalyzer()

			for k, v := range test.options {
				if err := analyzer.Flags.Set(k, v); err != nil {
					t.Error(err)
				}
			}

			results := analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), analyzer, test.dir)

			hasSuggestedFixes(t, results)
		})
	}
}

func hasSuggestedFixes(t *testing.T, results []*analysistest.Result) {
	t.Helper()

	for _, result := range results {
		for _, diagnostic := range result.Diagnostics {
			if len(diagnostic.SuggestedFixes) > 0 {
				return
			}
		}
	}

	t.Errorf("no suggested fixes found")
}
