package types

import "time"

type Series struct {
	Id           int       `json:"id"`
	Series_name  string    `json:"series_name"`
	Release_date time.Time `json:"release_date"`
	Director     *Director `json:"director"`
	Season       []Season  `json:"season"`
}
