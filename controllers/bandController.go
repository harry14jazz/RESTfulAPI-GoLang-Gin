package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harry/spotify/structs"
)

//get single band with id
func (idb *InDB) GetBand(c *gin.Context) {
	var (
		band   structs.Band
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&band).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": band,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

//get all band
func (idb *InDB) GetBands(c *gin.Context) {
	var (
		bands  []structs.Band
		result gin.H
	)

	idb.DB.Find(&bands)
	if len(bands) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": bands,
			"count":  len(bands),
		}
	}
	c.JSON(http.StatusOK, result)
}

//create new data to db
func (idb *InDB) CreateBand(c *gin.Context) {
	var (
		band   structs.Band
		result gin.H
	)

	bandName := c.PostForm("BandName")
	leader := c.PostForm("Leader")
	birthYear := c.PostForm("BirthYear")
	year, _ := strconv.Atoi(birthYear)
	band.BandName = bandName
	band.Leader = leader
	band.BirthYear = year
	idb.DB.Create(&band)
	result = gin.H{
		"result": band,
	}
	c.JSON(http.StatusOK, result)
}

//update data by id
func (idb *InDB) UpdateBand(c *gin.Context) {
	id := c.Query("id")
	BandName := c.PostForm("BandName")
	Leader := c.PostForm("leader")
	BirthYear := c.PostForm("BirthYear")
	var (
		band    structs.Band
		newBand structs.Band
		result  gin.H
	)

	err := idb.DB.First(&band, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	year, _ := strconv.Atoi(BirthYear)
	newBand.BandName = BandName
	newBand.Leader = Leader
	newBand.BirthYear = year
	err = idb.DB.Model(&band).Updates(newBand).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully update data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteBand(c *gin.Context) {
	var (
		band   structs.Band
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&band, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&band).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "delete successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}
