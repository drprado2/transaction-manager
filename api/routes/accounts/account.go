package accounts

import (
	dependencyInjection "github.com/drprado2/transaction-manager/pkg/dependency-injection"
	"github.com/drprado2/transaction-manager/pkg/entity"
	"github.com/drprado2/transaction-manager/pkg/models"
	"github.com/drprado2/transaction-manager/pkg/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var database storage.Database

type AccountRouter struct {
	storageServiceProvider dependencyInjection.StorageServiceProvider
	accountServiceProvider dependencyInjection.AccountServiceProvider
}

func NewAccountRouter(storageServiceProvider dependencyInjection.StorageServiceProvider,
	accountServiceProvider dependencyInjection.AccountServiceProvider) *AccountRouter {
	return &AccountRouter{
		storageServiceProvider: storageServiceProvider,
		accountServiceProvider: accountServiceProvider,
	}
}

func (ar *AccountRouter) ConfigureRoutes(router *gin.RouterGroup) {
	router.GET("/:accountId", ar.handleGetById)
	router.POST("/", ar.handleCreate)
}

// Get Account by ID
// @Summary Account by ID
// @Description Get an account by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} models.AccountModel
// @Failure 400 {array} models.Error
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /accounts/{id} [get]
func (ar *AccountRouter) handleGetById(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid accountId %v", accountId)
		return
	}

	db := ar.storageServiceProvider.ResolveDb()
	repository := ar.accountServiceProvider.ResolveAccountRepository(db)
	defer db.Close()

	data := repository.GetById(entity.ID(accountId))

	if err != nil {
		c.JSON(http.StatusBadRequest, []models.Error{{Message: err.Error()}})
		return
	}

	if data == nil {
		c.JSON(http.StatusNotFound, data)
		return
	}

	c.JSON(http.StatusOK, data)
}

// Create Account
// @Summary Create an account
// @Description Create an account to make transactions
// @Accept  json
// @Produce  json
// @Param account body models.CreateAccountModel true "Add account"
// @Success 201 {object} int
// @Failure 400 {array} models.Error
// @Failure 500 {object} string
// @Router /accounts [post]
func (ar *AccountRouter) handleCreate(c *gin.Context) {
	var model models.CreateAccountModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, []models.Error{{Message: err.Error()}})
		return
	}

	unitOfWork := ar.storageServiceProvider.ResolveUnitOfWork()
	repository := ar.accountServiceProvider.ResolveAccountRepository(unitOfWork.GetTxDb())
	service := ar.accountServiceProvider.ResolveCreateAccountService(repository)
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
		unitOfWork.Roolback()
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	unitOfWork.Commit()

	c.JSON(http.StatusCreated, id)
}
