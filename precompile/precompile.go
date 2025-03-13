// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package precompile

import (
	"errors"
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

// PrecompileMetaData contains all meta data concerning the Precompile contract.
var PrecompileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBLS12_G1Add\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBLS12_G1MSM\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBLS12_G2Add\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBLS12_G2MSM\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBLS12_MapFPtoG1\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBLS12_PairingCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBLSS12_MapFP2toG2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBlake2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBn256Add\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBn256Pairing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testBn256ScalarMul\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testEcrecover\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testIdentity\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testModexp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testPointEvaluation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testRipemd160\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"testSha256\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611ae58061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610109575f3560e01c8063aa63d4b6116100a0578063db77cf6d1161006f578063db77cf6d1461034d578063eb4a006b1461037d578063f5688603146103ad578063f5a7bad0146103dd578063f6b0bbf71461040d57610109565b8063aa63d4b61461028d578063d020aeb7146102bd578063d1ca680b146102ed578063d99b39e41461031d57610109565b80636fe7c1f0116100dc5780636fe7c1f0146101cd57806386c3f234146101fd5780639cce7cf91461022d578063a67ecca61461025d57610109565b806310a62c721461010d5780632a45bb521461013d5780634a6b46051461016d578063533dcda61461019d575b5f5ffd5b61012760048036038101906101229190611240565b61043d565b60405161013491906112e7565b60405180910390f35b61015760048036038101906101529190611240565b6104f3565b60405161016491906112e7565b60405180910390f35b61018760048036038101906101829190611240565b6105a9565b60405161019491906112e7565b60405180910390f35b6101b760048036038101906101b29190611240565b61065f565b6040516101c491906112e7565b60405180910390f35b6101e760048036038101906101e29190611240565b610715565b6040516101f491906112e7565b60405180910390f35b61021760048036038101906102129190611240565b6107cb565b60405161022491906112e7565b60405180910390f35b61024760048036038101906102429190611240565b610881565b60405161025491906112e7565b60405180910390f35b61027760048036038101906102729190611240565b610937565b60405161028491906112e7565b60405180910390f35b6102a760048036038101906102a29190611240565b6109ed565b6040516102b49190611321565b60405180910390f35b6102d760048036038101906102d29190611240565b610af3565b6040516102e491906112e7565b60405180910390f35b61030760048036038101906103029190611240565b610ba9565b60405161031491906112e7565b60405180910390f35b61033760048036038101906103329190611240565b610c5f565b60405161034491906112e7565b60405180910390f35b61036760048036038101906103629190611240565b610d15565b60405161037491906112e7565b60405180910390f35b61039760048036038101906103929190611240565b610dcb565b6040516103a491906112e7565b60405180910390f35b6103c760048036038101906103c29190611240565b610e81565b6040516103d49190611321565b60405180910390f35b6103f760048036038101906103f29190611240565b610f87565b60405161040491906112e7565b60405180910390f35b61042760048036038101906104229190611240565b61103d565b60405161043491906112e7565b60405180910390f35b60605f5f600d73ffffffffffffffffffffffffffffffffffffffff16846040516104679190611374565b5f60405180830381855afa9150503d805f811461049f576040519150601f19603f3d011682016040523d82523d5f602084013e6104a4565b606091505b5091509150816104e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e0906113e4565b60405180910390fd5b8092505050919050565b60605f5f600673ffffffffffffffffffffffffffffffffffffffff168460405161051d9190611374565b5f60405180830381855afa9150503d805f8114610555576040519150601f19603f3d011682016040523d82523d5f602084013e61055a565b606091505b50915091508161059f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105969061144c565b60405180910390fd5b8092505050919050565b60605f5f600c73ffffffffffffffffffffffffffffffffffffffff16846040516105d39190611374565b5f60405180830381855afa9150503d805f811461060b576040519150601f19603f3d011682016040523d82523d5f602084013e610610565b606091505b509150915081610655576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064c906114b4565b60405180910390fd5b8092505050919050565b60605f5f600973ffffffffffffffffffffffffffffffffffffffff16846040516106899190611374565b5f60405180830381855afa9150503d805f81146106c1576040519150601f19603f3d011682016040523d82523d5f602084013e6106c6565b606091505b50915091508161070b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107029061151c565b60405180910390fd5b8092505050919050565b60605f5f600573ffffffffffffffffffffffffffffffffffffffff168460405161073f9190611374565b5f60405180830381855afa9150503d805f8114610777576040519150601f19603f3d011682016040523d82523d5f602084013e61077c565b606091505b5091509150816107c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b890611584565b60405180910390fd5b8092505050919050565b60605f5f600173ffffffffffffffffffffffffffffffffffffffff16846040516107f59190611374565b5f60405180830381855afa9150503d805f811461082d576040519150601f19603f3d011682016040523d82523d5f602084013e610832565b606091505b509150915081610877576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086e906115ec565b60405180910390fd5b8092505050919050565b60605f5f600473ffffffffffffffffffffffffffffffffffffffff16846040516108ab9190611374565b5f60405180830381855afa9150503d805f81146108e3576040519150601f19603f3d011682016040523d82523d5f602084013e6108e8565b606091505b50915091508161092d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092490611654565b60405180910390fd5b8092505050919050565b60605f5f601073ffffffffffffffffffffffffffffffffffffffff16846040516109619190611374565b5f60405180830381855afa9150503d805f8114610999576040519150601f19603f3d011682016040523d82523d5f602084013e61099e565b606091505b5091509150816109e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109da906116bc565b60405180910390fd5b8092505050919050565b5f5f5f600873ffffffffffffffffffffffffffffffffffffffff1684604051610a169190611374565b5f60405180830381855afa9150503d805f8114610a4e576040519150601f19603f3d011682016040523d82523d5f602084013e610a53565b606091505b509150915081610a98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8f90611724565b60405180910390fd5b5f8151118015610aea57505f60f81b815f81518110610aba57610ab9611742565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b92505050919050565b60605f5f600273ffffffffffffffffffffffffffffffffffffffff1684604051610b1d9190611374565b5f60405180830381855afa9150503d805f8114610b55576040519150601f19603f3d011682016040523d82523d5f602084013e610b5a565b606091505b509150915081610b9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b96906117b9565b60405180910390fd5b8092505050919050565b60605f5f600b73ffffffffffffffffffffffffffffffffffffffff1684604051610bd39190611374565b5f60405180830381855afa9150503d805f8114610c0b576040519150601f19603f3d011682016040523d82523d5f602084013e610c10565b606091505b509150915081610c55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4c90611821565b60405180910390fd5b8092505050919050565b60605f5f600773ffffffffffffffffffffffffffffffffffffffff1684604051610c899190611374565b5f60405180830381855afa9150503d805f8114610cc1576040519150601f19603f3d011682016040523d82523d5f602084013e610cc6565b606091505b509150915081610d0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0290611889565b60405180910390fd5b8092505050919050565b60605f5f600a73ffffffffffffffffffffffffffffffffffffffff1684604051610d3f9190611374565b5f60405180830381855afa9150503d805f8114610d77576040519150601f19603f3d011682016040523d82523d5f602084013e610d7c565b606091505b509150915081610dc1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db8906118f1565b60405180910390fd5b8092505050919050565b60605f5f601173ffffffffffffffffffffffffffffffffffffffff1684604051610df59190611374565b5f60405180830381855afa9150503d805f8114610e2d576040519150601f19603f3d011682016040523d82523d5f602084013e610e32565b606091505b509150915081610e77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6e90611959565b60405180910390fd5b8092505050919050565b5f5f5f600f73ffffffffffffffffffffffffffffffffffffffff1684604051610eaa9190611374565b5f60405180830381855afa9150503d805f8114610ee2576040519150601f19603f3d011682016040523d82523d5f602084013e610ee7565b606091505b509150915081610f2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f23906119c1565b60405180910390fd5b5f8151118015610f7e57505f60f81b815f81518110610f4e57610f4d611742565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b92505050919050565b60605f5f600e73ffffffffffffffffffffffffffffffffffffffff1684604051610fb19190611374565b5f60405180830381855afa9150503d805f8114610fe9576040519150601f19603f3d011682016040523d82523d5f602084013e610fee565b606091505b509150915081611033576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102a90611a29565b60405180910390fd5b8092505050919050565b60605f5f600373ffffffffffffffffffffffffffffffffffffffff16846040516110679190611374565b5f60405180830381855afa9150503d805f811461109f576040519150601f19603f3d011682016040523d82523d5f602084013e6110a4565b606091505b5091509150816110e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110e090611a91565b60405180910390fd5b8092505050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6111528261110c565b810181811067ffffffffffffffff821117156111715761117061111c565b5b80604052505050565b5f6111836110f3565b905061118f8282611149565b919050565b5f67ffffffffffffffff8211156111ae576111ad61111c565b5b6111b78261110c565b9050602081019050919050565b828183375f83830152505050565b5f6111e46111df84611194565b61117a565b905082815260208101848484011115611200576111ff611108565b5b61120b8482856111c4565b509392505050565b5f82601f83011261122757611226611104565b5b81356112378482602086016111d2565b91505092915050565b5f60208284031215611255576112546110fc565b5b5f82013567ffffffffffffffff81111561127257611271611100565b5b61127e84828501611213565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6112b982611287565b6112c38185611291565b93506112d38185602086016112a1565b6112dc8161110c565b840191505092915050565b5f6020820190508181035f8301526112ff81846112af565b905092915050565b5f8115159050919050565b61131b81611307565b82525050565b5f6020820190506113345f830184611312565b92915050565b5f81905092915050565b5f61134e82611287565b611358818561133a565b93506113688185602086016112a1565b80840191505092915050565b5f61137f8284611344565b915081905092915050565b5f82825260208201905092915050565b7f424c5331325f47324164642063616c6c206661696c65640000000000000000005f82015250565b5f6113ce60178361138a565b91506113d98261139a565b602082019050919050565b5f6020820190508181035f8301526113fb816113c2565b9050919050565b7f626e3235364164642063616c6c206661696c65640000000000000000000000005f82015250565b5f61143660148361138a565b915061144182611402565b602082019050919050565b5f6020820190508181035f8301526114638161142a565b9050919050565b7f424c5331325f47314d534d2063616c6c206661696c65640000000000000000005f82015250565b5f61149e60178361138a565b91506114a98261146a565b602082019050919050565b5f6020820190508181035f8301526114cb81611492565b9050919050565b7f626c616b65322063616c6c206661696c656400000000000000000000000000005f82015250565b5f61150660128361138a565b9150611511826114d2565b602082019050919050565b5f6020820190508181035f830152611533816114fa565b9050919050565b7f6d6f646578702063616c6c206661696c656400000000000000000000000000005f82015250565b5f61156e60128361138a565b91506115798261153a565b602082019050919050565b5f6020820190508181035f83015261159b81611562565b9050919050565b7f65637265636f7665722063616c6c206661696c656400000000000000000000005f82015250565b5f6115d660158361138a565b91506115e1826115a2565b602082019050919050565b5f6020820190508181035f830152611603816115ca565b9050919050565b7f6964656e7469747920707265636f6d70696c65206661696c65640000000000005f82015250565b5f61163e601a8361138a565b91506116498261160a565b602082019050919050565b5f6020820190508181035f83015261166b81611632565b9050919050565b7f424c5331325f4d61704650746f47312063616c6c206661696c656400000000005f82015250565b5f6116a6601b8361138a565b91506116b182611672565b602082019050919050565b5f6020820190508181035f8301526116d38161169a565b9050919050565b7f626e32353650616972696e672063616c6c206661696c656400000000000000005f82015250565b5f61170e60188361138a565b9150611719826116da565b602082019050919050565b5f6020820190508181035f83015261173b81611702565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f7368613235362063616c6c206661696c656400000000000000000000000000005f82015250565b5f6117a360128361138a565b91506117ae8261176f565b602082019050919050565b5f6020820190508181035f8301526117d081611797565b9050919050565b7f424c5331325f47314164642063616c6c206661696c65640000000000000000005f82015250565b5f61180b60178361138a565b9150611816826117d7565b602082019050919050565b5f6020820190508181035f830152611838816117ff565b9050919050565b7f626e3235365363616c61724d756c2063616c6c206661696c65640000000000005f82015250565b5f611873601a8361138a565b915061187e8261183f565b602082019050919050565b5f6020820190508181035f8301526118a081611867565b9050919050565b7f706f696e744576616c756174696f6e2063616c6c206661696c656400000000005f82015250565b5f6118db601b8361138a565b91506118e6826118a7565b602082019050919050565b5f6020820190508181035f830152611908816118cf565b9050919050565b7f424c535331325f4d6170465032746f47322063616c6c206661696c65640000005f82015250565b5f611943601d8361138a565b915061194e8261190f565b602082019050919050565b5f6020820190508181035f83015261197081611937565b9050919050565b7f424c5331325f50616972696e67436865636b2063616c6c206661696c656400005f82015250565b5f6119ab601e8361138a565b91506119b682611977565b602082019050919050565b5f6020820190508181035f8301526119d88161199f565b9050919050565b7f424c5331325f47324d534d2063616c6c206661696c65640000000000000000005f82015250565b5f611a1360178361138a565b9150611a1e826119df565b602082019050919050565b5f6020820190508181035f830152611a4081611a07565b9050919050565b7f726970656d643136302063616c6c206661696c656400000000000000000000005f82015250565b5f611a7b60158361138a565b9150611a8682611a47565b602082019050919050565b5f6020820190508181035f830152611aa881611a6f565b905091905056fea2646970667358221220ab33c2e7269410fdf25caca11593df9e877f1c9cb0a27b18a57429398484681d64736f6c634300081c0033",
}

