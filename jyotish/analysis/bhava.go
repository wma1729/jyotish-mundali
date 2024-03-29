package analysis

type Graha struct {
	Name       string  `json:"name"`
	RashiNum   int     `json:"rashi"`
	Degree     float32 `json:"degrees"`
	Retrograde bool    `json:"retrograde"`
}

type Bhava struct {
	Number            int
	RashiNum          int
	RashiLord         string
	Grahas            []Graha
	FullAspect        []string
	ThreeQuaterAspect []string
	HalfAspect        []string
	QuaterAspect      []string
}

func (b *Bhava) ContainsGraha(name string) bool {
	for _, g := range b.Grahas {
		if g.Name == name {
			return true
		}
	}
	return false
}

func (b *Bhava) GrahaByName(name string) *Graha {
	for _, g := range b.Grahas {
		if g.Name == name {
			return &g
		}
	}
	return nil
}

func (b *Bhava) IsRetrograde(name string) bool {
	graha := b.GrahaByName(name)
	if graha != nil {
		return graha.Retrograde
	}
	return false
}
