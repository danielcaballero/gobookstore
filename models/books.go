package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Book represents a book in the bookstore
type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title  string             `bson:"title" json:"title"`
	Author string             `bson:"author" json:"author"`
	Year   int                `bson:"year" json:"year"`
}