// PrecompileABI is the input ABI used to generate the binding from.
// Deprecated: Use PrecompileMetaData.ABI instead.
var PrecompileABI = PrecompileMetaData.ABI

// PrecompileBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrecompileMetaData.Bin instead.
var PrecompileBin = PrecompileMetaData.Bin

// DeployPrecompile deploys a new Ethereum contract, binding an instance of Precompile to it.
func DeployPrecompile(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Precompile, error) {
	parsed, err := PrecompileMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrecompileBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Precompile{PrecompileCaller: PrecompileCaller{contract: contract}, PrecompileTransactor: PrecompileTransactor{contract: contract}, PrecompileFilterer: PrecompileFilterer{contract: contract}}, nil
}

// Precompile is an auto generated Go binding around an Ethereum contract.
type Precompile struct {
	PrecompileCaller     // Read-only binding to the contract
	PrecompileTransactor // Write-only binding to the contract
	PrecompileFilterer   // Log filterer for contract events
}

// PrecompileCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompileSession struct {
	Contract     *Precompile       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrecompileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompileCallerSession struct {
	Contract *PrecompileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PrecompileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompileTransactorSession struct {
	Contract     *PrecompileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PrecompileRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompileRaw struct {
	Contract *Precompile // Generic contract binding to access the raw methods on
}

// PrecompileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompileCallerRaw struct {
	Contract *PrecompileCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompileTransactorRaw struct {
	Contract *PrecompileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompile creates a new instance of Precompile, bound to a specific deployed contract.
func NewPrecompile(address common.Address, backend bind.ContractBackend) (*Precompile, error) {
	contract, err := bindPrecompile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Precompile{PrecompileCaller: PrecompileCaller{contract: contract}, PrecompileTransactor: PrecompileTransactor{contract: contract}, PrecompileFilterer: PrecompileFilterer{contract: contract}}, nil
}

// NewPrecompileCaller creates a new read-only instance of Precompile, bound to a specific deployed contract.
func NewPrecompileCaller(address common.Address, caller bind.ContractCaller) (*PrecompileCaller, error) {
	contract, err := bindPrecompile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileCaller{contract: contract}, nil
}

// NewPrecompileTransactor creates a new write-only instance of Precompile, bound to a specific deployed contract.
func NewPrecompileTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompileTransactor, error) {
	contract, err := bindPrecompile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileTransactor{contract: contract}, nil
}

// NewPrecompileFilterer creates a new log filterer instance of Precompile, bound to a specific deployed contract.
func NewPrecompileFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompileFilterer, error) {
	contract, err := bindPrecompile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompileFilterer{contract: contract}, nil
}

// bindPrecompile binds a generic wrapper to an already deployed contract.
func bindPrecompile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrecompileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompile *PrecompileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Precompile.Contract.PrecompileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompile *PrecompileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompile.Contract.PrecompileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompile *PrecompileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompile.Contract.PrecompileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Precompile *PrecompileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Precompile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Precompile *PrecompileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Precompile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Precompile *PrecompileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Precompile.Contract.contract.Transact(opts, method, params...)
}

