package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/binit2-1/golang-dojo/auth-project/internal/auth"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type LoginUserPayload struct {
	FullName string `json:"FullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Server struct {
	rdb *redis.Client
}

func NewServer(rdb *redis.Client) *Server{
	return &Server{rdb: rdb}
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginUserPayload

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	dummyHash, err := auth.HashPassword("secret")
	check := auth.CheckHashPassword(req.Password, dummyHash)

	if !check {
		http.Error(w, "Invalid credentials" , http.StatusUnauthorized)
		return
	}


	sessionToken := uuid.NewString()

	statusCmd := s.rdb.Set(r.Context(), sessionToken, req.Email, 24*time.Hour)
	if err := statusCmd.Err(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "session_id",
		Value:    sessionToken,
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		Path:     "/",
	})
	
	

	w.Write([]byte(`{"message":"Login Successful"}`))
}
