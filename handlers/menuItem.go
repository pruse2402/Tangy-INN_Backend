package handlers

import (
	"log"
	"net/http"
	"tangy-inn/tangy-inn_backend/datastore"
	"tangy-inn/tangy-inn_backend/errs"
	"tangy-inn/tangy-inn_backend/internal"
	"tangy-inn/tangy-inn_backend/models"
	"time"

	"github.com/go-chi/chi"
	"gopkg.in/mgo.v2/bson"
)

func (p *Provider) GetMenuItemList(rw http.ResponseWriter, r *http.Request) {
	dbSession := p.db.Copy()
	defer dbSession.Close()

	//ctx := r.Context()

	storeID_s := chi.URLParam(r, "id")

	if !bson.IsObjectIdHex(storeID_s) {
		log.Printf("ERROR: GetMenuItemList - Invalid Store ID(%s).\n", storeID_s)
		renderJson(rw, http.StatusBadRequest, `Invalid Store ID`)
		return
	}

	menuItems, err := datastore.NewMenuItem(dbSession).FindByStoreID(storeID_s)
	if err != nil {
		log.Printf("ERROR: GetMenuItemList %s", err)
		return
	}

	resp := struct {
		MenuItems *[]models.MenuItem `json:"menuItems"`
	}{
		menuItems,
	}

	// res := make(map[string]interface{})
	// res["user"] = menuItems

	renderJson(rw, http.StatusOK, resp)

}

func (p *Provider) InsertMenuItem(rw http.ResponseWriter, r *http.Request) {

	dbSession := p.db.Copy()
	defer dbSession.Close()

	menuItemIns := models.MenuItem{}

	if !parseJson(rw, r.Body, &menuItemIns) {
		return
	}

	hasErr, validationErr := internal.Validation(p.db, menuItemIns)

	if hasErr {
		log.Printf("ERROR: InsertMenuItem - %q\n", validationErr)
		err := &errs.AppError{
			Message: "Validation Error(s)",
			Errors:  validationErr,
		}
		respondError(rw, http.StatusBadRequest, err)
		return
	}

	menuItemIns.Id = bson.NewObjectId()
	menuItemIns.OriginalID = menuItemIns.Id

	menuItemIns.ActiveStatus = true
	menuItemIns.DateCreated = time.Now().UTC()
	menuItemIns.LastUpdated = time.Now().UTC()
	// menuItem.CreatedBy = bson.ObjectIdHex(session.Values["loggedInUser"].(string))
	// menuItem.ModifiedBy = bson.ObjectIdHex(session.Values["loggedInUser"].(string))
	// menuItem.OwnedBy = bson.ObjectIdHex(session.Values["accountID"].(string))

	c := datastore.NewMenuItem(dbSession).FindCount(menuItemIns)
	menuItemIns.SequenceNumber = c

	var resp string
	menuItem, err := datastore.NewMenuItem(dbSession).InsertMenuItem(menuItemIns)
	if err != nil {
		resp = "Error while saving menu item"
		renderJson(rw, http.StatusUnauthorized, resp)
		return
	}

	res := struct {
		ID      bson.ObjectId `json:"id"`
		Message string        `json:"message"`
	}{}
	res.ID = menuItem.Id
	res.Message = `Menu item saved successfully`

	renderJson(rw, http.StatusOK, res)
	return

}

func (p *Provider) UpdateMenuItem(rw http.ResponseWriter, r *http.Request) {

	dbSession := p.db.Copy()
	defer dbSession.Close()

	menuID := chi.URLParam(r, "id")

	if !isObjectIDValid(menuID) {
		p.log.Printf("ERROR: Handler - Update - %q\n", `Invalid Menu ID Supplied.`)
		err := &errs.UIErr{
			Code:    http.StatusBadRequest,
			Message: "Invalid Menu ID Supplied.",
		}
		renderError(rw, err)
		return
	}

	menuItem, err := datastore.NewMenuItem(dbSession).FindByID(bson.ObjectIdHex(menuID))

	if err != nil {
		p.log.Printf("ERROR: Handler - Update - %q\n", err)
		err := &errs.UIErr{
			Code:    http.StatusNotFound,
			Message: "Could not find category!.",
		}
		renderError(rw, err)
		return
	}

	if !parseJson(rw, r.Body, &menuItem) {
		return
	}

	hasErr, validationErr := internal.Validation(p.db, *menuItem)

	if hasErr {
		log.Printf("ERROR: InsertMenuItem - %q\n", validationErr)
		err := &errs.AppError{
			Message: "Validation Error(s)",
			Errors:  validationErr,
		}
		respondError(rw, http.StatusBadRequest, err)
		return
	}

	menuItem.LastUpdated = time.Now().UTC()

	if err = datastore.NewMenuItem(dbSession).Update(menuItem.Id, *menuItem); err != nil {
		p.log.Printf("ERROR: Handler - Update - %q\n", err)
		err := &errs.UIErr{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong!. Please try again!.",
		}
		renderError(rw, err)
		return
	}

	// Constructing response for client
	res := struct {
		ID      bson.ObjectId `json:"id"`
		Message string        `json:"message"`
	}{
		menuItem.Id,
		"MenuItem updated successfully.",
	}

	renderJson(rw, http.StatusOK, res)

}
