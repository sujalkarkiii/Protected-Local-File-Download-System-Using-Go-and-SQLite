package controllers

import (
	"centralserver/data"  
	"encoding/json"
	"net/http"

)



func Handleautth(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost && r.Header.Get("Content-Type") == "application/json" {
		var User data.Handlelingauthetication

		err := json.NewDecoder(r.Body).Decode(&User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}	
		var existingUser data.Handlelingauthetication
	err = DB.Where("phone_number = ?", User.RollNo).First(&existingUser).Error

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User is registered and okay to download file!"))
	} else {
		http.Error(w, "Cannot accept other than JSON", http.StatusBadRequest)
	}
}
