package businessContext_controller

import (
	"encoding/json"
	"net/http" // New import for conversion
	signUpModel "passport-mrs-go/business-context/user-sign-up/models"
	service "passport-mrs-go/business-context/user-sign-up/service"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var signUpReq signUpModel.SignUpReq
	if err := json.NewDecoder(r.Body).Decode(&signUpReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 1. Get the response from the Service
	response, err := service.SignUp(r.Context(), signUpReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Apply the Status Code managed by the Service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.GetIntStatusCode())
	json.NewEncoder(w).Encode(response)
}
