package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json: "id" form: "id"`
	Name     string `json: "name" form: "name"`
	Email    string `json: "email" form: "email"`
	Password string `json: "password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"user":     users,
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func GetUserController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	for i, _ := range users {
		if users[i].Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get single user",
				"user":     users[i],
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not found",
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, _ := range users {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages":            "deleted user",
				"delete user with id": id,
				"user":                users,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not found",
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	newUser := User{}
	c.Bind(&newUser)

	for i, _ := range users {
		if users[i].Id == id {
			users[i] = newUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "updated single user",
				"user":     users[i],
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not found",
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.Logger.Fatal(e.Start(":8000"))
}
