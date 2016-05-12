// Package charinfo implements the eve/CharacterInfo endpoint
package charinfo

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"fmt"
	"github.com/random-j-farmer/eveapi/types"
)

func ParseBody(id uint64, body []byte) (*types.CharacterInfo, error) {
	var charinfo types.CharacterInfo
	err := xml.Unmarshal(body, &charinfo)
	if err != nil {
		return &charinfo, errors.Wrap(err, fmt.Sprintf("Unmarshal(%d)", id))
	}

	return &charinfo, nil
}

