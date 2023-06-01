package intools

import (
	"github.com/hasemeneh/poc-material/poc-login/internal/definition"
	"github.com/hasemeneh/poc-material/poc-login/internal/repository"
)

type module struct {
	user       repository.User
	permission repository.Permission
	role       repository.Role
}
type Args struct {
	User       repository.User
	Permission repository.Permission
	Role       repository.Role
}

func New(o *Args) definition.Intools {
	return &module{
		user:       o.User,
		permission: o.Permission,
		role:       o.Role,
	}
}
