package parser

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Site struct {
	Name string
	URL  string
}

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
