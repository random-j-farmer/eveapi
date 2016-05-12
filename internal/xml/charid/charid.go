package charid

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func lookupError(names []string, err error) error {
	return errors.Wrap(err, fmt.Sprintf("Lookup(%s)", strings.Join(names, ", ")))
}

type row struct {
	Name        string `xml:"name,attr"`
	CharacterID uint64 `xml:"characterID,attr"`
}

type eveApi struct {
	XmlName     xml.Name `xml:"eveapi"`
	CurrentTime string   `xml:"currentTime"`
	CachedUntil string   `xml:"cachedUntil"`
	Row         []row    `xml:"result>rowset>row"`
}

func ParseBody(names []string, body []byte) (map[string]uint64, error) {
	var eveApi eveApi
	err := xml.Unmarshal(body, &eveApi)
	if err != nil {
		return nil, lookupError(names, err)
	}

	var cis = make(map[string]uint64, len(names))
	for _, row := range eveApi.Row {
		cis[row.Name] = row.CharacterID
	}

	return cis, nil
}
