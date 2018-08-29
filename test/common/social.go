package common

import (
	"fmt"
	"regexp"
)

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

// TestSocial checks a social site and value for correctness.
func TestSocial(site, value string) error {
	re, found := social[site]
	if !found {
		return fmt.Errorf("social site %q invalid", site)
	}
	if !re.MatchString(value) {
		return fmt.Errorf("social site %s value %q invalid", site, value)
	}
	return nil
}
