// Make a mountain trip plan and do model
// Save it to ElasticSearch (ES) so that it can be searched by keyword.
// The ES Application (for learning Rust) operates via a web server that uses Rust.

package domain

type PlandMountain struct {
	Title       string  `json:"title"`
	Period      int     `json:"period"`
	Description string  `json:"description"`
	Routes      []Route `json:"routes"`
	Spend       Spend   `json:"spend"`
}

type Route struct {
	Point      []string `json:"point"`
	CourseTime int      `json:"course_time"`
}

type Spend struct {
	Use   string `json:"use"`
	Price int    `json:"price"`
}
