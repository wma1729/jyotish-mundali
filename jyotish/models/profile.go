package models

import (
	"time"
)

type Profile struct {
	ID          string
	Name        string
	DateOfBirth time.Time
	City        string
	State       string
	Country     string
	Details     GrahasLocation
}
