package analysis

const MAX_BHAVA_NUM int = 12

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
	ARIES       = 1
	TAURUS      = 2
	GEMINI      = 3
	CANCER      = 4
	LEO         = 5
	VIRGO       = 6
	LIBRA       = 7
	SCORPIO     = 8
	SAGITTARIUS = 9
	CAPRICORN   = 10
	AQUARIUS    = 11
	PISCES      = 12
)

var RashiLordMap = map[int]string{
	ARIES:       MARS,
	TAURUS:      VENUS,
	GEMINI:      MERCURY,
	CANCER:      MOON,
	LEO:         SUN,
	VIRGO:       MERCURY,
	LIBRA:       VENUS,
	SCORPIO:     MARS,
	SAGITTARIUS: JUPITER,
	CAPRICORN:   SATURN,
	AQUARIUS:    SATURN,
	PISCES:      JUPITER,
}

type RashiRange struct {
	RashiNum int
	Min      int
	Max      int
}

type GrahaAttr struct {
	Exaltation   RashiRange
	Debilitation RashiRange
	Trinal       RashiRange
	Friends      []string
	Neutrals     []string
	Enemies      []string
}

var GrahaAttrMap = map[string]GrahaAttr{
	SUN: {
		Exaltation:   RashiRange{ARIES, 0, 10},
		Debilitation: RashiRange{LIBRA, 0, 10},
		Trinal:       RashiRange{LEO, 0, 20},
		Friends:      []string{MOON, MARS, JUPITER},
		Neutrals:     []string{MERCURY},
		Enemies:      []string{VENUS, SATURN},
	},
	MOON: {
		Exaltation:   RashiRange{TAURUS, 0, 3},
		Debilitation: RashiRange{SCORPIO, 0, 3},
		Trinal:       RashiRange{TAURUS, 4, 30},
		Friends:      []string{SUN, MERCURY},
		Neutrals:     []string{MARS, JUPITER, VENUS, SATURN},
		Enemies:      []string{},
	},
	MARS: {
		Exaltation:   RashiRange{CAPRICORN, 0, 28},
		Debilitation: RashiRange{CANCER, 0, 28},
		Trinal:       RashiRange{ARIES, 0, 12},
		Friends:      []string{SUN, MOON, JUPITER},
		Neutrals:     []string{VENUS, SATURN},
		Enemies:      []string{MERCURY},
	},
	MERCURY: {
		Exaltation:   RashiRange{VIRGO, 0, 15},
		Debilitation: RashiRange{PISCES, 0, 15},
		Trinal:       RashiRange{VIRGO, 16, 30},
		Friends:      []string{SUN, VENUS},
		Neutrals:     []string{MARS, JUPITER, SATURN},
		Enemies:      []string{MOON},
	},
	JUPITER: {
		Exaltation:   RashiRange{CANCER, 0, 5},
		Debilitation: RashiRange{CAPRICORN, 0, 5},
		Trinal:       RashiRange{SAGITTARIUS, 0, 10},
		Friends:      []string{SUN, MOON, MARS},
		Neutrals:     []string{SATURN},
		Enemies:      []string{MERCURY, VENUS},
	},
	VENUS: {
		Exaltation:   RashiRange{PISCES, 0, 27},
		Debilitation: RashiRange{VIRGO, 0, 27},
		Trinal:       RashiRange{LIBRA, 0, 15},
		Friends:      []string{MERCURY, SATURN},
		Neutrals:     []string{MARS, JUPITER},
		Enemies:      []string{SUN, MOON},
	},
	SATURN: {
		Exaltation:   RashiRange{LIBRA, 0, 20},
		Debilitation: RashiRange{ARIES, 0, 20},
		Trinal:       RashiRange{AQUARIUS, 0, 20},
		Friends:      []string{MERCURY, VENUS},
		Neutrals:     []string{JUPITER},
		Enemies:      []string{SUN, MOON, MARS},
	},
	RAHU: {
		Exaltation:   RashiRange{-1, 0, 0},
		Debilitation: RashiRange{-1, 0, 0},
		Trinal:       RashiRange{-1, 0, 0},
		Friends:      []string{MERCURY, VENUS, SATURN},
		Neutrals:     []string{MARS},
		Enemies:      []string{SUN, MOON, JUPITER},
	},
	KETU: {
		Exaltation:   RashiRange{-1, 0, 0},
		Debilitation: RashiRange{-1, 0, 0},
		Trinal:       RashiRange{-1, 0, 0},
		Friends:      []string{MERCURY, VENUS, SATURN},
		Neutrals:     []string{MARS},
		Enemies:      []string{SUN, MOON, JUPITER},
	},
}

const (
	BENEFIC = "benefic"
	MALEFIC = "malefic"
)
