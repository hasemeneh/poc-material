package initiator

import (
	"github.com/hasemeneh/poc-material/poc-login/config"
	"github.com/hasemeneh/poc-material/poc-login/internal/service"
)

func (i *initiator) InitInterface(cfg *config.MainConfig) *service.Interface {
	return &service.Interface{}
}
