// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main
//package htlc
//package main



import (
	
	 //"fmt"           // fmt.Println
    //"log"           // log.Fatalf
    
    //"os"            // os.Args, os.Exit
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	// "github.com/ethereum/go-ethereum/crypto"             // crypto.HexToECDSA
    //"github.com/ethereum/go-ethereum/ethclient"          // ethclient.Dial
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)
/*func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso:")
		fmt.Println("  htlc <PRIVATE_KEY_HEX> deploy <recipient> <token> <tokenId> <secret>")
		fmt.Println("  htlc <PRIVATE_KEY_HEX> refund")
		fmt.Println("  htlc <PRIVATE_KEY_HEX> withdraw <secret>")
		os.Exit(1)
	}

	privateKeyHex := os.Args[1]
	action := os.Args[2]

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Falha ao conectar Ethereum: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Erro ao carregar chave privada: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	if err != nil {
		log.Fatalf("Erro ao criar transactor: %v", err)
	}

	switch action {
	case "deploy":
		if len(os.Args) < 7 {
			log.Fatalf("Uso: htlc <PRIVATE_KEY_HEX> deploy <recipient> <token> <tokenId> <secret>")
		}
		recipient := common.HexToAddress(os.Args[3])
		token := common.HexToAddress(os.Args[4])
		tokenId, _ := new(big.Int).SetString(os.Args[5], 10)
		secret := os.Args[6]

		addr, tx, _, err := DeployHtlc(auth, client, recipient, token, tokenId, secret)
		if err != nil {
			log.Fatalf("Erro ao deploy HTLC: %v", err)
		}
		fmt.Println("HTLC implantado em:", addr.Hex())
		fmt.Println("Tx hash:", tx.Hash().Hex())

	case "refund":
		if len(os.Args) < 4 {
			log.Fatalf("Uso: htlc <PRIVATE_KEY_HEX> refund <contractAddress>")
		}
		contractAddr := common.HexToAddress(os.Args[3])
		instance, err := NewHtlc(contractAddr, client)
		if err != nil {
			log.Fatalf("Erro ao conectar contrato: %v", err)
		}
		tx, err := instance.Refund(auth)
		if err != nil {
			log.Fatalf("Erro no refund: %v", err)
		}
		fmt.Println("Refund enviado. Tx:", tx.Hash().Hex())

	case "withdraw":
		if len(os.Args) < 5 {
			log.Fatalf("Uso: htlc <PRIVATE_KEY_HEX> withdraw <contractAddress> <secret>")
		}
		contractAddr := common.HexToAddress(os.Args[3])
		secret := os.Args[4]
		instance, err := NewHtlc(contractAddr, client)
		if err != nil {
			log.Fatalf("Erro ao conectar contrato: %v", err)
		}
		tx, err := instance.Withdraw(auth, secret)
		if err != nil {
			log.Fatalf("Erro no withdraw: %v", err)
		}
		fmt.Println("Withdraw enviado. Tx:", tx.Hash().Hex())

	default:
		fmt.Println("Ação não reconhecida:", action)
	}
}*/

