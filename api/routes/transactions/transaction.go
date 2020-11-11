package transactions

import (
	dependencyInjection "github.com/drprado2/transaction-manager/pkg/dependency-injection"
	"github.com/drprado2/transaction-manager/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionRouter struct {
	storageServiceProvider     dependencyInjection.StorageServiceProvider
	transactionServiceProvider dependencyInjection.TransactionServiceProvider
}

func NewTransactionRouter(storageServiceProvider dependencyInjection.StorageServiceProvider,
	transactionServiceProvider dependencyInjection.TransactionServiceProvider) *TransactionRouter {
	return &TransactionRouter{
		storageServiceProvider:     storageServiceProvider,
		transactionServiceProvider: transactionServiceProvider,
	}
}

func (tr *TransactionRouter) ConfigureRoutes(router *gin.RouterGroup) {
	router.POST("/", tr.handleCreate)
}

// Create Transaction
// @Summary Create a transaction
// @Description Create a transaction
// @Accept  json
// @Produce  json
// @Param account body models.CreateTransactionModel true "Add transaction"
// @Success 201 {object} int
// @Failure 400 {array} models.Error
// @Failure 500 {object} string
// @Router /transactions [post]
func (tr *TransactionRouter) handleCreate(c *gin.Context) {
	var model models.CreateTransactionModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, []models.Error{{Message: err.Error()}})
		return
	}

	unitOfWork := tr.storageServiceProvider.ResolveUnitOfWork()
	repository := tr.transactionServiceProvider.ResolveTransactionRepository(unitOfWork.GetTxDb())
	service := tr.transactionServiceProvider.ResolveCreateTransactionService(repository)
	defer unitOfWork.Close()

	unitOfWork.BeginTran()

	defer func() {
		if error := recover(); error != nil {
			unitOfWork.Roolback()
			panic(error)
		}
	}()

	id, errors := service.Execute(&model)

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, errors)
		unitOfWork.Roolback()
		return
	}

	unitOfWork.Commit()

	c.JSON(http.StatusCreated, id)
}
