// Package mapkills parses map/kills body data
package mapkills

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/eveapi/types"
)

// ParseBody parses the map/kills endpoint
func ParseBody(body []byte) (*types.MapKills, error) {
	var kills types.MapKills
	err := xml.Unmarshal(body, &kills)
	if err != nil {
		return &kills, errors.Wrap(err, "Unmarshal")
	}

	return &kills, nil
}
