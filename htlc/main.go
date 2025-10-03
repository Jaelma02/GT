package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

func main() {
	// Cria instância do chaincode baseado no SmartContract definido em htlc.go
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Erro ao criar chaincode: %s\n", err.Error())
		return
	}

	// Inicia o chaincode
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Erro ao iniciar chaincode: %s\n", err.Error())
	}
}
