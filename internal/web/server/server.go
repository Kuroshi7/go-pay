package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kuroshi7/go-pay/internal/service"
	"github.com/kuroshi7/go-pay/internal/web/handlers"
)

type Server struct {
	//Server fields
	router *chi.Mux
	server *http.Server
	accountService *service.AccountService
	port string

}

func NewServer(accountService *service.AccountService, port string) *Server {
	return &Server{
		router: chi.NewRouter(),
		accountService: accountService,
		port: port,
	}
}

func (s *Server) ConfigureRoutes(){
	acocuntHandler := handlers.NewAccountHandler(s.accountService)

	s.router.Post("/accounts", acocuntHandler.Create)
	s.router.Get("/accounts", acocuntHandler.Get)

}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,

}
	return s.server.ListenAndServe()
}