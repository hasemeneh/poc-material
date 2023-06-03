package authorization

import (
	"github.com/hasemeneh/poc-material/poc-pricing/internal/definition"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/repository"
)

type module struct {
	user       repository.User
	permission repository.Customer
	role       repository.Role
}
type Args struct {
	User       repository.User
	Permission repository.Customer
	Role       repository.Role
}

func New(o *Args) definition.Authorization {
	return &module{
		user:       o.User,
		permission: o.Permission,
		role:       o.Role,
	}
}
