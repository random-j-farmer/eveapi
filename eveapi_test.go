package eveapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/eveapi/internal/url"
)

// testGetter is a getter with a fixed result
type testGetter struct {
	body   []byte
	err    error
	reader *bytes.Reader
}

func newTestGetter(fn string) *testGetter {
	b, err := ioutil.ReadFile(fn)
	return &testGetter{b, err, bytes.NewReader(b)}
}

func (tg *testGetter) Close() error { return nil }
func (tg *testGetter) Read(b []byte) (int, error) {
	return tg.reader.Read(b)
}

func (tg *testGetter) Get(url string) (resp *http.Response, err error) {
	resp = new(http.Response)
	resp.Body = tg
	resp.StatusCode = 200
	return resp, tg.err
}

func TestNewClient(t *testing.T) {
	c := NewClient(ClientConfig{Getter: newTestGetter("no_such_file.xxx")})
	if c == nil {
		t.Error("NewClient==nil")
	}

	// test getter with no such file - error
	_, err := c.getBody(url.URL("xxx"))
	if err == nil {
		t.Error("test getter with no file should have given error")
	}
}

func TestClient_CharacterID(t *testing.T) {
	c := NewClient(*new(ClientConfig))

	cis, err := c.CharacterID([]string{"Random J Farmer"})
	if err != nil {
		t.Errorf("ParseBody: %#v", err)
	}

	exp := map[string]uint64{"Random J Farmer": 95538430}

	if !reflect.DeepEqual(exp, cis) {
		t.Errorf("TestLookup: expexted=%#v, actual=%#v", exp, cis)
	}
}

func TestClient_CharacterInfo(t *testing.T) {
	c := NewClient(*new(ClientConfig))

	info, err := c.CharacterInfo(95538430)
	if err != nil {
		t.Errorf("ParseBody: %v cause %v", err, errors.Cause(err))
	}

	if info.CharacterName != "Random J Farmer" {
		t.Errorf("TestParseBody: expected=%#v, actual=%#v", "Random J Farmer", info.CharacterName)
	}

}

/*

this one takes a 15 seconds ... too long

func TestClient_MapKills(t *testing.T) {
	c := NewClient(*new(ClientConfig))

	_, err := c.MapKills()
	if err != nil {
		t.Errorf("MapKills: %#v cause %#v", err, errors.Cause(err))
	}
}
*/
