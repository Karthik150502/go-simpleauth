package main

import (
	"fmt"
	"net/http"

	"simple_auth/internal/handler"
	// "simple_auth/internal/lib/db"

	// pg "simple_auth/internal/lib/db"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handler.Handler(r)

	// db.InitGormDb()
	// pg, initErr := db.GetDb()
	// if initErr != nil {
	// 	log.Fatal(initErr)
	// } else {
	// 	db.InsertUser(pg)
	// }
	fmt.Println("Starting the server at PORT : 8000")
	var err = http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
