package app

import (
	"test-task/internal/app/endpoint"
	"test-task/internal/app/service"
)

type App struct {
	e *endpoint.Endpoint
	s *service.Service
}

func New() (*App, error) {
	a := &App{}
	a.s = service.NewService()
	a.e = endpoint.New(*a.s)

	return a, nil
}

func (a *App) Run() error {
	a.e.Print()
	return nil
}
