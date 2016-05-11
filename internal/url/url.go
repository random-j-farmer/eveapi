// Package url defines implementation details
package url

import (
	"fmt"
	"net/url"
	"strings"
)

// URL is a EVE Online XMLAPI URL
type URL string

// CharacterID gives the XMLAPI CharacterID Lookup URL for names.
func CharacterID(names []string) URL {
	return URL(fmt.Sprintf("%s?names=%s", urlHelper("eve/CharacterID"), quoteAndJoin(names, ",")))
}

const xmlApiUrl = "https://api.eveonline.com"

func urlHelper(s string) string {
	return fmt.Sprintf("%s/%s.xml.aspx", xmlApiUrl, s)
}

func quoteAndJoin(xs []string, sep string) string {
	esc := make([]string, len(xs))
	for i, s := range xs {
		esc[i] = url.QueryEscape(s)
	}
	return strings.Join(esc, sep)
}
