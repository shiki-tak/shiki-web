package domain

import "time"

type ClimbedMountain struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Height      int64     `json:"height"`
	ClimbedDate time.Time `json:"climbed_date"`
	Weather     string    `json:"weather"`
	Number      int       `json:"number"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
}
