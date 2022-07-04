package entities

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	Publication   string `json:"publication"`
	PublishedDate string `json:"published_date"`
}
