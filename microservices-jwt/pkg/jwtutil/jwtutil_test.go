package jwtutil

import (
	"testing"
)

func TestJWTGenerationAndValidation(t *testing.T){

	dummyEmail := "test@gmail.com"
	secret := "supersecret"

	token , err := GenerateToken(dummyEmail, secret)
	if err != nil{
		t.Fatalf("Generation failed: %v", err)
	}


	claims, err := ValidateToken(token, secret)
	if err !=nil{
		t.Fatalf("Validation failed: %v", err)
	}

	if claims.Email != dummyEmail {
		t.Errorf("Expected %s, got %s", dummyEmail, claims.Email)
	}


}