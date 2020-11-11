package main

import (
	"fmt"
	"github.com/drprado2/transaction-manager/api/routes"
	"github.com/drprado2/transaction-manager/configs"
	_ "github.com/drprado2/transaction-manager/docs"
	dependencyInjection "github.com/drprado2/transaction-manager/pkg/dependency-injection"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Transaction Manager
// @version 1.0
// @description Manage your transactions.
// @termsOfService http://swagger.io/terms/

// @contact.name Adriano Oliveira
// @contact.url https://github.com/drprado2
// @contact.email drprado2@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:9000
// @BasePath /api/v1
func main() {
	gin.ForceConsoleColor()

	log.Println("Starting http server")

	configuration := configs.GetConfig()
	router := gin.New()

	swaggerUrl := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", configuration.HttpServerPort))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	appServiceProvider := &dependencyInjection.AppServiceProvider{}

	migrateDb(configuration)

	routes.ConfigureRoutes(router, appServiceProvider)

	if error := router.Run(":" + configuration.HttpServerPort); error != nil {
		panic(error.Error())
	}
}

func migrateDb(config *configs.Configuration) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
		config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	m, err := migrate.New("github.com/drprado2/transaction-manager/storage/postgres-db/migrations", connString)
	if err != nil {
		panic(err)
	}
	m.Steps(2)
}
