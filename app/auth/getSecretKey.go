package auth

import (
	"sync"
	"tasks_company/app/utils"
)

var (
	secretKey string
	once      sync.Once
)

func LoadSecretKey() {
	secretKey = utils.InitENV("SECRET_KEY")
}

func GetSecretKey() string {
	once.Do(LoadSecretKey)
	return secretKey
}
