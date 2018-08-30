package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

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
	Name    string   `yaml:"name"` // Base name
	Artist  string   `yaml:"artist"`
	Artists []string `yaml:"artists"`
	Font    string   `yaml:"font"`
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
	"ibm vga":            true,
	"ibm vga50":          true,
	"ibm vga25g":         true,
	"ibm ega":            true,
	"ibm ega43":          true,
	"amiga topaz 1":      true,
	"amiga topaz 1+":     true,
	"amiga topaz 2":      true,
	"amiga topaz 2+":     true,
	"amiga p0t-noodle":   true,
	"amiga microknight":  true,
	"amiga microknight+": true,
	"amiga mosoul":       true,
	// Alias names
	"cp437":           true,
	"cp866":           true,
	"microknight":     true,
	"microknight+":    true,
	"microknightplus": true,
	"mo soul":         true,
	"mo'soul":         true,
	"mosoul":          true,
	"p0t noodle":      true,
	"p0tnoodle":       true,
	"topaz a1200":     true,
	"topaz a500":      true,
	"topaz":           true,
	"topaz+ a1200":    true,
	"topaz+ a500":     true,
	"topaz+":          true,
	"topaz2":          true,
	"topaz2+":         true,
	"topazplus a1200": true,
	"topazplus a500":  true,
}

func init() {
	for _, codePage := range []string{
		"437", // The character set of the original IBM PC. Also known as ‘MS-DOS Latin US’.
		"720", // Arabic. Also known as ‘Windows-1256’.
		"737", // Greek. Also known as ‘MS-DOS Greek’.
		"775", // Baltic Rim (Estonian, Lithuanian and Latvian). Also known as ‘MS-DOS Baltic Rim’.
		"819", // Latin-1 Supplemental. Also known as ‘Windows-28591’ and ‘ISO/IEC 8859-1’.
		"850", // Western Europe. Also known as ‘MS-DOS Latin 1’.
		"852", // Central Europe (Bosnian, Croatian, Czech, Hungarian, Polish, Romanian, Serbian and Slovak). Also known as ‘MS-DOS Latin 2’.
		"855", // Cyrillic (Serbian, Macedonian Bulgarian, Russian). Also known as ‘MS-DOS Cyrillic’.
		"857", // Turkish. Also known as ‘MS-DOS Turkish’.
		"858", // Western Europe.
		"860", // Portuguese. Also known as ‘MS-DOS Portuguese’.
		"861", // Icelandic. Also known as ‘MS-DOS Icelandic’.
		"862", // Hebrew. Also known as ‘MS-DOS Hebrew’.
		"863", // French Canada. Also known as ‘MS-DOS French Canada’.
		"864", // Arabic.
		"865", // Nordic.
		"866", // Cyrillic.
		"869", // Greek 2. Also known as ‘MS-DOS Greek 2’.
		"872", // Cyrillic.
		"KAM", // ‘Kamenický’ encoding. Also known as ‘KEYBCS2’.
		"MAZ", // ‘Mazovia’ encoding.
		"MIK", // Cyrillic.
	} {
		fonts["ibm vga "+strings.ToLower(codePage)] = true
	}
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
