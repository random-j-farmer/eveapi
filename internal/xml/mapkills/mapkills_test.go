package mapkills

import (
	"io/ioutil"
	"testing"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/eveapi/types"
)

func Test_ParseBody(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/mapkills.xml")
	if err != nil {
		t.Errorf("ReadFile: %#v", err)
	}

	kills, err := ParseBody(b)
	if err != nil {
		t.Errorf("ParseBody: %v cause %v", err, errors.Cause(err))
	}

	cachedUntil := "2016-09-08 21:23:08"
	if kills.CachedUntil != types.EveDatetime(cachedUntil) {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", cachedUntil, kills.CachedUntil)
	}

	id := uint64(30002671)
	if id != kills.Kills[0].SolarSystemID {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", id, kills.Kills[0].SolarSystemID)
	}

	ship := 0
	if ship != kills.Kills[0].ShipKills {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", ship, kills.Kills[0].ShipKills)
	}

	fac := 41
	if fac != kills.Kills[0].FactionKills {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", fac, kills.Kills[0].FactionKills)
	}

	pod := 1
	if pod != kills.Kills[0].PodKills {
		t.Errorf("Test_ParseBody: expected=%#v, actual=%#v", id, kills.Kills[0].PodKills)
	}

}
