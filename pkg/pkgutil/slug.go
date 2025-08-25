package pkgutil

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)

func CreateSlug(title string) string {
	if title == "" {
		return ""
	}

	slug := strings.ToLower(title)

	slug = removeAccents(slug)

	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	reg = regexp.MustCompile(`-+`)
	slug = reg.ReplaceAllString(slug, "-")

	return slug
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, _ := transform.String(t, s)
	return output
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
