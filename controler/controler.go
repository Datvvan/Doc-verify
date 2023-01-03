package controler

import (
	"log"
	"net/http"
	"os"

	"github.com/datvvan/doc-vertify/api"
	"github.com/datvvan/doc-vertify/config"
	"github.com/datvvan/doc-vertify/models"
	"github.com/datvvan/doc-vertify/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetLatestBlock(context *gin.Context) {
	block := service.GetHeader()
	context.JSON(http.StatusOK, gin.H{"block": block})
}

func StorageContract(context *gin.Context) {
	godotenv.Load(".env")
	arg := models.StoreContractArgument{}
	if err := context.ShouldBindJSON(&arg); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	add := common.HexToAddress(config.ContractAddress.Hex())
	client := config.GetClient()
	instance, err := api.NewApi(add, client)
	if err != nil {
		log.Fatal(err)
	}
	auth := service.GetAccountAuth(client, os.Getenv("PRV_KEY"))

	tx, err := instance.AddContract(auth, arg.Name, arg.Company, arg.IdentityCard, arg.HashDocs)
	if err != nil {
		log.Fatal(err)
	}

	txResponse := models.Transaction{
		Hash:     tx.Hash().Hex(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		Data:     string(tx.Data()),
		GasPrice: tx.GasPrice().Uint64(),
	}
	context.JSON(http.StatusCreated, gin.H{"transaction": txResponse})
}

func GetContract(context *gin.Context) {
	godotenv.Load(".env")
	arg := models.GetContractArgument{}
	if err := context.ShouldBindJSON(&arg); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	add := common.HexToAddress(config.ContractAddress.Hex())
	client := config.GetClient()
	instance, err := api.NewApi(add, client)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	tx, err := instance.GetContract(nil, arg.IdentityCard, arg.Company)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractResponse := models.GetContractResponse{
		HashDoc: tx,
	}
	context.JSON(http.StatusOK, gin.H{"hash_doc": contractResponse})

}
