// Package mapjumps parses map/jumps body data
package mapjumps

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/eveapi/types"
)

// ParseBody parses the map/kills endpoint
func ParseBody(body []byte) (*types.MapJumps, error) {
	var kills types.MapJumps
	err := xml.Unmarshal(body, &kills)
	if err != nil {
		return &kills, errors.Wrap(err, "Unmarshal")
	}

	return &kills, nil
}
