package url

import (
	"testing"
)

func TestCharacterID(t *testing.T) {
	u := CharacterID([]string{"Random J Farmer", "Rixx Javix"})
	exp := URL(xmlApiUrl + "/eve/CharacterID.xml.aspx?names=Random+J+Farmer,Rixx+Javix")
	if u != exp {
		t.Errorf("TestCharacterID: expected=%#v actual=%#v", exp, u)
	}
}

func TestCharacterInfo(t *testing.T) {
	u := CharacterInfo(666)
	exp := URL(xmlApiUrl + "/eve/CharacterInfo.xml.aspx?characterID=666")
	if u != exp {
		t.Errorf("TestCharacterID: expected=%#v actual=%#v", exp, u)
	}
}

func TestMapKills(t *testing.T) {
	u := MapKills()
	exp := URL(xmlApiUrl + "/map/kills.xml.aspx")
	if u != exp {
		t.Errorf("TestMapKills: exp=%#v actual=%#v", exp, u)
	}
}

func TestMapJumps(t *testing.T) {
	u := MapJumps()
	exp := URL(xmlApiUrl + "/map/jumps.xml.aspx")
	if u != exp {
		t.Errorf("TestMapJumps: exp=%#v actual=%#v", exp, u)
	}
}
