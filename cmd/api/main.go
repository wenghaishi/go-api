package main

import (
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/wenghaishi/go-api/internal/handlers"

	//"github.com/wenghaishi/go-api/internal/handlers"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting server...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
