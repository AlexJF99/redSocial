package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Birthday  time.Time          `bson:"birthday" json:"birthday"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	WebSite   string             `bson:"webSite" json:"webSite,omitempty"`
}
