package main

import (
	"fmt"
	"net/http"
	"order/pkg/handler"
	"order/pkg/initializer"
	"order/pkg/repository"
	"order/pkg/router"
	"order/pkg/service"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
)

var dbConn *gorm.DB
var validate *validator.Validate

func init() {
	fmt.Println("Loading Env vars ...")
	initializer.LoadEnvVars()
	fmt.Println("Env vars loaded successfully , initialization to DB started ...")
	dbConn = initializer.DBConnection()
	validate = validator.New()
	fmt.Println("DB connection completed. Migration Running ...")
	error := initializer.RunDBMigration()

	if error == migrate.ErrNoChange {
		fmt.Println("No change detected in migration ... ")
		return
	}

	if error != nil {
		fmt.Println("DB migration failed ... ")
		return
	}

	fmt.Println("Application started successfully ... ")
}

func main() {

	// Repository
	orderRepository := repository.NewOrderRepositoryImpl(dbConn)

	// Service
	orderService := service.NewOrderServiceImpl(orderRepository, validate)

	// Controller
	orderController := handler.NewOrderHandlerImpl(orderService)

	routes := router.NewRouter(orderController)

	port := os.Getenv("PORT")
	log.Printf("Server started at port : %s", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: routes,
	}

	err := server.ListenAndServe()

	panic(err)

}
