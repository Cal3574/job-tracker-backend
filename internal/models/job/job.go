package models

//Composite data type that groups together related fields under a single name. It's similar to a object in JavaScript
//The 'json' tags tell GO how to encode/decode the struct when it serializes/deserializes it to/from JSON

type Job struct {
	ID          int `json:"id"`
	JobTitle       string `json:"job_title"`
	Location string `json:"location"`
	Company     string `json:"company"`
	Salary      int `json:"salary"`
	URL 	   string `json:"url"`
}

