package models

import (
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Getter interface for to get all comics
type Getter interface {
	GetAll() []Comic
}

// Adder interface for to get all comics
type Adder interface {
	Add(comic Comic)
}

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

// Repo instance a object type Comic
type Repo struct {
	Comics []Comic
}

// SetStock set a random number for the comic stock
func (comic *Comic) SetStock() {
	min := 1
	max := 10

	comic.Stock = rand.Intn(max-min) + min
}

// New create a new repository
func New() *Repo {
	return &Repo{
		Comics: []Comic{},
	}
}

// Add a new comic in the comics repo
func (r *Repo) Add(comic Comic) {
	r.Comics = append(r.Comics, comic)
}

// GetAll return all comics in the comics repo
func (r *Repo) GetAll() []Comic {
	return r.Comics
}
