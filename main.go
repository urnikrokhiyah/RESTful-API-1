package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// // func HelloController(c echo.Context) error {
// // 	return c.String(http.StatusOK, "Hello world")
// // }

// // type User struct {
// // 	Id    int
// // 	Name  string
// // 	Email string
// // }

// // func getUser(c echo.Context) error {
// // 	user := User{Name: "urnik", Email: "urnik@gmail.com"}
// // 	return c.JSON(http.StatusOK, user)
// // }

// // func getUserController(c echo.Context) error {
// // 	id, _ := strconv.Atoi(c.Param("id"))
// // 	user := User{Id: id, Name: "urnik", Email: "urnik@gmail.com"}
// // 	return c.JSON(http.StatusOK, map[string]interface{}{"user": user})
// // }

// // func userSearch(c echo.Context) error {
// // 	match := c.QueryParam("match")
// // 	return c.JSON(http.StatusOK, map[string]interface{}{"match": match, "result": []string{"adi", "aan", "asif"}})
// // }

// // func createUser(c echo.Context) error {
// // 	name := c.FormValue("name")
// // 	email := c.FormValue("email")

// // 	var user Users
// // 	user.Name = name
// // 	user.Emailuser = email

// // 	return c.JSON(http.StatusOK, map[string]interface{}{"user": user})
// // }

// type Users struct {
// 	Name  string `json:"name" form:"name"`
// 	Email string `json:"email" form:"email`
// }

// func createUsers(c echo.Context) error {
// 	user := Users{}
// 	c.Bind(&user)

// 	return c.JSON(http.StatusOK, map[string]interface{}{"messages": "success", "user": user})
// }

// func main() {
// 	e := echo.New()
// 	// e.GET("/", HelloController)
// 	// e.GET("/user", getUser)
// 	// e.GET("/user/:id", getUserController)
// 	// e.GET("/user", userSearch)
// 	e.POST("/user", createUsers)
// 	e.Start(":8000")
// }

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
