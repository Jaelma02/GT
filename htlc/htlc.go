package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// ======================================
// Estrutura do token ERC721 simulado
// ======================================
type ERC721Token struct {
	TokenID uint64 `json:"tokenId"`
	Owner   string `json:"owner"`
}

// ======================================
// Estrutura do HTLC
// ======================================
type HTLC struct {
	HTLCID    string `json:"htlcId"`
	StartTime int64  `json:"startTime"`
	LockTime  int64  `json:"lockTime"` // segundos
	Secret    string `json:"secret"`   // armazenado quando withdraw é feito
	Hash      string `json:"hash"`     // hash do segredo
	Recipient string `json:"recipient"`
	Owner     string `json:"owner"`
	TokenID   uint64 `json:"tokenId"`
}

// ======================================
// SmartContract
// ======================================
type SmartContract struct {
	contractapi.Contract
}

// ======================================
// Função para criar token ERC721
// ======================================
func (s *SmartContract) MintToken(ctx contractapi.TransactionContextInterface, tokenID uint64, owner string) error {
	token := ERC721Token{
		TokenID: tokenID,
		Owner:   owner,
	}
	key := fmt.Sprintf("TOKEN_%d", tokenID)
	tokenJSON, _ := json.Marshal(token)
	return ctx.GetStub().PutState(key, tokenJSON)
}

// ======================================
// Função para criar HTLC
// ======================================
func (s *SmartContract) CreateHTLC(ctx contractapi.TransactionContextInterface, htlcID, owner, recipient, secret string, tokenID uint64, lockTime int64) error {
	// Checar se token existe e dono é o owner
	tokenKey := fmt.Sprintf("TOKEN_%d", tokenID)
	tokenBytes, err := ctx.GetStub().GetState(tokenKey)
	if err != nil || tokenBytes == nil {
		return fmt.Errorf("Token %d não existe", tokenID)
	}
	var token ERC721Token
	json.Unmarshal(tokenBytes, &token)
	if token.Owner != owner {
		return fmt.Errorf("Apenas o dono do token pode criar HTLC")
	}

	// Calcular hash do segredo
	hash := sha256.Sum256([]byte(secret))

	htlc := HTLC{
		HTLCID:    htlcID,
		StartTime: 0, // ainda não financiado
		LockTime:  lockTime,
		Secret:    "",
		Hash:      fmt.Sprintf("%x", hash),
		Recipient: recipient,
		Owner:     owner,
		TokenID:   tokenID,
	}

	htlcJSON, _ := json.Marshal(htlc)
	return ctx.GetStub().PutState("HTLC_"+htlcID, htlcJSON)
}

// ======================================
// Fund HTLC: transfere token para o HTLC
// ======================================
func (s *SmartContract) FundHTLC(ctx contractapi.TransactionContextInterface, htlcID, sender string) error {
	htlcKey := "HTLC_" + htlcID
	htlcBytes, err := ctx.GetStub().GetState(htlcKey)
	if err != nil || htlcBytes == nil {
		return fmt.Errorf("HTLC não encontrado")
	}

	var htlc HTLC
	json.Unmarshal(htlcBytes, &htlc)

	// Checar dono do token
	tokenKey := fmt.Sprintf("TOKEN_%d", htlc.TokenID)
	tokenBytes, _ := ctx.GetStub().GetState(tokenKey)
	var token ERC721Token
	json.Unmarshal(tokenBytes, &token)

	if token.Owner != sender {
		return fmt.Errorf("Apenas o dono do token pode financiar o HTLC")
	}

	// Transferir token para HTLC
	token.Owner = "HTLC_" + htlcID
	tokenJSON, _ := json.Marshal(token)
	ctx.GetStub().PutState(tokenKey, tokenJSON)

	// Atualizar HTLC
	htlc.StartTime = time.Now().Unix()
	htlcBytes, _ = json.Marshal(htlc)
	ctx.GetStub().PutState(htlcKey, htlcBytes)

	// Emitir evento
	ctx.GetStub().SetEvent("Funded", []byte(fmt.Sprintf("HTLC %s funded by %s", htlcID, sender)))
	return nil
}

// ======================================
// Withdraw: receptor resgata token se segredo correto
// ======================================
func (s *SmartContract) WithdrawHTLC(ctx contractapi.TransactionContextInterface, htlcID, sender, secret string) error {
	htlcKey := "HTLC_" + htlcID
	htlcBytes, err := ctx.GetStub().GetState(htlcKey)
	if err != nil || htlcBytes == nil {
		return fmt.Errorf("HTLC não encontrado")
	}

	var htlc HTLC
	json.Unmarshal(htlcBytes, &htlc)

	// Checar prazo
	if time.Now().Unix() > htlc.StartTime+htlc.LockTime {
		return s.RefundHTLC(ctx, htlcID)
	}

	// Checar hash
	hash := sha256.Sum256([]byte(secret))
	if fmt.Sprintf("%x", hash) != htlc.Hash {
		return fmt.Errorf("Segredo incorreto")
	}

	if sender != htlc.Recipient {
		return fmt.Errorf("Apenas o destinatário pode sacar")
	}

	// Transferir token
	tokenKey := fmt.Sprintf("TOKEN_%d", htlc.TokenID)
	tokenBytes, _ := ctx.GetStub().GetState(tokenKey)
	var token ERC721Token
	json.Unmarshal(tokenBytes, &token)
	token.Owner = sender
	tokenJSON, _ := json.Marshal(token)
	ctx.GetStub().PutState(tokenKey, tokenJSON)

	// Atualizar HTLC com segredo
	htlc.Secret = secret
	htlcBytes, _ = json.Marshal(htlc)
	ctx.GetStub().PutState(htlcKey, htlcBytes)

	// Emitir evento
	ctx.GetStub().SetEvent("Withdrawn", []byte(fmt.Sprintf("HTLC %s withdrawn by %s", htlcID, sender)))
	return nil
}

// ======================================
// Refund: devolve token para o dono se expirou
// ======================================
func (s *SmartContract) RefundHTLC(ctx contractapi.TransactionContextInterface, htlcID string) error {
	htlcKey := "HTLC_" + htlcID
	htlcBytes, err := ctx.GetStub().GetState(htlcKey)
	if err != nil || htlcBytes == nil {
		return fmt.Errorf("HTLC não encontrado")
	}

	var htlc HTLC
	json.Unmarshal(htlcBytes, &htlc)

	tokenKey := fmt.Sprintf("TOKEN_%d", htlc.TokenID)
	tokenBytes, _ := ctx.GetStub().GetState(tokenKey)
	var token ERC721Token
	json.Unmarshal(tokenBytes, &token)

	if token.Owner != "HTLC_"+htlcID {
		return fmt.Errorf("HTLC não possui o token")
	}

	// Devolver token
	token.Owner = htlc.Owner
	tokenJSON, _ := json.Marshal(token)
	ctx.GetStub().PutState(tokenKey, tokenJSON)

	// Emitir evento
	ctx.GetStub().SetEvent("Refunded", []byte(fmt.Sprintf("HTLC %s refunded to %s", htlcID, htlc.Owner)))
	return nil
}

// ======================================
// Main
// ======================================
