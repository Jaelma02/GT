// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main
//package notary
//package main

import (
	"errors"
	//"fmt"
	//"log"
	"math/big"
	//"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
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
		fmt.Println("  notary <PRIVATE_KEY_HEX> deploy")
		fmt.Println("  notary <PRIVATE_KEY_HEX> mint <contractAddr> <to> <senderInterChain>")
		fmt.Println("  notary <PRIVATE_KEY_HEX> transferFrom <contractAddr> <from> <to> <tokenId>")
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
		addr, tx, _, err := DeployNotary(auth, client)
		if err != nil {
			log.Fatalf("Erro no deploy Notary: %v", err)
		}
		fmt.Println("Notary implantado em:", addr.Hex())
		fmt.Println("Tx hash:", tx.Hash().Hex())

	case "mint":
		if len(os.Args) < 6 {
			log.Fatalf("Uso: notary <PRIVATE_KEY_HEX> mint <contractAddr> <to> <senderInterChain>")
		}
		contractAddr := common.HexToAddress(os.Args[3])
		to := common.HexToAddress(os.Args[4])
		senderInterChain := common.HexToAddress(os.Args[5])

		instance, err := NewNotary(contractAddr, client)
		if err != nil {
			log.Fatalf("Erro ao conectar contrato: %v", err)
		}
		tx, err := instance.MintNFT(auth, to, senderInterChain)
		if err != nil {
			log.Fatalf("Erro no mintNFT: %v", err)
		}
		fmt.Println("NFT Mintado. Tx:", tx.Hash().Hex())

	case "transferFrom":
		if len(os.Args) < 7 {
			log.Fatalf("Uso: notary <PRIVATE_KEY_HEX> transferFrom <contractAddr> <from> <to> <tokenId>")
		}
		contractAddr := common.HexToAddress(os.Args[3])
		from := common.HexToAddress(os.Args[4])
		to := common.HexToAddress(os.Args[5])
		tokenId, _ := new(big.Int).SetString(os.Args[6], 10)

		instance, err := NewNotary(contractAddr, client)
		if err != nil {
			log.Fatalf("Erro ao conectar contrato: %v", err)
		}
		tx, err := instance.TransferFrom(auth, from, to, tokenId)
		if err != nil {
			log.Fatalf("Erro no transferFrom: %v", err)
		}
		fmt.Println("NFT transferido. Tx:", tx.Hash().Hex())

	default:
		fmt.Println("Ação não reconhecida:", action)
	}
}
*/
// NotaryNFTTransfer is an auto generated low-level Go binding around an user-defined struct.
type NotaryNFTTransfer struct {
	NftContract        common.Address
	Sender             common.Address
	ReceiverOnChain    common.Address
	ReceiverInterChain common.Address
	TokenId            *big.Int
	TransferId         *big.Int
}

