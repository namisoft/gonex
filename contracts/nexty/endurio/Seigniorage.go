// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endurio

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbsorbableABI is the input ABI used to generate the binding from.
const AbsorbableABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// AbsorbableFuncSigs maps the 4-byte function signature to its string representation.
var AbsorbableFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"8aa3f897": "top(bool)",
}

// AbsorbableBin is the compiled bytecode used for deploying new contracts.
var AbsorbableBin = "0x608060405262049d4060035560026003548161001757fe5b0460045534801561002757600080fd5b50604051611cf9380380611cf98339818101604052604081101561004a57600080fd5b508051602090910151801561005f5760038190555b600082116100705760028104610072565b815b6004555050611c73806100866000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638aa3f897116100715780638aa3f897146101ba578063aa1c259c146101d9578063be91d72914610207578063ced4aac814610224578063ee1a68c614610264578063f318722b14610287576100b4565b806307c399a3146100b95780630d90b10a1461011357806343271d791461014a5780634ea097971461017157806369c07d31146101965780636e6452cb146101b2575b600080fd5b6100de600480360360408110156100cf57600080fd5b508035151590602001356102b3565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101386004803603604081101561012957600080fd5b508035151590602001356102fa565b60408051918252519081900360200190f35b61016f6004803603604081101561016057600080fd5b50803515159060200135610322565b005b6101386004803603604081101561018757600080fd5b508035151590602001356103a8565b61019e6103cc565b604080519115158252519081900360200190f35b61019e6103d1565b610138600480360360208110156101d057600080fd5b503515156103d6565b61016f600480360360408110156101ef57600080fd5b506001600160a01b03813581169160200135166103f9565b61016f6004803603602081101561021d57600080fd5b5035610548565b610138600480360360a081101561023a57600080fd5b5080351515906001600160a01b0360208201351690604081013590606081013590608001356106e2565b61026c61074a565b60408051921515835260208301919091528051918290030190f35b6101386004803603604081101561029d57600080fd5b506001600160a01b0381351690602001356107f4565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610392576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b6103a2828463ffffffff61080716565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b80151560009081526020819052604081206103f0816108c6565b9150505b919050565b6001546001600160a01b031615610445576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b6002546001600160a01b031615610491576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556104cb82826108df565b600254604080516318160ddd60e01b81529051610544926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561051157600080fd5b505afa158015610525573d6000803e3d6000fd5b505050506040513d602081101561053b57600080fd5b50516000610953565b5050565b331561058c576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6105966005610a1f565b156105b5576000600581905560068190556007556008805460ff191690555b6105bf6009610a38565b156105cc576105cc610a55565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561061157600080fd5b505afa158015610625573d6000803e3d6000fd5b505050506040513d602081101561063b57600080fd5b5051905081156106af5761064d610b20565b8061065d575061065d8183610b32565b1561066d5761066d826000610953565b6106776009610be0565b156106af5760006106888383610bfc565b905061069381610c16565b80156106a1575060085460ff165b156106ad5750506106df565b505b6106c060058263ffffffff610d0316565b156105445760006106cf610d38565b90506106da81610ded565b505050505b50565b84151560009081526020819052604081206106fb611beb565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261073e82828663ffffffff610e9416565b98975050505050505050565b6007546000908190610761575060009050806107f0565b60016107eb600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107ba57600080fd5b505afa1580156107ce573d6000803e3d6000fd5b505050506040513d60208110156107e457600080fd5b5051610bfc565b915091505b9091565b60006108008383610f06565b9392505050565b6000818152600283016020526040902060010154156108b6578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b15801561088957600080fd5b505af115801561089d573d6000803e3d6000fd5b505050506040513d60208110156108b357600080fd5b50505b610544828263ffffffff610fd316565b6000808052600282016020526040902060040154919050565b60008080526020526109187fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff61103316565b600160009081526020526105447fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff61103316565b604051806080016040528060035443018152602001600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156109b657600080fd5b505afa1580156109ca573d6000803e3d6000fd5b505050506040513d60208110156109e057600080fd5b50518152602080820194909452911515604092830152805160055591820151600655810151600755606001516008805460ff1916911515919091179055565b6000610a2a826111e0565b801561031c57505054431190565b6000610a43826111e6565b801561031c5750506004015443101590565b610a5f60096111e6565b610a6857610b1e565b600b5415610af957600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015610acc57600080fd5b505af1158015610ae0573d6000803e3d6000fd5b505050506040513d6020811015610af657600080fd5b50505b600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b6000610b2c6005610a1f565b90505b90565b600082821415610b445750600061031c565b6006546007541415610b585750600161031c565b82821115610bac57600654600754848403911015610b8e57600654600754036002818381610b8257fe5b0410159250505061031c565b600754600654036002828281610ba057fe5b0411159250505061031c565b600654600754838503911115610bce57600754600654036002818381610b8257fe5b600654600754036002828281610ba057fe5b6000610beb826111e6565b801561031c57505060040154431090565b6000818311610c1057828203600003610800565b50900390565b6000610c296009600101546000846111f5565b610c35575060006103f4565b600c54600a546000919084830381610c4957fe5b0581610c5157fe5b04905080610c5d575060015b600b54811115610c8c5750600b546000600581905560068190556007556008805460ff19169055610c8c610a55565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015610ce257600080fd5b505af1158015610cf6573d6000803e3d6000fd5b5060019695505050505050565b6000610d0e836111e0565b8015610d1e575082600201548214155b801561080057506108008360010154838560020154611225565b600080610d4f600560020154600560010154610bfc565b600754600254604080516318160ddd60e01b81529051939450600093610d9f93926001600160a01b0316916318160ddd916004808301926020929190829003018186803b1580156107ba57600080fd5b9050610dad600082846111f5565b610dbc57600092505050610b2f565b60006004548381610dc957fe5b059050610dd8600082846111f5565b610de657509150610b2f9050565b9250505090565b6000806000806000808613610e03576001610e06565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff1615610e6d57610e6281610e4787611250565b600954859291906001600160a01b031663ffffffff61126616565b935093505050610e8f565b610e8881610e7a87611250565b84919063ffffffff61161916565b9350935050505b915091565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610ecc848263ffffffff6118e616565b15610ea5575b6003015460008181526002860160205260409020909250610ef9848263ffffffff6118e616565b610ed25750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b60208310610f765780518252601f199092019160209182019101610f57565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610fb5573d6000803e3d6000fd5b5050506040513d6020811015610fca57600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b54151590565b546001600160a01b0316151590565b60008284131580156112075750818313155b8061121d575082841215801561121d5750818312155b949350505050565b60008284111580156112375750818311155b8061121d575082841015801561121d5750501115919050565b6000808213611262578160000361031c565b5090565b60008061127a86868663ffffffff61161916565b90925090506000808661128e578383611291565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b1580156112e757600080fd5b505afa1580156112fb573d6000803e3d6000fd5b505050506040513d602081101561131157600080fd5b5051821180159061139957508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b15801561136957600080fd5b505afa15801561137d573d6000803e3d6000fd5b505050506040513d602081101561139357600080fd5b50518211155b1561160e578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b1580156113ed57600080fd5b505af1158015611401573d6000803e3d6000fd5b505050506040513d602081101561141757600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b15801561147057600080fd5b505af1158015611484573d6000803e3d6000fd5b505050506040513d602081101561149a57600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156114e857600080fd5b505af11580156114fc573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561154e57600080fd5b505af1158015611562573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156115be57600080fd5b505af11580156115d2573d6000803e3d6000fd5b505050506040513d60208110156115e857600080fd5b508790506115f657806115f8565b815b83019250866116075781611609565b805b840193505b505094509492505050565b6000806000611627866108c6565b90505b600019811480159061163b57508382105b156118dd57600081815260028701602052604081209086611660578160020154611666565b81600101545b905060008761167957826001015461167f565b82600201545b905086611692868463ffffffff61190416565b1161178157885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b1580156116e857600080fd5b505af11580156116fc573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561175357600080fd5b505af1158015611767573d6000803e3d6000fd5b50505050600483015461177a8a8661195e565b93506118b1565b6000611793888763ffffffff6119e316565b905082818302816117a057fe5b0491508092506000896117b357826117b5565b835b905060008a6117c457846117c6565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b15801561181557600080fd5b505af1158015611829573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b15801561187b57600080fd5b505af115801561188f573d6000803e3d6000fd5b506118a892508e9150899050848463ffffffff611a3a16565b50600019955050505b6118c1858363ffffffff61190416565b94506118d3868263ffffffff61190416565b955050505061162a565b50935093915050565b60408201516001820154600290920154602090930151910291021190565b600082820183811015610800576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000818152600280840160205260409091200154156108b657600182015460008281526002808501602090815260408084208054930154815163a9059cbb60e01b81526001600160a01b03948516600482015260248101919091529051929094169363a9059cbb93604480830194928390030190829087803b15801561088957600080fd5b600082821115610c10576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b60008381526002850160205260409020611a53816111e6565b611a96576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b8060010154831115611aef576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b6001810154611b04908463ffffffff6119e316565b60018201556002810154821015611b25576002810180548390039055611b2d565b600060028201555b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015611b8757600080fd5b505af1158015611b9b573d6000803e3d6000fd5b505050506040513d6020811015611bb157600080fd5b50611bbd905081611bd2565b156106da576106da858563ffffffff61080716565b600081600101546000148061031c575050600201541590565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a7231582026e50a41c612785ba1d8747bec3dee2d6bc6e9e967b419de65c72fd1877b999c64736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployAbsorbable deploys a new Ethereum contract, binding an instance of Absorbable to it.
func DeployAbsorbable(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int) (common.Address, *types.Transaction, *Absorbable, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsorbableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbsorbableBin), backend, absorptionDuration, absorptionExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Absorbable{AbsorbableCaller: AbsorbableCaller{contract: contract}, AbsorbableTransactor: AbsorbableTransactor{contract: contract}, AbsorbableFilterer: AbsorbableFilterer{contract: contract}}, nil
}

// Absorbable is an auto generated Go binding around an Ethereum contract.
type Absorbable struct {
	AbsorbableCaller     // Read-only binding to the contract
	AbsorbableTransactor // Write-only binding to the contract
	AbsorbableFilterer   // Log filterer for contract events
}

// AbsorbableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsorbableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsorbableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsorbableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsorbableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsorbableSession struct {
	Contract     *Absorbable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsorbableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsorbableCallerSession struct {
	Contract *AbsorbableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AbsorbableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsorbableTransactorSession struct {
	Contract     *AbsorbableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AbsorbableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsorbableRaw struct {
	Contract *Absorbable // Generic contract binding to access the raw methods on
}

// AbsorbableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsorbableCallerRaw struct {
	Contract *AbsorbableCaller // Generic read-only contract binding to access the raw methods on
}

// AbsorbableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsorbableTransactorRaw struct {
	Contract *AbsorbableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsorbable creates a new instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbable(address common.Address, backend bind.ContractBackend) (*Absorbable, error) {
	contract, err := bindAbsorbable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Absorbable{AbsorbableCaller: AbsorbableCaller{contract: contract}, AbsorbableTransactor: AbsorbableTransactor{contract: contract}, AbsorbableFilterer: AbsorbableFilterer{contract: contract}}, nil
}

// NewAbsorbableCaller creates a new read-only instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableCaller(address common.Address, caller bind.ContractCaller) (*AbsorbableCaller, error) {
	contract, err := bindAbsorbable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsorbableCaller{contract: contract}, nil
}

// NewAbsorbableTransactor creates a new write-only instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsorbableTransactor, error) {
	contract, err := bindAbsorbable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsorbableTransactor{contract: contract}, nil
}

// NewAbsorbableFilterer creates a new log filterer instance of Absorbable, bound to a specific deployed contract.
func NewAbsorbableFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsorbableFilterer, error) {
	contract, err := bindAbsorbable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsorbableFilterer{contract: contract}, nil
}

