package internal

import (
	"database/sql"
	"fmt"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/handler"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/repository"
	"server-pulsa-app/internal/shared/service"
	"server-pulsa-app/internal/usecase"

	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "server-pulsa-app/docs"

	"github.com/gin-gonic/gin"
)

// @title Server Pulsa API
// @version 1.0
// @description API Server for Pulsa Application

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @host 167.172.91.111:8080
// @BasePath /api/v1
// @schemes http https
type Server struct {
	jwtService    service.JwtService
	authUc        usecase.AuthUseCase
	productUc     usecase.ProductUseCase
	merchantUc    usecase.MerchantUseCase
	transactionUc usecase.TransactionUseCase
	userUc        usecase.UserUsecase
	reportUc      usecase.ReportUseCase
	topupUc       usecase.TopupUseCase

	engine *gin.Engine
	host   string
}

var log = logger.NewLogger()

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)

	handler.NewMerchantHandler(s.merchantUc, authMiddleware, rg, &log).Route()
	handler.NewAuthController(s.authUc, rg, &log).Route()
	handler.NewProductController(s.productUc, rg, authMiddleware, &log).Route()
	handler.NewTransactionHandler(s.transactionUc, authMiddleware, rg, &log).Route()
	handler.NewUserHandler(s.userUc, authMiddleware, rg, &log).Route()
	handler.NewReportHandler(s.reportUc, authMiddleware, rg, &log).Route()
	handler.NewTopupHandler(s.topupUc, authMiddleware, rg, &log).Route()

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	userRepo := repository.NewUserRepository(db, &log)
	productRepo := repository.NewProductRepository(db, &log)
	merchantRepo := repository.NewMerchantRepository(db, &log)
	transactionRepo := repository.NewTransactionRepository(db, &log)
	reportRepo := repository.NewReportRepository(db, &log)
	topupRepo := repository.NewTopupRepository(db)

	//inject dependencies usecase layer
	jwtService := service.NewJwtService(cfg.TokenConfig)
	userUc := usecase.NewUserUsecase(userRepo, &log)
	authUc := usecase.NewAuthUseCase(userUc, jwtService, &log)
	productUc := usecase.NewProductUseCase(productRepo, &log)
	merchantUc := usecase.NewMerchantUseCase(merchantRepo, &log)
	transactionUc := usecase.NewTransactionUseCase(transactionRepo, &log)
	reportUc := usecase.NewReportUseCase(reportRepo, &log)
	topupUc := usecase.NewTopupUsecase(topupRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		jwtService:    jwtService,
		authUc:        authUc,
		productUc:     productUc,
		merchantUc:    merchantUc,
		transactionUc: transactionUc,
		userUc:        userUc,
		reportUc:      reportUc,
		topupUc:       topupUc,

		engine: engine,
		host:   host,
	}
}
