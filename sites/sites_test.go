package sites

import (
	"encoding/csv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tt := []struct {
		name           string
		input          string
		expectedResult []*Site
	}{
		{
			name:  "one line",
			input: "500px,https://500px.com/",
			expectedResult: []*Site{
				{
					Name: "500px",
					URL:  "https://500px.com/",
				},
			},
		},
		{
			name:  "multiple lines",
			input: "500px,https://500px.com/\nGoogle,https://google.com\nGitHub,https://github.com",
			expectedResult: []*Site{
				{
					Name: "500px",
					URL:  "https://500px.com/",
				},
				{
					Name: "Google",
					URL:  "https://google.com",
				},
				{
					Name: "GitHub",
					URL:  "https://github.com",
				},
			},
		},
		{
			name:           "empty",
			input:          "",
			expectedResult: []*Site{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := csv.NewReader(strings.NewReader(tc.input))
			sites, err := parse(r)
			if err != nil {
				t.Fatalf("while parsing csv input: %v", err)
			}

			if len(sites) != len(tc.expectedResult) {
				t.Fatalf("expected len of parsed sites to be %v. got=%v", len(tc.expectedResult), len(sites))
			}

			for i, es := range tc.expectedResult {
				if sites[i].Name != es.Name {
					t.Fatalf("expected a parsed site called %q. got a site called %q", es.Name, sites[i].Name)
				}

				if sites[i].URL != es.URL {
					t.Fatalf("expected a parsed site with URL %q. got a site with URL %q", es.URL, sites[i].URL)
				}
			}
		})
	}
}

func BenchmarkParse1Line(b *testing.B) {
	benchmarkParse(b, "Google,https://google.com/")
}

func BenchmarkParse50Lines(b *testing.B) {
	in := strings.Repeat("500px,https://500px.com\n", 50)
	benchmarkParse(b, in)
}

func BenchmarkParse100Lines(b *testing.B) {
	in := strings.Repeat("500px,https://500px.com/\n", 100)
	benchmarkParse(b, in)
}

func BenchmarkParse500Lines(b *testing.B) {
	in := strings.Repeat("500px,https://500px.com\n", 500)
	benchmarkParse(b, in)
}

func benchmarkParse(b *testing.B, input string) {
	r := csv.NewReader(strings.NewReader(input))
	for n := 0; n < b.N; n++ {
		parse(r)
	}
}
