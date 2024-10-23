package internal

import (
	"database/sql"
	"fmt"
	"server-pulsa-app/config"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type Server struct {
	// usecase
	engine *gin.Engine
	host   string
}

func (s *Server) initRoute() {
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	// change the _ prefix into db for injecting layer dependencies
	_, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		fmt.Println("connection error", err)
	}

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		engine: engine,
		host:   host,
	}
}
