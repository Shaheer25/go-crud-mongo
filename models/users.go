package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId string `json:"userId" bson:"userId"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"` 
	Password string `json:"password" bson:"password"`	
}