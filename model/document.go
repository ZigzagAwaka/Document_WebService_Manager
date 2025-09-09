package model

type Document struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	//URL         string `json:"url"`
}

// Locally stored documents, can be replaced by a database or any other storage solution
var Basic_documents = []Document{
	{ID: 1, Title: "Lab Chozu doc", Description: "Technical documentation for Lab Chozu project concerning the pharmaceutical industry."},
	{ID: 2, Title: "Anatomy sciences report", Description: "Detailed report on recent advancements in anatomy sciences."},
	{ID: 3, Title: "Biology research paper", Description: "In-depth research paper exploring various aspects of biology."},
	{ID: 4, Title: "Chemistry lab manual", Description: "Comprehensive manual for conducting experiments in a chemistry lab."},
	{ID: 5, Title: "Physics textbook", Description: "Textbook covering fundamental concepts in physics."},
}
