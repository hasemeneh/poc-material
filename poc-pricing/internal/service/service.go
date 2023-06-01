package service

import (
	"github.com/hasemeneh/poc-material/poc-pricing/config"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/definition"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/repository"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/redis"
)

type Service struct {
	Cfg     *config.MainConfig
	Usecase *Usecase
}
type Usecase struct {
	Authentication definition.Authentication
	Authorization  definition.Authorization
	Intools        definition.Intools
}
type Interface struct {
}
type Domain struct {
	User       repository.User
	Permission repository.Permission
	Role       repository.Role
	Token      repository.Token
}
type Basic struct {
	DB    *database.Store
	Cache redis.RedisInterface
}