// TestBLS12G1Add is a free data retrieval call binding the contract method 0xd1ca680b.
//
// Solidity: function testBLS12_G1Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBLS12G1Add(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBLS12_G1Add", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBLS12G1Add is a free data retrieval call binding the contract method 0xd1ca680b.
//
// Solidity: function testBLS12_G1Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBLS12G1Add(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G1Add(&_Precompile.CallOpts, data)
}

// TestBLS12G1Add is a free data retrieval call binding the contract method 0xd1ca680b.
//
// Solidity: function testBLS12_G1Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBLS12G1Add(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G1Add(&_Precompile.CallOpts, data)
}

// TestBLS12G1MSM is a free data retrieval call binding the contract method 0x4a6b4605.
//
// Solidity: function testBLS12_G1MSM(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBLS12G1MSM(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBLS12_G1MSM", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBLS12G1MSM is a free data retrieval call binding the contract method 0x4a6b4605.
//
// Solidity: function testBLS12_G1MSM(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBLS12G1MSM(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G1MSM(&_Precompile.CallOpts, data)
}

// TestBLS12G1MSM is a free data retrieval call binding the contract method 0x4a6b4605.
//
// Solidity: function testBLS12_G1MSM(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBLS12G1MSM(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G1MSM(&_Precompile.CallOpts, data)
}

