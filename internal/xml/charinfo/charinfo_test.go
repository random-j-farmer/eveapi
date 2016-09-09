package charinfo

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"testing"
)

func TestParseBody(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/random_j_farmer.xml")
	if err != nil {
		t.Errorf("ReadFile: %#v", err)
	}

	info, err := ParseBody(95538430, b)
	if err != nil {
		t.Errorf("ParseBody: %v cause %v", err, errors.Cause(err))
	}

	if info.CharacterName != "Random J Farmer" {
		t.Errorf("TestParseBody: expected=%#v, actual=%#v", "Random J Farmer", info.CharacterName)
	}

	if info.CorporationName != "Stay Frosty." {
		t.Errorf("TestParseBody: expected=%#v, actual=%#v", "Stay Frosty.", info.CorporationName)
	}

}
