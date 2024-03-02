package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"test-task/internal/app/endpoint"
	"test-task/internal/app/service"
	"text/template"
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

type Data struct {
	Curr string
}

func (a *App) handler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Curr: a.e.Print(),
	}
	t, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	_ = t.Execute(w, data)
}

func (a *App) Run() error {
	a.e.Print()
	r := chi.NewRouter()
	r.HandleFunc("/", a.handler)
	http.ListenAndServe("localhost:8080", r)
	return nil
}
