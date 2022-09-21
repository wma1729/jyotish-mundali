package models

import (
	"jyotish/analysis"
	"time"
)

type Profile struct {
	ID          string
	Name        string
	DateOfBirth time.Time
	City        string
	State       string
	Country     string
	Details     analysis.GrahaDetais
}
