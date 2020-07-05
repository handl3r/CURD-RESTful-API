package controllertests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignIn(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatalf("error when refresh user table: %v", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("can not seed user : %v", err)
	}
	samples := []struct {
		email        string
		password     string
		errorMessage string
	}{
		{
			email:        user.Email,
			password:     "thaibuixuan",
			errorMessage: "",
		},
		{
			email:        user.Email,
			password:     "wrong password",
			errorMessage: "crypto/bcrypt: hashedPassword is not the hash of the given password",
		},
		{
			email:        "wrongemail@gmail.com",
			password:     "thaibuixuan",
			errorMessage: "record not found",
		},
	}
	fmt.Println(user.Password)
	for _, v := range samples {
		token, err := server.SignIn(v.email, v.password)
		if err != nil {
			assert.Equal(t, err, errors.New(v.errorMessage))
		} else {
			assert.NotEqual(t, token, "")
		}
	}
}

func TestLogin(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatalf("err when refresh user table: %v", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("err when seed user: %v", err)
	}
	samples := []struct {
		inputJSON  string
		statusCode int
		email      string
		password   string
		errMessage string
	}{
		{
			inputJSON:  `{"email": "email@gmail.com", "password": "thaibuixuan"}`,
			statusCode: http.StatusOK,
			errMessage: "",
		},
		{
			inputJSON:  fmt.Sprintf(`{"email": "%s", "password": "wrong password"}`, user.Email),
			statusCode: http.StatusUnprocessableEntity,
			errMessage: "Incorrect Password",
		},
		{
			inputJSON:  fmt.Sprintf(`{"email": "%s", "password": "thaibuixuan"}`, "wrongemail@gmail.com"),
			statusCode: http.StatusUnprocessableEntity,
			errMessage: "Incorrect Details",
		},
		{
			inputJSON:  fmt.Sprintf(`{"email": "%s", "password": "thaibuixuan"}`, "invalid.gmail"),
			statusCode: http.StatusUnprocessableEntity,
			errMessage: "invalid email",
		},
		{
			inputJSON:  fmt.Sprintf(`{"email": "", "password": "thaibuixuan"}`),
			statusCode: http.StatusUnprocessableEntity,
			errMessage: "invalid email",
		},
		{
			inputJSON:  fmt.Sprintf(`{"email": "%s", "password": ""}`, user.Email),
			statusCode: http.StatusUnprocessableEntity,
			errMessage: "require password",
		},
	}
	for _, v := range samples {
		req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("err: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.Login)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == http.StatusOK {
			assert.NotEqual(t, rr.Body.String(), "")
		}
		if v.statusCode == http.StatusUnprocessableEntity && v.errMessage != "" {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			//dec := json.NewDecoder(bytes.NewReader([]byte(rr.Body.String())))
			//err = dec.Decode(&responseMap)
			if err != nil {
				t.Errorf("can not convert to json: %v", err)
			}
			assert.Equal(t, responseMap["error"], v.errMessage)
		}
	}
}
