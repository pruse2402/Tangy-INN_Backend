package db

// import (
// 	"encoding/json"
// 	"errors"
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/globalsign/mgo"
// )

// var ErrDup = errors.New("Duplicate field")

// type DbErrors struct {
// 	Message string `json:"message"`
// 	Error   string `json:"error"`
// 	Code    int    `json:"code"`
// }

// func renderError(w http.ResponseWriter, status int, err DbErrors) {
// 	resByte, e := json.Marshal(err)
// 	if e != nil {
// 		log.Printf("ERROR: renderError - %q\n", e)
// 	}
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(status)
// 	w.Write(resByte)
// }

// func HandleDbErrors(w http.ResponseWriter, err error, message string) int {
// 	log.Printf("DB ERROR: %s", err.Error())

// 	errIns := DbErrors{
// 		Message: message,
// 		Error:   err.Error(),
// 	}

// 	if mgo.IsDup(err) {
// 		err = ErrDup
// 	}

// 	switch err {

// 	case mgo.ErrNotFound:
// 		errIns.Code = 404

// 	case io.EOF:
// 		errIns.Code = 503
// 		errIns.Message = "EOF: Please check your DB connection"

// 	case mgo.ErrCursor:
// 		errIns.Code = 413
// 		errIns.Message = "DB Cursor exhausted"

// 	case ErrDup:
// 		errIns.Code = 409
// 		errIns.Message = "Duplicate data"

// 	default:
// 		errIns.Code = 400
// 	}

// 	if errIns.Message == "" {
// 		errIns.Message = errIns.Error
// 	}

// 	renderError(w, errIns.Code, errIns)
// 	return errIns.Code
// }
