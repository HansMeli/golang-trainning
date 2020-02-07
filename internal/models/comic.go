package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Response struct to map the Entire Response
type Response struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

// Data Struct to map every data.
type Data struct {
	Count   int       `json:"count"`
	Results []Results `json:"results"`
}

// Results to map our comics details
type Results struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// A Comic struct for db operation
type Comic struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IDMarvel    int                `json:"idMarvel,omitempty" bson:"idMarvel,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
}

func setStock() string {
	return "stock seted"
}
