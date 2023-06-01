package authentication

import (
	"github.com/hasemeneh/poc-material/poc-pricing/config"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/definition"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/repository"
)

type module struct {
	Cfg        *config.MainConfig
	UserRepo   repository.User
	permission repository.Permission
	role       repository.Role
	token      repository.Token
}
type Args struct {
	Cfg        *config.MainConfig
	UserRepo   repository.User
	Permission repository.Permission
	Role       repository.Role
	Token      repository.Token
}

func New(o *Args) definition.Authentication {
	return &module{
		UserRepo:   o.UserRepo,
		Cfg:        o.Cfg,
		permission: o.Permission,
		role:       o.Role,
		token:      o.Token,
	}
}
