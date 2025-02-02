package constants

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

const (
	RASHI_EXALTED     = "exalted"
	RASHI_DEBILITATED = "debilitated"
	RASHI_MOOLTRIKONA = "mool-trikona"
	RASHI_OWN         = "own-rashi"
	RASHI_FRIENDLY    = "friendly-rashi"
	RASHI_NEUTRAL     = "neutral-rashi"
	RASHI_ENEMY       = "enemy-rashi"
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

type RangeInRashi struct {
	RashiNum  int
	MinDegree int
	MaxDegree int
}

type GrahaBalaInRashiRules struct {
	Exaltation   RangeInRashi
	Debilitation RangeInRashi
	Trinal       RangeInRashi
	Owner        []int
	Friends      []string
	Neutrals     []string
	Enemies      []string
}

var GrahaBalaInRashiRulesMap = map[string]GrahaBalaInRashiRules{
	SUN: {
		Exaltation:   RangeInRashi{ARIES, 0, 10},
		Debilitation: RangeInRashi{LIBRA, 0, 10},
		Trinal:       RangeInRashi{LEO, 0, 20},
		Owner:        []int{LEO},
		Friends:      []string{MOON, MARS, JUPITER},
		Neutrals:     []string{MERCURY},
		Enemies:      []string{VENUS, SATURN},
	},
	MOON: {
		Exaltation:   RangeInRashi{TAURUS, 0, 3},
		Debilitation: RangeInRashi{SCORPIO, 0, 3},
		Trinal:       RangeInRashi{TAURUS, 4, 30},
		Owner:        []int{CANCER},
		Friends:      []string{SUN, MERCURY},
		Neutrals:     []string{MARS, JUPITER, VENUS, SATURN},
		Enemies:      []string{},
	},
	MARS: {
		Exaltation:   RangeInRashi{CAPRICORN, 0, 28},
		Debilitation: RangeInRashi{CANCER, 0, 28},
		Trinal:       RangeInRashi{ARIES, 0, 12},
		Owner:        []int{ARIES, SCORPIO},
		Friends:      []string{SUN, MOON, JUPITER},
		Neutrals:     []string{VENUS, SATURN},
		Enemies:      []string{MERCURY},
	},
	MERCURY: {
		Exaltation:   RangeInRashi{VIRGO, 0, 15},
		Debilitation: RangeInRashi{PISCES, 0, 15},
		Trinal:       RangeInRashi{VIRGO, 16, 30},
		Owner:        []int{GEMINI, VIRGO},
		Friends:      []string{SUN, VENUS},
		Neutrals:     []string{MARS, JUPITER, SATURN},
		Enemies:      []string{MOON},
	},
	JUPITER: {
		Exaltation:   RangeInRashi{CANCER, 0, 5},
		Debilitation: RangeInRashi{CAPRICORN, 0, 5},
		Trinal:       RangeInRashi{SAGITTARIUS, 0, 10},
		Owner:        []int{SAGITTARIUS, PISCES},
		Friends:      []string{SUN, MOON, MARS},
		Neutrals:     []string{SATURN},
		Enemies:      []string{MERCURY, VENUS},
	},
	VENUS: {
		Exaltation:   RangeInRashi{PISCES, 0, 27},
		Debilitation: RangeInRashi{VIRGO, 0, 27},
		Trinal:       RangeInRashi{LIBRA, 0, 15},
		Owner:        []int{TAURUS, LIBRA},
		Friends:      []string{MERCURY, SATURN},
		Neutrals:     []string{MARS, JUPITER},
		Enemies:      []string{SUN, MOON},
	},
	SATURN: {
		Exaltation:   RangeInRashi{LIBRA, 0, 20},
		Debilitation: RangeInRashi{ARIES, 0, 20},
		Trinal:       RangeInRashi{AQUARIUS, 0, 20},
		Owner:        []int{CAPRICORN, AQUARIUS},
		Friends:      []string{MERCURY, VENUS},
		Neutrals:     []string{JUPITER},
		Enemies:      []string{SUN, MOON, MARS},
	},
	RAHU: {
		Exaltation:   RangeInRashi{-1, 0, 0},
		Debilitation: RangeInRashi{-1, 0, 0},
		Trinal:       RangeInRashi{-1, 0, 0},
		Owner:        []int{},
		Friends:      []string{MERCURY, VENUS, SATURN},
		Neutrals:     []string{MARS},
		Enemies:      []string{SUN, MOON, JUPITER},
	},
	KETU: {
		Exaltation:   RangeInRashi{-1, 0, 0},
		Debilitation: RangeInRashi{-1, 0, 0},
		Trinal:       RangeInRashi{-1, 0, 0},
		Owner:        []int{},
		Friends:      []string{MERCURY, VENUS, SATURN},
		Neutrals:     []string{MARS},
		Enemies:      []string{SUN, MOON, JUPITER},
	},
}
