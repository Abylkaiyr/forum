package apiserver

import (
	"github.com/Abylkaiyr/forum/internal/app/sessions"
	"github.com/Abylkaiyr/forum/internal/app/store"
)

// APISERVER
type APIServer struct {
	config  *Config
	logger  *Logger
	store   store.Store
	session sessions.Sessions
}

// Initialization of APISERVER struct
func (s *APIServer) NewServer(config *Config, logger *Logger) *APIServer {
	return &APIServer{
		config:  config,
		logger:  logger,
		store:   store.Store{},
		session: sessions.Sessions{},
	}
}

// Starting the server
func (s *APIServer) Start() error {
	s.store.CreateTables()
	s.logger.InfoLog.Printf("Starting the server on adress %v", s.config.BindAddr)

	return nil
}