// TestBLS12G2Add is a free data retrieval call binding the contract method 0x10a62c72.
//
// Solidity: function testBLS12_G2Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBLS12G2Add(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBLS12_G2Add", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBLS12G2Add is a free data retrieval call binding the contract method 0x10a62c72.
//
// Solidity: function testBLS12_G2Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBLS12G2Add(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G2Add(&_Precompile.CallOpts, data)
}

// TestBLS12G2Add is a free data retrieval call binding the contract method 0x10a62c72.
//
// Solidity: function testBLS12_G2Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBLS12G2Add(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G2Add(&_Precompile.CallOpts, data)
}

// TestBLS12G2MSM is a free data retrieval call binding the contract method 0xf5a7bad0.
//
// Solidity: function testBLS12_G2MSM(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBLS12G2MSM(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBLS12_G2MSM", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBLS12G2MSM is a free data retrieval call binding the contract method 0xf5a7bad0.
//
// Solidity: function testBLS12_G2MSM(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBLS12G2MSM(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G2MSM(&_Precompile.CallOpts, data)
}

// TestBLS12G2MSM is a free data retrieval call binding the contract method 0xf5a7bad0.
//
// Solidity: function testBLS12_G2MSM(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBLS12G2MSM(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12G2MSM(&_Precompile.CallOpts, data)
}

