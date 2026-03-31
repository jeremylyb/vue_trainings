package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	/*
		Prevents attackers from sending huge payloads (DoS attack)
		Caps request body at 1 MB
	*/
	maxBytes := 1048576 //one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	/*
		Reads JSON directly from the request body
		More efficient than reading everything into memory first
	*/
	dec := json.NewDecoder(r.Body)

	/*
		- Use Decoder to decode into my data struct
		- Using interface{} benefits: can decode into any type of structs
			- if struct is type User struct { Name string `json;"name"` } and request body is {"name":"Jeremy"} then data = &User{}
			- Result => User{Name: "Jeremy"}
	*/
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	/*
		This is to ensure body only has single json value
		Multiple JSON objects in one request
		Hidden/malicious extra data
	*/
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}
	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()
	app.writeJSON(w, statusCode, payload)
}
