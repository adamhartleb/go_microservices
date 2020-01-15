package services

import (
	"go_microservices/domain"
	"go_microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
