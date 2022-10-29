package controllers

import (
	"net/http"
	"web-api-golang-myself/model"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var products []model.Product
	model.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})

}
func GetUserById(c *gin.Context) {
	var product model.Product
	id := c.Param("id")
	if err := model.DB.First(&product, id).Error; err != nil {
		panic(err)

	}
	c.JSON(http.StatusOK, gin.H{"product": product})

}
func CreateUser(c *gin.Context) {
	var product model.Product
	// catch JSON body and send to product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	model.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})

}
func UpdateUser(c *gin.Context) {
	var product model.Product
	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if model.DB.Model(&product).Where("id=?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "data not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})

}
func DeleteUser(c *gin.Context) {
	var product model.Product
	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if model.DB.Model(&product).Where("id=?", id).Delete(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "data tidak bisa di hapus"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})

}
