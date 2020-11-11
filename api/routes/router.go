package routes

import (
	"database/sql"
	"github.com/drprado2/transaction-manager/api/routes/accounts"
	"github.com/drprado2/transaction-manager/api/routes/transactions"
	dependencyInjection "github.com/drprado2/transaction-manager/pkg/dependency-injection"
	"github.com/gin-gonic/gin"
)

var dbConnection *sql.DB

func ConfigureRoutes(router *gin.Engine, serviceProvider dependencyInjection.ServiceProvider) {
	accountsRouter := accounts.NewAccountRouter(serviceProvider, serviceProvider)
	transactionsRouter := transactions.NewTransactionRouter(serviceProvider, serviceProvider)

	apiGroup := router.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			accountsGroup := v1Group.Group("/accounts")
			{
				accountsRouter.ConfigureRoutes(accountsGroup)
			}
			transactionsGroup := v1Group.Group("/transactions")
			{
				transactionsRouter.ConfigureRoutes(transactionsGroup)
			}
		}
	}
}
