package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wael-boudissaa/marquinoBackend/services/commande"
	"github.com/wael-boudissaa/marquinoBackend/services/product"
	"github.com/wael-boudissaa/marquinoBackend/services/user"
)

type APISERVER struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APISERVER {
	return &APISERVER{addr: addr, db: db}
}

func (s *APISERVER) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	commandeStore := commande.NewStore(s.db)
	commandeHanlder := commande.NewHandler(commandeStore)
	commandeHanlder.RegisterRoutes(subrouter)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
