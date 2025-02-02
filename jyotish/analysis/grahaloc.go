package analysis

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type GrahaLoc struct {
	Name       string  `json:"name"`
	RashiNum   int     `json:"rashi"`
	Degree     float32 `json:"degrees"`
	Retrograde bool    `json:"retrograde"`
}

type GrahasLocation struct {
	Grahas []GrahaLoc
}

/*
 * The receiver must be non-pointer for this.
 * Do not change.
 */
func (gl GrahasLocation) Value() (driver.Value, error) {
	return json.Marshal(gl)
}

func (gl *GrahasLocation) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unexpected value type: expected []byte, found %T", value)
	}
	json.Unmarshal(b, gl)
	return nil
}

func (gl *GrahasLocation) GetLagnaRashi() int {
	for _, graha := range gl.Grahas {
		if graha.Name == LAGNA {
			return graha.RashiNum
		}
	}
	return -1
}
