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
