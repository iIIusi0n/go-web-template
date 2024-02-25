package auth_test

import (
	"api-server/controllers/auth"

	"testing"
)

func TestBuildNewToken(t *testing.T) {
	token, err := auth.BuildNewToken("thisistestuid", "thisistestname")
	if err != nil {
		t.Error("Error while building new token: ", err)
		return
	}

	t.Log("New token: ", token)
}

func TestParseToken(t *testing.T) {
	tokenString, _ := auth.BuildNewToken("thisistestuid", "thisistestname")

	claims, err := auth.ParseToken(tokenString)
	if err != nil {
		t.Error("Error while parsing token: ", err)
		return
	}

	t.Log("Parsed token: ", claims)
}
