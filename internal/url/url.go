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

// CharacterInfo is the eve/CharacterInfo endpoint
func CharacterInfo(id uint64) URL {
	return URL(fmt.Sprintf("%s?characterID=%d", urlHelper("eve/CharacterInfo"), id))
}

// MapKills is the map/kills endpoint
func MapKills() URL {
	return URL(fmt.Sprintf("%s", urlHelper("map/kills")))
}

const xmlApiUrl = "https://api.eveonline.com"
const crestUrl = "https://crest-tq.eveonline.com"

func urlHelper(s string) string {
	return fmt.Sprintf("%s/%s.xml.aspx", xmlApiUrl, s)
}

func crestHelper(s string) string {
	return fmt.Sprintf("%s/%s", crestUrl, s)
}

func quoteAndJoin(xs []string, sep string) string {
	esc := make([]string, len(xs))
	for i, s := range xs {
		esc[i] = url.QueryEscape(s)
	}
	return strings.Join(esc, sep)
}
