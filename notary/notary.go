// SPDX-License-Identifier: MIT

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// =======================
// Estruturas de dados
// =======================

// Estrutura para armazenar informações da transferência
type NFTTransfer struct {
	NFTContract        string `json:"nftContract"`
	Sender             string `json:"sender"`
	ReceiverOnChain    string `json:"receiverOnChain"`
	ReceiverInterChain string `json:"receiverInterChain"`
	TokenId            int    `json:"tokenId"`
	TransferId         int    `json:"transferId"`
}

// Estrutura para armazenar informações de mint
type MintRecord struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	TokenId  int    `json:"tokenId"`
	MintId   int    `json:"mintId"`
}

// Chaincode principal
type NotaryContract struct {
	contractapi.Contract
}

// =======================
// Funções do contrato
// =======================

// Inicializa o contrato definindo o dono
func (c *NotaryContract) InitLedger(ctx contractapi.TransactionContextInterface, notaryOwner string) error {
	return ctx.GetStub().PutState("notaryOwner", []byte(notaryOwner))
}

// Retorna o dono do notary
func (c *NotaryContract) GetNotaryOwner(ctx contractapi.TransactionContextInterface) (string, error) {
	ownerBytes, err := ctx.GetStub().GetState("notaryOwner")
	if err != nil {
		return "", err
	}
	return string(ownerBytes), nil
}

// Função para registrar uma transferência inter-chain de NFT
func (c *NotaryContract) TransferNFTInterChain(ctx contractapi.TransactionContextInterface,
	nftContract string, tokenId int, sender string, receiverInterChain string) error {

	// Incrementa contador de transferências
	transferCounterKey := "transferCounter"
	counterBytes, err := ctx.GetStub().GetState(transferCounterKey)
	var counter int
	if err != nil {
		return err
	}
	if counterBytes != nil {
		counter, _ = strconv.Atoi(string(counterBytes))
	}
	counter++

	// Obtém o notary owner
	owner, err := c.GetNotaryOwner(ctx)
	if err != nil {
		return err
	}

	// Cria registro
	transfer := NFTTransfer{
		NFTContract:        nftContract,
		Sender:             sender,
		ReceiverOnChain:    owner,
		ReceiverInterChain: receiverInterChain,
		TokenId:            tokenId,
		TransferId:         counter,
	}

	transferJSON, _ := json.Marshal(transfer)

	// Salva no ledger
	key := "transfer_" + strconv.Itoa(counter)
	err = ctx.GetStub().PutState(key, transferJSON)
	if err != nil {
		return err
	}

	// Atualiza contador
	err = ctx.GetStub().PutState(transferCounterKey, []byte(strconv.Itoa(counter)))
	if err != nil {
		return err
	}

	// Também pode emitir evento se quiser
	ctx.GetStub().SetEvent("NFTTransferInterChain", transferJSON)

	return nil
}

// Função para mintar NFT
func (c *NotaryContract) MintNFT(ctx contractapi.TransactionContextInterface, to string, senderInterChain string) (int, error) {
	// Incrementa IDs
	tokenCounterKey := "tokenCounter"
	mintCounterKey := "mintCounter"

	tokenBytes, _ := ctx.GetStub().GetState(tokenCounterKey)
	mintBytes, _ := ctx.GetStub().GetState(mintCounterKey)

	var tokenCounter, mintCounter int
	if tokenBytes != nil {
		tokenCounter, _ = strconv.Atoi(string(tokenBytes))
	}
	if mintBytes != nil {
		mintCounter, _ = strconv.Atoi(string(mintBytes))
	}

	tokenCounter++
	mintCounter++

	// Cria registro
	record := MintRecord{
		Sender:   senderInterChain,
		Receiver: to,
		TokenId:  tokenCounter,
		MintId:   mintCounter,
	}
	recordJSON, _ := json.Marshal(record)

	// Salva no ledger
	key := "mint_" + strconv.Itoa(mintCounter)
	err := ctx.GetStub().PutState(key, recordJSON)
	if err != nil {
		return 0, err
	}

	// Atualiza contadores
	err = ctx.GetStub().PutState(tokenCounterKey, []byte(strconv.Itoa(tokenCounter)))
	if err != nil {
		return 0, err
	}
	err = ctx.GetStub().PutState(mintCounterKey, []byte(strconv.Itoa(mintCounter)))
	if err != nil {
		return 0, err
	}

	// Também pode emitir evento
	ctx.GetStub().SetEvent("NFTMinted", recordJSON)

	return tokenCounter, nil
}

// Recupera uma transferência pelo ID
func (c *NotaryContract) GetTransferById(ctx contractapi.TransactionContextInterface, transferId int) (*NFTTransfer, error) {
	key := "transfer_" + strconv.Itoa(transferId)
	bytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return nil, fmt.Errorf("transfer %d não encontrada", transferId)
	}

	var transfer NFTTransfer
	_ = json.Unmarshal(bytes, &transfer)
	return &transfer, nil
}

// Conta quantas transferências já foram feitas
func (c *NotaryContract) GetNftTransfers(ctx contractapi.TransactionContextInterface) (int, error) {
	counterBytes, err := ctx.GetStub().GetState("transferCounter")
	if err != nil {
		return 0, err
	}
	if counterBytes == nil {
		return 0, nil
	}
	counter, _ := strconv.Atoi(string(counterBytes))
	return counter, nil
}
