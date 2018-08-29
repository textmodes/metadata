package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// Pack represents an artpack
type Pack struct {
	Name    string            `yaml:"name"`    // Base name
	Year    int               `yaml:"year"`    // Year of publication
	Group   string            `yaml:"group"`   // Group slug (to be merged into Groups)
	Groups  []string          `yaml:"groups"`  // Group slugs
	Artist  string            `yaml:"artist"`  // Artist slug (to be merged into Artists)
	Artists []string          `yaml:"artists"` // Artist slugs
	Files   map[string]*File  `yaml:"files"`   // Directory contents
	Match   map[string]string `yaml:"match"`   // Globs to match filenames against artist slugs
}

// File represents a file in the archive
type File struct {
	Name    string    `yaml:"name"` // Base name
	ModTime time.Time `yaml:"date"`
	Artist  string    `yaml:"artist"`
	Artists []string  `yaml:"artists"`
	Font    string    `yaml:"font"`
}

func test(name string) bool {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL %s: %v\n", name, err)
		return false
	}

	pack := new(Pack)
	if err = yaml.UnmarshalStrict(b, pack); err != nil {
		fmt.Fprintf(os.Stderr, "FAIL %s: %v\n", name, err)
		return false
	}

	if pack.Match != nil {
		for pattern := range pack.Match {
			if !testGlob(pattern) {
				fmt.Fprintf(os.Stderr, "FAIL %s: invalid pattern %q\n", name, pattern)
				return false
			}
		}
	}

	if pack.Files != nil {
		for fileName, file := range pack.Files {
			if file.Font != "" && !testFont(file.Font) {
				fmt.Fprintf(os.Stderr, "FAIL %s: file %s has invalid font %q\n", name, fileName, file.Font)
				return false
			}
		}
	}

	if verbose {
		fmt.Printf("GOOD %s\n", name)
	}
	return true
}

var fonts = map[string]bool{
	// Official fonts
	"cp437 8x8":       true,
	"cp437 8x16":      true,
	"cp866 8x16":      true,
	"microknight":     true,
	"microknightplus": true,
	"mo'soul":         true,
	"p0t noodle":      true,
	"topaz a500":      true,
	"topazplus a500":  true,
	"topaz a1200":     true,
	"topazplus a1200": true,
	// Alias names
	"cp437":        true,
	"microknight+": true,
	"mosoul":       true,
	"mo soul":      true,
	"p0tnoodle":    true,
	"topaz":        true,
	"topaz+":       true,
	"topaz2":       true,
	"topaz2+":      true,
	"topaz+ a500":  true,
	"topaz+ a1200": true,
}

func testFont(name string) bool {
	name = strings.ToLower(name)
	for _, s := range []string{"_", "-"} {
		name = strings.Replace(name, s, " ", -1)
	}
	_, found := fonts[name]
	return found
}

func testGlob(pattern string) bool {
	pattern = regexp.QuoteMeta(pattern)
	pattern = strings.Replace(pattern, `\?`, ".", -1)
	pattern = strings.Replace(pattern, `\*`, ".*", -1)
	_, err := regexp.Compile(pattern)
	return err == nil
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
