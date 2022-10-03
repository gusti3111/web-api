package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/gusti3111/web-api/accounts"
)

type accountshandler struct {
	personService accounts.Services
}

func NewUserHandler(personService accounts.Services) *accountshandler {
	return &accountshandler{personService}
}

// GET
func (h *accountshandler) GetHandler(c *gin.Context) {
	user, err := h.personService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var personsResponse []accounts.PersonResponse
	for _, b := range user {
		personResponse := accounts.PersonResponse{
			Name:     b.Name,
			Password: b.Password,
		}
		personsResponse = append(personsResponse, personResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": personsResponse,
	})
}

func (h *accountshandler) PostHandler(c *gin.Context) {
	var User accounts.User

	err := c.ShouldBindJSON(&User)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			Message := fmt.Sprintf("Error pada %s,dengan kondisi %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, Message)
			return

		}

	}
	user, err := h.personService.Create(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
func (h *accountshandler) GetHandlerid(c *gin.Context) {
	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)
	user, err := h.personService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
func (h *accountshandler) UpdateHandler(c *gin.Context) {
	idstring := c.Param("id")
	id, _ := strconv.Atoi(idstring)
	user, err := h.personService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
