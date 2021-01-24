package main

import (
	"flag"
	_ "github.com/joho/godotenv/autoload"
	"github.com/newsolice/wx-admin/internal/app"
	"net/http"
	"time"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "启动端口")
}

func main() {

	application := app.New()

	application.Rout("/index", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(time.Now().String()))
	})

	application.Rout("/example", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(time.Now().String()))
	})
	application.Listen(port)
}
