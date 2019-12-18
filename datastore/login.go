package datastore

import (
	"log"
	"tangy-inn/tangy-inn_backend/models"
	"tangy-inn/tangy-inn_backend/services"
	"tangy-inn/tangy-inn_backend/utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// user is holds information for database
type user struct {
	// session    *mgo.Session
	collection *mgo.Collection
	// config     *config.Config
}

// NewUser creates a new user datastore
func NewUser(db *mgo.Session) services.Login {
	return &user{
		// session:    db,
		collection: db.DB("tangy-inn").C("user"),
		// config:     cfg,
	}
}

// FindByUser to get the user details
func (use *user) FindByUser(credential models.LoginDetail) (*models.User, error) {

	userIns := &models.User{}
	email := utils.AddEscapeString(credential.EmailID)

	query := bson.M{"emailId": bson.RegEx{`^` + email + `$`, `i`},
		"password": utils.SHAEncoding(credential.Password),
	}

	if err := use.collection.Find(query).One(&userIns); err != nil {
		log.Printf("ERROR: FindByUser(%s) - %q\n", credential.EmailID, err)
		return nil, err
	}

	return userIns, nil
}
