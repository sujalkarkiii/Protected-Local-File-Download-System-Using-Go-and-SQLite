package controllers

import (
	"centralserver/data"
	"encoding/json"
	"net/http"
	"os"
)

func HandleFetchingFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodGet{
	directoryPath := "./readingmatrials"

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}
	// i am using slices to store the names of the each file in filesnames
	fileNames := []data.File{}

	for _, file := range files {
  			fileNames = append(fileNames,data.File{Filename: file.Name()})
	}

	// we are saying our data is json type and sending the data using json encodeer 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fileNames)
	
}else{
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)}
	
}
