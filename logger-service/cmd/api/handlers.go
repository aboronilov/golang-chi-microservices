package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err = app.Models.LogEntry.Insert(event)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	res := jsonResponse{
		Error:   false,
		Message: "Log entry created successfully",
		Data:    event,
	}

	err = app.writeJSON(w, http.StatusCreated, res)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}
}
