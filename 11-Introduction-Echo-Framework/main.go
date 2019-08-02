package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var user User
	for _, value := range users {
		if value.Id == id {
			user = value
		}
	}
	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})

}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users = remove(users, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}
	c.Bind(&user)
	user.Id = id
	users = update(users, id, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
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

// ----------------------------- library --------------------------

// to remove user from slice
func remove(users []User, rmIdx int) []User {
	var idxRmv int
	for index, user := range users {
		if user.Id == rmIdx {
			idxRmv = index
		}
	}
	return append(users[:idxRmv], users[idxRmv+1:]...)
}

// to update user on slice
func update(users []User, id int, newUser User) []User {
	for index, user := range users {
		if user.Id == id {
			users[index] = newUser
		}
	}
	return users
}

// ----------------------- main function ----------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	e.Start(":8000")

}
