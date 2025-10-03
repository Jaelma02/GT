package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// =============================
// Configurações do cliente
// =============================
const (
	mspID        = "Org1MSP"
	cryptoPath   = "/mnt/c/Users/jaelm/Desktop/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
	certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem"
	keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint = "localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
	channelName  = "mychannel"
	chaincodeID  = "htlc"
)

// =============================
// Função auxiliar para conectar gRPC
// =============================
func newGrpcConnection() (*grpc.ClientConn, error) {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		return nil, err
	}
	cred := credentials.NewClientTLSFromCert(certificate, gatewayPeer)
	return grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(cred))
}

func loadCertificate(path string) (*x509.CertPool, error) {
	certData, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certData) {
		return nil, fmt.Errorf("falha ao adicionar certificado TLS")
	}
	return certPool, nil
}

// =============================
// Cria a identidade X509
// =============================
func newIdentity() (*identity.X509Identity, error) {
	certPEM, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, fmt.Errorf("falha ao decodificar certificado PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	id, err := identity.NewX509Identity(mspID, cert)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// =============================
// Cria o signer a partir da chave privada
// =============================
func newSign() (identity.Sign, error) {
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		return nil, err
	}

	var keyFile string
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), "_sk") {
			keyFile = filepath.Join(keyPath, f.Name())
			break
		}
	}
	if keyFile == "" {
		return nil, fmt.Errorf("arquivo de chave privada não encontrado")
	}

	privateKeyPEM, err := ioutil.ReadFile(filepath.Clean(keyFile))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return nil, fmt.Errorf("falha ao decodificar chave privada PEM")
	}

	var privateKey interface{}
	if pk, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		privateKey = pk
	} else if pk, err := x509.ParseECPrivateKey(block.Bytes); err == nil {
		privateKey = pk
	} else {
		return nil, fmt.Errorf("não foi possível parsear private key: %v", err)
	}

	return identity.NewPrivateKeySign(privateKey)
}

// =============================
// Função principal
// =============================
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run client.go <comando> [args...]")
		fmt.Println("Comandos disponíveis:")
		fmt.Println("  mint <tokenId> <owner>")
		fmt.Println("  createHTLC <htlcId> <owner> <recipient> <secret> <tokenId> <lockTime>")
		fmt.Println("  fundHTLC <htlcId> <sender>")
		fmt.Println("  withdrawHTLC <htlcId> <recipient> <secret>")
		fmt.Println("  refundHTLC <htlcId>")
		fmt.Println("  getTransfer <transferId>")
		os.Exit(1)
	}

	command := os.Args[1]

	grpcConn, err := newGrpcConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar gRPC: %v", err)
	}
	defer grpcConn.Close()

	id, err := newIdentity()
	if err != nil {
		log.Fatalf("Erro ao criar identidade: %v", err)
	}
	sign, err := newSign()
	if err != nil {
		log.Fatalf("Erro ao criar signer: %v", err)
	}

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(grpcConn),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(15*time.Second),
		client.WithCommitStatusTimeout(60*time.Second),
	)
	if err != nil {
		log.Fatalf("Erro ao conectar gateway: %v", err)
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeID)

	switch command {
	case "mint":
		if len(os.Args) != 4 {
			log.Fatal("Uso: mint <tokenId> <owner>")
		}
		tokenId := os.Args[2]
		owner := os.Args[3]
		_, err := contract.SubmitTransaction("MintToken", tokenId, owner)
		if err != nil {
			log.Fatalf("Erro MintToken: %v", err)
		}
		fmt.Println("Token mintado com sucesso")

	case "createHTLC":
		if len(os.Args) != 8 {
			log.Fatal("Uso: createHTLC <htlcId> <owner> <recipient> <secret> <tokenId> <lockTime>")
		}
		htlcId := os.Args[2]
		owner := os.Args[3]
		recipient := os.Args[4]
		secret := os.Args[5]
		tokenId := os.Args[6]
		lockTime := os.Args[7]
		_, err := contract.SubmitTransaction("CreateHTLC", htlcId, owner, recipient, secret, tokenId, lockTime)
		if err != nil {
			log.Fatalf("Erro CreateHTLC: %v", err)
		}
		fmt.Println("HTLC criado com sucesso")

	case "fundHTLC":
		if len(os.Args) != 4 {
			log.Fatal("Uso: fundHTLC <htlcId> <sender>")
		}
		htlcId := os.Args[2]
		sender := os.Args[3]
		_, err := contract.SubmitTransaction("FundHTLC", htlcId, sender)
		if err != nil {
			log.Fatalf("Erro FundHTLC: %v", err)
		}
		fmt.Println("HTLC financiado")

	case "withdrawHTLC":
		if len(os.Args) != 5 {
			log.Fatal("Uso: withdrawHTLC <htlcId> <recipient> <secret>")
		}
		htlcId := os.Args[2]
		recipient := os.Args[3]
		secret := os.Args[4]
		_, err := contract.SubmitTransaction("WithdrawHTLC", htlcId, recipient, secret)
		if err != nil {
			log.Fatalf("Erro WithdrawHTLC: %v", err)
		}
		fmt.Println("HTLC sacado")

	case "refundHTLC":
		if len(os.Args) != 3 {
			log.Fatal("Uso: refundHTLC <htlcId>")
		}
		htlcId := os.Args[2]
		_, err := contract.SubmitTransaction("RefundHTLC", htlcId)
		if err != nil {
			log.Fatalf("Erro RefundHTLC: %v", err)
		}
		fmt.Println("HTLC reembolsado")

	case "getTransfer":
		if len(os.Args) != 3 {
			log.Fatal("Uso: getTransfer <transferId>")
		}
		transferId := os.Args[2]
		result, err := contract.EvaluateTransaction("GetTransferById", transferId)
		if err != nil {
			log.Fatalf("Erro GetTransferById: %v", err)
		}
		var transfer map[string]interface{}
		json.Unmarshal(result, &transfer)
		fmt.Printf("Transfer: %+v\n", transfer)

	default:
		fmt.Println("Comando inválido")
	}
}
