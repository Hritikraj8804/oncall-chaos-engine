package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
)

type IncidentData struct {
	Incidents []string `json:"incidents"`
	Fixes []string `json:"fixes"`
	Tools []string `json:"tools"`
}

var data IncidentData

func loadIncidents() error {
	file, err := os.ReadFile("data/incidents.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	return nil
}

func randomItem(list []string) string {
	return list[rand.Intn(len(list))]
}

func incidentHandler(w http.ResponseWriter, r *http.Request) {

	visits := incrementVisits()
	
	response := map[string]interface{}{
		"incident": randomItem(data.Incidents),
		"tool": randomItem(data.Tools),
		"fix": randomItem(data.Fixes),
		"visits":  visits,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

