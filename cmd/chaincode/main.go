package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		log.Panicf("Erro criando chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Erro iniciando chaincode: %v", err)
	}
}