// HtlcMetaData contains all meta data concerning the Htlc contract.
var HtlcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_secret\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Funded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secret\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_secret\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052612710600155348015610015575f5ffd5b50604051611e30380380611e308339818101604052810190610037919061031e565b8360045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503360055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816006819055508260075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806002908161010d91906105a5565b508060405160200161011f91906106ae565b60405160208183030381529060405280519060200120600381905550505050506106c4565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61017e82610155565b9050919050565b61018e81610174565b8114610198575f5ffd5b50565b5f815190506101a981610185565b92915050565b5f819050919050565b6101c1816101af565b81146101cb575f5ffd5b50565b5f815190506101dc816101b8565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610230826101ea565b810181811067ffffffffffffffff8211171561024f5761024e6101fa565b5b80604052505050565b5f610261610144565b905061026d8282610227565b919050565b5f67ffffffffffffffff82111561028c5761028b6101fa565b5b610295826101ea565b9050602081019050919050565b8281835e5f83830152505050565b5f6102c26102bd84610272565b610258565b9050828152602081018484840111156102de576102dd6101e6565b5b6102e98482856102a2565b509392505050565b5f82601f830112610305576103046101e2565b5b81516103158482602086016102b0565b91505092915050565b5f5f5f5f608085870312156103365761033561014d565b5b5f6103438782880161019b565b94505060206103548782880161019b565b9350506040610365878288016101ce565b925050606085015167ffffffffffffffff81111561038657610385610151565b5b610392878288016102f1565b91505092959194509250565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103ec57607f821691505b6020821081036103ff576103fe6103a8565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104617fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610426565b61046b8683610426565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6104a66104a161049c846101af565b610483565b6101af565b9050919050565b5f819050919050565b6104bf8361048c565b6104d36104cb826104ad565b848454610432565b825550505050565b5f5f905090565b6104ea6104db565b6104f58184846104b6565b505050565b5b818110156105185761050d5f826104e2565b6001810190506104fb565b5050565b601f82111561055d5761052e81610405565b61053784610417565b81016020851015610546578190505b61055a61055285610417565b8301826104fa565b50505b505050565b5f82821c905092915050565b5f61057d5f1984600802610562565b1980831691505092915050565b5f610595838361056e565b9150826002028217905092915050565b6105ae8261039e565b67ffffffffffffffff8111156105c7576105c66101fa565b5b6105d182546103d5565b6105dc82828561051c565b5f60209050601f83116001811461060d575f84156105fb578287015190505b610605858261058a565b86555061066c565b601f19841661061b86610405565b5f5b828110156106425784890151825560018201915060208501945060208101905061061d565b8683101561065f578489015161065b601f89168261056e565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f6106888261039e565b6106928185610674565b93506106a28185602086016102a2565b80840191505092915050565b5f6106b9828461067e565b915081905092915050565b61175f806106d15f395ff3fe608060405234801561000f575f5ffd5b50600436106100b2575f3560e01c806366d003ac1161006f57806366d003ac1461016657806378e97925146101845780638da5cb5b146101a2578063b60d4288146101c0578063d1efd30d146101ca578063fc0c546a146101e8576100b2565b806309bd5a60146100b65780630d668087146100d4578063150b7a02146100f257806317d70f7c1461012257806331fb67c214610140578063590e1ae31461015c575b5f5ffd5b6100be610206565b6040516100cb9190610b52565b60405180910390f35b6100dc61020c565b6040516100e99190610b83565b60405180910390f35b61010c60048036038101906101079190610d6d565b610212565b6040516101199190610e27565b60405180910390f35b61012a610225565b6040516101379190610b83565b60405180910390f35b61015a60048036038101906101559190610ede565b61022b565b005b61016461047a565b005b61016e6106a3565b60405161017b9190610f34565b60405180910390f35b61018c6106c8565b6040516101999190610b83565b60405180910390f35b6101aa6106cd565b6040516101b79190610f34565b60405180910390f35b6101c86106f2565b005b6101d2610a89565b6040516101df9190610fad565b60405180910390f35b6101f0610b15565b6040516101fd9190611028565b60405180910390f35b60035481565b60015481565b5f63150b7a0260e01b9050949350505050565b60065481565b6001545f5461023a919061106e565b42111561024e5761024961047a565b610477565b6003548160405160200161026291906110db565b60405160208183030381529060405280519060200120146102b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102af9061113b565b60405180910390fd5b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610347576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033e906111c9565b60405180910390fd5b806002908161035691906113db565b5060075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e3060045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166006546040518463ffffffff1660e01b81526004016103d8939291906114aa565b5f604051808303815f87803b1580156103ef575f5ffd5b505af1158015610401573d5f5f3e3d5ffd5b5050505060045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d560065460405161046e9190610b83565b60405180910390a25b50565b3073ffffffffffffffffffffffffffffffffffffffff1660075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e6006546040518263ffffffff1660e01b81526004016104ed9190610b83565b602060405180830381865afa158015610508573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061052c91906114f3565b73ffffffffffffffffffffffffffffffffffffffff1614610582576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105799061158e565b60405180910390fd5b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e3060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166006546040518463ffffffff1660e01b8152600401610603939291906114aa565b5f604051808303815f87803b15801561061a575f5ffd5b505af115801561062c573d5f5f3e3d5ffd5b5050505060055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d06516006546040516106999190610b83565b60405180910390a2565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5481565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff1660075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e6006546040518263ffffffff1660e01b81526004016107659190610b83565b602060405180830381865afa158015610780573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a491906114f3565b73ffffffffffffffffffffffffffffffffffffffff16146107fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f1906115f6565b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff1660075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663081812fc6006546040518263ffffffff1660e01b815260040161086d9190610b83565b602060405180830381865afa158015610888573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ac91906114f3565b73ffffffffffffffffffffffffffffffffffffffff161480610965575060075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e985e9c533306040518363ffffffff1660e01b8152600401610925929190611614565b602060405180830381865afa158015610940573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109649190611670565b5b6109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099b9061170b565b60405180910390fd5b425f8190555060075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e33306006546040518463ffffffff1660e01b8152600401610a0a939291906114aa565b5f604051808303815f87803b158015610a21575f5ffd5b505af1158015610a33573d5f5f3e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff167f5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a524600654604051610a7f9190610b83565b60405180910390a2565b60028054610a9690611214565b80601f0160208091040260200160405190810160405280929190818152602001828054610ac290611214565b8015610b0d5780601f10610ae457610100808354040283529160200191610b0d565b820191905f5260205f20905b815481529060010190602001808311610af057829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f819050919050565b610b4c81610b3a565b82525050565b5f602082019050610b655f830184610b43565b92915050565b5f819050919050565b610b7d81610b6b565b82525050565b5f602082019050610b965f830184610b74565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610bd682610bad565b9050919050565b610be681610bcc565b8114610bf0575f5ffd5b50565b5f81359050610c0181610bdd565b92915050565b610c1081610b6b565b8114610c1a575f5ffd5b50565b5f81359050610c2b81610c07565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c7f82610c39565b810181811067ffffffffffffffff82111715610c9e57610c9d610c49565b5b80604052505050565b5f610cb0610b9c565b9050610cbc8282610c76565b919050565b5f67ffffffffffffffff821115610cdb57610cda610c49565b5b610ce482610c39565b9050602081019050919050565b828183375f83830152505050565b5f610d11610d0c84610cc1565b610ca7565b905082815260208101848484011115610d2d57610d2c610c35565b5b610d38848285610cf1565b509392505050565b5f82601f830112610d5457610d53610c31565b5b8135610d64848260208601610cff565b91505092915050565b5f5f5f5f60808587031215610d8557610d84610ba5565b5b5f610d9287828801610bf3565b9450506020610da387828801610bf3565b9350506040610db487828801610c1d565b925050606085013567ffffffffffffffff811115610dd557610dd4610ba9565b5b610de187828801610d40565b91505092959194509250565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610e2181610ded565b82525050565b5f602082019050610e3a5f830184610e18565b92915050565b5f67ffffffffffffffff821115610e5a57610e59610c49565b5b610e6382610c39565b9050602081019050919050565b5f610e82610e7d84610e40565b610ca7565b905082815260208101848484011115610e9e57610e9d610c35565b5b610ea9848285610cf1565b509392505050565b5f82601f830112610ec557610ec4610c31565b5b8135610ed5848260208601610e70565b91505092915050565b5f60208284031215610ef357610ef2610ba5565b5b5f82013567ffffffffffffffff811115610f1057610f0f610ba9565b5b610f1c84828501610eb1565b91505092915050565b610f2e81610bcc565b82525050565b5f602082019050610f475f830184610f25565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610f7f82610f4d565b610f898185610f57565b9350610f99818560208601610f67565b610fa281610c39565b840191505092915050565b5f6020820190508181035f830152610fc58184610f75565b905092915050565b5f819050919050565b5f610ff0610feb610fe684610bad565b610fcd565b610bad565b9050919050565b5f61100182610fd6565b9050919050565b5f61101282610ff7565b9050919050565b61102281611008565b82525050565b5f60208201905061103b5f830184611019565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61107882610b6b565b915061108383610b6b565b925082820190508082111561109b5761109a611041565b5b92915050565b5f81905092915050565b5f6110b582610f4d565b6110bf81856110a1565b93506110cf818560208601610f67565b80840191505092915050565b5f6110e682846110ab565b915081905092915050565b7f48544c433732313a2057726f6e672073656372657400000000000000000000005f82015250565b5f611125601583610f57565b9150611130826110f1565b602082019050919050565b5f6020820190508181035f83015261115281611119565b9050919050565b7f48544c433732313a204f6e6c7920726563697069656e742063616e20776974685f8201527f6472617700000000000000000000000000000000000000000000000000000000602082015250565b5f6111b3602483610f57565b91506111be82611159565b604082019050919050565b5f6020820190508181035f8301526111e0816111a7565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061122b57607f821691505b60208210810361123e5761123d6111e7565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026112a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611265565b6112aa8683611265565b95508019841693508086168417925050509392505050565b5f6112dc6112d76112d284610b6b565b610fcd565b610b6b565b9050919050565b5f819050919050565b6112f5836112c2565b611309611301826112e3565b848454611271565b825550505050565b5f5f905090565b611320611311565b61132b8184846112ec565b505050565b5b8181101561134e576113435f82611318565b600181019050611331565b5050565b601f8211156113935761136481611244565b61136d84611256565b8101602085101561137c578190505b61139061138885611256565b830182611330565b50505b505050565b5f82821c905092915050565b5f6113b35f1984600802611398565b1980831691505092915050565b5f6113cb83836113a4565b9150826002028217905092915050565b6113e482610f4d565b67ffffffffffffffff8111156113fd576113fc610c49565b5b6114078254611214565b611412828285611352565b5f60209050601f831160018114611443575f8415611431578287015190505b61143b85826113c0565b8655506114a2565b601f19841661145186611244565b5f5b8281101561147857848901518255600182019150602085019450602081019050611453565b868310156114955784890151611491601f8916826113a4565b8355505b6001600288020188555050505b505050505050565b5f6060820190506114bd5f830186610f25565b6114ca6020830185610f25565b6114d76040830184610b74565b949350505050565b5f815190506114ed81610bdd565b92915050565b5f6020828403121561150857611507610ba5565b5b5f611515848285016114df565b91505092915050565b7f48544c433732313a20436f6e747261637420646f6573206e6f74206f776e20745f8201527f686520746f6b656e000000000000000000000000000000000000000000000000602082015250565b5f611578602883610f57565b91506115838261151e565b604082019050919050565b5f6020820190508181035f8301526115a58161156c565b9050919050565b7f48544c433732313a204f6e6c7920746865206f776e65722063616e2066756e645f82015250565b5f6115e0602083610f57565b91506115eb826115ac565b602082019050919050565b5f6020820190508181035f83015261160d816115d4565b9050919050565b5f6040820190506116275f830185610f25565b6116346020830184610f25565b9392505050565b5f8115159050919050565b61164f8161163b565b8114611659575f5ffd5b50565b5f8151905061166a81611646565b92915050565b5f6020828403121561168557611684610ba5565b5b5f6116928482850161165c565b91505092915050565b7f48544c433732313a20436f6e7472616374206973206e6f7420617070726f76655f8201527f6420746f207472616e73666572207468697320746f6b656e0000000000000000602082015250565b5f6116f5603883610f57565b91506117008261169b565b604082019050919050565b5f6020820190508181035f830152611722816116e9565b905091905056fea2646970667358221220b88163aa8cacd2e087dc1c6341807498fbe3cd3e194918a6732c1eb4cc94d34e64736f6c634300081e0033",
}

