package constants

const (
	MAX_BHAVA_NUM = 12
)

const (
	BHAVA_PLACEMENT = 0
	BHAVA_OWNERSHIP = 1
	BHAVA_KARAKA    = 2
	BHAVA_ASPECT    = 3
)

const (
	SUBJECTS_ALL              = "all"
	SUBJECTS_LIVING_BEING     = "living-beings"
	SUBJECTS_NON_LIVING_BEING = "non-living-beings"
	SUBJECTS_YOUNGER_SIBLINGS = "younger-siblings"
	SUBJECTS_ELDER_SIBLINGS   = "elder-siblings"
	SUBJECTS_MOTHER           = "mother"
	SUBJECTS_FATHER           = "father"
	SUBJECTS_CHILDREN         = "children"
	SUBJECTS_SPOUSE           = "spouse"
	BHAVA_LORD                = "bhava-lord"
)

var BhavaKarakas = map[int][]string{
	1:  {SUN},
	2:  {JUPITER},
	3:  {MARS},
	4:  {MARS, MERCURY, MOON, VENUS},
	5:  {JUPITER},
	6:  {MARS, SATURN},
	7:  {VENUS},
	8:  {SATURN},
	9:  {JUPITER, SUN},
	10: {JUPITER, MERCURY, SATURN, SUN},
	11: {JUPITER},
	12: {SATURN},
}

func IsBadBhava(n int) bool {
	return (n == 3) || (n == 6) || (n == 8) || (n == 12)
}

func IsGoodBhava(n int) bool {
	return (n == 1) || (n == 4) || (n == 5) || (n == 7) || (n == 9) || (n == 10)
}
