package webservice

import "github.com/hasemeneh/poc-material/poc-pricing/pkg/router"

type WebService interface {
	Run() error
	GracefulStop() error
}
type WebRegistrator interface {
	Register(router router.Registrator)
}