// bindAbsorbable binds a generic wrapper to an already deployed contract.
func bindAbsorbable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsorbableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absorbable *AbsorbableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absorbable.Contract.AbsorbableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absorbable *AbsorbableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absorbable.Contract.AbsorbableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absorbable *AbsorbableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absorbable.Contract.AbsorbableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absorbable *AbsorbableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absorbable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absorbable *AbsorbableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absorbable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absorbable *AbsorbableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absorbable.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Absorbable *AbsorbableCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Absorbable *AbsorbableSession) Ask() (bool, error) {
	return _Absorbable.Contract.Ask(&_Absorbable.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Absorbable *AbsorbableCallerSession) Ask() (bool, error) {
	return _Absorbable.Contract.Ask(&_Absorbable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Absorbable *AbsorbableCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Absorbable *AbsorbableSession) Bid() (bool, error) {
	return _Absorbable.Contract.Bid(&_Absorbable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Absorbable *AbsorbableCallerSession) Bid() (bool, error) {
	return _Absorbable.Contract.Bid(&_Absorbable.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.CalcOrderID(&_Absorbable.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.CalcOrderID(&_Absorbable.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.FindAssistingID(&_Absorbable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.FindAssistingID(&_Absorbable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Absorbable *AbsorbableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Absorbable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Absorbable *AbsorbableSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Absorbable.Contract.GetOrder(&_Absorbable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Absorbable *AbsorbableCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Absorbable.Contract.GetOrder(&_Absorbable.CallOpts, _orderType, _id)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Absorbable *AbsorbableCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Absorbable.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Absorbable *AbsorbableSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Absorbable.Contract.GetRemainToAbsorb(&_Absorbable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Absorbable *AbsorbableCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Absorbable.Contract.GetRemainToAbsorb(&_Absorbable.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Next(&_Absorbable.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Next(&_Absorbable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Prev(&_Absorbable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Absorbable.Contract.Prev(&_Absorbable.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Absorbable *AbsorbableCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Absorbable.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Absorbable *AbsorbableSession) Top(orderType bool) ([32]byte, error) {
	return _Absorbable.Contract.Top(&_Absorbable.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Absorbable *AbsorbableCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Absorbable.Contract.Top(&_Absorbable.CallOpts, orderType)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.Contract.Cancel(&_Absorbable.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Absorbable *AbsorbableTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Absorbable.Contract.Cancel(&_Absorbable.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Absorbable.Contract.OnBlockInitialized(&_Absorbable.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Absorbable *AbsorbableTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Absorbable.Contract.OnBlockInitialized(&_Absorbable.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.Contract.RegisterTokens(&_Absorbable.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Absorbable *AbsorbableTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Absorbable.Contract.RegisterTokens(&_Absorbable.TransactOpts, volatileToken, stablizeToken)
}

// ITokenABI is the input ABI used to generate the binding from.
const ITokenABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"dexBurn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"dexMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ITokenFuncSigs maps the 4-byte function signature to its string representation.
var ITokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"692058c2": "dex()",
	"117f5a55": "dexBurn(uint256)",
	"bdfde911": "dexMint(uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IToken is an auto generated Go binding around an Ethereum contract.
type IToken struct {
	ITokenCaller     // Read-only binding to the contract
	ITokenTransactor // Write-only binding to the contract
	ITokenFilterer   // Log filterer for contract events
}

// ITokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSession struct {
	Contract     *IToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenCallerSession struct {
	Contract *ITokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransactorSession struct {
	Contract     *ITokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenRaw struct {
	Contract *IToken // Generic contract binding to access the raw methods on
}

// ITokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenCallerRaw struct {
	Contract *ITokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransactorRaw struct {
	Contract *ITokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIToken creates a new instance of IToken, bound to a specific deployed contract.
func NewIToken(address common.Address, backend bind.ContractBackend) (*IToken, error) {
	contract, err := bindIToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IToken{ITokenCaller: ITokenCaller{contract: contract}, ITokenTransactor: ITokenTransactor{contract: contract}, ITokenFilterer: ITokenFilterer{contract: contract}}, nil
}

// NewITokenCaller creates a new read-only instance of IToken, bound to a specific deployed contract.
func NewITokenCaller(address common.Address, caller bind.ContractCaller) (*ITokenCaller, error) {
	contract, err := bindIToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenCaller{contract: contract}, nil
}

// NewITokenTransactor creates a new write-only instance of IToken, bound to a specific deployed contract.
func NewITokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransactor, error) {
	contract, err := bindIToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransactor{contract: contract}, nil
}

// NewITokenFilterer creates a new log filterer instance of IToken, bound to a specific deployed contract.
func NewITokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFilterer, error) {
	contract, err := bindIToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFilterer{contract: contract}, nil
}

// bindIToken binds a generic wrapper to an already deployed contract.
func bindIToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.ITokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IToken *ITokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IToken *ITokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IToken *ITokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IToken *ITokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IToken *ITokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_IToken *ITokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IToken *ITokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IToken *ITokenSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IToken *ITokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, value)
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenTransactor) Dex(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dex")
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenSession) Dex() (*types.Transaction, error) {
	return _IToken.Contract.Dex(&_IToken.TransactOpts)
}

// Dex is a paid mutator transaction binding the contract method 0x692058c2.
//
// Solidity: function dex() returns(address)
func (_IToken *ITokenTransactorSession) Dex() (*types.Transaction, error) {
	return _IToken.Contract.Dex(&_IToken.TransactOpts)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 _amount) returns()
func (_IToken *ITokenTransactor) DexBurn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dexBurn", _amount)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 _amount) returns()
func (_IToken *ITokenSession) DexBurn(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexBurn(&_IToken.TransactOpts, _amount)
}

// DexBurn is a paid mutator transaction binding the contract method 0x117f5a55.
//
// Solidity: function dexBurn(uint256 _amount) returns()
func (_IToken *ITokenTransactorSession) DexBurn(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexBurn(&_IToken.TransactOpts, _amount)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 _amount) returns()
func (_IToken *ITokenTransactor) DexMint(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "dexMint", _amount)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 _amount) returns()
func (_IToken *ITokenSession) DexMint(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexMint(&_IToken.TransactOpts, _amount)
}

// DexMint is a paid mutator transaction binding the contract method 0xbdfde911.
//
// Solidity: function dexMint(uint256 _amount) returns()
func (_IToken *ITokenTransactorSession) DexMint(_amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DexMint(&_IToken.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IToken *ITokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, from, to, value)
}

// ITokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IToken contract.
type ITokenApprovalIterator struct {
	Event *ITokenApproval // Event containing the contract specifics and raw log

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
func (it *ITokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenApproval)
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
		it.Event = new(ITokenApproval)
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
func (it *ITokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenApproval represents a Approval event raised by the IToken contract.
type ITokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ITokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITokenApprovalIterator{contract: _IToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenApproval)
				if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) ParseApproval(log types.Log) (*ITokenApproval, error) {
	event := new(ITokenApproval)
	if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ITokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IToken contract.
type ITokenTransferIterator struct {
	Event *ITokenTransfer // Event containing the contract specifics and raw log

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
func (it *ITokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenTransfer)
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
		it.Event = new(ITokenTransfer)
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
func (it *ITokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenTransfer represents a Transfer event raised by the IToken contract.
type ITokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ITokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferIterator{contract: _IToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ITokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenTransfer)
				if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) ParseTransfer(log types.Log) (*ITokenTransfer, error) {
	event := new(ITokenTransfer)
	if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d95b9eea3487e49140e17da6d9bf93ebdf65b0a7324ff4934fcb59df821e193764736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OrderbookABI is the input ABI used to generate the binding from.
const OrderbookABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OrderbookFuncSigs maps the 4-byte function signature to its string representation.
var OrderbookFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"4ea09797": "next(bool,bytes32)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"8aa3f897": "top(bool)",
}

// OrderbookBin is the compiled bytecode used for deploying new contracts.
var OrderbookBin = "0x608060405234801561001057600080fd5b50610953806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636e6452cb116100665780636e6452cb1461019c5780638aa3f897146101a4578063aa1c259c146101c3578063ced4aac8146101f1578063f318722b146102315761009e565b806307c399a3146100a35780630d90b10a146100fd57806343271d79146101345780634ea097971461015b57806369c07d3114610180575b600080fd5b6100c8600480360360408110156100b957600080fd5b5080351515906020013561025d565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b6101226004803603604081101561011357600080fd5b508035151590602001356102a4565b60408051918252519081900360200190f35b6101596004803603604081101561014a57600080fd5b508035151590602001356102c8565b005b6101226004803603604081101561017157600080fd5b5080351515906020013561034e565b610188610372565b604080519115158252519081900360200190f35b610188610377565b610122600480360360208110156101ba57600080fd5b5035151561037c565b610159600480360360408110156101d957600080fd5b506001600160a01b038135811691602001351661039d565b610122600480360360a081101561020757600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610415565b6101226004803603604081101561024757600080fd5b506001600160a01b03813516906020013561047d565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b90151560009081526020818152604080832093835260029093019052206003015490565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610338576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610348828463ffffffff61048916565b50505050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b801515600090815260208190526040812061039681610548565b9392505050565b60008080526020526103d67fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5838363ffffffff61056116565b600160009081526020526104117fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d828463ffffffff61056116565b5050565b841515600090815260208190526040812061042e6108cb565b506040805160a0810182526001600160a01b038816815260208101879052908101859052600060608201819052608082015261047182828663ffffffff61070e16565b98975050505050505050565b60006103968383610780565b600081815260028301602052604090206001015415610538578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b15801561050b57600080fd5b505af115801561051f573d6000803e3d6000fd5b505050506040513d602081101561053557600080fd5b50505b610411828263ffffffff61084d16565b6000808052600282016020526040902060040154919050565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250610746848263ffffffff6108ad16565b1561071f575b6003015460008181526002860160205260409020909250610773848263ffffffff6108ad16565b61074c5750909392505050565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106107f05780518252601f1990920191602091820191016107d1565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561082f573d6000803e3d6000fd5b5050506040513d602081101561084457600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60408201516001820154600290920154602090930151910291021190565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fea265627a7a72315820b0b237c243311bd2df1d406a0c548026cf00e9c5d645cf095f4946d130263b3164736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployOrderbook deploys a new Ethereum contract, binding an instance of Orderbook to it.
func DeployOrderbook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Orderbook, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderbookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// Orderbook is an auto generated Go binding around an Ethereum contract.
type Orderbook struct {
	OrderbookCaller     // Read-only binding to the contract
	OrderbookTransactor // Write-only binding to the contract
	OrderbookFilterer   // Log filterer for contract events
}

// OrderbookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderbookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderbookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderbookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderbookSession struct {
	Contract     *Orderbook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderbookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderbookCallerSession struct {
	Contract *OrderbookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderbookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderbookTransactorSession struct {
	Contract     *OrderbookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderbookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderbookRaw struct {
	Contract *Orderbook // Generic contract binding to access the raw methods on
}

// OrderbookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderbookCallerRaw struct {
	Contract *OrderbookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderbookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderbookTransactorRaw struct {
	Contract *OrderbookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderbook creates a new instance of Orderbook, bound to a specific deployed contract.
func NewOrderbook(address common.Address, backend bind.ContractBackend) (*Orderbook, error) {
	contract, err := bindOrderbook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// NewOrderbookCaller creates a new read-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookCaller(address common.Address, caller bind.ContractCaller) (*OrderbookCaller, error) {
	contract, err := bindOrderbook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookCaller{contract: contract}, nil
}

// NewOrderbookTransactor creates a new write-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderbookTransactor, error) {
	contract, err := bindOrderbook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookTransactor{contract: contract}, nil
}

// NewOrderbookFilterer creates a new log filterer instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderbookFilterer, error) {
	contract, err := bindOrderbook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderbookFilterer{contract: contract}, nil
}

// bindOrderbook binds a generic wrapper to an already deployed contract.
func bindOrderbook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.OrderbookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Orderbook *OrderbookCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Orderbook *OrderbookSession) Ask() (bool, error) {
	return _Orderbook.Contract.Ask(&_Orderbook.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Orderbook *OrderbookCallerSession) Ask() (bool, error) {
	return _Orderbook.Contract.Ask(&_Orderbook.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Orderbook *OrderbookCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Orderbook *OrderbookSession) Bid() (bool, error) {
	return _Orderbook.Contract.Bid(&_Orderbook.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Orderbook *OrderbookCallerSession) Bid() (bool, error) {
	return _Orderbook.Contract.Bid(&_Orderbook.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Orderbook *OrderbookSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.CalcOrderID(&_Orderbook.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.CalcOrderID(&_Orderbook.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Orderbook *OrderbookSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.FindAssistingID(&_Orderbook.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.FindAssistingID(&_Orderbook.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Orderbook *OrderbookCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Orderbook.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Orderbook *OrderbookSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Orderbook *OrderbookCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _orderType, _id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Next(&_Orderbook.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Next(&_Orderbook.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Prev(&_Orderbook.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Orderbook.Contract.Prev(&_Orderbook.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Orderbook *OrderbookSession) Top(orderType bool) ([32]byte, error) {
	return _Orderbook.Contract.Top(&_Orderbook.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Orderbook.Contract.Top(&_Orderbook.CallOpts, orderType)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.Cancel(&_Orderbook.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Orderbook *OrderbookTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.Cancel(&_Orderbook.TransactOpts, orderType, id)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.RegisterTokens(&_Orderbook.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Orderbook *OrderbookTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.RegisterTokens(&_Orderbook.TransactOpts, volatileToken, stablizeToken)
}

// PreemptivableABI is the input ABI used to generate the binding from.
const PreemptivableABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PreemptivableFuncSigs maps the 4-byte function signature to its string representation.
var PreemptivableFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"c7f758a8": "getProposal(uint256)",
	"c08cc02d": "getProposalCount()",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"74a8f103": "revoke(address)",
	"c0ee0b8a": "tokenFallback(address,uint256,bytes)",
	"8aa3f897": "top(bool)",
	"4def5645": "totalVote(address)",
	"bd041c4d": "vote(address,bool)",
}

// PreemptivableBin is the compiled bytecode used for deploying new contracts.
var PreemptivableBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562127500600e556002600e54816200003057fe5b04600f5560006010553480156200004657600080fd5b506040516200361738038062003617833981810160405260808110156200006c57600080fd5b508051602082015160408301516060909301519192909183838015620000925760038190555b60008211620000a55760028104620000a7565b815b60045550508015620000b957600e8190555b60008211620000d6576002600e5481620000cf57fe5b04620000d8565b815b600f555050505061352880620000ef6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063aa1c259c116100a2578063c0ee0b8a11610071578063c0ee0b8a14610308578063c7f758a81461038d578063ced4aac8146103e7578063ee1a68c614610427578063f318722b1461044a57610116565b8063aa1c259c14610287578063bd041c4d146102b5578063be91d729146102e3578063c08cc02d1461030057610116565b80634ea09797116100e95780634ea09797146101f957806369c07d311461021e5780636e6452cb1461023a57806374a8f103146102425780638aa3f8971461026857610116565b806307c399a31461011b5780630d90b10a1461017557806343271d79146101ac5780634def5645146101d3575b600080fd5b6101406004803603604081101561013157600080fd5b50803515159060200135610476565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61019a6004803603604081101561018b57600080fd5b508035151590602001356104bd565b60408051918252519081900360200190f35b6101d1600480360360408110156101c257600080fd5b508035151590602001356104e5565b005b61019a600480360360208110156101e957600080fd5b50356001600160a01b031661056b565b61019a6004803603604081101561020f57600080fd5b50803515159060200135610593565b6102266105b7565b604080519115158252519081900360200190f35b6102266105bc565b6101d16004803603602081101561025857600080fd5b50356001600160a01b03166105c1565b61019a6004803603602081101561027e57600080fd5b50351515610729565b6101d16004803603604081101561029d57600080fd5b506001600160a01b0381358116916020013516610743565b6101d1600480360360408110156102cb57600080fd5b506001600160a01b0381351690602001351515610892565b6101d1600480360360208110156102f957600080fd5b503561090c565b61019a6109aa565b6101d16004803603606081101561031e57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561034e57600080fd5b82018360208201111561036057600080fd5b8035906020019184600183028401116401000000008311171561038257600080fd5b5090925090506109bc565b6103aa600480360360208110156103a357600080fd5b5035610adf565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61019a600480360360a08110156103fd57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610b33565b61042f610b9b565b60408051921515835260208301919091528051918290030190f35b61019a6004803603604081101561046057600080fd5b506001600160a01b038135169060200135610c45565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610555576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610565828463ffffffff610c5816565b50505050565b60008061057f60118463ffffffff610d1716565b905061058a81610d59565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b60006105d460118363ffffffff610d1716565b80549091506001600160a01b03838116911614610638576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b601380546001810180835560009290925260068301805490916003027f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a09001906106849082908490613219565b5050600180548454918501546040805163a9059cbb60e01b81526001600160a01b03948516600482015260248101929092525192909116935063a9059cbb92506044808201926020929091908290030181600087803b1580156106e657600080fd5b505af11580156106fa573d6000803e3d6000fd5b505050506040513d602081101561071057600080fd5b50610724905060118363ffffffff610e3f16565b505050565b801515600090815260208190526040812061058a81610e87565b6001546001600160a01b03161561078f576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b6002546001600160a01b0316156107db576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556108158282610ea0565b600254604080516318160ddd60e01b8152905161088e926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561085b57600080fd5b505afa15801561086f573d6000803e3d6000fd5b505050506040513d602081101561088557600080fd5b50516000610ef0565b5050565b6108a360118363ffffffff610fbc16565b6108e7576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b60006108fa60118463ffffffff610d1716565b9050610724818363ffffffff610fdd16565b3315610950576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b601354811015610988576109806013828154811061096d57fe5b9060005260206000209060030201610ff1565b600101610953565b5061099560136000613265565b61099d61105e565b506109a7816110ac565b50565b60006109b66011611245565b90505b90565b6080811480156109d657506001546001600160a01b031633145b15610a79576109ec60118563ffffffff610fbc16565b15610a37576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610a4d57600080fd5b50803594506020810135935060400135915060009050610a708888868686611249565b50505050610565565b600080806060841415610aad5784846060811015610a9657600080fd5b508035935060208101359250604001359050610ac9565b84846040811015610abd57600080fd5b50803593506020013591505b610ad6878488858561135c565b50505050505050565b6000808080808080610af860118963ffffffff6113b716565b805460018201546002830154600484015460038501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610b4c613286565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610b8f82828663ffffffff6113c816565b98975050505050505050565b6007546000908190610bb257506000905080610c41565b6001610c3c600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c0b57600080fd5b505afa158015610c1f573d6000803e3d6000fd5b505050506040513d6020811015610c3557600080fd5b505161143a565b915091505b9091565b6000610c518383611454565b9392505050565b600081815260028301602052604090206001015415610d07578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b158015610cda57600080fd5b505af1158015610cee573d6000803e3d6000fd5b505050506040513d6020811015610d0457600080fd5b50505b61088e828263ffffffff61152116565b6001600160a01b03811660009081526001830160205260408120548354849160001901908110610d4357fe5b9060005260206000209060090201905092915050565b600080805b610d6a84600601611245565b811015610e3857600080610d87600687018463ffffffff61158116565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610ddb57600080fd5b505afa158015610def573d6000803e3d6000fd5b505050506040513d6020811015610e0557600080fd5b50516001600160a01b038416310190508115610e245793840193610e2a565b80850394505b505050806001019050610d5e565b5092915050565b6001600160a01b038116600090815260018301602052604081205480610e695760009150506104df565b610e7d84600019830163ffffffff6115c016565b5060019392505050565b6000808052600282016020526040902060040154919050565b6000808052602052610ec76000805160206134af833981519152838363ffffffff61177616565b6001600090815260205261088e60008051602061348f833981519152828463ffffffff61177616565b604051806080016040528060035443018152602001600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f5357600080fd5b505afa158015610f67573d6000803e3d6000fd5b505050506040513d6020811015610f7d57600080fd5b50518152602080820194909452911515604092830152805160055591820151600655810151600755606001516008805460ff1916911515919091179055565b6001600160a01b031660009081526001919091016020526040902054151590565b61072460068301338363ffffffff61192316565b60005b815481101561105257600082600001828154811061100e57fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff19169055919091019050610ff4565b506109a78160006132b4565b600061106a6009611a3f565b15611077575060006109b9565b6000611081611a5b565b90506001600160a01b03811661109b5760009150506109b9565b6110a481611ae0565b600191505090565b33156110f0576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6110fa6005611c12565b15611119576000600581905560068190556007556008805460ff191690555b6111236009611c2b565b1561113057611130611c48565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561117557600080fd5b505afa158015611189573d6000803e3d6000fd5b505050506040513d602081101561119f57600080fd5b505190508115611213576111b1611d13565b806111c157506111c18183611d1f565b156111d1576111d1826000610ef0565b6111db6009611a3f565b156112135760006111ec838361143a565b90506111f781611dcd565b8015611205575060085460ff165b156112115750506109a7565b505b61122460058263ffffffff611eba16565b1561088e576000611233611eef565b905061123e81611fa4565b5050505050565b5490565b6112516132d2565b6001600160a01b038616815260208101859052604081018490524360a082015282156112e0576003600f548161128357fe5b04600f540183116112db576040805162461bcd60e51b815260206004820181905260248201527f736c617368696e67206475726174696f6e20706172616d20746f6f206c6f6e67604482015290519081900360640190fd5b6112e9565b600f5460808201525b8115611342576003600e54816112fb57fe5b04600e5403821061133d5760405162461bcd60e51b815260040180806020018281038252602181526020018061346e6021913960400191505060405180910390fd5b61134b565b600e5460608201525b610ad660118263ffffffff61204b16565b600061136733612198565b9050600061137e828888888863ffffffff61227b16565b905061139182828563ffffffff61241316565b506113ad8161139f33612451565b84919063ffffffff61250b16565b5050505050505050565b6000826000018281548110610d4357fe5b600081815260028401602052604081205b6004015460008181526002860160205260409020909250611400848263ffffffff6126a116565b156113d9575b600301546000818152600286016020526040902090925061142d848263ffffffff6126a116565b6114065750909392505050565b600081831161144e57828203600003610c51565b50900390565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106114c45780518252601f1990920191602091820191016114a5565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611503573d6000803e3d6000fd5b5050506040513d602081101561151857600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60008080611595858563ffffffff6126bf16565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b8160010160008360000183815481106115d557fe5b600091825260208083206009909202909101546001600160a01b0316835282019290925260400181205581546000190181146116fe5781548290600019810190811061161d57fe5b906000526020600020906009020182600001828154811061163a57fe5b60009182526020909120825460099092020180546001600160a01b0319166001600160a01b0390921691909117815560018083015490820155600280830154908201556003808301549082015560048083015490820155600580830154908201556006808301805490918301906116b49082908490613219565b505050905050806001018260010160008460000184815481106116d357fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040019020555b81548290600019810190811061171057fe5b60009182526020822060099091020180546001600160a01b0319168155600181018290556002810182905560038101829055600481018290556005810182905590600682018161176082826132b4565b505083549150610724905083600019830161331d565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff191685151517905560018601909152812054806119a6575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610c51565b84548111806119e45750836001600160a01b03168560000160018303815481106119cc57fe5b6000918252602090912001546001600160a01b031614155b15611a34575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610c51565b506000949350505050565b6000611a4a826126ec565b80156104df57505060040154431090565b60008080805b611a6b6011611245565b811015611ad9576000611a8560118363ffffffff6113b716565b90506004600e5481611a9357fe5b04816005015443031015611aa75750611ad1565b6000611ab2826126fb565b905084811315611ace5781549094506001600160a01b03169250835b50505b600101611a61565b5091505090565b6000611af360118363ffffffff610d1716565b9050611b0181600601610ff1565b6040805160a08101825282546001600160a01b03168082526002840154602083018190526001850154938301849052600485015460608401819052600386015443016080909401849052600980546001600160a01b031916909317909255600a55600b92909255600c91909155600d55611b8260118363ffffffff610e3f16565b5061088e611c0b600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015611bd757600080fd5b505afa158015611beb573d6000803e3d6000fd5b505050506040513d6020811015611c0157600080fd5b5051600a54612732565b6001610ef0565b6000611c1d8261274e565b80156104df57505054431190565b6000611c36826126ec565b80156104df5750506004015443101590565b611c5260096126ec565b611c5b57611d11565b600b5415611cec57600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015611cbf57600080fd5b505af1158015611cd3573d6000803e3d6000fd5b505050506040513d6020811015611ce957600080fd5b50505b600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b60006109b66005611c12565b600082821415611d31575060006104df565b6006546007541415611d45575060016104df565b82821115611d9957600654600754848403911015611d7b57600654600754036002818381611d6f57fe5b041015925050506104df565b600754600654036002828281611d8d57fe5b041115925050506104df565b600654600754838503911115611dbb57600754600654036002818381611d6f57fe5b600654600754036002828281611d8d57fe5b6000611de0600960010154600084612754565b611dec5750600061058e565b600c54600a546000919084830381611e0057fe5b0581611e0857fe5b04905080611e14575060015b600b54811115611e435750600b546000600581905560068190556007556008805460ff19169055611e43611c48565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015611e9957600080fd5b505af1158015611ead573d6000803e3d6000fd5b5060019695505050505050565b6000611ec58361274e565b8015611ed5575082600201548214155b8015610c515750610c518360010154838560020154612784565b600080611f0660056002015460056001015461143a565b600754600254604080516318160ddd60e01b81529051939450600093611f5693926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610c0b57600080fd5b9050611f6460008284612754565b611f73576000925050506109b9565b60006004548381611f8057fe5b059050611f8f60008284612754565b611f9d575091506109b99050565b9250505090565b6000806000806000808613611fba576001611fbd565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff16156120245761201981611ffe876127af565b600954859291906001600160a01b031663ffffffff6127c516565b935093505050612046565b61203f81612031876127af565b84919063ffffffff612b7816565b9350935050505b915091565b80516001600160a01b0316600090815260018301602052604081205480156120ba576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b835460018082018087556000878152602090819020875160099095020180546001600160a01b0319166001600160a01b03909516949094178455868101519284019290925560408601516002840155606086015160038401556080860151600484015560a0860151600584015560c0860151805180519294889490936006850192612149928492910190613349565b505086516001600160a01b03166000908152600189016020526040902084905550508554919250859160001984019150811061218157fe5b906000526020600020906009020191505092915050565b600080805260208190526000805160206134af833981519152546001600160a01b03838116911614156121e2575060008080526020526000805160206134af83398151915261058e565b6001600090815260205260008051602061348f833981519152546001600160a01b038381169116141561222e57506001600090815260205260008051602061348f83398151915261058e565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b6000808311801561228c5750600082115b6122ca576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b831080156122e05750600160801b82105b612324576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b60006123308686611454565b6000818152600289016020526040902090915061234c816126ec565b15612393576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b506040805160a0810182526001600160a01b03978816815260208082019687528183019586526000606083018181526080840182815286835260029c8d01909352939020915182546001600160a01b0319169916989098178155945160018601559251968401969096555051600382015591516004909201919091555090565b600082815260028401602052604081208161243586838663ffffffff612e4516565b905061244886868363ffffffff612eb716565b95945050505050565b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b03838116911614156124ad575060008080526020526000805160206134af83398151915261058e565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b038381169116141561222e57506001600090815260205260008051602061348f83398151915261058e565b600082815260028401602052604081208161252584610e87565b90505b60001981146124485761253a82612ef2565b1561254457612448565b600081815260028086016020908152604092839020835160a08101855286546001600160a01b0316815260018701549281019290925291850154928101929092526003840154606083015260048401546080830152906125aa908263ffffffff612f0b16565b6125b45750612448565b80600201548360010154101561266557600081600201548260010154856001015402816125dd57fe5b049050816001015481111561262b576040805162461bcd60e51b815260206004820152600f60248201526e66696c6c61626c65203e206861766560881b604482015290519081900360640190fd5b60018401546126459087908590849063ffffffff612f2a16565b600184015461265e90899089908463ffffffff612f2a16565b5050612448565b6002810154600182015461268391899189919063ffffffff612f2a16565b6004810154612698868463ffffffff6130c216565b91506125289050565b60408201516001820154600290920154602090930151910291021190565b60008260000182815481106126d057fe5b6000918252602090912001546001600160a01b03169392505050565b546001600160a01b0316151590565b60008061270783610d59565b90506000811361271b57600091505061058e565b61272483610d59565b836001015402915050919050565b6000808212156127495781600003830390506104df565b500190565b54151590565b60008284131580156127665750818313155b8061277c575082841215801561277c5750818312155b949350505050565b60008284111580156127965750818311155b8061277c575082841015801561277c5750501115919050565b60008082136127c157816000036104df565b5090565b6000806127d986868663ffffffff612b7816565b9092509050600080866127ed5783836127f0565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b15801561284657600080fd5b505afa15801561285a573d6000803e3d6000fd5b505050506040513d602081101561287057600080fd5b505182118015906128f857508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b1580156128c857600080fd5b505afa1580156128dc573d6000803e3d6000fd5b505050506040513d60208110156128f257600080fd5b50518211155b15612b6d578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b15801561294c57600080fd5b505af1158015612960573d6000803e3d6000fd5b505050506040513d602081101561297657600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b1580156129cf57600080fd5b505af11580156129e3573d6000803e3d6000fd5b505050506040513d60208110156129f957600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612a4757600080fd5b505af1158015612a5b573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612aad57600080fd5b505af1158015612ac1573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015612b1d57600080fd5b505af1158015612b31573d6000803e3d6000fd5b505050506040513d6020811015612b4757600080fd5b50879050612b555780612b57565b815b8301925086612b665781612b68565b805b840193505b505094509492505050565b6000806000612b8686610e87565b90505b6000198114801590612b9a57508382105b15612e3c57600081815260028701602052604081209086612bbf578160020154612bc5565b81600101545b9050600087612bd8578260010154612bde565b82600201545b905086612bf1868463ffffffff61314716565b11612ce057885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612c4757600080fd5b505af1158015612c5b573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612cb257600080fd5b505af1158015612cc6573d6000803e3d6000fd5b505050506004830154612cd98a866130c2565b9350612e10565b6000612cf2888763ffffffff6131a116565b90508281830281612cff57fe5b049150809250600089612d125782612d14565b835b905060008a612d235784612d25565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b158015612d7457600080fd5b505af1158015612d88573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612dda57600080fd5b505af1158015612dee573d6000803e3d6000fd5b50612e0792508e9150899050848463ffffffff612f2a16565b50600019955050505b612e20858363ffffffff61314716565b9450612e32868263ffffffff61314716565b9550505050612b89565b50935093915050565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250612e7d848263ffffffff6131f816565b15612e56575b6003015460008181526002860160205260409020909250612eaa848263ffffffff6131f816565b612e835750909392505050565b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b60008160010154600014806104df575050600201541590565b6040820151600282015460019092015460209093015191029102101590565b60008381526002850160205260409020612f43816126ec565b612f86576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b8060010154831115612fdf576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b6001810154612ff4908463ffffffff6131a116565b6001820155600281015482101561301557600281018054839003905561301d565b600060028201555b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561307757600080fd5b505af115801561308b573d6000803e3d6000fd5b505050506040513d60208110156130a157600080fd5b506130ad905081612ef2565b1561123e5761123e858563ffffffff610c5816565b600081815260028084016020526040909120015415610d0757600182015460008281526002808501602090815260408084208054930154815163a9059cbb60e01b81526001600160a01b03948516600482015260248101919091529051929094169363a9059cbb93604480830194928390030190829087803b158015610cda57600080fd5b600082820183811015610c51576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60008282111561144e576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b60008260020154826001015402826002015484600101540211905092915050565b8280548282559060005260206000209081019282156132595760005260206000209182015b8281111561325957825482559160010191906001019061323e565b506127c192915061339e565b50805460008255600302906000526020600020908101906109a791906133c2565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b50805460008255906000526020600020908101906109a791906133e5565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016133186133ff565b905290565b815481835581811115610724576009028160090283600052602060002091820191016107249190613412565b828054828255906000526020600020908101928215613259579160200282015b8281111561325957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613369565b6109b991905b808211156127c15780546001600160a01b03191681556001016133a4565b6109b991905b808211156127c15760006133dc82826132b4565b506003016133c8565b6109b991905b808211156127c157600081556001016133eb565b6040518060200160405280606081525090565b6109b991905b808211156127c15780546001600160a01b031916815560006001820181905560028201819055600382018190556004820181905560058201819055600682018161346282826132b4565b50505060090161341856fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a723158206f973c576824e6ccb5727db634a6553251fa8641d45a3cd8a1094f1535123b9a64736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployPreemptivable deploys a new Ethereum contract, binding an instance of Preemptivable to it.
func DeployPreemptivable(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int, initialSlashingDuration *big.Int, initialLockdownExpiration *big.Int) (common.Address, *types.Transaction, *Preemptivable, error) {
	parsed, err := abi.JSON(strings.NewReader(PreemptivableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PreemptivableBin), backend, absorptionDuration, absorptionExpiration, initialSlashingDuration, initialLockdownExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Preemptivable{PreemptivableCaller: PreemptivableCaller{contract: contract}, PreemptivableTransactor: PreemptivableTransactor{contract: contract}, PreemptivableFilterer: PreemptivableFilterer{contract: contract}}, nil
}

// Preemptivable is an auto generated Go binding around an Ethereum contract.
type Preemptivable struct {
	PreemptivableCaller     // Read-only binding to the contract
	PreemptivableTransactor // Write-only binding to the contract
	PreemptivableFilterer   // Log filterer for contract events
}

// PreemptivableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreemptivableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreemptivableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreemptivableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreemptivableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreemptivableSession struct {
	Contract     *Preemptivable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreemptivableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreemptivableCallerSession struct {
	Contract *PreemptivableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PreemptivableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreemptivableTransactorSession struct {
	Contract     *PreemptivableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PreemptivableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreemptivableRaw struct {
	Contract *Preemptivable // Generic contract binding to access the raw methods on
}

// PreemptivableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreemptivableCallerRaw struct {
	Contract *PreemptivableCaller // Generic read-only contract binding to access the raw methods on
}

// PreemptivableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreemptivableTransactorRaw struct {
	Contract *PreemptivableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreemptivable creates a new instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivable(address common.Address, backend bind.ContractBackend) (*Preemptivable, error) {
	contract, err := bindPreemptivable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Preemptivable{PreemptivableCaller: PreemptivableCaller{contract: contract}, PreemptivableTransactor: PreemptivableTransactor{contract: contract}, PreemptivableFilterer: PreemptivableFilterer{contract: contract}}, nil
}

// NewPreemptivableCaller creates a new read-only instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableCaller(address common.Address, caller bind.ContractCaller) (*PreemptivableCaller, error) {
	contract, err := bindPreemptivable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreemptivableCaller{contract: contract}, nil
}

// NewPreemptivableTransactor creates a new write-only instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableTransactor(address common.Address, transactor bind.ContractTransactor) (*PreemptivableTransactor, error) {
	contract, err := bindPreemptivable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreemptivableTransactor{contract: contract}, nil
}

// NewPreemptivableFilterer creates a new log filterer instance of Preemptivable, bound to a specific deployed contract.
func NewPreemptivableFilterer(address common.Address, filterer bind.ContractFilterer) (*PreemptivableFilterer, error) {
	contract, err := bindPreemptivable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreemptivableFilterer{contract: contract}, nil
}

// bindPreemptivable binds a generic wrapper to an already deployed contract.
func bindPreemptivable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PreemptivableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preemptivable *PreemptivableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Preemptivable.Contract.PreemptivableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preemptivable *PreemptivableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preemptivable.Contract.PreemptivableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preemptivable *PreemptivableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preemptivable.Contract.PreemptivableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preemptivable *PreemptivableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Preemptivable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preemptivable *PreemptivableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preemptivable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preemptivable *PreemptivableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preemptivable.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Preemptivable *PreemptivableCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Preemptivable *PreemptivableSession) Ask() (bool, error) {
	return _Preemptivable.Contract.Ask(&_Preemptivable.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Preemptivable *PreemptivableCallerSession) Ask() (bool, error) {
	return _Preemptivable.Contract.Ask(&_Preemptivable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Preemptivable *PreemptivableCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Preemptivable *PreemptivableSession) Bid() (bool, error) {
	return _Preemptivable.Contract.Bid(&_Preemptivable.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Preemptivable *PreemptivableCallerSession) Bid() (bool, error) {
	return _Preemptivable.Contract.Bid(&_Preemptivable.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.CalcOrderID(&_Preemptivable.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.CalcOrderID(&_Preemptivable.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.FindAssistingID(&_Preemptivable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.FindAssistingID(&_Preemptivable.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Preemptivable *PreemptivableCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Preemptivable.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Preemptivable *PreemptivableSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Preemptivable.Contract.GetOrder(&_Preemptivable.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Preemptivable *PreemptivableCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Preemptivable.Contract.GetOrder(&_Preemptivable.CallOpts, _orderType, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableCaller) GetProposal(opts *bind.CallOpts, idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	ret := new(struct {
		Maker              common.Address
		Stake              *big.Int
		Amount             *big.Int
		SlashingDuration   *big.Int
		LockdownExpiration *big.Int
		Number             *big.Int
	})
	out := ret
	err := _Preemptivable.contract.Call(opts, out, "getProposal", idx)
	return *ret, err
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Preemptivable.Contract.GetProposal(&_Preemptivable.CallOpts, idx)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Preemptivable *PreemptivableCallerSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Preemptivable.Contract.GetProposal(&_Preemptivable.CallOpts, idx)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Preemptivable *PreemptivableCaller) GetProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "getProposalCount")
	return *ret0, err
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Preemptivable *PreemptivableSession) GetProposalCount() (*big.Int, error) {
	return _Preemptivable.Contract.GetProposalCount(&_Preemptivable.CallOpts)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Preemptivable *PreemptivableCallerSession) GetProposalCount() (*big.Int, error) {
	return _Preemptivable.Contract.GetProposalCount(&_Preemptivable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Preemptivable *PreemptivableCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Preemptivable.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Preemptivable *PreemptivableSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Preemptivable.Contract.GetRemainToAbsorb(&_Preemptivable.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Preemptivable *PreemptivableCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Preemptivable.Contract.GetRemainToAbsorb(&_Preemptivable.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Next(&_Preemptivable.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Next(&_Preemptivable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Prev(&_Preemptivable.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Preemptivable.Contract.Prev(&_Preemptivable.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Preemptivable *PreemptivableCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Preemptivable *PreemptivableSession) Top(orderType bool) ([32]byte, error) {
	return _Preemptivable.Contract.Top(&_Preemptivable.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Preemptivable *PreemptivableCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Preemptivable.Contract.Top(&_Preemptivable.CallOpts, orderType)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Preemptivable *PreemptivableCaller) TotalVote(opts *bind.CallOpts, maker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Preemptivable.contract.Call(opts, out, "totalVote", maker)
	return *ret0, err
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Preemptivable *PreemptivableSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Preemptivable.Contract.TotalVote(&_Preemptivable.CallOpts, maker)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Preemptivable *PreemptivableCallerSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Preemptivable.Contract.TotalVote(&_Preemptivable.CallOpts, maker)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.Cancel(&_Preemptivable.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Preemptivable *PreemptivableTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.Cancel(&_Preemptivable.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.Contract.OnBlockInitialized(&_Preemptivable.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Preemptivable *PreemptivableTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Preemptivable.Contract.OnBlockInitialized(&_Preemptivable.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.RegisterTokens(&_Preemptivable.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Preemptivable *PreemptivableTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.RegisterTokens(&_Preemptivable.TransactOpts, volatileToken, stablizeToken)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableTransactor) Revoke(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "revoke", maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.Revoke(&_Preemptivable.TransactOpts, maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Preemptivable *PreemptivableTransactorSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Preemptivable.Contract.Revoke(&_Preemptivable.TransactOpts, maker)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableTransactor) TokenFallback(opts *bind.TransactOpts, maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "tokenFallback", maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.TokenFallback(&_Preemptivable.TransactOpts, maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Preemptivable *PreemptivableTransactorSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Preemptivable.Contract.TokenFallback(&_Preemptivable.TransactOpts, maker, value, data)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableTransactor) Vote(opts *bind.TransactOpts, maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.contract.Transact(opts, "vote", maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.Contract.Vote(&_Preemptivable.TransactOpts, maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Preemptivable *PreemptivableTransactorSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Preemptivable.Contract.Vote(&_Preemptivable.TransactOpts, maker, up)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158201aa9235f53fc564ce629d018739384ca1146bdc4cc3d008d4ffbe8bc28757b1764736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SeigniorageABI is the input ABI used to generate the binding from.
const SeigniorageABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"prev\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"totalVote\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"}],\"name\":\"top\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"volatileToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stablizeToken\",\"type\":\"address\"}],\"name\":\"registerTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"up\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"onBlockInitialized\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"slashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockdownExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"orderType\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wantAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"assistingID\",\"type\":\"bytes32\"}],\"name\":\"findAssistingID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainToAbsorb\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"calcOrderID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"absorptionDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"absorptionExpiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialSlashingDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLockdownExpiration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SeigniorageFuncSigs maps the 4-byte function signature to its string representation.
var SeigniorageFuncSigs = map[string]string{
	"69c07d31": "Ask()",
	"6e6452cb": "Bid()",
	"f318722b": "calcOrderID(address,bytes32)",
	"43271d79": "cancel(bool,bytes32)",
	"ced4aac8": "findAssistingID(bool,address,uint256,uint256,bytes32)",
	"07c399a3": "getOrder(bool,bytes32)",
	"c7f758a8": "getProposal(uint256)",
	"c08cc02d": "getProposalCount()",
	"ee1a68c6": "getRemainToAbsorb()",
	"4ea09797": "next(bool,bytes32)",
	"be91d729": "onBlockInitialized(uint256)",
	"0d90b10a": "prev(bool,bytes32)",
	"aa1c259c": "registerTokens(address,address)",
	"74a8f103": "revoke(address)",
	"c0ee0b8a": "tokenFallback(address,uint256,bytes)",
	"8aa3f897": "top(bool)",
	"4def5645": "totalVote(address)",
	"bd041c4d": "vote(address,bool)",
}

// SeigniorageBin is the compiled bytecode used for deploying new contracts.
var SeigniorageBin = "0x608060405262049d406003556002600354816200001857fe5b0460045562127500600e556002600e54816200003057fe5b04600f5560006010553480156200004657600080fd5b506040516200361f3803806200361f833981810160405260808110156200006c57600080fd5b50805160208201516040830151606090930151919290918383838383838015620000965760038190555b60008211620000a95760028104620000ab565b815b60045550508015620000bd57600e8190555b60008211620000da576002600e5481620000d357fe5b04620000dc565b815b600f55505050505050505061352880620000f76000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063aa1c259c116100a2578063c0ee0b8a11610071578063c0ee0b8a14610308578063c7f758a81461038d578063ced4aac8146103e7578063ee1a68c614610427578063f318722b1461044a57610116565b8063aa1c259c14610287578063bd041c4d146102b5578063be91d729146102e3578063c08cc02d1461030057610116565b80634ea09797116100e95780634ea09797146101f957806369c07d311461021e5780636e6452cb1461023a57806374a8f103146102425780638aa3f8971461026857610116565b806307c399a31461011b5780630d90b10a1461017557806343271d79146101ac5780634def5645146101d3575b600080fd5b6101406004803603604081101561013157600080fd5b50803515159060200135610476565b604080516001600160a01b03909616865260208601949094528484019290925260608401526080830152519081900360a00190f35b61019a6004803603604081101561018b57600080fd5b508035151590602001356104bd565b60408051918252519081900360200190f35b6101d1600480360360408110156101c257600080fd5b508035151590602001356104e5565b005b61019a600480360360208110156101e957600080fd5b50356001600160a01b031661056b565b61019a6004803603604081101561020f57600080fd5b50803515159060200135610593565b6102266105b7565b604080519115158252519081900360200190f35b6102266105bc565b6101d16004803603602081101561025857600080fd5b50356001600160a01b03166105c1565b61019a6004803603602081101561027e57600080fd5b50351515610729565b6101d16004803603604081101561029d57600080fd5b506001600160a01b0381358116916020013516610743565b6101d1600480360360408110156102cb57600080fd5b506001600160a01b0381351690602001351515610892565b6101d1600480360360208110156102f957600080fd5b503561090c565b61019a6109aa565b6101d16004803603606081101561031e57600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561034e57600080fd5b82018360208201111561036057600080fd5b8035906020019184600183028401116401000000008311171561038257600080fd5b5090925090506109bc565b6103aa600480360360208110156103a357600080fd5b5035610adf565b604080516001600160a01b0390971687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61019a600480360360a08110156103fd57600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359060800135610b33565b61042f610b9b565b60408051921515835260208301919091528051918290030190f35b61019a6004803603604081101561046057600080fd5b506001600160a01b038135169060200135610c45565b90151560009081526020818152604080832093835260029384019091529020805460018201549282015460038301546004909301546001600160a01b039092169490929190565b8115156000908152602081815260408083208484526002019091529020600301545b92915050565b8115156000908152602081815260408083208484526002810190925290912080546001600160a01b03163314610555576040805162461bcd60e51b815260206004820152601060248201526f37b7363c9037b93232b91036b0b5b2b960811b604482015290519081900360640190fd5b610565828463ffffffff610c5816565b50505050565b60008061057f60118463ffffffff610d1716565b905061058a81610d59565b9150505b919050565b90151560009081526020818152604080832093835260029093019052206004015490565b600081565b600181565b60006105d460118363ffffffff610d1716565b80549091506001600160a01b03838116911614610638576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c79206d616b65722063616e207265766f6b652070726f706f73616c0000604482015290519081900360640190fd5b601380546001810180835560009290925260068301805490916003027f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a09001906106849082908490613219565b5050600180548454918501546040805163a9059cbb60e01b81526001600160a01b03948516600482015260248101929092525192909116935063a9059cbb92506044808201926020929091908290030181600087803b1580156106e657600080fd5b505af11580156106fa573d6000803e3d6000fd5b505050506040513d602081101561071057600080fd5b50610724905060118363ffffffff610e3f16565b505050565b801515600090815260208190526040812061058a81610e87565b6001546001600160a01b03161561078f576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b6002546001600160a01b0316156107db576040805162461bcd60e51b815260206004820152600b60248201526a185b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b600180546001600160a01b038085166001600160a01b03199283161790925560028054928416929091169190911790556108158282610ea0565b600254604080516318160ddd60e01b8152905161088e926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561085b57600080fd5b505afa15801561086f573d6000803e3d6000fd5b505050506040513d602081101561088557600080fd5b50516000610ef0565b5050565b6108a360118363ffffffff610fbc16565b6108e7576040805162461bcd60e51b815260206004820152601060248201526f1b9bc81cdd58da081c1c9bdc1bdcd85b60821b604482015290519081900360640190fd5b60006108fa60118463ffffffff610d1716565b9050610724818363ffffffff610fdd16565b3315610950576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b60005b601354811015610988576109806013828154811061096d57fe5b9060005260206000209060030201610ff1565b600101610953565b5061099560136000613265565b61099d61105e565b506109a7816110ac565b50565b60006109b66011611245565b90505b90565b6080811480156109d657506001546001600160a01b031633145b15610a79576109ec60118563ffffffff610fbc16565b15610a37576040805162461bcd60e51b8152602060048201526016602482015275185b1c9958591e481a185cc818481c1c9bdc1bdcd85b60521b604482015290519081900360640190fd5b60008060008085856080811015610a4d57600080fd5b50803594506020810135935060400135915060009050610a708888868686611249565b50505050610565565b600080806060841415610aad5784846060811015610a9657600080fd5b508035935060208101359250604001359050610ac9565b84846040811015610abd57600080fd5b50803593506020013591505b610ad6878488858561135c565b50505050505050565b6000808080808080610af860118963ffffffff6113b716565b805460018201546002830154600484015460038501546005909501546001600160a01b039094169d929c50909a509850919650945092505050565b8415156000908152602081905260408120610b4c613286565b506040805160a0810182526001600160a01b0388168152602081018790529081018590526000606082018190526080820152610b8f82828663ffffffff6113c816565b98975050505050505050565b6007546000908190610bb257506000905080610c41565b6001610c3c600560020154600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c0b57600080fd5b505afa158015610c1f573d6000803e3d6000fd5b505050506040513d6020811015610c3557600080fd5b505161143a565b915091505b9091565b6000610c518383611454565b9392505050565b600081815260028301602052604090206001015415610d07578154600082815260028401602090815260408083208054600190910154825163a9059cbb60e01b81526001600160a01b0392831660048201526024810191909152915194169363a9059cbb93604480840194938390030190829087803b158015610cda57600080fd5b505af1158015610cee573d6000803e3d6000fd5b505050506040513d6020811015610d0457600080fd5b50505b61088e828263ffffffff61152116565b6001600160a01b03811660009081526001830160205260408120548354849160001901908110610d4357fe5b9060005260206000209060090201905092915050565b600080805b610d6a84600601611245565b811015610e3857600080610d87600687018463ffffffff61158116565b600154604080516370a0823160e01b81526001600160a01b03808616600483015291519496509294506000939116916370a08231916024808301926020929190829003018186803b158015610ddb57600080fd5b505afa158015610def573d6000803e3d6000fd5b505050506040513d6020811015610e0557600080fd5b50516001600160a01b038416310190508115610e245793840193610e2a565b80850394505b505050806001019050610d5e565b5092915050565b6001600160a01b038116600090815260018301602052604081205480610e695760009150506104df565b610e7d84600019830163ffffffff6115c016565b5060019392505050565b6000808052600282016020526040902060040154919050565b6000808052602052610ec76000805160206134af833981519152838363ffffffff61177616565b6001600090815260205261088e60008051602061348f833981519152828463ffffffff61177616565b604051806080016040528060035443018152602001600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f5357600080fd5b505afa158015610f67573d6000803e3d6000fd5b505050506040513d6020811015610f7d57600080fd5b50518152602080820194909452911515604092830152805160055591820151600655810151600755606001516008805460ff1916911515919091179055565b6001600160a01b031660009081526001919091016020526040902054151590565b61072460068301338363ffffffff61192316565b60005b815481101561105257600082600001828154811061100e57fe5b60009182526020808320909101546001600160a01b03168252600185810182526040808420849055600287019092529120805460ff19169055919091019050610ff4565b506109a78160006132b4565b600061106a6009611a3f565b15611077575060006109b9565b6000611081611a5b565b90506001600160a01b03811661109b5760009150506109b9565b6110a481611ae0565b600191505090565b33156110f0576040805162461bcd60e51b815260206004820152600e60248201526d636f6e73656e737573206f6e6c7960901b604482015290519081900360640190fd5b6110fa6005611c12565b15611119576000600581905560068190556007556008805460ff191690555b6111236009611c2b565b1561113057611130611c48565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561117557600080fd5b505afa158015611189573d6000803e3d6000fd5b505050506040513d602081101561119f57600080fd5b505190508115611213576111b1611d13565b806111c157506111c18183611d1f565b156111d1576111d1826000610ef0565b6111db6009611a3f565b156112135760006111ec838361143a565b90506111f781611dcd565b8015611205575060085460ff165b156112115750506109a7565b505b61122460058263ffffffff611eba16565b1561088e576000611233611eef565b905061123e81611fa4565b5050505050565b5490565b6112516132d2565b6001600160a01b038616815260208101859052604081018490524360a082015282156112e0576003600f548161128357fe5b04600f540183116112db576040805162461bcd60e51b815260206004820181905260248201527f736c617368696e67206475726174696f6e20706172616d20746f6f206c6f6e67604482015290519081900360640190fd5b6112e9565b600f5460808201525b8115611342576003600e54816112fb57fe5b04600e5403821061133d5760405162461bcd60e51b815260040180806020018281038252602181526020018061346e6021913960400191505060405180910390fd5b61134b565b600e5460608201525b610ad660118263ffffffff61204b16565b600061136733612198565b9050600061137e828888888863ffffffff61227b16565b905061139182828563ffffffff61241316565b506113ad8161139f33612451565b84919063ffffffff61250b16565b5050505050505050565b6000826000018281548110610d4357fe5b600081815260028401602052604081205b6004015460008181526002860160205260409020909250611400848263ffffffff6126a116565b156113d9575b600301546000818152600286016020526040902090925061142d848263ffffffff6126a116565b6114065750909392505050565b600081831161144e57828203600003610c51565b50900390565b60006002838360405160200180836001600160a01b03166001600160a01b031660601b8152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083106114c45780518252601f1990920191602091820191016114a5565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015611503573d6000803e3d6000fd5b5050506040513d602081101561151857600080fd5b50519392505050565b6000818152600292830160205260408082206004808201805460038085018054885286882090940182905583549187529486209094019390935593835280546001600160a01b031916815560018101839055909301819055908190559055565b60008080611595858563ffffffff6126bf16565b6001600160a01b038116600090815260028701602052604090205490935060ff169150509250929050565b8160010160008360000183815481106115d557fe5b600091825260208083206009909202909101546001600160a01b0316835282019290925260400181205581546000190181146116fe5781548290600019810190811061161d57fe5b906000526020600020906009020182600001828154811061163a57fe5b60009182526020909120825460099092020180546001600160a01b0319166001600160a01b0390921691909117815560018083015490820155600280830154908201556003808301549082015560048083015490820155600580830154908201556006808301805490918301906116b49082908490613219565b505050905050806001018260010160008460000184815481106116d357fe5b600091825260208083206009909202909101546001600160a01b031683528201929092526040019020555b81548290600019810190811061171057fe5b60009182526020822060099091020180546001600160a01b0319168155600181018290556002810182905560038101829055600481018290556005810182905590600682018161176082826132b4565b505083549150610724905083600019830161331d565b818360000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550808360010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060a00160405280306001600160a01b0316815260200160008152602001600081526020016000801b815260200160001960001b8152508360020160008060001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030155608082015181600401559050506040518060a00160405280306001600160a01b0316815260200160008152602001600181526020016000801b815260200160001960001b81525083600201600060001960001b815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040155905050505050565b6001600160a01b03821660009081526002840160209081526040808320805460ff191685151517905560018601909152812054806119a6575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610c51565b84548111806119e45750836001600160a01b03168560000160018303815481106119cc57fe5b6000918252602090912001546001600160a01b031614155b15611a34575050825460018082018086556000868152602080822090940180546001600160a01b0319166001600160a01b0388169081179091558152828701909352604090922091909155610c51565b506000949350505050565b6000611a4a826126ec565b80156104df57505060040154431090565b60008080805b611a6b6011611245565b811015611ad9576000611a8560118363ffffffff6113b716565b90506004600e5481611a9357fe5b04816005015443031015611aa75750611ad1565b6000611ab2826126fb565b905084811315611ace5781549094506001600160a01b03169250835b50505b600101611a61565b5091505090565b6000611af360118363ffffffff610d1716565b9050611b0181600601610ff1565b6040805160a08101825282546001600160a01b03168082526002840154602083018190526001850154938301849052600485015460608401819052600386015443016080909401849052600980546001600160a01b031916909317909255600a55600b92909255600c91909155600d55611b8260118363ffffffff610e3f16565b5061088e611c0b600260009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015611bd757600080fd5b505afa158015611beb573d6000803e3d6000fd5b505050506040513d6020811015611c0157600080fd5b5051600a54612732565b6001610ef0565b6000611c1d8261274e565b80156104df57505054431190565b6000611c36826126ec565b80156104df5750506004015443101590565b611c5260096126ec565b611c5b57611d11565b600b5415611cec57600154600954600b546040805163a9059cbb60e01b81526001600160a01b039384166004820152602481019290925251919092169163a9059cbb9160448083019260209291908290030181600087803b158015611cbf57600080fd5b505af1158015611cd3573d6000803e3d6000fd5b505050506040513d6020811015611ce957600080fd5b50505b600980546001600160a01b03191690556000600a819055600b819055600c819055600d555b565b60006109b66005611c12565b600082821415611d31575060006104df565b6006546007541415611d45575060016104df565b82821115611d9957600654600754848403911015611d7b57600654600754036002818381611d6f57fe5b041015925050506104df565b600754600654036002828281611d8d57fe5b041115925050506104df565b600654600754838503911115611dbb57600754600654036002818381611d6f57fe5b600654600754036002828281611d8d57fe5b6000611de0600960010154600084612754565b611dec5750600061058e565b600c54600a546000919084830381611e0057fe5b0581611e0857fe5b04905080611e14575060015b600b54811115611e435750600b546000600581905560068190556007556008805460ff19169055611e43611c48565b600b805482900390556001546040805163117f5a5560e01b81526004810184905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015611e9957600080fd5b505af1158015611ead573d6000803e3d6000fd5b5060019695505050505050565b6000611ec58361274e565b8015611ed5575082600201548214155b8015610c515750610c518360010154838560020154612784565b600080611f0660056002015460056001015461143a565b600754600254604080516318160ddd60e01b81529051939450600093611f5693926001600160a01b0316916318160ddd916004808301926020929190829003018186803b158015610c0b57600080fd5b9050611f6460008284612754565b611f73576000925050506109b9565b60006004548381611f8057fe5b059050611f8f60008284612754565b611f9d575091506109b99050565b9250505090565b6000806000806000808613611fba576001611fbd565b60005b15158152602081019190915260400160002060025481546008549293506001600160a01b039081169116149060ff16156120245761201981611ffe876127af565b600954859291906001600160a01b031663ffffffff6127c516565b935093505050612046565b61203f81612031876127af565b84919063ffffffff612b7816565b9350935050505b915091565b80516001600160a01b0316600090815260018301602052604081205480156120ba576040805162461bcd60e51b815260206004820152601c60248201527f6d616b657220616c72656164792068617320612070726f706f73616c00000000604482015290519081900360640190fd5b835460018082018087556000878152602090819020875160099095020180546001600160a01b0319166001600160a01b03909516949094178455868101519284019290925560408601516002840155606086015160038401556080860151600484015560a0860151600584015560c0860151805180519294889490936006850192612149928492910190613349565b505086516001600160a01b03166000908152600189016020526040902084905550508554919250859160001984019150811061218157fe5b906000526020600020906009020191505092915050565b600080805260208190526000805160206134af833981519152546001600160a01b03838116911614156121e2575060008080526020526000805160206134af83398151915261058e565b6001600090815260205260008051602061348f833981519152546001600160a01b038381169116141561222e57506001600090815260205260008051602061348f83398151915261058e565b6040805162461bcd60e51b815260206004820152601760248201527f6e6f206f7264657220626f6f6b20666f7220746f6b656e000000000000000000604482015290519081900360640190fd5b6000808311801561228c5750600082115b6122ca576040805162461bcd60e51b815260206004820152600a6024820152691e995c9bc81a5b9c1d5d60b21b604482015290519081900360640190fd5b600160801b831080156122e05750600160801b82105b612324576040805162461bcd60e51b815260206004820152601060248201526f1a5b9c1d5d081bdd995c881b1a5b5a5d60821b604482015290519081900360640190fd5b60006123308686611454565b6000818152600289016020526040902090915061234c816126ec565b15612393576040805162461bcd60e51b81526020600482015260126024820152716f7264657220696e6465782065786973747360701b604482015290519081900360640190fd5b506040805160a0810182526001600160a01b03978816815260208082019687528183019586526000606083018181526080840182815286835260029c8d01909352939020915182546001600160a01b0319169916989098178155945160018601559251968401969096555051600382015591516004909201919091555090565b600082815260028401602052604081208161243586838663ffffffff612e4516565b905061244886868363ffffffff612eb716565b95945050505050565b600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb6546001600160a01b03838116911614156124ad575060008080526020526000805160206134af83398151915261058e565b600160009081526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7e546001600160a01b038381169116141561222e57506001600090815260205260008051602061348f83398151915261058e565b600082815260028401602052604081208161252584610e87565b90505b60001981146124485761253a82612ef2565b1561254457612448565b600081815260028086016020908152604092839020835160a08101855286546001600160a01b0316815260018701549281019290925291850154928101929092526003840154606083015260048401546080830152906125aa908263ffffffff612f0b16565b6125b45750612448565b80600201548360010154101561266557600081600201548260010154856001015402816125dd57fe5b049050816001015481111561262b576040805162461bcd60e51b815260206004820152600f60248201526e66696c6c61626c65203e206861766560881b604482015290519081900360640190fd5b60018401546126459087908590849063ffffffff612f2a16565b600184015461265e90899089908463ffffffff612f2a16565b5050612448565b6002810154600182015461268391899189919063ffffffff612f2a16565b6004810154612698868463ffffffff6130c216565b91506125289050565b60408201516001820154600290920154602090930151910291021190565b60008260000182815481106126d057fe5b6000918252602090912001546001600160a01b03169392505050565b546001600160a01b0316151590565b60008061270783610d59565b90506000811361271b57600091505061058e565b61272483610d59565b836001015402915050919050565b6000808212156127495781600003830390506104df565b500190565b54151590565b60008284131580156127665750818313155b8061277c575082841215801561277c5750818312155b949350505050565b60008284111580156127965750818311155b8061277c575082841015801561277c5750501115919050565b60008082136127c157816000036104df565b5090565b6000806127d986868663ffffffff612b7816565b9092509050600080866127ed5783836127f0565b82845b895460408051636eb1769f60e11b81526001600160a01b038a811660048301523060248301529151949650929450169163dd62ed3e91604480820192602092909190829003018186803b15801561284657600080fd5b505afa15801561285a573d6000803e3d6000fd5b505050506040513d602081101561287057600080fd5b505182118015906128f857508754604080516370a0823160e01b81526001600160a01b038881166004830152915191909216916370a08231916024808301926020929190829003018186803b1580156128c857600080fd5b505afa1580156128dc573d6000803e3d6000fd5b505050506040513d60208110156128f257600080fd5b50518211155b15612b6d578754604080516334902c6160e11b815290516001600160a01b03909216916323b872dd918891849163692058c29160048083019260209291908290030181600087803b15801561294c57600080fd5b505af1158015612960573d6000803e3d6000fd5b505050506040513d602081101561297657600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482018690525160648083019260209291908290030181600087803b1580156129cf57600080fd5b505af11580156129e3573d6000803e3d6000fd5b505050506040513d60208110156129f957600080fd5b505087546040805163117f5a5560e01b81526004810185905290516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612a4757600080fd5b505af1158015612a5b573d6000803e3d6000fd5b50505060018901546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612aad57600080fd5b505af1158015612ac1573d6000803e3d6000fd5b5050505060018801546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015612b1d57600080fd5b505af1158015612b31573d6000803e3d6000fd5b505050506040513d6020811015612b4757600080fd5b50879050612b555780612b57565b815b8301925086612b665781612b68565b805b840193505b505094509492505050565b6000806000612b8686610e87565b90505b6000198114801590612b9a57508382105b15612e3c57600081815260028701602052604081209086612bbf578160020154612bc5565b81600101545b9050600087612bd8578260010154612bde565b82600201545b905086612bf1868463ffffffff61314716565b11612ce057885460018401546040805163117f5a5560e01b81526004810192909252516001600160a01b039092169163117f5a559160248082019260009290919082900301818387803b158015612c4757600080fd5b505af1158015612c5b573d6000803e3d6000fd5b50505060018a015460028501546040805163bdfde91160e01b81526004810192909252516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612cb257600080fd5b505af1158015612cc6573d6000803e3d6000fd5b505050506004830154612cd98a866130c2565b9350612e10565b6000612cf2888763ffffffff6131a116565b90508281830281612cff57fe5b049150809250600089612d125782612d14565b835b905060008a612d235784612d25565b835b8c546040805163117f5a5560e01b81526004810186905290519293506001600160a01b039091169163117f5a559160248082019260009290919082900301818387803b158015612d7457600080fd5b505af1158015612d88573d6000803e3d6000fd5b50505060018d01546040805163bdfde91160e01b81526004810185905290516001600160a01b03909216925063bdfde91191602480830192600092919082900301818387803b158015612dda57600080fd5b505af1158015612dee573d6000803e3d6000fd5b50612e0792508e9150899050848463ffffffff612f2a16565b50600019955050505b612e20858363ffffffff61314716565b9450612e32868263ffffffff61314716565b9550505050612b89565b50935093915050565b600081815260028401602052604081205b6004015460008181526002860160205260409020909250612e7d848263ffffffff6131f816565b15612e56575b6003015460008181526002860160205260409020909250612eaa848263ffffffff6131f816565b612e835750909392505050565b600081815260029093016020526040808420600490810180548587528387206003808201879055930181905586529185200183905592529055565b60008160010154600014806104df575050600201541590565b6040820151600282015460019092015460209093015191029102101590565b60008381526002850160205260409020612f43816126ec565b612f86576040805162461bcd60e51b815260206004820152600f60248201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604482015290519081900360640190fd5b8060010154831115612fdf576040805162461bcd60e51b815260206004820152601a60248201527f66696c6c206d6f7265207468616e206861766520616d6f756e74000000000000604482015290519081900360640190fd5b6001810154612ff4908463ffffffff6131a116565b6001820155600281015482101561301557600281018054839003905561301d565b600060028201555b600185015481546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018690529051919092169163a9059cbb9160448083019260209291908290030181600087803b15801561307757600080fd5b505af115801561308b573d6000803e3d6000fd5b505050506040513d60208110156130a157600080fd5b506130ad905081612ef2565b1561123e5761123e858563ffffffff610c5816565b600081815260028084016020526040909120015415610d0757600182015460008281526002808501602090815260408084208054930154815163a9059cbb60e01b81526001600160a01b03948516600482015260248101919091529051929094169363a9059cbb93604480830194928390030190829087803b158015610cda57600080fd5b600082820183811015610c51576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60008282111561144e576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b60008260020154826001015402826002015484600101540211905092915050565b8280548282559060005260206000209081019282156132595760005260206000209182015b8281111561325957825482559160010191906001019061323e565b506127c192915061339e565b50805460008255600302906000526020600020908101906109a791906133c2565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b50805460008255906000526020600020908101906109a791906133e5565b6040518060e0016040528060006001600160a01b0316815260200160008152602001600081526020016000815260200160008152602001600081526020016133186133ff565b905290565b815481835581811115610724576009028160090283600052602060002091820191016107249190613412565b828054828255906000526020600020908101928215613259579160200282015b8281111561325957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613369565b6109b991905b808211156127c15780546001600160a01b03191681556001016133a4565b6109b991905b808211156127c15760006133dc82826132b4565b506003016133c8565b6109b991905b808211156127c157600081556001016133eb565b6040518060200160405280606081525090565b6109b991905b808211156127c15780546001600160a01b031916815560006001820181905560028201819055600382018190556004820181905560058201819055600682018161346282826132b4565b50505060090161341856fe6c6f636b646f776e206475726174696f6e20706172616d20746f6f2073686f7274ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7dad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5a265627a7a72315820c37702e3aeed4575a5a8d2c0777710408fb0690156cafe904d2dc94319130d6064736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeploySeigniorage deploys a new Ethereum contract, binding an instance of Seigniorage to it.
func DeploySeigniorage(auth *bind.TransactOpts, backend bind.ContractBackend, absorptionDuration *big.Int, absorptionExpiration *big.Int, initialSlashingDuration *big.Int, initialLockdownExpiration *big.Int) (common.Address, *types.Transaction, *Seigniorage, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigniorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SeigniorageBin), backend, absorptionDuration, absorptionExpiration, initialSlashingDuration, initialLockdownExpiration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Seigniorage{SeigniorageCaller: SeigniorageCaller{contract: contract}, SeigniorageTransactor: SeigniorageTransactor{contract: contract}, SeigniorageFilterer: SeigniorageFilterer{contract: contract}}, nil
}

// Seigniorage is an auto generated Go binding around an Ethereum contract.
type Seigniorage struct {
	SeigniorageCaller     // Read-only binding to the contract
	SeigniorageTransactor // Write-only binding to the contract
	SeigniorageFilterer   // Log filterer for contract events
}

// SeigniorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeigniorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeigniorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeigniorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeigniorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeigniorageSession struct {
	Contract     *Seigniorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeigniorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeigniorageCallerSession struct {
	Contract *SeigniorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SeigniorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeigniorageTransactorSession struct {
	Contract     *SeigniorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SeigniorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeigniorageRaw struct {
	Contract *Seigniorage // Generic contract binding to access the raw methods on
}

// SeigniorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeigniorageCallerRaw struct {
	Contract *SeigniorageCaller // Generic read-only contract binding to access the raw methods on
}

// SeigniorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeigniorageTransactorRaw struct {
	Contract *SeigniorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeigniorage creates a new instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorage(address common.Address, backend bind.ContractBackend) (*Seigniorage, error) {
	contract, err := bindSeigniorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Seigniorage{SeigniorageCaller: SeigniorageCaller{contract: contract}, SeigniorageTransactor: SeigniorageTransactor{contract: contract}, SeigniorageFilterer: SeigniorageFilterer{contract: contract}}, nil
}

// NewSeigniorageCaller creates a new read-only instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageCaller(address common.Address, caller bind.ContractCaller) (*SeigniorageCaller, error) {
	contract, err := bindSeigniorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeigniorageCaller{contract: contract}, nil
}

// NewSeigniorageTransactor creates a new write-only instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SeigniorageTransactor, error) {
	contract, err := bindSeigniorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeigniorageTransactor{contract: contract}, nil
}

// NewSeigniorageFilterer creates a new log filterer instance of Seigniorage, bound to a specific deployed contract.
func NewSeigniorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SeigniorageFilterer, error) {
	contract, err := bindSeigniorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeigniorageFilterer{contract: contract}, nil
}

// bindSeigniorage binds a generic wrapper to an already deployed contract.
func bindSeigniorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeigniorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seigniorage *SeigniorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Seigniorage.Contract.SeigniorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seigniorage *SeigniorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seigniorage.Contract.SeigniorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seigniorage *SeigniorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seigniorage.Contract.SeigniorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Seigniorage *SeigniorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Seigniorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Seigniorage *SeigniorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Seigniorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Seigniorage *SeigniorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Seigniorage.Contract.contract.Transact(opts, method, params...)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Seigniorage *SeigniorageCaller) Ask(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "Ask")
	return *ret0, err
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Seigniorage *SeigniorageSession) Ask() (bool, error) {
	return _Seigniorage.Contract.Ask(&_Seigniorage.CallOpts)
}

// Ask is a free data retrieval call binding the contract method 0x69c07d31.
//
// Solidity: function Ask() constant returns(bool)
func (_Seigniorage *SeigniorageCallerSession) Ask() (bool, error) {
	return _Seigniorage.Contract.Ask(&_Seigniorage.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Seigniorage *SeigniorageCaller) Bid(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "Bid")
	return *ret0, err
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Seigniorage *SeigniorageSession) Bid() (bool, error) {
	return _Seigniorage.Contract.Bid(&_Seigniorage.CallOpts)
}

// Bid is a free data retrieval call binding the contract method 0x6e6452cb.
//
// Solidity: function Bid() constant returns(bool)
func (_Seigniorage *SeigniorageCallerSession) Bid() (bool, error) {
	return _Seigniorage.Contract.Bid(&_Seigniorage.CallOpts)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) CalcOrderID(opts *bind.CallOpts, maker common.Address, index [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "calcOrderID", maker, index)
	return *ret0, err
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.CalcOrderID(&_Seigniorage.CallOpts, maker, index)
}

// CalcOrderID is a free data retrieval call binding the contract method 0xf318722b.
//
// Solidity: function calcOrderID(address maker, bytes32 index) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) CalcOrderID(maker common.Address, index [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.CalcOrderID(&_Seigniorage.CallOpts, maker, index)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) FindAssistingID(opts *bind.CallOpts, orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "findAssistingID", orderType, maker, haveAmount, wantAmount, assistingID)
	return *ret0, err
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.FindAssistingID(&_Seigniorage.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// FindAssistingID is a free data retrieval call binding the contract method 0xced4aac8.
//
// Solidity: function findAssistingID(bool orderType, address maker, uint256 haveAmount, uint256 wantAmount, bytes32 assistingID) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) FindAssistingID(orderType bool, maker common.Address, haveAmount *big.Int, wantAmount *big.Int, assistingID [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.FindAssistingID(&_Seigniorage.CallOpts, orderType, maker, haveAmount, wantAmount, assistingID)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Seigniorage *SeigniorageCaller) GetOrder(opts *bind.CallOpts, _orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new([32]byte)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Seigniorage.contract.Call(opts, out, "getOrder", _orderType, _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Seigniorage *SeigniorageSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Seigniorage.Contract.GetOrder(&_Seigniorage.CallOpts, _orderType, _id)
}

// GetOrder is a free data retrieval call binding the contract method 0x07c399a3.
//
// Solidity: function getOrder(bool _orderType, bytes32 _id) constant returns(address, uint256, uint256, bytes32, bytes32)
func (_Seigniorage *SeigniorageCallerSession) GetOrder(_orderType bool, _id [32]byte) (common.Address, *big.Int, *big.Int, [32]byte, [32]byte, error) {
	return _Seigniorage.Contract.GetOrder(&_Seigniorage.CallOpts, _orderType, _id)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageCaller) GetProposal(opts *bind.CallOpts, idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	ret := new(struct {
		Maker              common.Address
		Stake              *big.Int
		Amount             *big.Int
		SlashingDuration   *big.Int
		LockdownExpiration *big.Int
		Number             *big.Int
	})
	out := ret
	err := _Seigniorage.contract.Call(opts, out, "getProposal", idx)
	return *ret, err
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Seigniorage.Contract.GetProposal(&_Seigniorage.CallOpts, idx)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 idx) constant returns(address maker, uint256 stake, int256 amount, uint256 slashingDuration, uint256 lockdownExpiration, uint256 number)
func (_Seigniorage *SeigniorageCallerSession) GetProposal(idx *big.Int) (struct {
	Maker              common.Address
	Stake              *big.Int
	Amount             *big.Int
	SlashingDuration   *big.Int
	LockdownExpiration *big.Int
	Number             *big.Int
}, error) {
	return _Seigniorage.Contract.GetProposal(&_Seigniorage.CallOpts, idx)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Seigniorage *SeigniorageCaller) GetProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "getProposalCount")
	return *ret0, err
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Seigniorage *SeigniorageSession) GetProposalCount() (*big.Int, error) {
	return _Seigniorage.Contract.GetProposalCount(&_Seigniorage.CallOpts)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() constant returns(uint256)
func (_Seigniorage *SeigniorageCallerSession) GetProposalCount() (*big.Int, error) {
	return _Seigniorage.Contract.GetProposalCount(&_Seigniorage.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Seigniorage *SeigniorageCaller) GetRemainToAbsorb(opts *bind.CallOpts) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Seigniorage.contract.Call(opts, out, "getRemainToAbsorb")
	return *ret0, *ret1, err
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Seigniorage *SeigniorageSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Seigniorage.Contract.GetRemainToAbsorb(&_Seigniorage.CallOpts)
}

// GetRemainToAbsorb is a free data retrieval call binding the contract method 0xee1a68c6.
//
// Solidity: function getRemainToAbsorb() constant returns(bool, int256)
func (_Seigniorage *SeigniorageCallerSession) GetRemainToAbsorb() (bool, *big.Int, error) {
	return _Seigniorage.Contract.GetRemainToAbsorb(&_Seigniorage.CallOpts)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Next(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "next", orderType, id)
	return *ret0, err
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Next(&_Seigniorage.CallOpts, orderType, id)
}

// Next is a free data retrieval call binding the contract method 0x4ea09797.
//
// Solidity: function next(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Next(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Next(&_Seigniorage.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Prev(opts *bind.CallOpts, orderType bool, id [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "prev", orderType, id)
	return *ret0, err
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Prev(&_Seigniorage.CallOpts, orderType, id)
}

// Prev is a free data retrieval call binding the contract method 0x0d90b10a.
//
// Solidity: function prev(bool orderType, bytes32 id) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Prev(orderType bool, id [32]byte) ([32]byte, error) {
	return _Seigniorage.Contract.Prev(&_Seigniorage.CallOpts, orderType, id)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Seigniorage *SeigniorageCaller) Top(opts *bind.CallOpts, orderType bool) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "top", orderType)
	return *ret0, err
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Seigniorage *SeigniorageSession) Top(orderType bool) ([32]byte, error) {
	return _Seigniorage.Contract.Top(&_Seigniorage.CallOpts, orderType)
}

// Top is a free data retrieval call binding the contract method 0x8aa3f897.
//
// Solidity: function top(bool orderType) constant returns(bytes32)
func (_Seigniorage *SeigniorageCallerSession) Top(orderType bool) ([32]byte, error) {
	return _Seigniorage.Contract.Top(&_Seigniorage.CallOpts, orderType)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Seigniorage *SeigniorageCaller) TotalVote(opts *bind.CallOpts, maker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Seigniorage.contract.Call(opts, out, "totalVote", maker)
	return *ret0, err
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Seigniorage *SeigniorageSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Seigniorage.Contract.TotalVote(&_Seigniorage.CallOpts, maker)
}

// TotalVote is a free data retrieval call binding the contract method 0x4def5645.
//
// Solidity: function totalVote(address maker) constant returns(int256)
func (_Seigniorage *SeigniorageCallerSession) TotalVote(maker common.Address) (*big.Int, error) {
	return _Seigniorage.Contract.TotalVote(&_Seigniorage.CallOpts, maker)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageTransactor) Cancel(opts *bind.TransactOpts, orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "cancel", orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.Cancel(&_Seigniorage.TransactOpts, orderType, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x43271d79.
//
// Solidity: function cancel(bool orderType, bytes32 id) returns()
func (_Seigniorage *SeigniorageTransactorSession) Cancel(orderType bool, id [32]byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.Cancel(&_Seigniorage.TransactOpts, orderType, id)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageTransactor) OnBlockInitialized(opts *bind.TransactOpts, target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "onBlockInitialized", target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.Contract.OnBlockInitialized(&_Seigniorage.TransactOpts, target)
}

// OnBlockInitialized is a paid mutator transaction binding the contract method 0xbe91d729.
//
// Solidity: function onBlockInitialized(uint256 target) returns()
func (_Seigniorage *SeigniorageTransactorSession) OnBlockInitialized(target *big.Int) (*types.Transaction, error) {
	return _Seigniorage.Contract.OnBlockInitialized(&_Seigniorage.TransactOpts, target)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageTransactor) RegisterTokens(opts *bind.TransactOpts, volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "registerTokens", volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.RegisterTokens(&_Seigniorage.TransactOpts, volatileToken, stablizeToken)
}

// RegisterTokens is a paid mutator transaction binding the contract method 0xaa1c259c.
//
// Solidity: function registerTokens(address volatileToken, address stablizeToken) returns()
func (_Seigniorage *SeigniorageTransactorSession) RegisterTokens(volatileToken common.Address, stablizeToken common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.RegisterTokens(&_Seigniorage.TransactOpts, volatileToken, stablizeToken)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageTransactor) Revoke(opts *bind.TransactOpts, maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "revoke", maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.Revoke(&_Seigniorage.TransactOpts, maker)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address maker) returns()
func (_Seigniorage *SeigniorageTransactorSession) Revoke(maker common.Address) (*types.Transaction, error) {
	return _Seigniorage.Contract.Revoke(&_Seigniorage.TransactOpts, maker)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageTransactor) TokenFallback(opts *bind.TransactOpts, maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "tokenFallback", maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.TokenFallback(&_Seigniorage.TransactOpts, maker, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address maker, uint256 value, bytes data) returns()
func (_Seigniorage *SeigniorageTransactorSession) TokenFallback(maker common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Seigniorage.Contract.TokenFallback(&_Seigniorage.TransactOpts, maker, value, data)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageTransactor) Vote(opts *bind.TransactOpts, maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.contract.Transact(opts, "vote", maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.Contract.Vote(&_Seigniorage.TransactOpts, maker, up)
}

// Vote is a paid mutator transaction binding the contract method 0xbd041c4d.
//
// Solidity: function vote(address maker, bool up) returns()
func (_Seigniorage *SeigniorageTransactorSession) Vote(maker common.Address, up bool) (*types.Transaction, error) {
	return _Seigniorage.Contract.Vote(&_Seigniorage.TransactOpts, maker, up)
}

// AbsnABI is the input ABI used to generate the binding from.
const AbsnABI = "[]"

// AbsnBin is the compiled bytecode used for deploying new contracts.
var AbsnBin = "0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582069e688955fdef0c43830f2adcae2c2114ac6f6fd34a2cb86fb04071a5748979c64736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployAbsn deploys a new Ethereum contract, binding an instance of Absn to it.
func DeployAbsn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Absn, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbsnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Absn{AbsnCaller: AbsnCaller{contract: contract}, AbsnTransactor: AbsnTransactor{contract: contract}, AbsnFilterer: AbsnFilterer{contract: contract}}, nil
}

// Absn is an auto generated Go binding around an Ethereum contract.
type Absn struct {
	AbsnCaller     // Read-only binding to the contract
	AbsnTransactor // Write-only binding to the contract
	AbsnFilterer   // Log filterer for contract events
}

// AbsnCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsnSession struct {
	Contract     *Absn             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsnCallerSession struct {
	Contract *AbsnCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbsnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsnTransactorSession struct {
	Contract     *AbsnTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbsnRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsnRaw struct {
	Contract *Absn // Generic contract binding to access the raw methods on
}

// AbsnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsnCallerRaw struct {
	Contract *AbsnCaller // Generic read-only contract binding to access the raw methods on
}

// AbsnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsnTransactorRaw struct {
	Contract *AbsnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsn creates a new instance of Absn, bound to a specific deployed contract.
func NewAbsn(address common.Address, backend bind.ContractBackend) (*Absn, error) {
	contract, err := bindAbsn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Absn{AbsnCaller: AbsnCaller{contract: contract}, AbsnTransactor: AbsnTransactor{contract: contract}, AbsnFilterer: AbsnFilterer{contract: contract}}, nil
}

// NewAbsnCaller creates a new read-only instance of Absn, bound to a specific deployed contract.
func NewAbsnCaller(address common.Address, caller bind.ContractCaller) (*AbsnCaller, error) {
	contract, err := bindAbsn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsnCaller{contract: contract}, nil
}

// NewAbsnTransactor creates a new write-only instance of Absn, bound to a specific deployed contract.
func NewAbsnTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsnTransactor, error) {
	contract, err := bindAbsn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsnTransactor{contract: contract}, nil
}

// NewAbsnFilterer creates a new log filterer instance of Absn, bound to a specific deployed contract.
func NewAbsnFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsnFilterer, error) {
	contract, err := bindAbsn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsnFilterer{contract: contract}, nil
}

// bindAbsn binds a generic wrapper to an already deployed contract.
func bindAbsn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbsnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absn *AbsnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absn.Contract.AbsnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absn *AbsnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absn.Contract.AbsnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absn *AbsnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absn.Contract.AbsnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Absn *AbsnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Absn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Absn *AbsnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Absn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Absn *AbsnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Absn.Contract.contract.Transact(opts, method, params...)
}

// DexABI is the input ABI used to generate the binding from.
const DexABI = "[]"

// DexBin is the compiled bytecode used for deploying new contracts.
var DexBin = "0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582091ce0c38a72cb8c507208dd88dcf47af364db331b32169cdd889e6337905d78c64736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployDex deploys a new Ethereum contract, binding an instance of Dex to it.
func DeployDex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dex, error) {
	parsed, err := abi.JSON(strings.NewReader(DexABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// Dex is an auto generated Go binding around an Ethereum contract.
type Dex struct {
	DexCaller     // Read-only binding to the contract
	DexTransactor // Write-only binding to the contract
	DexFilterer   // Log filterer for contract events
}

// DexCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexSession struct {
	Contract     *Dex              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexCallerSession struct {
	Contract *DexCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexTransactorSession struct {
	Contract     *DexTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexRaw struct {
	Contract *Dex // Generic contract binding to access the raw methods on
}

// DexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexCallerRaw struct {
	Contract *DexCaller // Generic read-only contract binding to access the raw methods on
}

// DexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexTransactorRaw struct {
	Contract *DexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDex creates a new instance of Dex, bound to a specific deployed contract.
func NewDex(address common.Address, backend bind.ContractBackend) (*Dex, error) {
	contract, err := bindDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dex{DexCaller: DexCaller{contract: contract}, DexTransactor: DexTransactor{contract: contract}, DexFilterer: DexFilterer{contract: contract}}, nil
}

// NewDexCaller creates a new read-only instance of Dex, bound to a specific deployed contract.
func NewDexCaller(address common.Address, caller bind.ContractCaller) (*DexCaller, error) {
	contract, err := bindDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexCaller{contract: contract}, nil
}

// NewDexTransactor creates a new write-only instance of Dex, bound to a specific deployed contract.
func NewDexTransactor(address common.Address, transactor bind.ContractTransactor) (*DexTransactor, error) {
	contract, err := bindDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexTransactor{contract: contract}, nil
}

// NewDexFilterer creates a new log filterer instance of Dex, bound to a specific deployed contract.
func NewDexFilterer(address common.Address, filterer bind.ContractFilterer) (*DexFilterer, error) {
	contract, err := bindDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexFilterer{contract: contract}, nil
}

// bindDex binds a generic wrapper to an already deployed contract.
func bindDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.DexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.DexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dex *DexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dex *DexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dex *DexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dex.Contract.contract.Transact(opts, method, params...)
}

// MapABI is the input ABI used to generate the binding from.
const MapABI = "[]"

// MapBin is the compiled bytecode used for deploying new contracts.
var MapBin = "0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820dfe3c45afba83a7a8446c7e89a3cbc534a91c672e1fbc50f2831d9e7fa707e2964736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployMap deploys a new Ethereum contract, binding an instance of Map to it.
func DeployMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Map, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// Map is an auto generated Go binding around an Ethereum contract.
type Map struct {
	MapCaller     // Read-only binding to the contract
	MapTransactor // Write-only binding to the contract
	MapFilterer   // Log filterer for contract events
}

// MapCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapSession struct {
	Contract     *Map              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapCallerSession struct {
	Contract *MapCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapTransactorSession struct {
	Contract     *MapTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapRaw struct {
	Contract *Map // Generic contract binding to access the raw methods on
}

// MapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapCallerRaw struct {
	Contract *MapCaller // Generic read-only contract binding to access the raw methods on
}

// MapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapTransactorRaw struct {
	Contract *MapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMap creates a new instance of Map, bound to a specific deployed contract.
func NewMap(address common.Address, backend bind.ContractBackend) (*Map, error) {
	contract, err := bindMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Map{MapCaller: MapCaller{contract: contract}, MapTransactor: MapTransactor{contract: contract}, MapFilterer: MapFilterer{contract: contract}}, nil
}

// NewMapCaller creates a new read-only instance of Map, bound to a specific deployed contract.
func NewMapCaller(address common.Address, caller bind.ContractCaller) (*MapCaller, error) {
	contract, err := bindMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapCaller{contract: contract}, nil
}

// NewMapTransactor creates a new write-only instance of Map, bound to a specific deployed contract.
func NewMapTransactor(address common.Address, transactor bind.ContractTransactor) (*MapTransactor, error) {
	contract, err := bindMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapTransactor{contract: contract}, nil
}

// NewMapFilterer creates a new log filterer instance of Map, bound to a specific deployed contract.
func NewMapFilterer(address common.Address, filterer bind.ContractFilterer) (*MapFilterer, error) {
	contract, err := bindMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapFilterer{contract: contract}, nil
}

// bindMap binds a generic wrapper to an already deployed contract.
func bindMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.MapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.MapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Map *MapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Map.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Map *MapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Map.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Map *MapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Map.Contract.contract.Transact(opts, method, params...)
}

// UtilABI is the input ABI used to generate the binding from.
const UtilABI = "[]"

// UtilBin is the compiled bytecode used for deploying new contracts.
var UtilBin = "0x607a6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a5a337489b3fe71a6e37719e0f9e39d2965c3223eda8767029665dcbbd45699764736f6c637827302e352e31312d646576656c6f702e323031392e382e382b636f6d6d69742e36393536656131650057"

// DeployUtil deploys a new Ethereum contract, binding an instance of Util to it.
func DeployUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Util, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Util{UtilCaller: UtilCaller{contract: contract}, UtilTransactor: UtilTransactor{contract: contract}, UtilFilterer: UtilFilterer{contract: contract}}, nil
}

// Util is an auto generated Go binding around an Ethereum contract.
type Util struct {
	UtilCaller     // Read-only binding to the contract
	UtilTransactor // Write-only binding to the contract
	UtilFilterer   // Log filterer for contract events
}

// UtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilSession struct {
	Contract     *Util             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilCallerSession struct {
	Contract *UtilCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilTransactorSession struct {
	Contract     *UtilTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilRaw struct {
	Contract *Util // Generic contract binding to access the raw methods on
}

// UtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilCallerRaw struct {
	Contract *UtilCaller // Generic read-only contract binding to access the raw methods on
}

// UtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilTransactorRaw struct {
	Contract *UtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtil creates a new instance of Util, bound to a specific deployed contract.
func NewUtil(address common.Address, backend bind.ContractBackend) (*Util, error) {
	contract, err := bindUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Util{UtilCaller: UtilCaller{contract: contract}, UtilTransactor: UtilTransactor{contract: contract}, UtilFilterer: UtilFilterer{contract: contract}}, nil
}

// NewUtilCaller creates a new read-only instance of Util, bound to a specific deployed contract.
func NewUtilCaller(address common.Address, caller bind.ContractCaller) (*UtilCaller, error) {
	contract, err := bindUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilCaller{contract: contract}, nil
}

// NewUtilTransactor creates a new write-only instance of Util, bound to a specific deployed contract.
func NewUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilTransactor, error) {
	contract, err := bindUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilTransactor{contract: contract}, nil
}

// NewUtilFilterer creates a new log filterer instance of Util, bound to a specific deployed contract.
func NewUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilFilterer, error) {
	contract, err := bindUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilFilterer{contract: contract}, nil
}

// bindUtil binds a generic wrapper to an already deployed contract.
func bindUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Util *UtilRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Util.Contract.UtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Util *UtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Util.Contract.UtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Util *UtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Util.Contract.UtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Util *UtilCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Util.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Util *UtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Util.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Util *UtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Util.Contract.contract.Transact(opts, method, params...)
}
