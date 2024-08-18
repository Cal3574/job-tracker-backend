package models

//Composite data type that groups together related fields under a single name. It's similar to a object in JavaScript
//The 'json' tags tell GO how to encode/decode the struct when it serializes/deserializes it to/from JSON

type User struct {
	ID          int `json:"id"`
	CreatedAt  string `json:"created_at"`
	Email 	 string `json:"email"`
	Name 	 string `json:"name"` 
}

