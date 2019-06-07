package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/harry/spotify/config"
	"github.com/harry/spotify/controllers"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/band/:id", inDB.GetBand)
	router.GET("/bands", inDB.GetBands)
	router.POST("/band", inDB.CreateBand)
	router.PUT("/band", inDB.UpdateBand)
	router.DELETE("/band/:id", inDB.DeleteBand)
	router.Run(":3000")
}
