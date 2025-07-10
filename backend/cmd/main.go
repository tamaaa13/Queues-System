package main

import (
	// "fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tamaaa13/Queue/internal/configs"
	"github.com/tamaaa13/Queue/internal/handlers/admin"
	"github.com/tamaaa13/Queue/internal/handlers/public"
	adminRepo "github.com/tamaaa13/Queue/internal/repository/admin"
	publicRepo "github.com/tamaaa13/Queue/internal/repository/public"
	adminSvc "github.com/tamaaa13/Queue/internal/service/admin"
	publicSvc "github.com/tamaaa13/Queue/internal/service/public"
	"github.com/tamaaa13/Queue/pkg/internalsql"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Println("No .env file found")
	} else {
		log.Println("Environment Variables Loaded")
	}

	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Failed Config Initiation", err)
	}

	cfg = configs.Get()
	log.Println("=== CONFIG LOADED")
	log.Println("config", cfg)
	log.Println("Service Port:", cfg.Service.Port)
	log.Println("JWT Secret:", cfg.Service.SecretJWT)
	log.Println("Database DSN:", cfg.Database.DataSourceName)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed Database Initiation", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public
	publicRepo := publicRepo.NewRepository(db)
	publicService := publicSvc.NewService(cfg, publicRepo)
	publicHandler := public.NewHandler(r, publicService)
	publicHandler.RegisterRoute()

	// Admin
	adminRepo := adminRepo.NewRepository(db)
	adminService := adminSvc.NewService(cfg, adminRepo)
	adminHandler := admin.NewHandler(r, adminService)
	adminHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
