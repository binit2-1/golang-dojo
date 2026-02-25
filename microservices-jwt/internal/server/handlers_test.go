package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHander(t *testing.T){
	srv := NewServer("secret")

	req := &LoginRequest{
		Email: "test@gmail.com",
		Password: "secret",
	}

	var res LoginResponse

	body, _ := json.Marshal(req)

	fakeReq := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	//fake response recorder
	rr := httptest.NewRecorder()

	srv.LoginHandler(rr, fakeReq)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}

	err := json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
        t.Fatalf("Failed to parse response JSON: %v", err)
    }

	if res.Token == ""{
		t.Errorf("Expected a JWT token, but got an empty string")
	}

}

