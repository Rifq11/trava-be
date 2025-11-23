package controller

import (
	"net/http"

	config "github.com/Rifq11/Trava-be/config"
	models "github.com/Rifq11/Trava-be/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTransportTypes(c *gin.Context) {
	var types []models.TransportType

	if err := config.DB.Find(&types).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if types == nil {
		types = []models.TransportType{}
	}

	c.JSON(http.StatusOK, gin.H{"data": types})
}

func CreateTransportType(c *gin.Context) {
	var req models.CreateTransportTypeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tType := models.TransportType{Name: req.Name}

	if err := config.DB.Create(&tType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transport type created successfully",
		"data":    tType,
	})
}

func UpdateTransportType(c *gin.Context) {
	id := c.Param("id")

	var tType models.TransportType
	if err := config.DB.First(&tType, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transport type not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req models.UpdateTransportTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tType.Name = req.Name

	if err := config.DB.Save(&tType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transport type updated successfully",
		"data":    tType,
	})
}

func DeleteTransportType(c *gin.Context) {
	id := c.Param("id")

	var tType models.TransportType
	if err := config.DB.First(&tType, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transport type not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&tType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transport type deleted successfully"})
}
