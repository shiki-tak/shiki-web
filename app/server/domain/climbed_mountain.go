package domain

import "time"

type ClimbedMountain struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Height    int64     `json:"height"`
	Timestamp time.Time `json:"timestamp"`
	Weather   string    `json:"weather"`
	Number    int       `json:"number"`
	ImageURL  string    `json:"image_url"`
}
