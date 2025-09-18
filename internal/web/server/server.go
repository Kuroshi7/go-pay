package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kuroshi7/go-pay/internal/service"
	"github.com/kuroshi7/go-pay/internal/web/handlers"
	"github.com/kuroshi7/go-pay/internal/web/middleware"
)

type Server struct {
	//Server fields
	router *chi.Mux
	server *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port string

}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router: chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port: port,
	}
}

func (s *Server) ConfigureRoutes(){
	acocuntHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middleware.NewAuthMiddleware(s.accountService)



	s.router.Post("/accounts", acocuntHandler.Create)
	s.router.Get("/accounts", acocuntHandler.Get)
	
	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		r.Post("/invoices", invoiceHandler.Create)
		r.Get("/invoices/{id}", invoiceHandler.GetByID)
		r.Get("/invoices", invoiceHandler.ListByAccount)

	})

}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,

}
	return s.server.ListenAndServe()
}