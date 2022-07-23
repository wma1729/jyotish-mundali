package models

import "time"

type ProfileDetails struct {
	Planet  string  `json:"planet"`
	Rashi   int     `json:"rashi"`
	Degrees float32 `json:"degrees"`
}

type Profile struct {
	ID          string
	Name        string
	DateOfBirth time.Time
	City        string
	State       string
	Country     string
	Details     []ProfileDetails
}
