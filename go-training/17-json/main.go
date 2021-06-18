package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"` //use tags on struct field to customize the JSON keys
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio,omitempty"` // The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value(zero-value).
}

func main() {
	userData := []byte(`{"id":1, "username":"Bob", "email":"bob@gmail.com"}`)
	var user User
	if err := json.Unmarshal(userData, &user); err != nil {
		panic(err)
	}
	fmt.Println(user) 

	var userMap map[string]interface{}
	if err := json.Unmarshal(userData, &userMap); err != nil {
		panic(err)
	}
	userID := userMap["id"].(float64) 
	fmt.Println(userID)
	fmt.Println(user)

	user = User{ID: 1, Username: "John", Email: "johny@foo.bar"}
	userData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(userData))

	usersData := []byte(`[{"id":1, "username":"Bob", "email":"bob@gmail.com"}, {"id":1, "username":"Bob", "email":"bob@gmail.com"}]`)
	var users []User
	if err = json.Unmarshal(usersData, &users); err != nil {
		panic(err)
	}
	fmt.Println(users)

	type Err struct {
		Code    int    `json:"error"`
		Message string `json:"message"`
	}

	type AppError struct {
		Error Err `json:"error"`
	}

	errData := []byte(`{"error":{"code":200, "message":"oops, something went wrong"}}`)
	var appErr AppError
	if err := json.Unmarshal(errData, &appErr); err != nil {
		panic(err)
	}
	fmt.Println(appErr)

	appErr = AppError{
		Error: Err{
			Code:    200,
			Message: "Some error message",
		},
	}
	errData, err = json.Marshal(appErr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(errData))
}
