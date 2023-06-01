module github.com/hasemeneh/poc-material/poc-login

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/go-chi/chi v1.5.4
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/golang/mock v1.6.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/julienschmidt/httprouter v1.3.0
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/onsi/gomega v1.24.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.14.0
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/viper v1.14.0 // indirect
	github.com/stretchr/testify v1.8.1
	golang.org/x/crypto v0.1.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab
