package initiator

import (
	"log"

	"github.com/hasemeneh/poc-material/poc-login/config"
	"github.com/hasemeneh/poc-material/poc-login/internal/service"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/hasemeneh/poc-material/poc-login/pkg/logger"
)

type initiator struct {
	configReader      config.ConfigManager
	databaseConnector database.Connector
	fatalln           logger.LogFunc
}

type Initiator interface {
	InitConfig() *config.MainConfig
	InitBasic(cfg *config.MainConfig) *service.Basic
	InitDomain(basic *service.Basic) *service.Domain
	InitUsecase(domain *service.Domain, extInterface *service.Interface, cfg *config.MainConfig) *service.Usecase
	InitInterface(cfg *config.MainConfig) *service.Interface
}

func NewInitiator() Initiator {
	return &initiator{
		configReader:      config.New(),
		fatalln:           log.Fatalln,
		databaseConnector: database.NewConnector(),
	}
}

func (i *initiator) InitConfig() *config.MainConfig {
	return i.configReader.Read()
}