// NotaryMetaData contains all meta data concerning the Notary contract.
var NotaryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"NFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiverOnChain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiverInterChain\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferId\",\"type\":\"uint256\"}],\"name\":\"NFTTransferInterChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNftTransfers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNotaryOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transferId\",\"type\":\"uint256\"}],\"name\":\"getTransferById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverOnChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverInterChain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferId\",\"type\":\"uint256\"}],\"internalType\":\"structNotary.NFTTransfer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderInterChain\",\"type\":\"address\"}],\"name\":\"mintNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverOnChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverInterChain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notaryOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiverInterChain\",\"type\":\"address\"}],\"name\":\"transferNFTInterChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50336040518060400160405280600981526020017f4e6f746172794e465400000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4e54590000000000000000000000000000000000000000000000000000000000815250815f908161008b9190610463565b50806001908161009b9190610463565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361010e575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016101059190610571565b60405180910390fd5b61011d8161016360201b60201c565b503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061058a565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806102a157607f821691505b6020821081036102b4576102b361025d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826102db565b61032086836102db565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61036461035f61035a84610338565b610341565b610338565b9050919050565b5f819050919050565b61037d8361034a565b6103916103898261036b565b8484546102e7565b825550505050565b5f5f905090565b6103a8610399565b6103b3818484610374565b505050565b5b818110156103d6576103cb5f826103a0565b6001810190506103b9565b5050565b601f82111561041b576103ec816102ba565b6103f5846102cc565b81016020851015610404578190505b610418610410856102cc565b8301826103b8565b50505b505050565b5f82821c905092915050565b5f61043b5f1984600802610420565b1980831691505092915050565b5f610453838361042c565b9150826002028217905092915050565b61046c82610226565b67ffffffffffffffff81111561048557610484610230565b5b61048f825461028a565b61049a8282856103da565b5f60209050601f8311600181146104cb575f84156104b9578287015190505b6104c38582610448565b86555061052a565b601f1984166104d9866102ba565b5f5b82811015610500578489015182556001820191506020850194506020810190506104db565b8683101561051d5784890151610519601f89168261042c565b8355505b6001600288020188555050505b505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61055b82610532565b9050919050565b61056b81610551565b82525050565b5f6020820190506105845f830184610562565b92915050565b612d54806105975f395ff3fe608060405234801561000f575f5ffd5b5060043610610156575f3560e01c806370a08231116100c1578063af04951c1161007a578063af04951c146103dd578063b6f8f404146103f9578063b88d4fde1461042c578063c87b56dd14610448578063e985e9c514610478578063f2fde38b146104a857610156565b806370a0823114610316578063715018a6146103465780638da5cb5b1461035057806395d89b411461036e578063a22cb4651461038c578063a4d743b6146103a857610156565b80630a2ac745116101135780630a2ac7451461024257806323b872dd146102605780632f994c6f1461027c57806342842e0e1461029a578063465dcabb146102b65780636352211e146102e657610156565b8063010028ce1461015a57806301ffc9a71461017857806305ffec51146101a857806306fdde03146101d8578063081812fc146101f6578063095ea7b314610226575b5f5ffd5b6101626104c4565b60405161016f9190612269565b60405180910390f35b610192600480360381019061018d91906122e8565b6104d0565b60405161019f919061232d565b60405180910390f35b6101c260048036038101906101bd9190612370565b6105b1565b6040516101cf9190612462565b60405180910390f35b6101e06107ae565b6040516101ed91906124eb565b60405180910390f35b610210600480360381019061020b9190612370565b61083d565b60405161021d919061251a565b60405180910390f35b610240600480360381019061023b919061255d565b610858565b005b61024a61086e565b604051610257919061251a565b60405180910390f35b61027a6004803603810190610275919061259b565b610893565b005b610284610992565b604051610291919061251a565b60405180910390f35b6102b460048036038101906102af919061259b565b6109ba565b005b6102d060048036038101906102cb91906125eb565b6109d9565b6040516102dd9190612269565b60405180910390f35b61030060048036038101906102fb9190612370565b610bae565b60405161030d919061251a565b60405180910390f35b610330600480360381019061032b9190612629565b610bbf565b60405161033d9190612269565b60405180910390f35b61034e610c75565b005b610358610c88565b604051610365919061251a565b60405180910390f35b610376610cb0565b60405161038391906124eb565b60405180910390f35b6103a660048036038101906103a1919061267e565b610d40565b005b6103c260048036038101906103bd9190612370565b610d56565b6040516103d4969594939291906126bc565b60405180910390f35b6103f760048036038101906103f2919061271b565b610e19565b005b610413600480360381019061040e9190612370565b611169565b604051610423949392919061276b565b60405180910390f35b610446600480360381019061044191906128da565b6111e2565b005b610462600480360381019061045d9190612370565b611207565b60405161046f91906124eb565b60405180910390f35b610492600480360381019061048d91906125eb565b61126d565b60405161049f919061232d565b60405180910390f35b6104c260048036038101906104bd9190612629565b6112fb565b005b5f600c80549050905090565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061059a57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806105aa57506105a98261137f565b5b9050919050565b6105b96121c9565b5f821180156105ca57506009548211155b610609576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610600906129a4565b60405180910390fd5b600c60018361061891906129ef565b8154811061062957610628612a22565b5b905f5260205f2090600602016040518060c00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b60605f80546107bc90612a7c565b80601f01602080910402602001604051908101604052809291908181526020018280546107e890612a7c565b80156108335780601f1061080a57610100808354040283529160200191610833565b820191905f5260205f20905b81548152906001019060200180831161081657829003601f168201915b5050505050905090565b5f610847826113e8565b506108518261146e565b9050919050565b61086a82826108656114a7565b6114ae565b5050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610903575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016108fa919061251a565b60405180910390fd5b5f61091683836109116114a7565b6114c0565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461098c578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161098393929190612aac565b60405180910390fd5b50505050565b5f60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6109d483838360405180602001604052805f8152506111e2565b505050565b5f6109e26116cb565b60085f8154809291906109f490612ae1565b9190505550600a5f815480929190610a0b90612ae1565b91905055505f6008549050610a208482611752565b5f60405180608001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001600a548152509050600b81908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301555050818573ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f9e8a709a89b83b7026f9ebab1f1b5b04c5edf5339369ea5ae302868f443aead7600a54604051610b9b9190612269565b60405180910390a4819250505092915050565b5f610bb8826113e8565b9050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c30575f6040517f89c62b64000000000000000000000000000000000000000000000000000000008152600401610c27919061251a565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b610c7d6116cb565b610c865f611845565b565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060018054610cbf90612a7c565b80601f0160208091040260200160405190810160405280929190818152602001828054610ceb90612a7c565b8015610d365780601f10610d0d57610100808354040283529160200191610d36565b820191905f5260205f20905b815481529060010190602001808311610d1957829003601f168201915b5050505050905090565b610d52610d4b6114a7565b8383611908565b5050565b600c8181548110610d65575f80fd5b905f5260205f2090600602015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050154905086565b5f8390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd3360075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16866040518463ffffffff1660e01b8152600401610e7b93929190612b28565b5f604051808303815f87803b158015610e92575f5ffd5b505af1158015610ea4573d5f5f3e3d5ffd5b5050505060095f815480929190610eba90612ae1565b91905055505f6040518060c001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff16815260200160075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018581526020016009548152509050600c81908060018154018082558091505060019003905f5260205f2090600602015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050155505060075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f891979b6954d2045870d32e490afb9469603bc17f2447c48403df3b1fa5c58f0868860095460405161115a93929190612b5d565b60405180910390a45050505050565b600b8181548110611178575f80fd5b905f5260205f2090600402015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154905084565b6111ed848484610893565b6112016111f86114a7565b85858585611a71565b50505050565b6060611212826113e8565b505f61121c611c1d565b90505f81511161123a5760405180602001604052805f815250611265565b8061124484611c33565b604051602001611255929190612bcc565b6040516020818303038152906040525b915050919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b6113036116cb565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611373575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161136a919061251a565b60405180910390fd5b61137c81611845565b50565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f5f6113f383611cfd565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361146557826040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161145c9190612269565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b6114bb8383836001611d36565b505050565b5f5f6114cb84611cfd565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461150c5761150b818486611ef5565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146115975761154b5f855f5f611d36565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461161657600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b6116d36114a7565b73ffffffffffffffffffffffffffffffffffffffff166116f1610c88565b73ffffffffffffffffffffffffffffffffffffffff1614611750576117146114a7565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401611747919061251a565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036117c2575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016117b9919061251a565b60405180910390fd5b5f6117ce83835f6114c0565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611840575f6040517f73c6ac6e000000000000000000000000000000000000000000000000000000008152600401611837919061251a565b60405180910390fd5b505050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361197857816040517f5b08ba1800000000000000000000000000000000000000000000000000000000815260040161196f919061251a565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051611a64919061232d565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115611c16578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401611acf9493929190612c41565b6020604051808303815f875af1925050508015611b0a57506040513d601f19601f82011682018060405250810190611b079190612c9f565b60015b611b8b573d805f8114611b38576040519150601f19603f3d011682016040523d82523d5f602084013e611b3d565b606091505b505f815103611b8357836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401611b7a919061251a565b60405180910390fd5b805160208201fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614611c1457836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401611c0b919061251a565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60605f6001611c4184611fb8565b0190505f8167ffffffffffffffff811115611c5f57611c5e6127b6565b5b6040519080825280601f01601f191660200182016040528015611c915781602001600182028036833780820191505090505b5090505f82602083010190505b600115611cf2578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8581611ce757611ce6612cca565b5b0494505f8503611c9e575b819350505050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8080611d6e57505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611ea0575f611d7d846113e8565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015611de757508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b8015611dfa5750611df8818461126d565b155b15611e3c57826040517fa9fbf51f000000000000000000000000000000000000000000000000000000008152600401611e33919061251a565b60405180910390fd5b8115611e9e57838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b611f00838383612109565b611fb3575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611f7457806040517f7e273289000000000000000000000000000000000000000000000000000000008152600401611f6b9190612269565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611faa929190612cf7565b60405180910390fd5b505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310612014577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000838161200a57612009612cca565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310612051576d04ee2d6d415b85acef8100000000838161204757612046612cca565b5b0492506020810190505b662386f26fc10000831061208057662386f26fc10000838161207657612075612cca565b5b0492506010810190505b6305f5e10083106120a9576305f5e100838161209f5761209e612cca565b5b0492506008810190505b61271083106120ce5761271083816120c4576120c3612cca565b5b0492506004810190505b606483106120f157606483816120e7576120e6612cca565b5b0492506002810190505b600a8310612100576001810190505b80915050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156121c057508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806121815750612180848461126d565b5b806121bf57508273ffffffffffffffffffffffffffffffffffffffff166121a78361146e565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b6040518060c001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b5f819050919050565b61226381612251565b82525050565b5f60208201905061227c5f83018461225a565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6122c781612293565b81146122d1575f5ffd5b50565b5f813590506122e2816122be565b92915050565b5f602082840312156122fd576122fc61228b565b5b5f61230a848285016122d4565b91505092915050565b5f8115159050919050565b61232781612313565b82525050565b5f6020820190506123405f83018461231e565b92915050565b61234f81612251565b8114612359575f5ffd5b50565b5f8135905061236a81612346565b92915050565b5f602082840312156123855761238461228b565b5b5f6123928482850161235c565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6123c48261239b565b9050919050565b6123d4816123ba565b82525050565b6123e381612251565b82525050565b60c082015f8201516123fd5f8501826123cb565b50602082015161241060208501826123cb565b50604082015161242360408501826123cb565b50606082015161243660608501826123cb565b50608082015161244960808501826123da565b5060a082015161245c60a08501826123da565b50505050565b5f60c0820190506124755f8301846123e9565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6124bd8261247b565b6124c78185612485565b93506124d7818560208601612495565b6124e0816124a3565b840191505092915050565b5f6020820190508181035f83015261250381846124b3565b905092915050565b612514816123ba565b82525050565b5f60208201905061252d5f83018461250b565b92915050565b61253c816123ba565b8114612546575f5ffd5b50565b5f8135905061255781612533565b92915050565b5f5f604083850312156125735761257261228b565b5b5f61258085828601612549565b92505060206125918582860161235c565b9150509250929050565b5f5f5f606084860312156125b2576125b161228b565b5b5f6125bf86828701612549565b93505060206125d086828701612549565b92505060406125e18682870161235c565b9150509250925092565b5f5f604083850312156126015761260061228b565b5b5f61260e85828601612549565b925050602061261f85828601612549565b9150509250929050565b5f6020828403121561263e5761263d61228b565b5b5f61264b84828501612549565b91505092915050565b61265d81612313565b8114612667575f5ffd5b50565b5f8135905061267881612654565b92915050565b5f5f604083850312156126945761269361228b565b5b5f6126a185828601612549565b92505060206126b28582860161266a565b9150509250929050565b5f60c0820190506126cf5f83018961250b565b6126dc602083018861250b565b6126e9604083018761250b565b6126f6606083018661250b565b612703608083018561225a565b61271060a083018461225a565b979650505050505050565b5f5f5f606084860312156127325761273161228b565b5b5f61273f86828701612549565b93505060206127508682870161235c565b925050604061276186828701612549565b9150509250925092565b5f60808201905061277e5f83018761250b565b61278b602083018661250b565b612798604083018561225a565b6127a5606083018461225a565b95945050505050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6127ec826124a3565b810181811067ffffffffffffffff8211171561280b5761280a6127b6565b5b80604052505050565b5f61281d612282565b905061282982826127e3565b919050565b5f67ffffffffffffffff821115612848576128476127b6565b5b612851826124a3565b9050602081019050919050565b828183375f83830152505050565b5f61287e6128798461282e565b612814565b90508281526020810184848401111561289a576128996127b2565b5b6128a584828561285e565b509392505050565b5f82601f8301126128c1576128c06127ae565b5b81356128d184826020860161286c565b91505092915050565b5f5f5f5f608085870312156128f2576128f161228b565b5b5f6128ff87828801612549565b945050602061291087828801612549565b93505060406129218782880161235c565b925050606085013567ffffffffffffffff8111156129425761294161228f565b5b61294e878288016128ad565b91505092959194509250565b7f5472616e7366657220494420696e76616c69646f0000000000000000000000005f82015250565b5f61298e601483612485565b91506129998261295a565b602082019050919050565b5f6020820190508181035f8301526129bb81612982565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6129f982612251565b9150612a0483612251565b9250828203905081811115612a1c57612a1b6129c2565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680612a9357607f821691505b602082108103612aa657612aa5612a4f565b5b50919050565b5f606082019050612abf5f83018661250b565b612acc602083018561225a565b612ad9604083018461250b565b949350505050565b5f612aeb82612251565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612b1d57612b1c6129c2565b5b600182019050919050565b5f606082019050612b3b5f83018661250b565b612b48602083018561250b565b612b55604083018461225a565b949350505050565b5f606082019050612b705f83018661250b565b612b7d602083018561225a565b612b8a604083018461225a565b949350505050565b5f81905092915050565b5f612ba68261247b565b612bb08185612b92565b9350612bc0818560208601612495565b80840191505092915050565b5f612bd78285612b9c565b9150612be38284612b9c565b91508190509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f612c1382612bef565b612c1d8185612bf9565b9350612c2d818560208601612495565b612c36816124a3565b840191505092915050565b5f608082019050612c545f83018761250b565b612c61602083018661250b565b612c6e604083018561225a565b8181036060830152612c808184612c09565b905095945050505050565b5f81519050612c99816122be565b92915050565b5f60208284031215612cb457612cb361228b565b5b5f612cc184828501612c8b565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f604082019050612d0a5f83018561250b565b612d17602083018461225a565b939250505056fea2646970667358221220967349d3842f2cc052512f04a370f71788409f133503c0358d59831c095b7e8764736f6c634300081e0033",
}

