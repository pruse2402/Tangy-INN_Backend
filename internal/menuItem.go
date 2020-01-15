package internal

import (
	"tangy-inn/tangy-inn_backend/datastore"
	"tangy-inn/tangy-inn_backend/models"

	"gopkg.in/mgo.v2"
)

func Validation(dbSession *mgo.Session, menuItem models.MenuItem) (bool, map[string]interface{}) {

	menuItemIns := datastore.NewMenuItem(dbSession)

	hasErr, validationErr := menuItemIns.Validate(menuItem)

	return hasErr, validationErr
}
