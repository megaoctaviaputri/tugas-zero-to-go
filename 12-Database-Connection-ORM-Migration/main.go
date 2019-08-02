package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	db.AutoMigrate(&User{})
}

// database connection
func InitDB(connectionString string) {
	var err error
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
}

func main() {
	// create a new echo instance
	e := echo.New()
	InitDB("root@/go_db?charset=utf8&parseTime=True&loc=Local")
	InitialMigration()

	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/user/:id", GetUserController)
	e.PUT("/user/:id", UpdateUserController)
	e.DELETE("/user/:id", DeleteUserController)

	// start the server
	e.Start(":8000")
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// Create User
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := db.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}

	db.First(&user, id)

	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

// update user controller
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	newUser := User{}
	user := User{}
	c.Bind(&newUser)

	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Tidak lengkap guys",
		})
	}

	db.First(&user, id)
	db.Save(&newUser)

	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    newUser,
	})

}

// delete user controller
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}

	db.First(&user, id)
	db.Delete(&user)

	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}
