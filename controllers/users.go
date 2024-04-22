package controllers

import (
	helper "awesomeProject/helpers"
	users "awesomeProject/models"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"strings"
	"regexp"
)


func CreateUser(c *gin.Context){

	var user users.Users

	if err := c.BindJSON(&user); 
	err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"data": nil,
			"status" : false,
			"message":"Error Binding to JSON",
		})
		return
	}
	nameMatch , _ := regexp.MatchString("^[a-zA-Z\\s]+$", user.Name)
	if !nameMatch {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"status" : false,
			"message":"Name can only contain alphabets and spaces.",
		})
		return
	}
	var email string =  strings.ToLower(user.Email) 
	response , status := helper.CheckValidEmailLength(email); {
		if !status{
			c.JSON(http.StatusBadRequest, gin.H{
				"data": nil,
				"status" : false,
				"message": string(response),
			})
			return
		}
	}
	ValidResp , statutes := helper.ValidateEmail(email);{
		if !statutes{
			c.JSON(http.StatusBadRequest, gin.H{
				"data": nil,
				"status" : false,
				"message": string(ValidResp),
			})
			return
		}
	}
	
	client := helper.ConnectToDB()
	Collection := client.Database("awesomeProject").Collection("users")
	resp , err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			"status" : false,
			"message":"DB Error",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
	"data": gin.H{
		"id":resp.InsertedID,
	},
	"status" : true,
	"message":"User Created Successfully",
})
}