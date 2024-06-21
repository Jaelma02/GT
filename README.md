# HLF-PET

Este projeto, chamado HLF-PET, demonstra como interagir com uma rede Hyperledger Fabric usando um cliente escrito em Go. O código inclui operações básicas como inicializar o ledger, criar ativos, transferir ativos e consultar o ledger.

## Pré-requisitos

1. Go 1.16+ instalado.
2. Hyperledger Fabric test network configurada e em execução.
3. Certificados e chaves configurados no diretório correto.

## Estrutura do Projeto

```plaintext
.
├── client.go
├── go.mod
├── go.sum
└── README.md
```
## Instalação

1. Clone este repositório:

        git clone https://github.com/seu-usuario/hlf-pet.git
        cd hlf-pet

2. Baixe as dependências do Go:
        go mod tidy

## Configuração

Certifique-se de que os caminhos para os certificados e chaves TLS estão corretos no arquivo client.go:

    const (
            mspID        = "Org1MSP"
            cryptoPath   = "../../test-network/organizations/peerOrganizations/org1.example.com"
            certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts"
            keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore"
            tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
            peerEndpoint = "dns:///localhost:7051"
            gatewayPeer  = "peer0.org1.example.com"
    )

## Compilação e Execução
Para compilar e executar o código, use os seguintes comandos:

    go build -o fabric-client client.go
    ./fabric-client -action=<ação> [opções]

#### Ações Disponíveis:

initLedger: Inicializa o ledger com um conjunto de dados de ativos.

    ./fabric-client -action=initLedger

transferAsset: Transfere a propriedade de um ativo.

    ./fabric-client -action=transferAsset -id=<AssetID> -newOwner=<NovoProprietário>
 
createAsset: Cria um novo ativo no ledger.
    ./fabric-client -action=createAsset -num=<Número>

createAssetEndorse: Cria um novo ativo mostrando todas as etapas da criação.

    ./fabric-client -action=createAssetEndorse -num=<Número> -benchmark=<Tipo>

getAll: Retorna todos os ativos atuais no ledger.

     ./fabric-client -action=getAll

getByKey: Obtém os detalhes do ativo por ID.

    ./fabric-client -action=getByKey -key=<Chave>

updateAsset: Atualiza um ativo que não existe.

    ./fabric-client -action=updateAsset -key=<Chave>

## Exemplo de Uso

Para inicializar o ledger:

        ./fabric-client -action=initLedger

Para criar um novo ativo:

    ./fabric-client -action=createAsset -num=1

Para transferir um ativo:

    ./fabric-client -action=transferAsset -id=asset1 -newOwner=JohnDoe
