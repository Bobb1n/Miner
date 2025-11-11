package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	httphandlers *HttpHandlers
}

func NewHttpServer(httpHandlers *HttpHandlers) *HttpServer {
	return &HttpServer{
		httphandlers: httpHandlers,
	}
}

func (h *HttpServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/api/user").Methods("GET").HandlerFunc(h.httphandlers.HandleGetInfoUser)
	router.Path("/api/user/stop").Methods("POST").HandlerFunc(h.httphandlers.HandleStopGame)
	router.Path("/api/miners/types").Methods("GET").HandlerFunc(h.httphandlers.HandleInfoSalryMiner)
	router.Path("/api/miners").Methods("POST").HandlerFunc(h.httphandlers.HandleAddNewMiner)
	router.Path("/api/miners").Methods("GET").HandlerFunc(h.httphandlers.HandleGetInfoMiner)
	router.Path("/api/upgrades").Methods("GET").HandlerFunc(h.httphandlers.HandleGetInfoQueryUpdrades)
	router.Path("/api/upgrades").Methods("POST").HandlerFunc(h.httphandlers.HandleAddUpgrades)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
