package main

import (
    "bytes"
    "crypto/x509"
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "path"
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
    cryptoPath   = "../../test-network/organizations/peerOrganizations/org1.example.com"
    certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts"
    keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore"
    tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
    peerEndpoint = "localhost:7051" // Corrigido para o formato correto para DNS
    gatewayPeer  = "peer0.org1.example.com"
)

var now = time.Now()
var assetID = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)

func main() {
    // Parse command-line arguments
    action := flag.String("action", "", "Ação a ser realizada: initLedger, transferAsset, createAsset, createAssetEndorse, getAll, getByKey, updateAsset")
    id := flag.String("id", "", "ID do ativo")
    newOwner := flag.String("newOwner", "", "Novo proprietário para o ativo")
    num := flag.Int("num", 0, "Número para createAsset")
    key := flag.String("key", "", "Chave para getByKey ou updateAsset")
    benchmark := flag.String("benchmark", "", "Tipo de benchmark para createAssetEndorse")
    flag.Parse()

    // A conexão do cliente gRPC deve ser compartilhada por todas as conexões do Gateway neste endpoint
    clientConnection := newGrpcConnection()
    defer clientConnection.Close()

    id := newIdentity()
    sign := newSign()

    // Crie uma conexão Gateway para uma identidade de cliente específica
    gw, err := client.Connect(
        id,
        client.WithSign(sign),
        client.WithClientConnection(clientConnection),
        // Padrões de tempo limite para chamadas gRPC diferentes
        client.WithEvaluateTimeout(5*time.Second),
        client.WithEndorseTimeout(15*time.Second),
        client.WithSubmitTimeout(5*time.Second),
        client.WithCommitStatusTimeout(1*time.Minute),
    )
    if err != nil {
        panic(err)
    }
    defer gw.Close()

    // Substitua os valores padrão para nome de chaincode e nome do canal, pois podem ser diferentes em contextos de teste.
    chaincodeName := "basic"
    if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
        chaincodeName = ccname
    }

    channelName := "mychannel"
    if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
        channelName = cname
    }

    network := gw.GetNetwork(channelName)
    contract := network.GetContract(chaincodeName)

    // Execute a ação com base no argumento da linha de comando
    switch *action {
    case "initLedger":
        initLedger(contract)
    case "transferAsset":
        if *id == "" || *newOwner == "" {
            fmt.Println("Parâmetros obrigatórios ausentes para transferAsset")
            return
        }
        transferAssetAsync(contract, *id, *newOwner)
    case "createAsset":
        createAsset(contract, *num)
    /* 
    case "createAssetEndorse":
        if *benchmark == "B" {
            createAssetEndorseBenchmarks(contract, *num)
        } else {
            createAssetEndorse(contract, *num)
        }
    */
    case "getAll":
        getAllAssets(contract)
    case "getByKey":
        if *key == "" {
            fmt.Println("Parâmetro obrigatório 'key' ausente para getByKey")
            return
        }
        readAssetByID(contract, *key)
    case "updateAsset":
        if *key == "" {
            fmt.Println("Parâmetro obrigatório 'key' ausente para updateAsset")
            return
        }
        updateNonExistentAsset(contract, *key)
    default:
        fmt.Printf("Argumento Inválido: %s\n", *action)
    }
}

// newGrpcConnection cria uma conexão gRPC para o servidor Gateway.
func newGrpcConnection() *grpc.ClientConn {
    certificatePEM, err := os.ReadFile(tlsCertPath)
    if err != nil {
        panic(fmt.Errorf("falha ao ler arquivo de certificado TLS: %w", err))
    }

    certificate, err := identity.CertificateFromPEM(certificatePEM)
    if err != nil {
        panic(err)
    }

    certPool := x509.NewCertPool()
    certPool.AddCert(certificate)
    transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

    connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
    if err != nil {
        panic(fmt.Errorf("falha ao criar conexão gRPC: %w", err))
    }

    return connection
}

