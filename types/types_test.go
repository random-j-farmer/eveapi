package types

import "testing"

func TestParseEveTime(t *testing.T) {
	ts := ParseTime("2016-05-11 20:15:00")

	y, m, d := ts.Date()
	hh, mm, ss := ts.Clock()

	if y != 2016 { t.Errorf("%v != %v", y, 2016) }
	if m != 5 { t.Errorf("%v != %v", m, 5) }
	if d != 11 { t.Errorf("%v != %v", d, 11) }
	if hh != 20 { t.Errorf("%v != %v", hh, 20) }
	if mm != 15 { t.Errorf("%v != %v", mm, 15) }
	if ss != 0 { t.Errorf("%v != %v", ss, 0) }
}
