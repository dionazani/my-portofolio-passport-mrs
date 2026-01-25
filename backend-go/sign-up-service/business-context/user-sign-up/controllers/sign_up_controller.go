package businessContext_controller_signUp

import (
	"encoding/json"
	"net/http" // New import for conversion
	signUpReqModel "passport-mrs-go-sign-up-service/business-context/user-sign-up/models"
	signUpService "passport-mrs-go-sign-up-service/business-context/user-sign-up/service"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	var signUpReq signUpReqModel.SignUpReqModel
	if err := json.NewDecoder(r.Body).Decode(&signUpReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 1. Get the response from the Service
	response, err := signUpService.SignUp(r.Context(), signUpReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Apply the Status Code managed by the Service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.GetIntStatusCode())
	json.NewEncoder(w).Encode(response)

}
