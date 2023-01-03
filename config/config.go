package config

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ContractAddress common.Address

var Client *ethclient.Client

func ConnectClient() *ethclient.Client {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	return Client
}

func GetClient() *ethclient.Client {
	return Client
}