// TestBLS12MapFPtoG1 is a free data retrieval call binding the contract method 0xa67ecca6.
//
// Solidity: function testBLS12_MapFPtoG1(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBLS12MapFPtoG1(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBLS12_MapFPtoG1", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBLS12MapFPtoG1 is a free data retrieval call binding the contract method 0xa67ecca6.
//
// Solidity: function testBLS12_MapFPtoG1(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBLS12MapFPtoG1(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12MapFPtoG1(&_Precompile.CallOpts, data)
}

// TestBLS12MapFPtoG1 is a free data retrieval call binding the contract method 0xa67ecca6.
//
// Solidity: function testBLS12_MapFPtoG1(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBLS12MapFPtoG1(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLS12MapFPtoG1(&_Precompile.CallOpts, data)
}

// TestBLS12PairingCheck is a free data retrieval call binding the contract method 0xf5688603.
//
// Solidity: function testBLS12_PairingCheck(bytes data) view returns(bool)
func (_Precompile *PrecompileCaller) TestBLS12PairingCheck(opts *bind.CallOpts, data []byte) (bool, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBLS12_PairingCheck", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestBLS12PairingCheck is a free data retrieval call binding the contract method 0xf5688603.
//
// Solidity: function testBLS12_PairingCheck(bytes data) view returns(bool)
func (_Precompile *PrecompileSession) TestBLS12PairingCheck(data []byte) (bool, error) {
	return _Precompile.Contract.TestBLS12PairingCheck(&_Precompile.CallOpts, data)
}

// TestBLS12PairingCheck is a free data retrieval call binding the contract method 0xf5688603.
//
// Solidity: function testBLS12_PairingCheck(bytes data) view returns(bool)
func (_Precompile *PrecompileCallerSession) TestBLS12PairingCheck(data []byte) (bool, error) {
	return _Precompile.Contract.TestBLS12PairingCheck(&_Precompile.CallOpts, data)
}

// TestBLSS12MapFP2toG2 is a free data retrieval call binding the contract method 0xeb4a006b.
//
// Solidity: function testBLSS12_MapFP2toG2(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBLSS12MapFP2toG2(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBLSS12_MapFP2toG2", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBLSS12MapFP2toG2 is a free data retrieval call binding the contract method 0xeb4a006b.
//
// Solidity: function testBLSS12_MapFP2toG2(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBLSS12MapFP2toG2(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLSS12MapFP2toG2(&_Precompile.CallOpts, data)
}

