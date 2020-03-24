package handlers

import (
	"net/http"
	"tangy-inn/tangy-inn_backend/internal"
	"tangy-inn/tangy-inn_backend/models"
)

func (p *Provider) Authenticate(rw http.ResponseWriter, r *http.Request) {

	dbSession := p.db.Copy()
	defer dbSession.Close()

	res := make(map[string]interface{})
	credential := models.LoginDetail{}

	if !parseJson(rw, r.Body, &credential) {
		return
	}

	user, err := internal.UserHasAccess(p.db, credential)
	if err != nil {
		res["message"] = "Invalid username or password please check it"
		renderJson(rw, http.StatusUnauthorized, res)
		return
	}

	res["message"] = "Logged in Successfully"
	res["user"] = user

	renderJson(rw, http.StatusOK, res)

}
