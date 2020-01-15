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
type menuItem struct {
	// session    *mgo.Session
	collection *mgo.Collection
	// config     *config.Config
}

// NewUser creates a new user datastore
func NewMenuItem(db *mgo.Session) services.MenuItem {
	return &menuItem{
		// session:    db,
		collection: db.DB("tangy-inn").C("menuItem"),
		// config:     cfg,
	}
}

func (menu *menuItem) FindCount(menuItem models.MenuItem) int {

	query := bson.M{"createdFor": menuItem.CreatedFor}
	c, _ := menu.collection.Find(query).Count()

	return c
}

func (menu *menuItem) FindByID(id bson.ObjectId) (*models.MenuItem, error) {

	menuItemIns := &models.MenuItem{}
	err := menu.collection.FindId(id).One(menuItemIns)
	if err != nil && err != mgo.ErrNotFound {
		log.Printf("ERROR: FindByID(%s) - %s\n", id, err)
		return nil, err
	}
	return menuItemIns, nil

}

func (menu *menuItem) FindByStoreID(storeID string) (*[]models.MenuItem, error) {

	menuItems := &[]models.MenuItem{}
	query := bson.M{"createdFor": bson.ObjectIdHex(storeID)}

	err := menu.collection.Find(query).All(menuItems)
	if err != nil {
		log.Printf("ERROR: FindMenuItemsByStoreID(%s) - %q\n", storeID, err)
		return nil, err
	}

	return menuItems, nil
}

func (menu *menuItem) InsertMenuItem(menuItem models.MenuItem) (*models.MenuItem, error) {

	if err := menu.collection.Insert(&menuItem); err != nil {
		log.Printf("ERROR: InsertMenuItem(%s) - %q\n", menuItem.Name, err)
		return nil, err
	}

	return &menuItem, nil

}

func (menu *menuItem) Update(id bson.ObjectId, menuItemIns models.MenuItem) error {
	err := menu.collection.UpdateId(id, &menuItemIns)
	if err != nil {
		log.Printf("ERROR: Update(%s, %s) - %s\n", id, menuItemIns.Id, err)
		return err
	}
	return nil
}

func (menu *menuItem) Validate(menuItem models.MenuItem) (bool, map[string]interface{}) {

	v := &utils.Validation{}

	//Escape special characters to avoid duplicates
	menuItemName := utils.AddEscapeString(menuItem.Name)

	res := v.Required(menuItem.Name).Key("name").Message("Enter name")
	if res.Ok {
		res = v.MaxSize(menuItem.Name, 50).Key("name").Message("Name should not be more than 50 characters")
		if res.Ok {
			query := bson.M{"name": bson.RegEx{`^` + menuItemName + `$`, `i`}}
			if menuItem.Id.Hex() != "" {
				query["_id"] = bson.M{"$ne": menuItem.Id}
			}

			count, _ := menu.collection.Find(query).Count()
			if count > 0 {
				v.Error("Name already exists").Key("name")
			}
		}
	}

	return v.HasErrors(), v.ErrorMap()

}
