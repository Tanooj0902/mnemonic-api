package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID       string
	Password string
}

type InputValidationError struct {
	Message string `json:"message"`
	Field   string `json: "field"`
}

type ServerError struct {
	Message string `json:"message"`
}

type Validations struct {
	Errors []InputValidationError `json:"errors"`
}

func validateData(user User) Validations {

	errors := []InputValidationError{}

	if user.Password == "" {
		errors = append(errors, InputValidationError{"Password is required.", "password"})
	} else if len(user.Password) < 8 {
		errors = append(errors, InputValidationError{"Password must be at least 8 characters.", "password"})
	} else if len(user.Password) > 64 {
		errors = append(errors, InputValidationError{"Password must be at most 64 characters.", "password"})
	}

	if user.ID == "" {
		errors = append(errors, InputValidationError{"Id is required.", "id"})
	}

	return Validations{errors}
}

func fetchJsonParams(r *http.Request) (User, error) {
	user := User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func privateKeyAction(w http.ResponseWriter, user User) {

	validation := validateData(user)

	if len(validation.Errors) > 0 {
		errorJson, _ := json.Marshal(validation)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(errorJson)
		Info.Println("/private_key responded with", http.StatusUnprocessableEntity)
		return
	}

	privateKey, error := getMnemonicPrivateKey(user.Password)

	if error != nil {
		errorJson, _ := json.Marshal(ServerError{"Server Error"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorJson)
		Info.Println("/private_key responded with", http.StatusInternalServerError)
		return
	}

	privateKeyJson, _ := json.Marshal(privateKey)
	w.WriteHeader(http.StatusOK)
	w.Write(privateKeyJson)
	Info.Println("/private_key responded with", http.StatusOK)
}
