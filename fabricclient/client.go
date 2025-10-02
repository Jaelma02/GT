package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	mspID        = "Org1MSP"
	peerEndpoint = "localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
)

var (
	mspPath     = "/mnt/c/Users/jaelm/Desktop/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp"
	certPath    = filepath.Join(mspPath, "signcerts", "User1@org1.example.com-cert.pem")
	keyPath     = filepath.Join(mspPath, "keystore")
	tlsCertPath = "/mnt/c/Users/jaelm/Desktop/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: ./fabric-client <operacao> [args...]")
		os.Exit(1)
	}
	operacao := os.Args[1]

	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	check(err)
	defer gw.Close()

	network := gw.GetNetwork("mychannel")

	htlc := network.GetContract("htlc")
	notary := network.GetContract("notary")

	switch operacao {
	// ------------------- HTLC -------------------
	case "fund":
		if len(os.Args) < 6 {
			log.Fatalln("Uso: ./fabric-client fund <from> <to> <tokenId> <secret>")
		}
		result, err := htlc.SubmitTransaction("Fund", os.Args[2], os.Args[3], os.Args[4], os.Args[5])
		check(err)
		fmt.Println("Fund:", string(result))

	case "withdraw":
		if len(os.Args) < 4 {
			log.Fatalln("Uso: ./fabric-client withdraw <tokenId> <secret>")
		}
		result, err := htlc.SubmitTransaction("Withdraw", os.Args[2], os.Args[3])
		check(err)
		fmt.Println("Withdraw:", string(result))

	case "refund":
		if len(os.Args) < 3 {
			log.Fatalln("Uso: ./fabric-client refund <tokenId>")
		}
		result, err := htlc.SubmitTransaction("Refund", os.Args[2])
		check(err)
		fmt.Println("Refund:", string(result))

	case "owner":
		if len(os.Args) < 3 {
			log.Fatalln("Uso: ./fabric-client owner <tokenId>")
		}
		result, err := htlc.EvaluateTransaction("Owner", os.Args[2])
		check(err)
		fmt.Println("Owner:", string(result))

	case "recipient":
		if len(os.Args) < 3 {
			log.Fatalln("Uso: ./fabric-client recipient <tokenId>")
		}
		result, err := htlc.EvaluateTransaction("Recipient", os.Args[2])
		check(err)
		fmt.Println("Recipient:", string(result))

	// ------------------- NOTARY -------------------
	case "mintNFT":
		if len(os.Args) < 4 {
			log.Fatalln("Uso: ./fabric-client mintNFT <to> <senderInterChain>")
		}
		result, err := notary.SubmitTransaction("MintNFT", os.Args[2], os.Args[3])
		check(err)
		fmt.Println("MintNFT:", string(result))

	case "transferNFTInterChain":
		if len(os.Args) < 5 {
			log.Fatalln("Uso: ./fabric-client transferNFTInterChain <nftContract> <tokenId> <receiver>")
		}
		result, err := notary.SubmitTransaction("TransferNFTInterChain", os.Args[2], os.Args[3], os.Args[4])
		check(err)
		fmt.Println("TransferNFTInterChain:", string(result))

	case "transferFrom":
		if len(os.Args) < 5 {
			log.Fatalln("Uso: ./fabric-client transferFrom <from> <to> <tokenId>")
		}
		result, err := notary.SubmitTransaction("TransferFrom", os.Args[2], os.Args[3], os.Args[4])
		check(err)
		fmt.Println("TransferFrom:", string(result))

	case "safeTransferFrom":
		if len(os.Args) < 5 {
			log.Fatalln("Uso: ./fabric-client safeTransferFrom <from> <to> <tokenId>")
		}
		result, err := notary.SubmitTransaction("SafeTransferFrom", os.Args[2], os.Args[3], os.Args[4])
		check(err)
		fmt.Println("SafeTransferFrom:", string(result))

	case "approve":
		if len(os.Args) < 4 {
			log.Fatalln("Uso: ./fabric-client approve <to> <tokenId>")
		}
		result, err := notary.SubmitTransaction("Approve", os.Args[2], os.Args[3])
		check(err)
		fmt.Println("Approve:", string(result))

	case "setApprovalForAll":
		if len(os.Args) < 4 {
			log.Fatalln("Uso: ./fabric-client setApprovalForAll <operator> <true|false>")
		}
		result, err := notary.SubmitTransaction("SetApprovalForAll", os.Args[2], os.Args[3])
		check(err)
		fmt.Println("SetApprovalForAll:", string(result))

	case "getNftTransfers":
		result, err := notary.EvaluateTransaction("GetNftTransfers")
		check(err)
		fmt.Println("NFT Transfers:", string(result))

	case "getTransferById":
		if len(os.Args) < 3 {
			log.Fatalln("Uso: ./fabric-client getTransferById <transferId>")
		}
		result, err := notary.EvaluateTransaction("GetTransferById", os.Args[2])
		check(err)
		fmt.Println("TransferById:", string(result))

	case "getNotaryOwner":
		result, err := notary.EvaluateTransaction("GetNotaryOwner")
		check(err)
		fmt.Println("Notary Owner:", string(result))

	case "transferOwnership":
		if len(os.Args) < 3 {
			log.Fatalln("Uso: ./fabric-client transferOwnership <newOwner>")
		}
		result, err := notary.SubmitTransaction("TransferOwnership", os.Args[2])
		check(err)
		fmt.Println("Ownership transferred:", string(result))

	/*
		case "renounceOwnership":
			// TODO: implementar corretamente
	*/
	default:
		fmt.Println("Operação não reconhecida:", operacao)
	}
}

// ---------------- AUX -------------------

func check(err error) {
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}
}

func newGrpcConnection() *grpc.ClientConn {
	certificate, err := os.ReadFile(tlsCertPath)
	check(err)
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	check(err)
	return connection
}

func newIdentity() *identity.X509Identity {
	certPEM, err := os.ReadFile(certPath)
	check(err)

	block, _ := pem.Decode(certPEM)
	if block == nil {
		log.Fatalf("failed to decode PEM certificate from %s", certPath)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	check(err)

	id, err := identity.NewX509Identity(mspID, cert)
	check(err)

	return id
}

func newSign() identity.Sign {
	files, err := os.ReadDir(keyPath)
	check(err)
	if len(files) == 0 {
		log.Fatalf("Nenhuma chave encontrada em %s", keyPath)
	}
	keyPEM, err := os.ReadFile(filepath.Join(keyPath, files[0].Name()))
	check(err)
	privateKey, err := identity.PrivateKeyFromPEM(keyPEM)
	check(err)
	sign, err := identity.NewPrivateKeySign(privateKey)
	check(err)
	return sign
}
