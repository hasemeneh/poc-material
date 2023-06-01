package runner

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/handler/rest"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/service"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/service/initiator"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/logger"
)

type runner struct {
	initinitiator initiator.Initiator
	restBuilder   rest.Builder
	fatalln       logger.LogFunc
}

type Runner interface {
	RunWebService()
}

func NewRunner() Runner {
	return &runner{
		initinitiator: initiator.NewInitiator(),
		restBuilder:   rest.GetDefaultBuilder(),
	}
}

func (a *runner) AssembleService() *service.Service {
	cfg := a.initinitiator.InitConfig()
	basic := a.initinitiator.InitBasic(cfg)
	domain := a.initinitiator.InitDomain(basic)
	extInterface := a.initinitiator.InitInterface(cfg)
	usecase := a.initinitiator.InitUsecase(domain, extInterface, cfg)
	return &service.Service{
		Cfg:     cfg,
		Usecase: usecase,
	}
}

func (a *runner) RunWebService() {
	service := a.AssembleService()
	webService := a.restBuilder.Build(service)
	go func() {
		<-a.terminateSignal()
		log.Println("Shutting down gracefully....")
		if err := webService.GracefulStop(); err != nil {
			a.fatalln("server shutdown failed because of ", err.Error())
		}
	}()
	log.Println("Running Server on ", service.Cfg.RestPort)
	if err := webService.Run(); err != nil && err != http.ErrServerClosed {
		a.fatalln("Failed to run web server because of ", err.Error())
	}

}

func (a *runner) terminateSignal() chan os.Signal {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	return term
}
