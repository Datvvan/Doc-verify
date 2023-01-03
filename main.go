package main

import (
	"fmt"
	"os"

	"github.com/datvvan/doc-vertify/api"
	"github.com/datvvan/doc-vertify/config"
	"github.com/datvvan/doc-vertify/controler"
	"github.com/datvvan/doc-vertify/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	client := config.ConnectClient()

	auth := service.GetAccountAuth(client, os.Getenv("PRV_KEY"))
	deployedContractAddress, tx, _, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}
	config.ContractAddress = deployedContractAddress
	fmt.Println(config.ContractAddress.Hex())
	fmt.Println(tx.Hash().Hex())

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/block", controler.GetLatestBlock)
		api.POST("/storagecontract", controler.StorageContract)
		api.GET("/getcontract", controler.GetContract)

	}
	router.Run(":8080")
}
