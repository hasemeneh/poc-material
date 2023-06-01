package initiator

import (
	"github.com/hasemeneh/poc-material/poc-pricing/config"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/service"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/redis"
)

func (i *initiator) InitBasic(cfg *config.MainConfig) *service.Basic {
	db, err := i.databaseConnector.Connect(&database.Option{
		MasterConnectionString: cfg.DBConnectionString.Master,
		SlaveConnectionString:  cfg.DBConnectionString.Slave,
	})
	if err != nil {
		i.fatalln("failed to connect to database ", err)
	}
	cache := redis.New(cfg.CacheConnection.CacheAddress, cfg.CacheConnection.CachePassword)
	return &service.Basic{
		DB:    db,
		Cache: cache,
	}
}