// NotaryABI is the input ABI used to generate the binding from.
// Deprecated: Use NotaryMetaData.ABI instead.
var NotaryABI = NotaryMetaData.ABI

// NotaryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NotaryMetaData.Bin instead.
var NotaryBin = NotaryMetaData.Bin

// DeployNotary deploys a new Ethereum contract, binding an instance of Notary to it.
func DeployNotary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Notary, error) {
	parsed, err := NotaryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NotaryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Notary{NotaryCaller: NotaryCaller{contract: contract}, NotaryTransactor: NotaryTransactor{contract: contract}, NotaryFilterer: NotaryFilterer{contract: contract}}, nil
}

// Notary is an auto generated Go binding around an Ethereum contract.
type Notary struct {
	NotaryCaller     // Read-only binding to the contract
	NotaryTransactor // Write-only binding to the contract
	NotaryFilterer   // Log filterer for contract events
}

// NotaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotaryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotarySession struct {
	Contract     *Notary           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotaryCallerSession struct {
	Contract *NotaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NotaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotaryTransactorSession struct {
	Contract     *NotaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotaryRaw struct {
	Contract *Notary // Generic contract binding to access the raw methods on
}

// NotaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotaryCallerRaw struct {
	Contract *NotaryCaller // Generic read-only contract binding to access the raw methods on
}

// NotaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotaryTransactorRaw struct {
	Contract *NotaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotary creates a new instance of Notary, bound to a specific deployed contract.
func NewNotary(address common.Address, backend bind.ContractBackend) (*Notary, error) {
	contract, err := bindNotary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Notary{NotaryCaller: NotaryCaller{contract: contract}, NotaryTransactor: NotaryTransactor{contract: contract}, NotaryFilterer: NotaryFilterer{contract: contract}}, nil
}

// NewNotaryCaller creates a new read-only instance of Notary, bound to a specific deployed contract.
func NewNotaryCaller(address common.Address, caller bind.ContractCaller) (*NotaryCaller, error) {
	contract, err := bindNotary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryCaller{contract: contract}, nil
}

// NewNotaryTransactor creates a new write-only instance of Notary, bound to a specific deployed contract.
func NewNotaryTransactor(address common.Address, transactor bind.ContractTransactor) (*NotaryTransactor, error) {
	contract, err := bindNotary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryTransactor{contract: contract}, nil
}

// NewNotaryFilterer creates a new log filterer instance of Notary, bound to a specific deployed contract.
func NewNotaryFilterer(address common.Address, filterer bind.ContractFilterer) (*NotaryFilterer, error) {
	contract, err := bindNotary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotaryFilterer{contract: contract}, nil
}

// bindNotary binds a generic wrapper to an already deployed contract.
func bindNotary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NotaryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notary *NotaryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notary.Contract.NotaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notary *NotaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notary.Contract.NotaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notary *NotaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notary.Contract.NotaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notary *NotaryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notary *NotaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notary *NotaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notary.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Notary *NotaryCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Notary *NotarySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Notary.Contract.BalanceOf(&_Notary.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Notary *NotaryCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Notary.Contract.BalanceOf(&_Notary.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Notary *NotaryCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Notary *NotarySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Notary.Contract.GetApproved(&_Notary.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Notary *NotaryCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Notary.Contract.GetApproved(&_Notary.CallOpts, tokenId)
}

// GetNftTransfers is a free data retrieval call binding the contract method 0x010028ce.
//
// Solidity: function getNftTransfers() view returns(uint256)
func (_Notary *NotaryCaller) GetNftTransfers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "getNftTransfers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNftTransfers is a free data retrieval call binding the contract method 0x010028ce.
//
// Solidity: function getNftTransfers() view returns(uint256)
func (_Notary *NotarySession) GetNftTransfers() (*big.Int, error) {
	return _Notary.Contract.GetNftTransfers(&_Notary.CallOpts)
}

// GetNftTransfers is a free data retrieval call binding the contract method 0x010028ce.
//
// Solidity: function getNftTransfers() view returns(uint256)
func (_Notary *NotaryCallerSession) GetNftTransfers() (*big.Int, error) {
	return _Notary.Contract.GetNftTransfers(&_Notary.CallOpts)
}

// GetNotaryOwner is a free data retrieval call binding the contract method 0x2f994c6f.
//
// Solidity: function getNotaryOwner() view returns(address)
func (_Notary *NotaryCaller) GetNotaryOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "getNotaryOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNotaryOwner is a free data retrieval call binding the contract method 0x2f994c6f.
//
// Solidity: function getNotaryOwner() view returns(address)
func (_Notary *NotarySession) GetNotaryOwner() (common.Address, error) {
	return _Notary.Contract.GetNotaryOwner(&_Notary.CallOpts)
}

// GetNotaryOwner is a free data retrieval call binding the contract method 0x2f994c6f.
//
// Solidity: function getNotaryOwner() view returns(address)
func (_Notary *NotaryCallerSession) GetNotaryOwner() (common.Address, error) {
	return _Notary.Contract.GetNotaryOwner(&_Notary.CallOpts)
}

// GetTransferById is a free data retrieval call binding the contract method 0x05ffec51.
//
// Solidity: function getTransferById(uint256 transferId) view returns((address,address,address,address,uint256,uint256))
func (_Notary *NotaryCaller) GetTransferById(opts *bind.CallOpts, transferId *big.Int) (NotaryNFTTransfer, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "getTransferById", transferId)

	if err != nil {
		return *new(NotaryNFTTransfer), err
	}

	out0 := *abi.ConvertType(out[0], new(NotaryNFTTransfer)).(*NotaryNFTTransfer)

	return out0, err

}

// GetTransferById is a free data retrieval call binding the contract method 0x05ffec51.
//
// Solidity: function getTransferById(uint256 transferId) view returns((address,address,address,address,uint256,uint256))
func (_Notary *NotarySession) GetTransferById(transferId *big.Int) (NotaryNFTTransfer, error) {
	return _Notary.Contract.GetTransferById(&_Notary.CallOpts, transferId)
}

// GetTransferById is a free data retrieval call binding the contract method 0x05ffec51.
//
// Solidity: function getTransferById(uint256 transferId) view returns((address,address,address,address,uint256,uint256))
func (_Notary *NotaryCallerSession) GetTransferById(transferId *big.Int) (NotaryNFTTransfer, error) {
	return _Notary.Contract.GetTransferById(&_Notary.CallOpts, transferId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Notary *NotaryCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Notary *NotarySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Notary.Contract.IsApprovedForAll(&_Notary.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Notary *NotaryCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Notary.Contract.IsApprovedForAll(&_Notary.CallOpts, owner, operator)
}

// MintRecords is a free data retrieval call binding the contract method 0xb6f8f404.
//
// Solidity: function mintRecords(uint256 ) view returns(address sender, address receiver, uint256 tokenId, uint256 mintId)
func (_Notary *NotaryCaller) MintRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sender   common.Address
	Receiver common.Address
	TokenId  *big.Int
	MintId   *big.Int
}, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "mintRecords", arg0)

	outstruct := new(struct {
		Sender   common.Address
		Receiver common.Address
		TokenId  *big.Int
		MintId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MintId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MintRecords is a free data retrieval call binding the contract method 0xb6f8f404.
//
// Solidity: function mintRecords(uint256 ) view returns(address sender, address receiver, uint256 tokenId, uint256 mintId)
func (_Notary *NotarySession) MintRecords(arg0 *big.Int) (struct {
	Sender   common.Address
	Receiver common.Address
	TokenId  *big.Int
	MintId   *big.Int
}, error) {
	return _Notary.Contract.MintRecords(&_Notary.CallOpts, arg0)
}

// MintRecords is a free data retrieval call binding the contract method 0xb6f8f404.
//
// Solidity: function mintRecords(uint256 ) view returns(address sender, address receiver, uint256 tokenId, uint256 mintId)
func (_Notary *NotaryCallerSession) MintRecords(arg0 *big.Int) (struct {
	Sender   common.Address
	Receiver common.Address
	TokenId  *big.Int
	MintId   *big.Int
}, error) {
	return _Notary.Contract.MintRecords(&_Notary.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Notary *NotaryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Notary *NotarySession) Name() (string, error) {
	return _Notary.Contract.Name(&_Notary.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Notary *NotaryCallerSession) Name() (string, error) {
	return _Notary.Contract.Name(&_Notary.CallOpts)
}

// NftTransfers is a free data retrieval call binding the contract method 0xa4d743b6.
//
// Solidity: function nftTransfers(uint256 ) view returns(address nftContract, address sender, address receiverOnChain, address receiverInterChain, uint256 tokenId, uint256 transferId)
func (_Notary *NotaryCaller) NftTransfers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftContract        common.Address
	Sender             common.Address
	ReceiverOnChain    common.Address
	ReceiverInterChain common.Address
	TokenId            *big.Int
	TransferId         *big.Int
}, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "nftTransfers", arg0)

	outstruct := new(struct {
		NftContract        common.Address
		Sender             common.Address
		ReceiverOnChain    common.Address
		ReceiverInterChain common.Address
		TokenId            *big.Int
		TransferId         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Sender = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReceiverOnChain = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ReceiverInterChain = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TransferId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NftTransfers is a free data retrieval call binding the contract method 0xa4d743b6.
//
// Solidity: function nftTransfers(uint256 ) view returns(address nftContract, address sender, address receiverOnChain, address receiverInterChain, uint256 tokenId, uint256 transferId)
func (_Notary *NotarySession) NftTransfers(arg0 *big.Int) (struct {
	NftContract        common.Address
	Sender             common.Address
	ReceiverOnChain    common.Address
	ReceiverInterChain common.Address
	TokenId            *big.Int
	TransferId         *big.Int
}, error) {
	return _Notary.Contract.NftTransfers(&_Notary.CallOpts, arg0)
}

// NftTransfers is a free data retrieval call binding the contract method 0xa4d743b6.
//
// Solidity: function nftTransfers(uint256 ) view returns(address nftContract, address sender, address receiverOnChain, address receiverInterChain, uint256 tokenId, uint256 transferId)
func (_Notary *NotaryCallerSession) NftTransfers(arg0 *big.Int) (struct {
	NftContract        common.Address
	Sender             common.Address
	ReceiverOnChain    common.Address
	ReceiverInterChain common.Address
	TokenId            *big.Int
	TransferId         *big.Int
}, error) {
	return _Notary.Contract.NftTransfers(&_Notary.CallOpts, arg0)
}

// NotaryOwner is a free data retrieval call binding the contract method 0x0a2ac745.
//
// Solidity: function notaryOwner() view returns(address)
func (_Notary *NotaryCaller) NotaryOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "notaryOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NotaryOwner is a free data retrieval call binding the contract method 0x0a2ac745.
//
// Solidity: function notaryOwner() view returns(address)
func (_Notary *NotarySession) NotaryOwner() (common.Address, error) {
	return _Notary.Contract.NotaryOwner(&_Notary.CallOpts)
}

// NotaryOwner is a free data retrieval call binding the contract method 0x0a2ac745.
//
// Solidity: function notaryOwner() view returns(address)
func (_Notary *NotaryCallerSession) NotaryOwner() (common.Address, error) {
	return _Notary.Contract.NotaryOwner(&_Notary.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Notary *NotaryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Notary *NotarySession) Owner() (common.Address, error) {
	return _Notary.Contract.Owner(&_Notary.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Notary *NotaryCallerSession) Owner() (common.Address, error) {
	return _Notary.Contract.Owner(&_Notary.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Notary *NotaryCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Notary *NotarySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Notary.Contract.OwnerOf(&_Notary.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Notary *NotaryCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Notary.Contract.OwnerOf(&_Notary.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Notary *NotaryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Notary *NotarySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Notary.Contract.SupportsInterface(&_Notary.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Notary *NotaryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Notary.Contract.SupportsInterface(&_Notary.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Notary *NotaryCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Notary *NotarySession) Symbol() (string, error) {
	return _Notary.Contract.Symbol(&_Notary.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Notary *NotaryCallerSession) Symbol() (string, error) {
	return _Notary.Contract.Symbol(&_Notary.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Notary *NotaryCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Notary.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Notary *NotarySession) TokenURI(tokenId *big.Int) (string, error) {
	return _Notary.Contract.TokenURI(&_Notary.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Notary *NotaryCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Notary.Contract.TokenURI(&_Notary.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Notary *NotaryTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Notary *NotarySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.Contract.Approve(&_Notary.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Notary *NotaryTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.Contract.Approve(&_Notary.TransactOpts, to, tokenId)
}

// MintNFT is a paid mutator transaction binding the contract method 0x465dcabb.
//
// Solidity: function mintNFT(address to, address senderInterChain) returns(uint256)
func (_Notary *NotaryTransactor) MintNFT(opts *bind.TransactOpts, to common.Address, senderInterChain common.Address) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "mintNFT", to, senderInterChain)
}

// MintNFT is a paid mutator transaction binding the contract method 0x465dcabb.
//
// Solidity: function mintNFT(address to, address senderInterChain) returns(uint256)
func (_Notary *NotarySession) MintNFT(to common.Address, senderInterChain common.Address) (*types.Transaction, error) {
	return _Notary.Contract.MintNFT(&_Notary.TransactOpts, to, senderInterChain)
}

// MintNFT is a paid mutator transaction binding the contract method 0x465dcabb.
//
// Solidity: function mintNFT(address to, address senderInterChain) returns(uint256)
func (_Notary *NotaryTransactorSession) MintNFT(to common.Address, senderInterChain common.Address) (*types.Transaction, error) {
	return _Notary.Contract.MintNFT(&_Notary.TransactOpts, to, senderInterChain)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Notary *NotaryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Notary *NotarySession) RenounceOwnership() (*types.Transaction, error) {
	return _Notary.Contract.RenounceOwnership(&_Notary.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Notary *NotaryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Notary.Contract.RenounceOwnership(&_Notary.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Notary *NotaryTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Notary *NotarySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.Contract.SafeTransferFrom(&_Notary.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Notary *NotaryTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.Contract.SafeTransferFrom(&_Notary.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Notary *NotaryTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Notary *NotarySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Notary.Contract.SafeTransferFrom0(&_Notary.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Notary *NotaryTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Notary.Contract.SafeTransferFrom0(&_Notary.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Notary *NotaryTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Notary *NotarySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Notary.Contract.SetApprovalForAll(&_Notary.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Notary *NotaryTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Notary.Contract.SetApprovalForAll(&_Notary.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Notary *NotaryTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Notary *NotarySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.Contract.TransferFrom(&_Notary.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Notary *NotaryTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Notary.Contract.TransferFrom(&_Notary.TransactOpts, from, to, tokenId)
}

// TransferNFTInterChain is a paid mutator transaction binding the contract method 0xaf04951c.
//
// Solidity: function transferNFTInterChain(address nftContract, uint256 tokenId, address receiverInterChain) returns()
func (_Notary *NotaryTransactor) TransferNFTInterChain(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, receiverInterChain common.Address) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "transferNFTInterChain", nftContract, tokenId, receiverInterChain)
}

// TransferNFTInterChain is a paid mutator transaction binding the contract method 0xaf04951c.
//
// Solidity: function transferNFTInterChain(address nftContract, uint256 tokenId, address receiverInterChain) returns()
func (_Notary *NotarySession) TransferNFTInterChain(nftContract common.Address, tokenId *big.Int, receiverInterChain common.Address) (*types.Transaction, error) {
	return _Notary.Contract.TransferNFTInterChain(&_Notary.TransactOpts, nftContract, tokenId, receiverInterChain)
}

// TransferNFTInterChain is a paid mutator transaction binding the contract method 0xaf04951c.
//
// Solidity: function transferNFTInterChain(address nftContract, uint256 tokenId, address receiverInterChain) returns()
func (_Notary *NotaryTransactorSession) TransferNFTInterChain(nftContract common.Address, tokenId *big.Int, receiverInterChain common.Address) (*types.Transaction, error) {
	return _Notary.Contract.TransferNFTInterChain(&_Notary.TransactOpts, nftContract, tokenId, receiverInterChain)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Notary *NotaryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Notary.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Notary *NotarySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Notary.Contract.TransferOwnership(&_Notary.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Notary *NotaryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Notary.Contract.TransferOwnership(&_Notary.TransactOpts, newOwner)
}

// NotaryApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Notary contract.
type NotaryApprovalIterator struct {
	Event *NotaryApproval // Event containing the contract specifics and raw log

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
func (it *NotaryApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryApproval)
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
		it.Event = new(NotaryApproval)
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
func (it *NotaryApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryApproval represents a Approval event raised by the Notary contract.
type NotaryApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Notary *NotaryFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NotaryApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NotaryApprovalIterator{contract: _Notary.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Notary *NotaryFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NotaryApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryApproval)
				if err := _Notary.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Notary *NotaryFilterer) ParseApproval(log types.Log) (*NotaryApproval, error) {
	event := new(NotaryApproval)
	if err := _Notary.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Notary contract.
type NotaryApprovalForAllIterator struct {
	Event *NotaryApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NotaryApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryApprovalForAll)
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
		it.Event = new(NotaryApprovalForAll)
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
func (it *NotaryApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryApprovalForAll represents a ApprovalForAll event raised by the Notary contract.
type NotaryApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Notary *NotaryFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NotaryApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NotaryApprovalForAllIterator{contract: _Notary.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Notary *NotaryFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NotaryApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryApprovalForAll)
				if err := _Notary.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Notary *NotaryFilterer) ParseApprovalForAll(log types.Log) (*NotaryApprovalForAll, error) {
	event := new(NotaryApprovalForAll)
	if err := _Notary.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the Notary contract.
type NotaryNFTMintedIterator struct {
	Event *NotaryNFTMinted // Event containing the contract specifics and raw log

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
func (it *NotaryNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryNFTMinted)
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
		it.Event = new(NotaryNFTMinted)
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
func (it *NotaryNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryNFTMinted represents a NFTMinted event raised by the Notary contract.
type NotaryNFTMinted struct {
	Sender   common.Address
	Receiver common.Address
	TokenId  *big.Int
	MintId   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0x9e8a709a89b83b7026f9ebab1f1b5b04c5edf5339369ea5ae302868f443aead7.
//
// Solidity: event NFTMinted(address indexed sender, address indexed receiver, uint256 indexed tokenId, uint256 mintId)
func (_Notary *NotaryFilterer) FilterNFTMinted(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, tokenId []*big.Int) (*NotaryNFTMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "NFTMinted", senderRule, receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NotaryNFTMintedIterator{contract: _Notary.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0x9e8a709a89b83b7026f9ebab1f1b5b04c5edf5339369ea5ae302868f443aead7.
//
// Solidity: event NFTMinted(address indexed sender, address indexed receiver, uint256 indexed tokenId, uint256 mintId)
func (_Notary *NotaryFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *NotaryNFTMinted, sender []common.Address, receiver []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "NFTMinted", senderRule, receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryNFTMinted)
				if err := _Notary.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0x9e8a709a89b83b7026f9ebab1f1b5b04c5edf5339369ea5ae302868f443aead7.
//
// Solidity: event NFTMinted(address indexed sender, address indexed receiver, uint256 indexed tokenId, uint256 mintId)
func (_Notary *NotaryFilterer) ParseNFTMinted(log types.Log) (*NotaryNFTMinted, error) {
	event := new(NotaryNFTMinted)
	if err := _Notary.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryNFTTransferInterChainIterator is returned from FilterNFTTransferInterChain and is used to iterate over the raw logs and unpacked data for NFTTransferInterChain events raised by the Notary contract.
type NotaryNFTTransferInterChainIterator struct {
	Event *NotaryNFTTransferInterChain // Event containing the contract specifics and raw log

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
func (it *NotaryNFTTransferInterChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryNFTTransferInterChain)
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
		it.Event = new(NotaryNFTTransferInterChain)
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
func (it *NotaryNFTTransferInterChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryNFTTransferInterChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryNFTTransferInterChain represents a NFTTransferInterChain event raised by the Notary contract.
type NotaryNFTTransferInterChain struct {
	NftContract        common.Address
	Sender             common.Address
	ReceiverOnChain    common.Address
	ReceiverInterChain common.Address
	TokenId            *big.Int
	TransferId         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNFTTransferInterChain is a free log retrieval operation binding the contract event 0x891979b6954d2045870d32e490afb9469603bc17f2447c48403df3b1fa5c58f0.
//
// Solidity: event NFTTransferInterChain(address indexed nftContract, address indexed sender, address indexed receiverOnChain, address receiverInterChain, uint256 tokenId, uint256 transferId)
func (_Notary *NotaryFilterer) FilterNFTTransferInterChain(opts *bind.FilterOpts, nftContract []common.Address, sender []common.Address, receiverOnChain []common.Address) (*NotaryNFTTransferInterChainIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverOnChainRule []interface{}
	for _, receiverOnChainItem := range receiverOnChain {
		receiverOnChainRule = append(receiverOnChainRule, receiverOnChainItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "NFTTransferInterChain", nftContractRule, senderRule, receiverOnChainRule)
	if err != nil {
		return nil, err
	}
	return &NotaryNFTTransferInterChainIterator{contract: _Notary.contract, event: "NFTTransferInterChain", logs: logs, sub: sub}, nil
}

// WatchNFTTransferInterChain is a free log subscription operation binding the contract event 0x891979b6954d2045870d32e490afb9469603bc17f2447c48403df3b1fa5c58f0.
//
// Solidity: event NFTTransferInterChain(address indexed nftContract, address indexed sender, address indexed receiverOnChain, address receiverInterChain, uint256 tokenId, uint256 transferId)
func (_Notary *NotaryFilterer) WatchNFTTransferInterChain(opts *bind.WatchOpts, sink chan<- *NotaryNFTTransferInterChain, nftContract []common.Address, sender []common.Address, receiverOnChain []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverOnChainRule []interface{}
	for _, receiverOnChainItem := range receiverOnChain {
		receiverOnChainRule = append(receiverOnChainRule, receiverOnChainItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "NFTTransferInterChain", nftContractRule, senderRule, receiverOnChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryNFTTransferInterChain)
				if err := _Notary.contract.UnpackLog(event, "NFTTransferInterChain", log); err != nil {
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

// ParseNFTTransferInterChain is a log parse operation binding the contract event 0x891979b6954d2045870d32e490afb9469603bc17f2447c48403df3b1fa5c58f0.
//
// Solidity: event NFTTransferInterChain(address indexed nftContract, address indexed sender, address indexed receiverOnChain, address receiverInterChain, uint256 tokenId, uint256 transferId)
func (_Notary *NotaryFilterer) ParseNFTTransferInterChain(log types.Log) (*NotaryNFTTransferInterChain, error) {
	event := new(NotaryNFTTransferInterChain)
	if err := _Notary.contract.UnpackLog(event, "NFTTransferInterChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Notary contract.
type NotaryOwnershipTransferredIterator struct {
	Event *NotaryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NotaryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryOwnershipTransferred)
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
		it.Event = new(NotaryOwnershipTransferred)
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
func (it *NotaryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryOwnershipTransferred represents a OwnershipTransferred event raised by the Notary contract.
type NotaryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Notary *NotaryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NotaryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NotaryOwnershipTransferredIterator{contract: _Notary.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Notary *NotaryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NotaryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryOwnershipTransferred)
				if err := _Notary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Notary *NotaryFilterer) ParseOwnershipTransferred(log types.Log) (*NotaryOwnershipTransferred, error) {
	event := new(NotaryOwnershipTransferred)
	if err := _Notary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Notary contract.
type NotaryTransferIterator struct {
	Event *NotaryTransfer // Event containing the contract specifics and raw log

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
func (it *NotaryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryTransfer)
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
		it.Event = new(NotaryTransfer)
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
func (it *NotaryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryTransfer represents a Transfer event raised by the Notary contract.
type NotaryTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Notary *NotaryFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NotaryTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NotaryTransferIterator{contract: _Notary.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Notary *NotaryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NotaryTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryTransfer)
				if err := _Notary.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Notary *NotaryFilterer) ParseTransfer(log types.Log) (*NotaryTransfer, error) {
	event := new(NotaryTransfer)
	if err := _Notary.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