// HtlcABI is the input ABI used to generate the binding from.
// Deprecated: Use HtlcMetaData.ABI instead.
var HtlcABI = HtlcMetaData.ABI

// HtlcBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HtlcMetaData.Bin instead.
var HtlcBin = HtlcMetaData.Bin

// DeployHtlc deploys a new Ethereum contract, binding an instance of Htlc to it.
func DeployHtlc(auth *bind.TransactOpts, backend bind.ContractBackend, _recipient common.Address, _token common.Address, _tokenId *big.Int, _secret string) (common.Address, *types.Transaction, *Htlc, error) {
	parsed, err := HtlcMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HtlcBin), backend, _recipient, _token, _tokenId, _secret)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Htlc{HtlcCaller: HtlcCaller{contract: contract}, HtlcTransactor: HtlcTransactor{contract: contract}, HtlcFilterer: HtlcFilterer{contract: contract}}, nil
}

// Htlc is an auto generated Go binding around an Ethereum contract.
type Htlc struct {
	HtlcCaller     // Read-only binding to the contract
	HtlcTransactor // Write-only binding to the contract
	HtlcFilterer   // Log filterer for contract events
}

// HtlcCaller is an auto generated read-only Go binding around an Ethereum contract.
type HtlcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HtlcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HtlcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HtlcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HtlcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HtlcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HtlcSession struct {
	Contract     *Htlc             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HtlcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HtlcCallerSession struct {
	Contract *HtlcCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HtlcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HtlcTransactorSession struct {
	Contract     *HtlcTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HtlcRaw is an auto generated low-level Go binding around an Ethereum contract.
type HtlcRaw struct {
	Contract *Htlc // Generic contract binding to access the raw methods on
}

// HtlcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HtlcCallerRaw struct {
	Contract *HtlcCaller // Generic read-only contract binding to access the raw methods on
}

// HtlcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HtlcTransactorRaw struct {
	Contract *HtlcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHtlc creates a new instance of Htlc, bound to a specific deployed contract.
func NewHtlc(address common.Address, backend bind.ContractBackend) (*Htlc, error) {
	contract, err := bindHtlc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Htlc{HtlcCaller: HtlcCaller{contract: contract}, HtlcTransactor: HtlcTransactor{contract: contract}, HtlcFilterer: HtlcFilterer{contract: contract}}, nil
}

// NewHtlcCaller creates a new read-only instance of Htlc, bound to a specific deployed contract.
func NewHtlcCaller(address common.Address, caller bind.ContractCaller) (*HtlcCaller, error) {
	contract, err := bindHtlc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HtlcCaller{contract: contract}, nil
}

// NewHtlcTransactor creates a new write-only instance of Htlc, bound to a specific deployed contract.
func NewHtlcTransactor(address common.Address, transactor bind.ContractTransactor) (*HtlcTransactor, error) {
	contract, err := bindHtlc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HtlcTransactor{contract: contract}, nil
}

// NewHtlcFilterer creates a new log filterer instance of Htlc, bound to a specific deployed contract.
func NewHtlcFilterer(address common.Address, filterer bind.ContractFilterer) (*HtlcFilterer, error) {
	contract, err := bindHtlc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HtlcFilterer{contract: contract}, nil
}

// bindHtlc binds a generic wrapper to an already deployed contract.
func bindHtlc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HtlcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Htlc *HtlcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Htlc.Contract.HtlcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Htlc *HtlcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Htlc.Contract.HtlcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Htlc *HtlcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Htlc.Contract.HtlcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Htlc *HtlcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Htlc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Htlc *HtlcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Htlc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Htlc *HtlcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Htlc.Contract.contract.Transact(opts, method, params...)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_Htlc *HtlcCaller) Hash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "hash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_Htlc *HtlcSession) Hash() ([32]byte, error) {
	return _Htlc.Contract.Hash(&_Htlc.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_Htlc *HtlcCallerSession) Hash() ([32]byte, error) {
	return _Htlc.Contract.Hash(&_Htlc.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_Htlc *HtlcCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_Htlc *HtlcSession) LockTime() (*big.Int, error) {
	return _Htlc.Contract.LockTime(&_Htlc.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_Htlc *HtlcCallerSession) LockTime() (*big.Int, error) {
	return _Htlc.Contract.LockTime(&_Htlc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Htlc *HtlcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Htlc *HtlcSession) Owner() (common.Address, error) {
	return _Htlc.Contract.Owner(&_Htlc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Htlc *HtlcCallerSession) Owner() (common.Address, error) {
	return _Htlc.Contract.Owner(&_Htlc.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_Htlc *HtlcCaller) Recipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "recipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_Htlc *HtlcSession) Recipient() (common.Address, error) {
	return _Htlc.Contract.Recipient(&_Htlc.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_Htlc *HtlcCallerSession) Recipient() (common.Address, error) {
	return _Htlc.Contract.Recipient(&_Htlc.CallOpts)
}

// Secret is a free data retrieval call binding the contract method 0xd1efd30d.
//
// Solidity: function secret() view returns(string)
func (_Htlc *HtlcCaller) Secret(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "secret")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Secret is a free data retrieval call binding the contract method 0xd1efd30d.
//
// Solidity: function secret() view returns(string)
func (_Htlc *HtlcSession) Secret() (string, error) {
	return _Htlc.Contract.Secret(&_Htlc.CallOpts)
}

// Secret is a free data retrieval call binding the contract method 0xd1efd30d.
//
// Solidity: function secret() view returns(string)
func (_Htlc *HtlcCallerSession) Secret() (string, error) {
	return _Htlc.Contract.Secret(&_Htlc.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Htlc *HtlcCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Htlc *HtlcSession) StartTime() (*big.Int, error) {
	return _Htlc.Contract.StartTime(&_Htlc.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Htlc *HtlcCallerSession) StartTime() (*big.Int, error) {
	return _Htlc.Contract.StartTime(&_Htlc.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Htlc *HtlcCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Htlc *HtlcSession) Token() (common.Address, error) {
	return _Htlc.Contract.Token(&_Htlc.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Htlc *HtlcCallerSession) Token() (common.Address, error) {
	return _Htlc.Contract.Token(&_Htlc.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Htlc *HtlcCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Htlc.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Htlc *HtlcSession) TokenId() (*big.Int, error) {
	return _Htlc.Contract.TokenId(&_Htlc.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Htlc *HtlcCallerSession) TokenId() (*big.Int, error) {
	return _Htlc.Contract.TokenId(&_Htlc.CallOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Htlc *HtlcTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Htlc.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Htlc *HtlcSession) Fund() (*types.Transaction, error) {
	return _Htlc.Contract.Fund(&_Htlc.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() returns()
func (_Htlc *HtlcTransactorSession) Fund() (*types.Transaction, error) {
	return _Htlc.Contract.Fund(&_Htlc.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Htlc *HtlcTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Htlc.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Htlc *HtlcSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Htlc.Contract.OnERC721Received(&_Htlc.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Htlc *HtlcTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Htlc.Contract.OnERC721Received(&_Htlc.TransactOpts, arg0, arg1, arg2, arg3)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Htlc *HtlcTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Htlc.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Htlc *HtlcSession) Refund() (*types.Transaction, error) {
	return _Htlc.Contract.Refund(&_Htlc.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Htlc *HtlcTransactorSession) Refund() (*types.Transaction, error) {
	return _Htlc.Contract.Refund(&_Htlc.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string _secret) returns()
func (_Htlc *HtlcTransactor) Withdraw(opts *bind.TransactOpts, _secret string) (*types.Transaction, error) {
	return _Htlc.contract.Transact(opts, "withdraw", _secret)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string _secret) returns()
func (_Htlc *HtlcSession) Withdraw(_secret string) (*types.Transaction, error) {
	return _Htlc.Contract.Withdraw(&_Htlc.TransactOpts, _secret)
}

// Withdraw is a paid mutator transaction binding the contract method 0x31fb67c2.
//
// Solidity: function withdraw(string _secret) returns()
func (_Htlc *HtlcTransactorSession) Withdraw(_secret string) (*types.Transaction, error) {
	return _Htlc.Contract.Withdraw(&_Htlc.TransactOpts, _secret)
}

// HtlcFundedIterator is returned from FilterFunded and is used to iterate over the raw logs and unpacked data for Funded events raised by the Htlc contract.
type HtlcFundedIterator struct {
	Event *HtlcFunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HtlcFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HtlcFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HtlcFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HtlcFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HtlcFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HtlcFunded represents a Funded event raised by the Htlc contract.
type HtlcFunded struct {
	Funder  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFunded is a free log retrieval operation binding the contract event 0x5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a524.
//
// Solidity: event Funded(address indexed funder, uint256 tokenId)
func (_Htlc *HtlcFilterer) FilterFunded(opts *bind.FilterOpts, funder []common.Address) (*HtlcFundedIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _Htlc.contract.FilterLogs(opts, "Funded", funderRule)
	if err != nil {
		return nil, err
	}
	return &HtlcFundedIterator{contract: _Htlc.contract, event: "Funded", logs: logs, sub: sub}, nil
}

// WatchFunded is a free log subscription operation binding the contract event 0x5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a524.
//
// Solidity: event Funded(address indexed funder, uint256 tokenId)
func (_Htlc *HtlcFilterer) WatchFunded(opts *bind.WatchOpts, sink chan<- *HtlcFunded, funder []common.Address) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _Htlc.contract.WatchLogs(opts, "Funded", funderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HtlcFunded)
				if err := _Htlc.contract.UnpackLog(event, "Funded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFunded is a log parse operation binding the contract event 0x5af8184bef8e4b45eb9f6ed7734d04da38ced226495548f46e0c8ff8d7d9a524.
//
// Solidity: event Funded(address indexed funder, uint256 tokenId)
func (_Htlc *HtlcFilterer) ParseFunded(log types.Log) (*HtlcFunded, error) {
	event := new(HtlcFunded)
	if err := _Htlc.contract.UnpackLog(event, "Funded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HtlcRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Htlc contract.
type HtlcRefundedIterator struct {
	Event *HtlcRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HtlcRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HtlcRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HtlcRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HtlcRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HtlcRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HtlcRefunded represents a Refunded event raised by the Htlc contract.
type HtlcRefunded struct {
	Owner   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed owner, uint256 tokenId)
func (_Htlc *HtlcFilterer) FilterRefunded(opts *bind.FilterOpts, owner []common.Address) (*HtlcRefundedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Htlc.contract.FilterLogs(opts, "Refunded", ownerRule)
	if err != nil {
		return nil, err
	}
	return &HtlcRefundedIterator{contract: _Htlc.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed owner, uint256 tokenId)
func (_Htlc *HtlcFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *HtlcRefunded, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Htlc.contract.WatchLogs(opts, "Refunded", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HtlcRefunded)
				if err := _Htlc.contract.UnpackLog(event, "Refunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefunded is a log parse operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed owner, uint256 tokenId)
func (_Htlc *HtlcFilterer) ParseRefunded(log types.Log) (*HtlcRefunded, error) {
	event := new(HtlcRefunded)
	if err := _Htlc.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HtlcWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Htlc contract.
type HtlcWithdrawnIterator struct {
	Event *HtlcWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HtlcWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HtlcWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HtlcWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HtlcWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HtlcWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HtlcWithdrawn represents a Withdrawn event raised by the Htlc contract.
type HtlcWithdrawn struct {
	Recipient common.Address
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed recipient, uint256 tokenId)
func (_Htlc *HtlcFilterer) FilterWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*HtlcWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Htlc.contract.FilterLogs(opts, "Withdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &HtlcWithdrawnIterator{contract: _Htlc.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed recipient, uint256 tokenId)
func (_Htlc *HtlcFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *HtlcWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Htlc.contract.WatchLogs(opts, "Withdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HtlcWithdrawn)
				if err := _Htlc.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed recipient, uint256 tokenId)
func (_Htlc *HtlcFilterer) ParseWithdrawn(log types.Log) (*HtlcWithdrawn, error) {
	event := new(HtlcWithdrawn)
	if err := _Htlc.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
