package app

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	*http.ServeMux
}

type Handler func(w http.ResponseWriter, r *http.Request)

func New() *App {
	return &App{http.DefaultServeMux}
}

func (a *App) Rout(path string, fn Handler) {
	a.HandleFunc(path, a.recoverPanic(fn))
}

func (a *App) recoverPanic(fn Handler) Handler {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				writer.WriteHeader(500)
				writer.Write([]byte(fmt.Sprintf("%v", err)))
			}
		}()
		fn(writer, request)
	}
}

func (a *App) Listen(port string) error {
	log.Println("server running at http://localhost:" + port)
	return http.ListenAndServe(":"+port, a)
}
