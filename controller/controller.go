package controller

import (
	"golang-chap49/config"
	"golang-chap49/database"
	"golang-chap49/service"

	"go.uber.org/zap"
)

type Controller struct {
	User UserController
}

func NewController(service service.Service, logger *zap.Logger, cacher database.Cacher, config config.Configuration) *Controller {
	return &Controller{
		User: *NewUserController(service, logger, cacher, config),
	}
}
