// Package eveapi provides access to the EVE Online XML API
package eveapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/random-j-farmer/eveapi/internal/url"
	"github.com/random-j-farmer/eveapi/internal/xml/charid"
	"github.com/random-j-farmer/eveapi/internal/xml/charinfo"
	"github.com/random-j-farmer/eveapi/internal/xml/mapkills"
	"github.com/random-j-farmer/eveapi/types"
)

// Getter gets an URL
type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

// ClientConfig is getter & looger
type ClientConfig struct {
	Getter Getter
	Log    *log.Logger
}

// Client is the same, but unexported
type Client struct {
	getter Getter
	log    *log.Logger
}

// NewClient returns a client for the config
func NewClient(cfg ClientConfig) *Client {
	c := Client{getter: cfg.Getter, log: cfg.Log}
	if c.getter == nil {
		c.getter = http.DefaultClient
	}
	if c.log == nil {
		c.log = log.New(ioutil.Discard, "", 0)
	}
	return &c
}

// CharacterID performs a character id lookup for the given names.
func (c *Client) CharacterID(names []string) (map[string]uint64, error) {

	url := url.CharacterID(names)

	body, err := c.getBody(url)
	if err != nil {
		return nil, errors.Wrap(err, "getBody")
	}

	return charid.ParseBody(names, body)
}

// CharacterInfo gives the character info for a character id.
func (c *Client) CharacterInfo(id uint64) (*types.CharacterInfo, error) {

	url := url.CharacterInfo(id)

	body, err := c.getBody(url)
	if err != nil {
		return nil, errors.Wrap(nil, "getBody")
	}

	return charinfo.ParseBody(id, body)
}

// MapKills queries the map/kills endpoint
func (c *Client) MapKills() (*types.MapKills, error) {
	url := url.MapKills()

	body, err := c.getBody(url)
	if err != nil {
		return nil, errors.Wrap(nil, "getBody")
	}

	return mapkills.ParseBody(body)
}

// Get a EVE Online XMLAPI URL
func (c *Client) getBody(url url.URL) ([]byte, error) {
	resp, err := c.getter.Get(string(url))
	if err != nil {
		return nil, errors.Wrap(err, "http.Get failed")
	}
	defer resp.Body.Close()

	if (resp.StatusCode / 100) != 2 {
		return nil, fmt.Errorf("StatusCode %d for %s", resp.StatusCode, url)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("ReadAll(%s)", url))
	}

	return body, nil
}
