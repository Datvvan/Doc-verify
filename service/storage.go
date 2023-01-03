package service

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/datvvan/doc-vertify/config"
	"github.com/datvvan/doc-vertify/models"
)

func GetHeader() *models.Block {
	client := config.GetClient()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	_block := &models.Block{
		BlockNumber:      block.Number().Int64(),
		Timestamp:        block.Time(),
		Hash:             block.Hash().String(),
		TransactionCount: len(block.Transactions()),
		Transactions:     []models.Transaction{},
	}
	for _, tx := range block.Transactions() {
		_block.Transactions = append(_block.Transactions, models.Transaction{
			Hash:     tx.Hash().String(),
			Value:    tx.Value().String(),
			Data:     hex.EncodeToString(tx.Data()),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice().Uint64(),
		})
	}
	return _block

}

// func StorageDoc() {
// 	client := config.GetClient()
// 	add := common.HexToAddress("0x4a77eE4264a7535852F72ff9C018E8207BfB94D4")
// 	instance, err := api.NewApi(add, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	name := [32]byte{}
// 	company := [32]byte{}
// 	identity_card := [32]byte{}
// 	hash_docs := [32]byte{}

// 	copy(name[:], byte(""))

// 	tx, err := instance.AddContract()
// }
