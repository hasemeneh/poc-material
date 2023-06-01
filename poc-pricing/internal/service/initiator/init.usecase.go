package initiator

import (
	"github.com/hasemeneh/poc-material/poc-pricing/config"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/service"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/usecase/authentication"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/usecase/authorization"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/usecase/intools"
)

func (i *initiator) InitUsecase(domain *service.Domain, extInterface *service.Interface, Cfg *config.MainConfig) *service.Usecase {
	auth := authentication.New(&authentication.Args{
		Cfg:        Cfg,
		UserRepo:   domain.User,
		Permission: domain.Permission,
		Role:       domain.Role,
		Token:      domain.Token,
	})
	authorization := authorization.New(&authorization.Args{
		User:       domain.User,
		Permission: domain.Permission,
		Role:       domain.Role,
	})
	intools := intools.New(&intools.Args{
		User:       domain.User,
		Permission: domain.Permission,
		Role:       domain.Role,
	})
	return &service.Usecase{
		Authentication: auth,
		Authorization:  authorization,
		Intools:        intools,
	}
}
