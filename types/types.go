// Package types defines types exported to the client.
package types

import (
	"time"

	"github.com/pkg/errors"
)

// EveDatetime is a string in eve specific format
type EveDatetime string

// CharacterInfo from CharacterInfo endpoint
type CharacterInfo struct {
	CharacterID     uint64       `xml:"result>characterID"`
	CharacterName   string       `xml:"result>characterName"`
	CorporationID   uint64       `xml:"result>corporationID"`
	CorporationName string       `xml:"result>corporation"`
	CorporationDate EveDatetime  `xml:"result>corporationDate"`
	AllianceID      uint64       `xml:"result>allianceID"`
	AllianceName    string       `xml:"result>alliance"`
	AllianceDate    EveDatetime  `xml:"result>allianceDate"`
	SecurityStatus  float64      `xml:"result>securityStatus"`
	Employment      []Employment `xml:"result>rowset>row"`
	CurrentTime     EveDatetime  `xml:"currentTime"`
	CachedUntil     EveDatetime  `xml:"cachedUntil"`
}

// Employment gives historic employment information for the CharacterInfo endpoint
type Employment struct {
	CorporationID   string      `xml:"corporationID,attr"`
	CorporationName string      `xml:"corporationName,attr"`
	StartDate       EveDatetime `xml:"startDate,attr"`
}

// MapKills is the data from the /map/kills endpoint
type MapKills struct {
	DataTime    EveDatetime        `xml:"result>dataTime"`
	Kills       []SolarSystemKills `xml:"result>rowset>row"`
	CurrentTime EveDatetime        `xml:"currentTime"`
	CachedUntil EveDatetime        `xml:"cachedUntil"`
}

// SolarSystemKills is a row from the /map/kills endpoint
type SolarSystemKills struct {
	SolarSystemID uint64 `xml:"solarSystemID,attr"`
	ShipKills     int    `xml:"shipKills,attr"`
	FactionKills  int    `xml:"factionKills,attr"`
	PodKills      int    `xml:"podKills,attr"`
}

// ParseTime parses an Eve Timestamp.
// EVE Times are always given in UTC
func ParseTime(dt EveDatetime) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", string(dt))
	if err != nil {
		panic(errors.Wrap(err, "ParseEveDateTime"))
	}
	return t
}
