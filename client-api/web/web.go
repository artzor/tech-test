// Package web provides REST API
package web

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/artzor/tech-test/client-api/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Web provides web server with REST API
type Web struct {
	portDomainSvc portDomainSvc
	port          string
	server        *http.Server
}

type portDomainSvc interface {
	Get(ctx context.Context, portID string) (entity.PortDetails, error)
}

// New creates Web instance
func New(portDomainSvc portDomainSvc, port string) *Web {
	web := &Web{portDomainSvc: portDomainSvc, port: port}
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Get("/port/{portID}", web.portDetails)

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	web.server = server
	return web
}

// Start triggers web service start
func (web *Web) Start() error {
	log.Printf("[info] starting web server on port %s", web.port)
	return web.server.ListenAndServe()
}

// Stop shuts down server
func (web *Web) Stop() error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return web.server.Shutdown(ctx)
}

func (web *Web) portDetails(w http.ResponseWriter, r *http.Request) {
	portID := chi.URLParam(r, "portID")

	if portID == "" {
		respErr(w, http.StatusBadRequest, "port id missing")
		return
	}

	portDetails, err := web.portDomainSvc.Get(r.Context(), portID)
	if err != nil {
		respErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp(w, portDetails)
}

func respErr(w http.ResponseWriter, code int, text string) {
	response := map[string]string{
		"error": text,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("[error] failed to encode error response: %v", err)
	}
	w.WriteHeader(code)
}

func resp(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("[error] failed to encode response: %v", err)
	}
	w.WriteHeader(http.StatusOK)
}
