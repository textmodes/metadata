package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"

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

	if artist.Aliases != nil {
		for _, slug := range artist.Aliases {
			if !testSlug(slug) {
				fmt.Fprintf(os.Stderr, "FAIL %s: invalid slug %q\n", name, slug)
				return false
			}
		}
	}

	if artist.Social != nil {
		for site, value := range artist.Social {
			if err = testSocial(site, value); err != nil {
				fmt.Fprintf(os.Stderr, "FAIL %s: %v\n", name, err)
				return false
			}
		}
	}

	if verbose {
		fmt.Printf("GOOD %s\n", name)
	}
	return true
}

func testSlug(slug string) bool {
	// TODO(maze): more elaborate testing
	return url.PathEscape(slug) == slug
}

var (
	nonEmpty = regexp.MustCompile(`.+`)
	numeric  = regexp.MustCompile(`^\d+$`)
	social   = map[string]*regexp.Regexp{
		"artcity":         numeric,
		"behance":         nonEmpty,
		"csdb":            numeric,
		"demozoo":         numeric,
		"deviantart":      nonEmpty,
		"facebook":        nonEmpty,
		"flickr":          regexp.MustCompile(`^\d+@N\d+$`),
		"github":          regexp.MustCompile(`^\w[-\w]*$`),
		"google+":         regexp.MustCompile(`^(?:\+\w*|\d+)$`),
		"instagram":       nonEmpty,
		"linkedin":        regexp.MustCompile(`^\w[-\w]*$`),
		"pinterest":       nonEmpty,
		"pouet":           numeric,
		"twitter":         nonEmpty,
		"vimeo":           regexp.MustCompile(`^[a-zA-Z]\w+$`),
		"youtube":         regexp.MustCompile(`^[a-zA-Z]\w+$`),
		"youtube-channel": regexp.MustCompile(`^[-A-Za-z0-9+/]+$`),
	}
)

func testSocial(site, value string) error {
	re, found := social[site]
	if !found {
		return fmt.Errorf("social site %q invalid", site)
	}
	if !re.MatchString(value) {
		return fmt.Errorf("social site %s value %q invalid", site, value)
	}
	return nil
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
