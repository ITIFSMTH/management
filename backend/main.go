package main

import (
	"flag"
	"fmt"
	"log"
	"management-backend/cron"
	"management-backend/db"
	grpcServer "management-backend/grpc"
	"management-backend/middlewares"
	"management-backend/routes"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nleeper/goment"
)

func main() {
	time.Sleep(30 * time.Second)
	goment.SetLocale("ru")

	// Define flags
	var env string

	// Get flag env
	flag.StringVar(&env, "env", "dev", "Default environment: dev")
	flag.Parse()

	// If env == dev load a env file
	if env == "dev" {
		if err := godotenv.Load("../env/.env.backend"); err != nil {
			log.Panic("No .env file found")
		}
	}

	// Get ENV variables
	JwtSecret, _ := os.LookupEnv("BACKEND_JWT_SECRET")
	DBPath, _ := os.LookupEnv("BACKEND_DB_PATH")
	RestPort, _ := os.LookupEnv("BACKEND_REST_PORT")
	gRPCPort, _ := os.LookupEnv("BACKEND_GRPC_PORT")

	// Init JWT Middleware
	middlewares.InitJWT(JwtSecret)

	// Init DB
	db.InitDB(DBPath)

	// Init Gin
	r := gin.Default()

	if env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.SetTrustedProxies(nil)
	routes.InitRouter(r)

	// Run gRPC Server
	go grpcServer.Run(fmt.Sprintf(":%s", gRPCPort))

	// Init cron
	go cron.InitCronTasks()

	// Run RESTful API Server
	r.Run(fmt.Sprintf(":%s", RestPort))
}
