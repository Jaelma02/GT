/*
Copyright 2021 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"github.com/hyperledger/fabric-protos-go-apiv2/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	mspID        = "Org1MSP"
	cryptoPath   = "../peerOrganizations/org1.example.com"
	certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts"
	keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint = "dns:///localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
)

var now = time.Now()

//var assetId = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)

func main() {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "fabcar"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	// Switch baseado no argumento passado
	operacao := os.Args[1]
	switch operacao {
	case "initLedger":
		initLedger(contract)
	case "getAllAssets":
		getAllAssets(contract)
	case "createAsset":
		n := 1 // Valor padrão para criar um asset
		if len(os.Args) >= 3 {
			numAssets, err := strconv.Atoi(os.Args[2])
			if err == nil {
				n = numAssets
			} else {
				fmt.Println("Erro ao converter número de assets, usando o valor padrão de 1.")
			}
		}
		createAssets(contract, n)
	case "readAssetByID":
		if len(os.Args) < 3 {
			fmt.Println("Uso: go run main.go readAssetByID <assetId>")
			return
		}
		assetId := os.Args[2]
		readAssetByID(contract, assetId)
	case "transferAsset":
		if len(os.Args) < 4 {
			fmt.Println("Uso: go run main.go transferAssetAsync <assetId> <newOwner>")
			return
		}
		assetId := os.Args[2]
		newOwner := os.Args[3]
		transferAssetAsync(contract, assetId, newOwner)
	case "createAssetBench":
		tps := 10        // Valor padrão para TPS
		numAssets := 100 // Número padrão de assets a serem criados

		// Verifica se foi fornecido o número de TPS como argumento
		if len(os.Args) >= 3 {
			tpsVal, err := strconv.Atoi(os.Args[2])
			if err == nil {
				tps = tpsVal
			} else {
				fmt.Println("Erro ao converter TPS, usando o valor padrão de 10.")
			}
		}

		// Verifica se foi fornecido o número de assets como argumento
		if len(os.Args) >= 4 {
			numAssetsVal, err := strconv.Atoi(os.Args[3])
			if err == nil {
				numAssets = numAssetsVal
			} else {
				fmt.Println("Erro ao converter número de assets, usando o valor padrão de 100.")
			}
		}

		createAssetBench(contract, tps, numAssets)
	case "exampleErrorHandling":
		exampleErrorHandling(contract)
	default:
		fmt.Println("Operation not recognized.")
	}
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	certificatePEM, err := os.ReadFile(tlsCertPath)
	if err != nil {
		panic(fmt.Errorf("failed to read TLS certifcate file: %w", err))
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.NewClient(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	certificatePEM, err := readFirstFile(certPath)
	if err != nil {
		panic(fmt.Errorf("failed to read certificate file: %w", err))
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	privateKeyPEM, err := readFirstFile(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func readFirstFile(dirPath string) ([]byte, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fileNames, err := dir.Readdirnames(1)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path.Join(dirPath, fileNames[0]))
}

var methods = []string{
	"InitLedger",
	"CreateCar",
	"QueryAllCars",
	"QueryCar",
	"ChangeCarOwner",
}

func generateRandomHash() string {
	// Gerar uma string aleatória de 8 bytes
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(fmt.Errorf("erro ao gerar bytes aleatórios: %v", err))
	}
	randomString := fmt.Sprintf("%x", randomBytes)

	// Calcular o hash SHA-256 da string aleatória
	hash := sha256.New()
	hash.Write([]byte(randomString))
	hashInBytes := hash.Sum(nil)

	// Converter o hash em uma string hexadecimal
	hashString := fmt.Sprintf("%x", hashInBytes)

	return hashString
}

// This type of transaction would typically only be run once by an application the first time it was started after its
// initial deployment. A new version of the chaincode deployed later would likely not need to run an "init" function.
func initLedger(contract *client.Contract) {
	fmt.Printf("\n--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger \n")

	_, err := contract.SubmitTransaction(methods[0])
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

// Evaluate a transaction to query ledger state.
func getAllAssets(contract *client.Contract) {
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction(methods[2])
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

// Submit transactions synchronously, blocking until each has been committed to the ledger.
func createAssets(contract *client.Contract, n int) {
	if n <= 0 {
		n = 1 // Set n to 1 if it's zero or negative
	}

	fmt.Printf("\n--> Submit Transactions: CreateAsset, creates %d new assets with ID, Color, Size, Owner, and AppraisedValue arguments\n", n)

	for i := 0; i < n; i++ {
		// Assuming generateRandomHash() is defined elsewhere to generate a random hash
		hash := generateRandomHash()

		startTime := time.Now()

		_, err := contract.SubmitTransaction(methods[1], hash, "yellow", "5", "Tom", "1300")
		if err != nil {
			panic(fmt.Errorf("failed to submit transaction: %w", err))
		}

		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)

		fmt.Printf("*** Transaction %s committed successfully\n", hash)
		fmt.Printf("Time taken: %v\n", elapsedTime)
	}
}

func createAssetBench(contract *client.Contract, tps int, numAssets int) {
	fmt.Printf("\n--> Benchmarking CreateAsset, TPS: %d\n", tps)

	// Canal para controlar o TPS
	tpsCh := make(chan struct{}, tps)

	// Canal para controlar o término da criação de assets
	doneCh := make(chan struct{})

	// Contador para transações concluídas
	var completedTransactions int

	// Início do tempo para medição
	startTime := time.Now()

	// Função para criação assíncrona de assets
	createAsync := func(assetId string) {
		defer func() {
			// Sinaliza que a transação foi concluída
			<-tpsCh
		}()

		// Submete a transação para criar um asset
		_, err := contract.SubmitTransaction(methods[1], assetId, "yellow", "5", "Tom", "1300")
		if err != nil {
			fmt.Printf("*** Failed to submit transaction for asset %s: %v\n", assetId, err)
			return
		}

		// Incrementa o contador de transações concluídas de forma segura
		completedTransactions++
		if completedTransactions >= numAssets {
			doneCh <- struct{}{}
		}
	}

	// Loop para criar assets até atingir o número desejado
	for i := 1; i <= numAssets; i++ {
		assetId := generateRandomHash()
		tpsCh <- struct{}{} // Envia um sinal para o canal, controlando o TPS

		// Cria o asset assincronamente
		go createAsync(assetId)

		// Intervalo para controlar o TPS
		time.Sleep(time.Second / time.Duration(tps))
	}

	// Espera até que todas as transações sejam concluídas
	<-doneCh

	// Tempo decorrido desde o início
	elapsedTime := time.Since(startTime)

	// Exibição dos resultados do benchmark
	fmt.Println("\n*** Benchmarking Complete ***")
	fmt.Printf("Transactions executed: %d\n", completedTransactions)
	fmt.Printf("Elapsed time: %s\n", elapsedTime)
	fmt.Printf("TPS achieved: %.2f\n", float64(completedTransactions)/elapsedTime.Seconds())
}

// Evaluate a transaction by assetID to query ledger state.
func readAssetByID(contract *client.Contract, assetId string) {
	fmt.Printf("\n--> Evaluate Transaction: ReadAsset, function returns asset attributes for asset ID: %s\n", assetId)

	evaluateResult, err := contract.EvaluateTransaction(methods[3], assetId)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

// Submit transaction asynchronously, blocking until the transaction has been sent to the orderer, and allowing
// this thread to process the chaincode response (e.g. update a UI) without waiting for the commit notification
func transferAssetAsync(contract *client.Contract, assetId, newOwner string) {
	fmt.Printf("\n--> Async Submit Transaction: TransferAsset, updates existing asset owner")

	submitResult, commit, err := contract.SubmitAsync(methods[4], client.WithArguments(assetId, newOwner))
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction asynchronously: %w", err))
	}

	fmt.Printf("\n*** Successfully submitted transaction to transfer ownership from %s to Mark. \n", string(submitResult))
	fmt.Println("*** Waiting for transaction commit.")

	if commitStatus, err := commit.Status(); err != nil {
		panic(fmt.Errorf("failed to get commit status: %w", err))
	} else if !commitStatus.Successful {
		panic(fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code)))
	}

	fmt.Printf("*** Transaction committed successfully\n")
}

// Submit transaction, passing in the wrong number of arguments ,expected to throw an error containing details of any error responses from the smart contract.
func exampleErrorHandling(contract *client.Contract) {
	fmt.Println("\n--> Submit Transaction: UpdateAsset asset70, asset70 does not exist and should return an error")

	_, err := contract.SubmitTransaction("UpdateAsset", "asset70", "blue", "5", "Tomoko", "300")
	if err == nil {
		panic("******** FAILED to return an error")
	}

	fmt.Println("*** Successfully caught the error:")

	var endorseErr *client.EndorseError
	var submitErr *client.SubmitError
	var commitStatusErr *client.CommitStatusError
	var commitErr *client.CommitError

	if errors.As(err, &endorseErr) {
		fmt.Printf("Endorse error for transaction %s with gRPC status %v: %s\n", endorseErr.TransactionID, status.Code(endorseErr), endorseErr)
	} else if errors.As(err, &submitErr) {
		fmt.Printf("Submit error for transaction %s with gRPC status %v: %s\n", submitErr.TransactionID, status.Code(submitErr), submitErr)
	} else if errors.As(err, &commitStatusErr) {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Printf("Timeout waiting for transaction %s commit status: %s", commitStatusErr.TransactionID, commitStatusErr)
		} else {
			fmt.Printf("Error obtaining commit status for transaction %s with gRPC status %v: %s\n", commitStatusErr.TransactionID, status.Code(commitStatusErr), commitStatusErr)
		}
	} else if errors.As(err, &commitErr) {
		fmt.Printf("Transaction %s failed to commit with status %d: %s\n", commitErr.TransactionID, int32(commitErr.Code), err)
	} else {
		panic(fmt.Errorf("unexpected error type %T: %w", err, err))
	}

	// Any error that originates from a peer or orderer node external to the gateway will have its details
	// embedded within the gRPC status error. The following code shows how to extract that.
	statusErr := status.Convert(err)

	details := statusErr.Details()
	if len(details) > 0 {
		fmt.Println("Error Details:")

		for _, detail := range details {
			switch detail := detail.(type) {
			case *gateway.ErrorDetail:
				fmt.Printf("- address: %s, mspId: %s, message: %s\n", detail.Address, detail.MspId, detail.Message)
			}
		}
	}
}

// Format JSON data
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}
