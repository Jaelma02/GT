// SPDX-License-Identifier: MIT
package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&NotaryContract{})
	if err != nil {
		log.Panicf("Erro ao criar chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Erro ao iniciar chaincode: %v", err)
	}
}
