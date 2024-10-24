package internal

import (
	"database/sql"
	"fmt"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/handler"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/repository"
	"server-pulsa-app/internal/shared/service"
	"server-pulsa-app/internal/usecase"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type Server struct {
	jwtService    service.JwtService
	authUc        usecase.AuthUseCase
	productUc     usecase.ProductUseCase
	merchantUc    usecase.MerchantUseCase
	transactionUc usecase.TransactionUseCase
	engine        *gin.Engine
	host          string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)

	handler.NewMerchantHandler(s.merchantUc, authMiddleware, rg).Route()
	handler.NewAuthController(s.authUc, rg).Route()
	handler.NewProductController(s.productUc, rg, authMiddleware).Route()
	handler.NewTransactionHandler(s.transactionUc, authMiddleware, rg).Route()
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

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		fmt.Println("connection error", err)
	}

	//inject dependencies repo layer
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	merchantRepo := repository.NewMerchantRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	//inject dependencies usecase layer
	jwtService := service.NewJwtService(cfg.TokenConfig)
	userUc := usecase.NewUserUsecase(userRepo)
	authUc := usecase.NewAuthUseCase(userUc, jwtService)
	productUc := usecase.NewProductUseCase(productRepo)
	merchantUc := usecase.NewMerchantUseCase(merchantRepo)
	transactionUc := usecase.NewTransactionUseCase(transactionRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		jwtService:    jwtService,
		authUc:        authUc,
		productUc:     productUc,
		merchantUc:    merchantUc,
		transactionUc: transactionUc,

		engine: engine,
		host:   host,
	}
}
