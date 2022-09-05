package models

const MAX_BHAVA_NUM int = 12

type Graha struct {
	Name       string  `json:"name"`
	RashiNum   int     `json:"rashi"`
	Degree     float32 `json:"degrees"`
	Retrograde bool    `json:"retrograde"`
}

type Bhava struct {
	Number    int
	RashiNum  int
	RashiLord string
	Grahas    []Graha
}
