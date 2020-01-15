package services

import (
	"tangy-inn/tangy-inn_backend/models"

	"gopkg.in/mgo.v2/bson"
)

type MenuItem interface {
	Validate(models.MenuItem) (bool, map[string]interface{})
	FindCount(models.MenuItem) int
	InsertMenuItem(models.MenuItem) (*models.MenuItem, error)
	FindByStoreID(string) (*[]models.MenuItem, error)
	FindByID(bson.ObjectId) (*models.MenuItem, error)
	Update(bson.ObjectId, models.MenuItem) error
}
