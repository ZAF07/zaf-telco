package httpserver

import (
	"log"
	"net"
	"net/http"

	"github.com/ZAF07/telco/config"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Config      *config.Config
	Server      *http.Server
	HTTPHandler *gin.Engine
	Listener    net.Listener
}

func NewHTTPServer(listener net.Listener, config *config.Config) *HTTPServer {
	g := gin.Default()
	server := &http.Server{
		ReadTimeout:  config.HTTPServerConfig.ReadTimeout,
		WriteTimeout: config.HTTPServerConfig.WriteTimeout,
		Handler:      g,
	}

	return &HTTPServer{
		Server:      server,
		HTTPHandler: g,
		Listener:    listener,
		Config:      config,
	}
}

func (h *HTTPServer) Start() {
	go h.start()
	log.Println("🔥🔥🔥 Telco App has started successfully 🔥🔥🔥")
}

func (h *HTTPServer) start() {
	if err := h.Server.Serve(h.Listener); err != nil {
		log.Fatalf("error serving HTTP : %+v", err)
	}
}
