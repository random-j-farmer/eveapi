package charid

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseBody(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/character_id_random_j_farmer.xml")
	if err != nil {
		t.Errorf("ReadFile: %#v", err)
	}

	m, err := ParseBody([]string{"Random J Farmer"}, b)
	if err != nil {
		t.Errorf("ParseBody: %#v", err)
	}

	exp := map[string]uint64{"Random J Farmer": 95538430}

	if !reflect.DeepEqual(exp, m) {
		t.Errorf("TestLookup: expexted=%#v, actual=%#v", exp, m)
	}

}