// newIdentity cria uma identidade de cliente para esta conexão Gateway usando um certificado X.509.
func newIdentity() *identity.X509Identity {
    certificatePEM, err := readFirstFile(certPath)
    if err != nil {
        panic(fmt.Errorf("falha ao ler arquivo de certificado: %w", err))
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

// newSign cria uma função que gera uma assinatura digital a partir de um digest de mensagem usando uma chave privada.
func newSign() identity.Sign {
    privateKeyPEM, err := readFirstFile(keyPath)
    if err != nil {
        panic(fmt.Errorf("falha ao ler arquivo de chave privada: %w", err))
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
    defer dir.Close()

    fileNames, err := dir.Readdirnames(1)
    if err != nil {
        return nil, err
    }

    return os.ReadFile(path.Join(dirPath, fileNames[0]))
}

func initLedger(contract *client.Contract) {
    fmt.Printf("\n--> Submeter Transação: InitLedger, a função cria o conjunto inicial de ativos no ledger \n")

    _, err := contract.SubmitTransaction("InitLedger")
    if err != nil {
        panic(fmt.Errorf("falha ao submeter transação: %w", err))
    }

    fmt.Printf("*** Transação confirmada com sucesso\n")
}

func getAllAssets(contract *client.Contract) {
    fmt.Println("\n--> Avaliar Transação: GetAllAssets, a função retorna todos os ativos atuais no ledger")

    evaluateResult, err := contract.EvaluateTransaction("QueryAllCars")
    if err != nil {
        panic(fmt.Errorf("falha ao avaliar transação: %w", err))
    }
    result := formatJSON(evaluateResult)

    fmt.Printf("*** Resultado:%s\n", result)
}

func createAsset(contract *client.Contract, num int) {
    fmt.Printf("\n--> Submeter Transação: CreateAsset, cria um novo ativo com ID, Cor, Tamanho, Proprietário e ValorAvaliado\n")

    assetID := fmt.Sprintf("asset%d", num)
    _, err := contract.SubmitTransaction("CreateCar", assetID, "yellow", "5", "Tom", "1300")
    if err != nil {
        panic(fmt.Errorf("falha ao submeter transação: %w", err))
    }

    fmt.Printf("*** Transação confirmada com sucesso\n")
}

func readAssetByID(contract *client.Contract, key string) {
    fmt.Printf("\n--> Avaliar Transação: ReadAsset, a função retorna os atributos do ativo\n")

    evaluateResult, err := contract.EvaluateTransaction("QueryCar", key)
    if err != nil {
        panic(fmt.Errorf("falha ao avaliar transação: %w", err))
    }
    result := formatJSON(evaluateResult)

    fmt.Printf("*** Resultado:%s\n", result)
}

func transferAssetAsync(contract *client.Contract, id string, newOwner string) {
    fmt.Printf("\n--> Submeter Transação Assíncrona: TransferAsset, atualiza o proprietário existente do ativo")

    submitResult, commit, err := contract.SubmitAsync("ChangeCarOwner", client.WithArguments(id, newOwner))
    if err != nil {
        panic(fmt.Errorf("falha ao submeter transação assíncrona: %w", err))
    }

    fmt.Printf("\n*** Transação submetida com sucesso para transferência de propriedade de %s para %s. \n", id, newOwner)
    fmt.Println("*** Aguardando o commit da transação.")

    if commitStatus, err := commit.Status(); err != nil {
        panic(fmt.Errorf("falha ao obter status de commit: %w", err))
    } else if !commitStatus.Successful {
        panic(fmt.Errorf("a transação %s falhou em ser confirmada com status: %d", commitStatus.TransactionID, int32(commitStatus.Code)))
    }

    fmt.Printf("*** Transação confirmada com sucesso\n")
}

func updateNonExistentAsset(contract *client.Contract, key string) {
    fmt.Println("\n--> Submeter Transação: UpdateAsset asset70, asset70 não existe e deve retornar um erro")

    _, err := contract.SubmitTransaction("UpdateAsset", key, "blue", "5", "Tomoko", "300")
    if err == nil {
        panic("******** FALHOU em retornar um erro")
    }

    fmt.Println("*** Capturado com sucesso o erro:")

    var endorseErr *client.EndorseError
    var submitErr *client.SubmitError
    var commitStatusErr *client.CommitStatusError
    var commitErr *client.CommitError

    if errors.As(err, &endorseErr) {
        fmt.Printf("Erro de Endosse para a transação %s com status gRPC %v: %s\n", endorseErr.TransactionID, status.Code(endorseErr), endorseErr)
    } else if errors.As(err, &submitErr) {
        fmt.Printf("Erro de Submissão para a transação %s com status gRPC %v: %s\n", submitErr.TransactionID, status.Code(submitErr), submitErr)
    } else if errors.As(err, &commitStatusErr) {
        if errors.Is(err, context.DeadlineExceeded) {
            fmt.Printf("Tempo esgotado aguardando o status de commit da transação %s: %s", commitStatusErr.TransactionID, commitStatusErr)
        } else {
            fmt.Printf("Erro ao obter status de commit para a transação %s com status gRPC %v: %s\n", commitStatusErr.TransactionID, status.Code(commitStatusErr), commitStatusErr)
        }
    } else if errors.As(err, &commitErr) {
        fmt.Printf("A transação %s falhou em ser confirmada com status %d: %s\n", commitErr.TransactionID, int32(commitErr.Code), err)
    } else {
        panic(fmt.Errorf("tipo de erro inesperado %T: %w", err, err))
    }

    statusErr := status.Convert(err)

    details := statusErr.Details()
    if len(details) > 0 {
        fmt.Println("Detalhes do Erro:")

        for _, detail := range details {
            switch detail := detail.(type) {
            case *gateway.ErrorDetail:
                fmt.Printf("- endereço: %s, mspId: %s, mensagem: %s\n", detail.Address, detail.MspId, detail.Message)
            }
        }
    }
}

func formatJSON(data []byte) string {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
        panic(fmt.Errorf("falha ao analisar JSON: %w", err))
    }
    return prettyJSON.String()
}