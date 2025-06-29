package constants

const NEUTRAL = 0

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

const (
	BENEFIC = 1
	MALEFIC = 2
)

const (
	IN_EXALTATION_RASHI = iota
	IN_DEBILITATION_RASHI
	IN_MOOLTRIKONA_RASHI
	IN_OWN_RASHI
	IN_FRIENDLY_RASHI
	IN_NEUTRAL_RASHI
	IN_INIMICAL_RASHI
)

const (
	FRIEND = 1
	ENEMY  = 2
)

const (
	CHILD = iota
	YOUTH
	ADULT
	OLD
	DEAD
)

var GrahaNames = []string{
	SUN,
	MOON,
	MARS,
	MERCURY,
	JUPITER,
	VENUS,
	SATURN,
	RAHU,
	KETU,
}
