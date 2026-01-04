package controllers

import (
	"centralserver/data"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

// This will be hidden in real case when deployed

func Handleautth(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	if r.Method == http.MethodPost {
		var User data.Handlelingauthetication
		Department := r.FormValue("Department")
		RollNo := r.FormValue("RollNo")

		res := db.Where("roll_no = ? AND department = ?", RollNo, Department).First(&User)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				fmt.Println("User not found")
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			} else {
				fmt.Println("DB error:", res.Error)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"roll_no":    User.RollNo,
			"department": User.Department,
			"exp":        time.Now().Add(time.Hour).Unix(), // expires in 1 hour
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Send single JSON response
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"status":  "ok",
			"message": "Login successful!",
			"token":   tokenString,
		}
		json.NewEncoder(w).Encode(response)

	} else {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}
