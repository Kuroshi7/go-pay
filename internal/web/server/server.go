package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/kuroshi7/go-pay/internal/service"
	"github.com/kuroshi7/go-pay/internal/web/handlers"
	authMiddleware "github.com/kuroshi7/go-pay/internal/web/middleware"
)

type Server struct {
	//Server fields
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	acocuntHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	authMidd := authMiddleware.NewAuthMiddleware(s.accountService)

	// Configure CORS
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-API-Key"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s.router.Post("/accounts", acocuntHandler.Create)
	s.router.Get("/accounts", acocuntHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(authMidd.Authenticate)
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
