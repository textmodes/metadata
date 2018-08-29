package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Artist represents an artist
type Artist struct {
	Name      string            `yaml:"name"`
	Aliases   []string          `yaml:"aliases"`
	Country   string            `yaml:"country"`
	Biography string            `yaml:"biography"`
	Website   string            `yaml:"website"`
	Social    map[string]string `yaml:"social"`
}

func test(name string) bool {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL %s: %v\n", name, err)
		return false
	}

	artist := new(Artist)
	if err = yaml.UnmarshalStrict(b, artist); err != nil {
		fmt.Fprintf(os.Stderr, "FAIL %s: %v\n", name, err)
		return false
	}

	if verbose {
		fmt.Printf("GOOD %s\n", name)
	}
	return true
}

var verbose bool

func main() {
	flag.BoolVar(&verbose, "v", false, "be verbose")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "%s <yaml file(s)>\n", os.Args[0])
		os.Exit(1)
	}

	var failed int
	for _, name := range args {
		if !test(name) {
			failed++
		}
	}

	if failed == 0 {
		fmt.Printf("GOOD %d files passed\n", len(args))
	}

	os.Exit(failed)
}
