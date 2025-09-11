package model

type Document struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	//URL         string `json:"url"`
}
