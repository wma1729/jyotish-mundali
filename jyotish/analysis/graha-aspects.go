package analysis

type GrahaAspects struct {
	Name               string
	FullAspect         []string
	ThreeQuarterAspect []string
	HalfAspect         []string
	QuarterAspect      []string
}

func (aspects *GrahaAspects) EvaluateGrahaAspects(name string, chart *Chart) {
	aspects.Name = name
	_, b := chart.GetGrahaBhava(name)
	if b == nil {
		return
	}

	aspects.FullAspect = b.FullAspect
	aspects.ThreeQuarterAspect = b.ThreeQuarterAspect
	aspects.HalfAspect = b.HalfAspect
	aspects.QuarterAspect = b.QuarterAspect
}
