package initiator

import (
	"github.com/hasemeneh/poc-material/poc-pricing/config"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/service"
)

func (i *initiator) InitInterface(cfg *config.MainConfig) *service.Interface {
	return &service.Interface{}
}
