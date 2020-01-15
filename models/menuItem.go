package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MenuItem struct {
	Id             bson.ObjectId `bson:"_id" json:"id"`
	OriginalID     bson.ObjectId `json:"originalID" bson:"originalID"`
	Name           string        `bson:"name" json:"name"`
	Description    string        `bson:"description,omitempty" json:"description"`
	ActiveStatus   bool          `bson:"activeStatus" json:"activeStatus"`
	CostPrice      int           `bson:"costPrice,omitempty" json:"costPrice,omitempty"`
	CostPriceStr   string        `bson:"costPriceStr,omitempty" json:"costPriceStr,omitempty"`
	DateCreated    time.Time     `bson:"dateCreated" json:"dateCreated"`
	LastUpdated    time.Time     `bson:"lastUpdated" json:"lastUpdated"`
	SequenceNumber int           `json:"sequenceNumber" bson:"sequenceNumber"`
	CreatedFor     bson.ObjectId `bson:"createdFor" json:"createdFor"`
}
