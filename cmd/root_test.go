package cmd

import (
	"encoding/csv"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadAndParseCSV(t *testing.T) {
	var fakeCSV []string
	var i int
	for ; i <= 1000; i++ {
		fakeCSV = append(fakeCSV, ".com,https://$.com,https://$.com")
	}

	r := csv.NewReader(strings.NewReader(strings.Join(fakeCSV, "\n")))
	sites, err := readAndParseCSV(r, "m")
	if err != nil {
		t.Fatalf("while reading and parsing fake .csv: %v", err)
	}

	if len(sites) != i {
		t.Fatalf("expected sites to have a length of %v. got=%v", i, len(sites))
	}
}

func TestMakeRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		switch r.URL.String() {
		case "/ok":
			w.WriteHeader(http.StatusOK)
		case "/notfound":
			w.WriteHeader(http.StatusNotFound)
		case "/internalerror":
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer ts.Close()

	tt := []struct {
		name               string
		url                string
		expectedToFail     bool
		expectedStatusCode int
	}{
		{
			name:           "no URL",
			url:            "",
			expectedToFail: true,
		},
		{
			name:               "ok",
			url:                ts.URL + "/ok",
			expectedToFail:     false,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "404",
			url:                ts.URL + "/notfound",
			expectedToFail:     false,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "500",
			url:                ts.URL + "/internalerror",
			expectedToFail:     false,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	c := &http.Client{}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			status, err := makeRequest(c, tc.url, "")
			if err != nil && !tc.expectedToFail {
				t.Fatalf("not expected to fail: %v", err)
			}

			if status != tc.expectedStatusCode {
				t.Fatalf("expected a %v as status code. got=%v", tc.expectedStatusCode, status)
			}
		})
	}
}

func TestReplaceURL(t *testing.T) {
	tt := []struct {
		old      string
		new      string
		expected string
	}{
		{
			old:      "$",
			new:      "me",
			expected: "me",
		},
		{
			old:      "$$$y",
			new:      "me",
			expected: "me$$y",
		},
		{
			old:      "https://$.com",
			new:      "me",
			expected: "https://me.com",
		},
		{
			old:      "",
			new:      "",
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.expected, func(t *testing.T) {
			url := replaceURL(tc.old, tc.new)
			if url != tc.expected {
				t.Fatalf("expected %q as result. got=%q", tc.expected, url)
			}
		})
	}
}
