package models

const (
	LAGNA   = "lagna"
	SUN     = "sun"
	MOON    = "moon"
	MARS    = "mars"
	MERCURY = "mercury"
	JUPITER = "jupiter"
	VENUS   = "venus"
	SATURN  = "saturn"
	RAHU    = "rahu"
	KETU    = "ketu"
)

var RashiLordMap = map[int]string{
	1:  MARS,
	2:  VENUS,
	3:  MERCURY,
	4:  MOON,
	5:  SUN,
	6:  MERCURY,
	7:  VENUS,
	8:  MARS,
	9:  JUPITER,
	10: SATURN,
	11: SATURN,
	12: JUPITER,
}