// TestBLSS12MapFP2toG2 is a free data retrieval call binding the contract method 0xeb4a006b.
//
// Solidity: function testBLSS12_MapFP2toG2(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBLSS12MapFP2toG2(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBLSS12MapFP2toG2(&_Precompile.CallOpts, data)
}

// TestBlake2 is a free data retrieval call binding the contract method 0x533dcda6.
//
// Solidity: function testBlake2(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBlake2(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBlake2", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBlake2 is a free data retrieval call binding the contract method 0x533dcda6.
//
// Solidity: function testBlake2(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBlake2(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBlake2(&_Precompile.CallOpts, data)
}

// TestBlake2 is a free data retrieval call binding the contract method 0x533dcda6.
//
// Solidity: function testBlake2(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBlake2(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBlake2(&_Precompile.CallOpts, data)
}

// TestBn256Add is a free data retrieval call binding the contract method 0x2a45bb52.
//
// Solidity: function testBn256Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBn256Add(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBn256Add", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBn256Add is a free data retrieval call binding the contract method 0x2a45bb52.
//
// Solidity: function testBn256Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBn256Add(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBn256Add(&_Precompile.CallOpts, data)
}

// TestBn256Add is a free data retrieval call binding the contract method 0x2a45bb52.
//
// Solidity: function testBn256Add(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBn256Add(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBn256Add(&_Precompile.CallOpts, data)
}

// TestBn256Pairing is a free data retrieval call binding the contract method 0xaa63d4b6.
//
// Solidity: function testBn256Pairing(bytes data) view returns(bool)
func (_Precompile *PrecompileCaller) TestBn256Pairing(opts *bind.CallOpts, data []byte) (bool, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBn256Pairing", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestBn256Pairing is a free data retrieval call binding the contract method 0xaa63d4b6.
//
// Solidity: function testBn256Pairing(bytes data) view returns(bool)
func (_Precompile *PrecompileSession) TestBn256Pairing(data []byte) (bool, error) {
	return _Precompile.Contract.TestBn256Pairing(&_Precompile.CallOpts, data)
}

// TestBn256Pairing is a free data retrieval call binding the contract method 0xaa63d4b6.
//
// Solidity: function testBn256Pairing(bytes data) view returns(bool)
func (_Precompile *PrecompileCallerSession) TestBn256Pairing(data []byte) (bool, error) {
	return _Precompile.Contract.TestBn256Pairing(&_Precompile.CallOpts, data)
}

// TestBn256ScalarMul is a free data retrieval call binding the contract method 0xd99b39e4.
//
// Solidity: function testBn256ScalarMul(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestBn256ScalarMul(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testBn256ScalarMul", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestBn256ScalarMul is a free data retrieval call binding the contract method 0xd99b39e4.
//
// Solidity: function testBn256ScalarMul(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestBn256ScalarMul(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBn256ScalarMul(&_Precompile.CallOpts, data)
}

