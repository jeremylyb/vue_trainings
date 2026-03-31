package main

import (
	"errors"
	"net/http"
	"time"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type envelope map[string]interface{}

// func (app *application) Login(w http.ResponseWriter, r *http.Request) {
// 	type credentials struct {
// 		UserName string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	var requestCreds credentials
// 	var responsePayload jsonResponse

/*
	Old Custom approach. Now use new readJson and writeJson approach
*/
// err := json.NewDecoder(r.Body).Decode(&requestCreds)
// if err != nil {
// 	app.errorLog.Println("invalid json")
// 	responsePayload.Error = true
// 	responsePayload.Message = "Invalid json"

// 	out, err := json.MarshalIndent(responsePayload, "", "\t")
// 	if err != nil {
// 		app.errorLog.Println(err)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// 	return
// }
// /*
// 	Authenticate Mechanism
// 	Currently just print log
// */
// app.infoLog.Println(requestCreds.UserName, requestCreds.Password)

// // send backa response
// responsePayload.Error = false
// responsePayload.Message = "Signed in"
// responseJSON, err := json.MarshalIndent(responsePayload, "", "\t")
// if err != nil {
// 	app.errorLog.Println(err)
// }
// w.WriteHeader(http.StatusOK)
// w.Write(responseJSON)

/*
	New approach
*/
// 	err := app.readJSON(w, r, &requestCreds)
// 	if err != nil {
// 		app.errorLog.Println(err)
// 		responsePayload.Error = true
// 		responsePayload.Message = "Invalid json supplied, or json missing entirely"
// 		_ = app.writeJSON(w, http.StatusBadRequest, responsePayload)
// 	}

// 	// Authentication mechanism
// 	app.infoLog.Println(requestCreds.UserName, requestCreds.Password)

// 	// send back response
// 	responsePayload.Error = false
// 	responsePayload.Message = "Signed in"
// 	err = app.writeJSON(w, http.StatusOK, responsePayload)
// 	if err != nil {
// 		app.errorLog.Println(err)
// 	}
// }

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// TODO authenticate
	app.infoLog.Println(creds.UserName, creds.Password)

	// look up the user by email
	user, err := app.models.User.GetByEmail(creds.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	// validate the user's password
	validPassword, err := user.PasswordMatches(creds.Password)
	if err != nil || !validPassword {
		app.errorJSON(w, errors.New("invalid username/password"))
		return
	}

	// we have a valid user, so generate a token
	token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// save it to the database
	err = app.models.Token.Insert(*token, *user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// send back a response
	payload = jsonResponse{
		Error:   false,
		Message: "logged in",
		Data:    envelope{"token": token, "user": user},
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json"))
		return
	}

	err = app.models.Token.DeleteByToken(requestPayload.Token)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json"))
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "logged out",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
