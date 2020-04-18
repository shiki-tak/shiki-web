package domain

type ClimbedMountain struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Height      int    `json:"height"`
	ClimbedDate string `json:"climbed_date"`
	Weather     string `json:"weather"`
	Number      int    `json:"number"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}
