package config

type MainConfig struct {
	DBConnectionString DatabaseConnectionString `json:"db_connection"`
	RestPort           string                   `json:"rest_port"`
	CacheConnection    CacheConnection          `json:"cache_connection"`
	AppConfig          AppConfig                `json:"app_config"`
}
type CacheConnection struct {
	CacheAddress  string `json:"cache_address"`
	CachePassword string `json:"cache_password"`
}
type AppConfig struct {
	Secret string `json:"secret"`
}
type DatabaseConnectionString struct {
	Master string `json:"master"`
	Slave  string `json:"slave"`
}
