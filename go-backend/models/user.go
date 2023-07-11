package models

import "fmt"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"user_name"`
}

var MockedUsersTable = []User{
	{Id: 1, Username: "first_user_123"},
	{Id: 2, Username: "second_user_123"},
	{Id: 3, Username: "third_user_123"},
}

func GetUserById(id int) (User, error){
	for _, user := range MockedUsersTable{
		if user.Id == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("there is no user with such id")
}