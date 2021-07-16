package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	log "github.com/sirupsen/logrus"
)

type Store struct {
	DB *gorm.DB
}

var GameBootCamp Store

func InitDB(c *Configs){
	var err error

	log.Infoln(c.Mssql.ConnectionMasterDB)

	db, err := gorm.Open("mssql", c.Mssql.ConnectionMasterDB)

	if err != nil {
		log.Errorln("Failed to connect database : ", c.Mssql.ConnectionMasterDB)
		log.Errorln("Error : ", err)
		panic("Failed to connect database MS SQL")
	}

	GameBootCamp.DB = db
}
