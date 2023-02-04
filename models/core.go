package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Business struct {
	ID        primitive.ObjectID
	City      string
	State     string
	Country   string
	Latitude  float64
	Longitude float64
}
