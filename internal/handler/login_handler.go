package handler

import (
	"encoding/json"
	"net/http"
	"pd_pritani/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID uint   `json:"user_id"`
		Email  string `json:"email"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	token, err := auth.GenerateJWT(req.UserID, req.Email)
	if err != nil {
		http.Error(w, "Gagal membuat token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type ", "application/json")
	
	json.NewEncoder(w).Encode(map[string]string{"token": token}) 
}