// TestBn256ScalarMul is a free data retrieval call binding the contract method 0xd99b39e4.
//
// Solidity: function testBn256ScalarMul(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestBn256ScalarMul(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestBn256ScalarMul(&_Precompile.CallOpts, data)
}

// TestEcrecover is a free data retrieval call binding the contract method 0x86c3f234.
//
// Solidity: function testEcrecover(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestEcrecover(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testEcrecover", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestEcrecover is a free data retrieval call binding the contract method 0x86c3f234.
//
// Solidity: function testEcrecover(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestEcrecover(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestEcrecover(&_Precompile.CallOpts, data)
}

// TestEcrecover is a free data retrieval call binding the contract method 0x86c3f234.
//
// Solidity: function testEcrecover(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestEcrecover(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestEcrecover(&_Precompile.CallOpts, data)
}

// TestIdentity is a free data retrieval call binding the contract method 0x9cce7cf9.
//
// Solidity: function testIdentity(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestIdentity(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testIdentity", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestIdentity is a free data retrieval call binding the contract method 0x9cce7cf9.
//
// Solidity: function testIdentity(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestIdentity(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestIdentity(&_Precompile.CallOpts, data)
}

// TestIdentity is a free data retrieval call binding the contract method 0x9cce7cf9.
//
// Solidity: function testIdentity(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestIdentity(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestIdentity(&_Precompile.CallOpts, data)
}

// TestModexp is a free data retrieval call binding the contract method 0x6fe7c1f0.
//
// Solidity: function testModexp(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestModexp(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testModexp", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestModexp is a free data retrieval call binding the contract method 0x6fe7c1f0.
//
// Solidity: function testModexp(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestModexp(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestModexp(&_Precompile.CallOpts, data)
}

// TestModexp is a free data retrieval call binding the contract method 0x6fe7c1f0.
//
// Solidity: function testModexp(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestModexp(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestModexp(&_Precompile.CallOpts, data)
}

// TestPointEvaluation is a free data retrieval call binding the contract method 0xdb77cf6d.
//
// Solidity: function testPointEvaluation(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestPointEvaluation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testPointEvaluation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestPointEvaluation is a free data retrieval call binding the contract method 0xdb77cf6d.
//
// Solidity: function testPointEvaluation(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestPointEvaluation(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestPointEvaluation(&_Precompile.CallOpts, data)
}

// TestPointEvaluation is a free data retrieval call binding the contract method 0xdb77cf6d.
//
// Solidity: function testPointEvaluation(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestPointEvaluation(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestPointEvaluation(&_Precompile.CallOpts, data)
}

// TestRipemd160 is a free data retrieval call binding the contract method 0xf6b0bbf7.
//
// Solidity: function testRipemd160(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestRipemd160(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testRipemd160", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestRipemd160 is a free data retrieval call binding the contract method 0xf6b0bbf7.
//
// Solidity: function testRipemd160(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestRipemd160(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestRipemd160(&_Precompile.CallOpts, data)
}

// TestRipemd160 is a free data retrieval call binding the contract method 0xf6b0bbf7.
//
// Solidity: function testRipemd160(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestRipemd160(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestRipemd160(&_Precompile.CallOpts, data)
}

// TestSha256 is a free data retrieval call binding the contract method 0xd020aeb7.
//
// Solidity: function testSha256(bytes data) view returns(bytes)
func (_Precompile *PrecompileCaller) TestSha256(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Precompile.contract.Call(opts, &out, "testSha256", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestSha256 is a free data retrieval call binding the contract method 0xd020aeb7.
//
// Solidity: function testSha256(bytes data) view returns(bytes)
func (_Precompile *PrecompileSession) TestSha256(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestSha256(&_Precompile.CallOpts, data)
}

// TestSha256 is a free data retrieval call binding the contract method 0xd020aeb7.
//
// Solidity: function testSha256(bytes data) view returns(bytes)
func (_Precompile *PrecompileCallerSession) TestSha256(data []byte) ([]byte, error) {
	return _Precompile.Contract.TestSha256(&_Precompile.CallOpts, data)
}
