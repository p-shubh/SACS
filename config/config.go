package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type app struct {
	DB     *gorm.DB
	Router *gin.Engine
	Env    Enviroment
}

type Enviroment struct {
	Name     string `mapstrcucture:"PSQL_DB_NAME"`
	Host     string `mapstrcucture:"PSQL_HOST"`
	Port     string `mapstrcucture:"PSQL_PORT"`
	Username string `mapstrcucture:"PSQL_USERNAME"`
	Password string `mapstrcucture:"PSQL_PASSWORD"`
}

//    App := &app{
// 	Env
//    } 