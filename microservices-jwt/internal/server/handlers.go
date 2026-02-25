package server

import (
	"encoding/json"
	"net/http"

	"github.com/binit2-1/golang-dojo/microservices-jwt/pkg/jwtutil"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Server struct{
	JWTSecret string 
}

func NewServer(secret string) *Server{
	return &Server{
		JWTSecret: secret,
	}
}

func(s *Server) LoginHandler(w http.ResponseWriter, r *http.Request){

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err!=nil{
		http.Error(w, "Invalid JSON Payload", http.StatusBadRequest)
		return
	}


	if req.Email != "test@gmail.com" || req.Password != "secret"{
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := jwtutil.GenerateToken(req.Email, s.JWTSecret)
	if err != nil {
    	http.Error(w, "Failed to generate token", http.StatusInternalServerError)
    	return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
