package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func (a *App) handlerDefault(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	_ = t.Execute(w, nil)
}

func (a *App) handlerStatus(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(a.e.Status()))
}

func (a *App) handlerPair(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PAIR")

	_ = r.ParseForm()
	curr1 := r.Form.Get("curr1")
	curr2 := r.Form.Get("curr2")

	_, _ = w.Write([]byte(a.e.Pair(curr1, curr2)))
}

func (a *App) Run() error {
	r := chi.NewRouter()

	// Добавляем middleware для обработки POST запросов
	r.Use(middleware.AllowContentType("application/x-www-form-urlencoded"))

	r.HandleFunc("/", a.handlerDefault)
	r.HandleFunc("/status", a.handlerStatus)
	r.HandleFunc("/pair", a.handlerPair)
	http.ListenAndServe("localhost:8080", r)
	return nil
}
