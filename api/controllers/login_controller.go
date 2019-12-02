package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/NguyenHuy1812/forum/api/auth"
	"github.com/NguyenHuy1812/forum/api/models"
	"github.com/NguyenHuy1812/forum/api/security"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(c *gin.Context) {

	//clear previous error if any

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":      http.StatusUnprocessableEntity,
			"first error": "Unable to get request",
		})
		return
	}
	user := models.User{}
	json.Unmarshal(body, &user)
	
	user.Prepare()
	
	userData, _ := server.SignIn(user.Email, user.Password)
	
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userData,
	})
}

func (server *Server) SignIn(email, password string) (map[string]interface{}, error) {

	var err error

	userData := make(map[string]interface{})

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		fmt.Println("this is the error getting the user: ", err)
		return nil, err
	}
	err = security.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("this is the error hashing the password: ", err)
		return nil, err
	}
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		fmt.Println("this is the error creating the token: ", err)
		return nil, err
	}
	userData["token"] = token
	userData["id"] = user.ID
	userData["email"] = user.Email
	userData["username"] = user.Username

	return userData, nil
}
