package config

import (
	"github.com/harry/spotify/structs"
	"github.com/jinzhu/gorm"
)

//db init create conn to db

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:kurkur14@/spotifygo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Band{})
	return db
}
