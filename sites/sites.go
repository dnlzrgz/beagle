// Package sites defines utilities to parse and manage
// sites from a .csv file with a name,url format.
package sites

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strings"
)

// Site defines a site with a name and a URL.
type Site struct {
	Name string
	URL  string
}

// ReplaceURL receives a piece of the URL that
// has to be replaced by a new string.
func (s *Site) ReplaceURL(old string, new string) {
	s.URL = strings.Replace(s.URL, old, new, 1)
}

// Parse receives the name of a .csv file
// from which to extrar a list of sites and add
// them to the []*Site returned if no errors occurs.
func Parse(file string) ([]*Site, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bufio.NewReader(f))
	return parse(r)
}

func parse(r *csv.Reader) ([]*Site, error) {
	var sites []*Site

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		sites = append(sites, &Site{Name: line[0], URL: line[1]})
	}

	return sites, nil
}
