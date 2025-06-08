package db

import (
	"sync"
	"tasks_company/app/utils"
)

var (
	dbPassword string
	once       sync.Once
)

func LoadDBPassword() {
	dbPassword = utils.InitENV("SECRET_KEY")
}

func GetDBPassword() string {
	once.Do(LoadDBPassword)
	return dbPassword
}
