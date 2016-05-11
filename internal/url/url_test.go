package url

import (
	"testing"
)

func TestCharacterID(t *testing.T) {
	s := CharacterID([]string{"Random J Farmer", "Rixx Javix"})
	exp := URL(xmlApiUrl + "/eve/CharacterID.xml.aspx?names=Random+J+Farmer,Rixx+Javix")
	if s != exp {
		t.Errorf("TestCharacterID: expected=%#v actual=%#v", exp, s)
	}
}
