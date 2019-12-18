package internal

import (
	"log"
	"tangy-inn/tangy-inn_backend/datastore"
	"tangy-inn/tangy-inn_backend/models"

	"gopkg.in/mgo.v2"
)

// UserHasAccess checks whether user has the access
func UserHasAccess(dbSession *mgo.Session, credential models.LoginDetail) (*models.User, error) {

	user, err := datastore.NewUser(dbSession).FindByUser(credential)
	if err != nil {
		log.Printf("Error in login: %s", err.Error())
		return user, err
	}

	return user, nil

}
