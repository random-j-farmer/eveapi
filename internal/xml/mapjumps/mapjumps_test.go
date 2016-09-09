package mapjumps

import (
	"io/ioutil"
	"testing"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/eveapi/types"
)

func Test_ParseBody(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/mapjumps.xml")
	if err != nil {
		t.Errorf("ReadFile: %#v", err)
	}

	jumps, err := ParseBody(b)
	if err != nil {
		t.Errorf("ParseBody: %v cause %v", err, errors.Cause(err))
	}

	cachedUntil := "2016-09-09 10:55:00"
	if jumps.CachedUntil != types.EveDatetime(cachedUntil) {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", cachedUntil, jumps.CachedUntil)
	}

	id := uint64(30003999)
	if id != jumps.Jumps[0].SolarSystemID {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", id, jumps.Jumps[0].SolarSystemID)
	}

	numJumps := 3
	if numJumps != jumps.Jumps[0].ShipJumps {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", numJumps, jumps.Jumps[0].ShipJumps)
	}
}
