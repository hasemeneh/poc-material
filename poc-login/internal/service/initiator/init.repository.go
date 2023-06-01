package initiator

import (
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/permission"
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/role"
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/token"
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/user"
	"github.com/hasemeneh/poc-material/poc-login/internal/service"
)

func (i *initiator) InitDomain(basic *service.Basic) *service.Domain {
	user := user.New(&user.Args{
		Redis: basic.Cache,
		DB:    basic.DB,
	})
	permission := permission.New(&permission.Args{
		Redis: basic.Cache,
		DB:    basic.DB,
	})
	role := role.New(&role.Args{
		Redis: basic.Cache,
		DB:    basic.DB,
	})
	token := token.New(&token.Args{
		Redis: basic.Cache,
		DB:    basic.DB,
	})
	return &service.Domain{
		User:       user,
		Permission: permission,
		Role:       role,
		Token:      token,
	}
}
