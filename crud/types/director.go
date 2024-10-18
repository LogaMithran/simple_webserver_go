package types

import "time"

type Director struct {
	Director_name string    `json:"director_name"`
	Date_of_birth time.Time `json:"date_of_birth"`
}
