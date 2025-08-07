package analysis

import (
	"jyotish/constants"
	"log"
)

func (c *Chart) GetGrahaAttributes(name string) *GrahaAttributes {
	for _, grahaAttr := range c.GrahasAttr {
		if grahaAttr.Name == name {
			return &grahaAttr
		}
	}
	log.Printf("unable to get attributes of graha %s", name)
	return nil
}

func (c *Chart) EvaluateGrahasAttributes() {
	c.GrahasAttr = make([]GrahaAttributes, 9)
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Name = graha
		_, bhava := c.GetGrahaBhava(graha)
		if bhava != nil {
			c.GrahasAttr[i].AbsoluteDegree = bhava.GrahaDegree(graha, true)
		}
	}
	c.evaluateGrahaRelations()
	c.evaluateGrahaAspects()
	c.evaluateGrahaNature()
	c.evaluateGrahaStrength()
}

func (c *Chart) evaluateGrahaRelations() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Relations.EvaluateGrahaRelations(graha, c)
	}
}

func (c *Chart) evaluateGrahaAspects() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Aspects.EvaluateGrahaAspects(graha, c)
	}
}

func (c *Chart) evaluateGrahaNature() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Nature.EvaluateGrahaNature(graha, c)
	}

	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Nature.EvaluateGrahaFunctionalNature(graha, c)
	}
}

func (c *Chart) evaluateGrahaStrength() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Strength.EvaluateGrahaStrength(graha, c)
	}
}
