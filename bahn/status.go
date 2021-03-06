package bahn

import (
	"encoding/json"
	"net/http"
	"time"
)

// StatusURL is the URL of the status API call.
const StatusURL = "http://ice.portal/jetty/api/v1/status"

// Status holds the information returned by the status API call.
type Status struct {
	Connection   bool
	ServiceLevel string
	Speed        float64
	Longitude    float64
	Latitude     float64
	ServerTime   time.Time
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface.
func (s *Status) UnmarshalJSON(b []byte) error {
	var parsed struct {
		Connection   bool
		ServiceLevel string
		Speed        float64
		Longitude    float64
		Latitude     float64
		ServerTime   int
	}
	if err := json.Unmarshal(b, &parsed); err != nil {
		return err
	}

	*s = Status{
		Connection:   parsed.Connection,
		ServiceLevel: parsed.ServiceLevel,
		Speed:        parsed.Speed,
		Longitude:    parsed.Longitude,
		Latitude:     parsed.Latitude,
		ServerTime:   time.Unix(int64(parsed.ServerTime/1000), 0),
	}

	return nil
}

// StatusInfo calls the status API and returns the parsed data.
func StatusInfo() (*Status, error) {
	res, err := http.Get(StatusURL)
	if err != nil {
		return nil, err
	}

	var s Status
	if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
		return nil, err
	}

	return &s, nil
}
