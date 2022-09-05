package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type GrahaDetails struct {
	Grahas []Graha
}

func (chart GrahaDetails) Value() (driver.Value, error) {
	return json.Marshal(chart)
}

func (chart *GrahaDetails) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("unexpected value type: expected []byte, found %T", value))
	}
	json.Unmarshal(b, chart)
	return nil
}

type Profile struct {
	ID          string
	Name        string
	DateOfBirth time.Time
	City        string
	State       string
	Country     string
	Details     GrahaDetails
}
