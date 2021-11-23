// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// SmartChefInitializableMetaData contains all meta data concerning the SmartChefInitializable contract.
var SmartChefInitializableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenRecovered\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AdminTokenRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"NewPoolLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"NewRewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"NewStartAndEndBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RewardsStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SMART_CHEF_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accTokenPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasUserLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLimitPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"_stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRewardWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverWrongTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_hasUserLimit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"updatePoolLimitPerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"updateRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"name\":\"updateStartAndEndBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600061002161010c60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506001808190555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610114565b600033905090565b612c9880620001246000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80638da5cb5b116100f9578063bd61719111610097578063db2e21bc11610071578063db2e21bc146105f3578063f2fde38b146105fd578063f40f0f5214610641578063f7c618c114610699576101a9565b8063bd6171911461056d578063cc7a262e146105a1578063ccd34cd5146105d5576101a9565b80639513997f116100d35780639513997f146104af578063a0b40905146104e7578063a9f8d18114610521578063b6b55f251461053f576101a9565b80638da5cb5b1461043d5780638f6629151461047157806392e8990e1461048f576101a9565b80633f138d4b11610166578063715018a611610140578063715018a61461035f57806380dc06721461036957806383a5041c146103735780638ae39cac1461041f576101a9565b80633f138d4b146102d557806348cd4cb11461032357806366fe9f8a14610341576101a9565b806301f8a976146101ae5780631959a002146101dc5780631aed65531461023b5780632e1a7d4d146102595780633279beab14610287578063392e53cd146102b5575b600080fd5b6101da600480360360208110156101c457600080fd5b81019080803590602001909291905050506106cd565b005b61021e600480360360208110156101f257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610834565b604051808381526020018281526020019250505060405180910390f35b610243610858565b6040518082815260200191505060405180910390f35b6102856004803603602081101561026f57600080fd5b810190808035906020019092919050505061085e565b005b6102b36004803603602081101561029d57600080fd5b8101908080359060200190929190505050610b47565b005b6102bd610c46565b60405180821515815260200191505060405180910390f35b610321600480360360408110156102eb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c59565b005b61032b610f14565b6040518082815260200191505060405180910390f35b610349610f1a565b6040518082815260200191505060405180910390f35b610367610f20565b005b61037161108d565b005b61041d600480360360e081101561038957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611145565b005b6104276114c2565b6040518082815260200191505060405180910390f35b6104456114c8565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104796114f1565b6040518082815260200191505060405180910390f35b6104976114f7565b60405180821515815260200191505060405180910390f35b6104e5600480360360408110156104c557600080fd5b81019080803590602001909291908035906020019092919050505061150a565b005b61051f600480360360408110156104fd57600080fd5b810190808035151590602001909291908035906020019092919050505061173a565b005b610529611954565b6040518082815260200191505060405180910390f35b61056b6004803603602081101561055557600080fd5b810190808035906020019092919050505061195a565b005b610575611c7d565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105a9611ca3565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105dd611cc9565b6040518082815260200191505060405180910390f35b6105fb611ccf565b005b61063f6004803603602081101561061357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e64565b005b6106836004803603602081101561065757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612056565b6040518082815260200191505060405180910390f35b6106a161227b565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6106d56122a1565b73ffffffffffffffffffffffffffffffffffffffff166106f36114c8565b73ffffffffffffffffffffffffffffffffffffffff161461077c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b60055443106107f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f506f6f6c2068617320737461727465640000000000000000000000000000000081525060200191505060405180910390fd5b806008819055507f0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df816040518082815260200191505060405180910390a150565b600c6020528060005260406000206000915090508060000154908060010154905082565b60045481565b600260015414156108d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0081525060200191505060405180910390fd5b60026001819055506000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050818160000154101561099c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f416d6f756e7420746f20776974686472617720746f6f2068696768000000000081525060200191505060405180910390fd5b6109a46122a9565b60006109e982600101546109db6009546109cd600354876000015461240c90919063ffffffff16565b61249290919063ffffffff16565b61251b90919063ffffffff16565b90506000831115610a6157610a0b83836000015461251b90919063ffffffff16565b8260000181905550610a603384600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661259e9092919063ffffffff16565b5b6000811115610ab857610ab73382600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661259e9092919063ffffffff16565b5b610ae5600954610ad7600354856000015461240c90919063ffffffff16565b61249290919063ffffffff16565b82600101819055503373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364846040518082815260200191505060405180910390a250506001808190555050565b610b4f6122a1565b73ffffffffffffffffffffffffffffffffffffffff16610b6d6114c8565b73ffffffffffffffffffffffffffffffffffffffff1614610bf6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b610c433382600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661259e9092919063ffffffff16565b50565b600260159054906101000a900460ff1681565b610c616122a1565b73ffffffffffffffffffffffffffffffffffffffff16610c7f6114c8565b73ffffffffffffffffffffffffffffffffffffffff1614610d08576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610dcc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f43616e6e6f74206265207374616b656420746f6b656e0000000000000000000081525060200191505060405180910390fd5b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e90576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f43616e6e6f742062652072657761726420746f6b656e0000000000000000000081525060200191505060405180910390fd5b610ebb33828473ffffffffffffffffffffffffffffffffffffffff1661259e9092919063ffffffff16565b7f74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab781298282604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15050565b60055481565b60075481565b610f286122a1565b73ffffffffffffffffffffffffffffffffffffffff16610f466114c8565b73ffffffffffffffffffffffffffffffffffffffff1614610fcf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6110956122a1565b73ffffffffffffffffffffffffffffffffffffffff166110b36114c8565b73ffffffffffffffffffffffffffffffffffffffff161461113c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b43600481905550565b600260159054906101000a900460ff16156111c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f416c726561647920696e697469616c697a65640000000000000000000000000081525060200191505060405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461128b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f4e6f7420666163746f727900000000000000000000000000000000000000000081525060200191505060405180910390fd5b6001600260156101000a81548160ff02191690831515021790555086600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460088190555083600581905550826004819055506000821115611369576001600260146101000a81548160ff021916908315150217905550816007819055505b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156113d357600080fd5b505afa1580156113e7573d6000803e3d6000fd5b505050506040513d60208110156113fd57600080fd5b810190808051906020019092919050505060ff169050601e8110611489576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f4d75737420626520696e666572696f7220746f2033300000000000000000000081525060200191505060405180910390fd5b61149d81601e61251b90919063ffffffff16565b600a0a6009819055506005546006819055506114b882611e64565b5050505050505050565b60085481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60035481565b600260149054906101000a900460ff1681565b6115126122a1565b73ffffffffffffffffffffffffffffffffffffffff166115306114c8565b73ffffffffffffffffffffffffffffffffffffffff16146115b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6005544310611630576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f506f6f6c2068617320737461727465640000000000000000000000000000000081525060200191505060405180910390fd5b808210611688576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180612bbe602e913960400191505060405180910390fd5b8143106116e0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180612c126030913960400191505060405180910390fd5b81600581905550806004819055506005546006819055507f7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce068282604051808381526020018281526020019250505060405180910390a15050565b6117426122a1565b73ffffffffffffffffffffffffffffffffffffffff166117606114c8565b73ffffffffffffffffffffffffffffffffffffffff16146117e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600260149054906101000a900460ff1661186b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f4d7573742062652073657400000000000000000000000000000000000000000081525060200191505060405180910390fd5b81156118f45760075481116118e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f4e6577206c696d6974206d75737420626520686967686572000000000000000081525060200191505060405180910390fd5b80600781905550611917565b81600260146101000a81548160ff02191690831515021790555060006007819055505b7f241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c6007546040518082815260200191505060405180910390a15050565b60065481565b600260015414156119d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0081525060200191505060405180910390fd5b60026001819055506000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600260149054906101000a900460ff1615611ac257600754611a4d82600001548461264090919063ffffffff16565b1115611ac1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f5573657220616d6f756e742061626f7665206c696d697400000000000000000081525060200191505060405180910390fd5b5b611aca6122a9565b600081600001541115611b77576000611b1c8260010154611b0e600954611b00600354876000015461240c90919063ffffffff16565b61249290919063ffffffff16565b61251b90919063ffffffff16565b90506000811115611b7557611b743382600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661259e9092919063ffffffff16565b5b505b6000821115611bef57611b9782826000015461264090919063ffffffff16565b8160000181905550611bee333084600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166126c8909392919063ffffffff16565b5b611c1c600954611c0e600354846000015461240c90919063ffffffff16565b61249290919063ffffffff16565b81600101819055503373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c836040518082815260200191505060405180910390a2506001808190555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60026001541415611d48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0081525060200191505060405180910390fd5b60026001819055506000600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160000154905060008260000181905550600082600101819055506000811115611e0757611e063382600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661259e9092919063ffffffff16565b5b3373ffffffffffffffffffffffffffffffffffffffff167f5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd969583600001546040518082815260200191505060405180910390a2505060018081905550565b611e6c6122a1565b73ffffffffffffffffffffffffffffffffffffffff16611e8a6114c8565b73ffffffffffffffffffffffffffffffffffffffff1614611f13576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611f99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612b986026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080600c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561212557600080fd5b505afa158015612139573d6000803e3d6000fd5b505050506040513d602081101561214f57600080fd5b8101908080519060200190929190505050905060065443118015612174575060008114155b1561222e57600061218760065443612789565b905060006121a06008548361240c90919063ffffffff16565b905060006121df6121ce856121c06009548661240c90919063ffffffff16565b61249290919063ffffffff16565b60035461264090919063ffffffff16565b90506122228560010154612214600954612206858a6000015461240c90919063ffffffff16565b61249290919063ffffffff16565b61251b90919063ffffffff16565b95505050505050612276565b6122718260010154612263600954612255600354876000015461240c90919063ffffffff16565b61249290919063ffffffff16565b61251b90919063ffffffff16565b925050505b919050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b60065443116122b75761240a565b6000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561234257600080fd5b505afa158015612356573d6000803e3d6000fd5b505050506040513d602081101561236c57600080fd5b81019080805190602001909291905050509050600081141561239557436006819055505061240a565b60006123a360065443612789565b905060006123bc6008548361240c90919063ffffffff16565b90506123f96123e8846123da6009548561240c90919063ffffffff16565b61249290919063ffffffff16565b60035461264090919063ffffffff16565b600381905550436006819055505050505b565b60008083141561241f576000905061248c565b600082840290508284828161243057fe5b0414612487576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612c426021913960400191505060405180910390fd5b809150505b92915050565b6000808211612509576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525060200191505060405180910390fd5b81838161251257fe5b04905092915050565b600082821115612593576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525060200191505060405180910390fd5b818303905092915050565b61263b8363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506127de565b505050565b6000808284019050838110156126be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b612783846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506127de565b50505050565b600060045482116127ae576127a7838361251b90919063ffffffff16565b90506127d8565b60045483106127c057600090506127d8565b6127d58360045461251b90919063ffffffff16565b90505b92915050565b6060612840826040518060400160405280602081526020017f5361666542455032303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166128cd9092919063ffffffff16565b90506000815111156128c85780806020019051602081101561286157600080fd5b81019080805190602001909291905050506128c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612b6e602a913960400191505060405180910390fd5b5b505050565b60606128dc84846000856128e5565b90509392505050565b606082471015612940576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180612bec6026913960400191505060405180910390fd5b61294985612a8e565b6129bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b60208310612a0b57805182526020820191506020810190506020830392506129e8565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114612a6d576040519150601f19603f3d011682016040523d82523d6000602084013e612a72565b606091505b5091509150612a82828286612aa1565b92505050949350505050565b600080823b905060008111915050919050565b60608315612ab157829050612b66565b600083511115612ac45782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015612b2b578082015181840152602081019050612b10565b50505050905090810190601f168015612b585780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b939250505056fe5361666542455032303a204245503230206f7065726174696f6e20646964206e6f7420737563636565644f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734e6577207374617274426c6f636b206d757374206265206c6f776572207468616e206e657720656e64426c6f636b416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c4e6577207374617274426c6f636b206d75737420626520686967686572207468616e2063757272656e7420626c6f636b536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a264697066735822122074fc55607a0bc2f237d4f11275c833dba709f0fa8c5aadceaa49caf7c28cd1cc64736f6c634300060c0033",
}

// SmartChefInitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartChefInitializableMetaData.ABI instead.
var SmartChefInitializableABI = SmartChefInitializableMetaData.ABI

// SmartChefInitializableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SmartChefInitializableMetaData.Bin instead.
var SmartChefInitializableBin = SmartChefInitializableMetaData.Bin

// DeploySmartChefInitializable deploys a new Ethereum contract, binding an instance of SmartChefInitializable to it.
func DeploySmartChefInitializable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SmartChefInitializable, error) {
	parsed, err := SmartChefInitializableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SmartChefInitializableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmartChefInitializable{SmartChefInitializableCaller: SmartChefInitializableCaller{contract: contract}, SmartChefInitializableTransactor: SmartChefInitializableTransactor{contract: contract}, SmartChefInitializableFilterer: SmartChefInitializableFilterer{contract: contract}}, nil
}

// SmartChefInitializable is an auto generated Go binding around an Ethereum contract.
type SmartChefInitializable struct {
	SmartChefInitializableCaller     // Read-only binding to the contract
	SmartChefInitializableTransactor // Write-only binding to the contract
	SmartChefInitializableFilterer   // Log filterer for contract events
}

// SmartChefInitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartChefInitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartChefInitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartChefInitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartChefInitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartChefInitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartChefInitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartChefInitializableSession struct {
	Contract     *SmartChefInitializable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SmartChefInitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartChefInitializableCallerSession struct {
	Contract *SmartChefInitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// SmartChefInitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartChefInitializableTransactorSession struct {
	Contract     *SmartChefInitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SmartChefInitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartChefInitializableRaw struct {
	Contract *SmartChefInitializable // Generic contract binding to access the raw methods on
}

// SmartChefInitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartChefInitializableCallerRaw struct {
	Contract *SmartChefInitializableCaller // Generic read-only contract binding to access the raw methods on
}

// SmartChefInitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartChefInitializableTransactorRaw struct {
	Contract *SmartChefInitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartChefInitializable creates a new instance of SmartChefInitializable, bound to a specific deployed contract.
func NewSmartChefInitializable(address common.Address, backend bind.ContractBackend) (*SmartChefInitializable, error) {
	contract, err := bindSmartChefInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializable{SmartChefInitializableCaller: SmartChefInitializableCaller{contract: contract}, SmartChefInitializableTransactor: SmartChefInitializableTransactor{contract: contract}, SmartChefInitializableFilterer: SmartChefInitializableFilterer{contract: contract}}, nil
}

// NewSmartChefInitializableCaller creates a new read-only instance of SmartChefInitializable, bound to a specific deployed contract.
func NewSmartChefInitializableCaller(address common.Address, caller bind.ContractCaller) (*SmartChefInitializableCaller, error) {
	contract, err := bindSmartChefInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableCaller{contract: contract}, nil
}

// NewSmartChefInitializableTransactor creates a new write-only instance of SmartChefInitializable, bound to a specific deployed contract.
func NewSmartChefInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartChefInitializableTransactor, error) {
	contract, err := bindSmartChefInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableTransactor{contract: contract}, nil
}

// NewSmartChefInitializableFilterer creates a new log filterer instance of SmartChefInitializable, bound to a specific deployed contract.
func NewSmartChefInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartChefInitializableFilterer, error) {
	contract, err := bindSmartChefInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableFilterer{contract: contract}, nil
}

// bindSmartChefInitializable binds a generic wrapper to an already deployed contract.
func bindSmartChefInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartChefInitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartChefInitializable *SmartChefInitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartChefInitializable.Contract.SmartChefInitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartChefInitializable *SmartChefInitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.SmartChefInitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartChefInitializable *SmartChefInitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.SmartChefInitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartChefInitializable *SmartChefInitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmartChefInitializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartChefInitializable *SmartChefInitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartChefInitializable *SmartChefInitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.contract.Transact(opts, method, params...)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SmartChefInitializable.Contract.PRECISIONFACTOR(&_SmartChefInitializable.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SmartChefInitializable.Contract.PRECISIONFACTOR(&_SmartChefInitializable.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCaller) SMARTCHEFFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "SMART_CHEF_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _SmartChefInitializable.Contract.SMARTCHEFFACTORY(&_SmartChefInitializable.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _SmartChefInitializable.Contract.SMARTCHEFFACTORY(&_SmartChefInitializable.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) AccTokenPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "accTokenPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) AccTokenPerShare() (*big.Int, error) {
	return _SmartChefInitializable.Contract.AccTokenPerShare(&_SmartChefInitializable.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) AccTokenPerShare() (*big.Int, error) {
	return _SmartChefInitializable.Contract.AccTokenPerShare(&_SmartChefInitializable.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) BonusEndBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.BonusEndBlock(&_SmartChefInitializable.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) BonusEndBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.BonusEndBlock(&_SmartChefInitializable.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_SmartChefInitializable *SmartChefInitializableCaller) HasUserLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "hasUserLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_SmartChefInitializable *SmartChefInitializableSession) HasUserLimit() (bool, error) {
	return _SmartChefInitializable.Contract.HasUserLimit(&_SmartChefInitializable.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) HasUserLimit() (bool, error) {
	return _SmartChefInitializable.Contract.HasUserLimit(&_SmartChefInitializable.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SmartChefInitializable *SmartChefInitializableCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SmartChefInitializable *SmartChefInitializableSession) IsInitialized() (bool, error) {
	return _SmartChefInitializable.Contract.IsInitialized(&_SmartChefInitializable.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) IsInitialized() (bool, error) {
	return _SmartChefInitializable.Contract.IsInitialized(&_SmartChefInitializable.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) LastRewardBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.LastRewardBlock(&_SmartChefInitializable.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) LastRewardBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.LastRewardBlock(&_SmartChefInitializable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableSession) Owner() (common.Address, error) {
	return _SmartChefInitializable.Contract.Owner(&_SmartChefInitializable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) Owner() (common.Address, error) {
	return _SmartChefInitializable.Contract.Owner(&_SmartChefInitializable.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) PendingReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "pendingReward", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _SmartChefInitializable.Contract.PendingReward(&_SmartChefInitializable.CallOpts, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _SmartChefInitializable.Contract.PendingReward(&_SmartChefInitializable.CallOpts, _user)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) PoolLimitPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "poolLimitPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) PoolLimitPerUser() (*big.Int, error) {
	return _SmartChefInitializable.Contract.PoolLimitPerUser(&_SmartChefInitializable.CallOpts)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) PoolLimitPerUser() (*big.Int, error) {
	return _SmartChefInitializable.Contract.PoolLimitPerUser(&_SmartChefInitializable.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "rewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) RewardPerBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.RewardPerBlock(&_SmartChefInitializable.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) RewardPerBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.RewardPerBlock(&_SmartChefInitializable.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableSession) RewardToken() (common.Address, error) {
	return _SmartChefInitializable.Contract.RewardToken(&_SmartChefInitializable.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) RewardToken() (common.Address, error) {
	return _SmartChefInitializable.Contract.RewardToken(&_SmartChefInitializable.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCaller) StakedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "stakedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableSession) StakedToken() (common.Address, error) {
	return _SmartChefInitializable.Contract.StakedToken(&_SmartChefInitializable.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) StakedToken() (common.Address, error) {
	return _SmartChefInitializable.Contract.StakedToken(&_SmartChefInitializable.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableSession) StartBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.StartBlock(&_SmartChefInitializable.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) StartBlock() (*big.Int, error) {
	return _SmartChefInitializable.Contract.StartBlock(&_SmartChefInitializable.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SmartChefInitializable *SmartChefInitializableCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _SmartChefInitializable.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SmartChefInitializable *SmartChefInitializableSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SmartChefInitializable.Contract.UserInfo(&_SmartChefInitializable.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SmartChefInitializable *SmartChefInitializableCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SmartChefInitializable.Contract.UserInfo(&_SmartChefInitializable.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.Deposit(&_SmartChefInitializable.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.Deposit(&_SmartChefInitializable.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) EmergencyRewardWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "emergencyRewardWithdraw", _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.EmergencyRewardWithdraw(&_SmartChefInitializable.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.EmergencyRewardWithdraw(&_SmartChefInitializable.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SmartChefInitializable *SmartChefInitializableSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.EmergencyWithdraw(&_SmartChefInitializable.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.EmergencyWithdraw(&_SmartChefInitializable.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) Initialize(opts *bind.TransactOpts, _stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "initialize", _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.Initialize(&_SmartChefInitializable.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x83a5041c.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, address _admin) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.Initialize(&_SmartChefInitializable.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _admin)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) RecoverWrongTokens(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "recoverWrongTokens", _tokenAddress, _tokenAmount)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) RecoverWrongTokens(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.RecoverWrongTokens(&_SmartChefInitializable.TransactOpts, _tokenAddress, _tokenAmount)
}

// RecoverWrongTokens is a paid mutator transaction binding the contract method 0x3f138d4b.
//
// Solidity: function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) RecoverWrongTokens(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.RecoverWrongTokens(&_SmartChefInitializable.TransactOpts, _tokenAddress, _tokenAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartChefInitializable *SmartChefInitializableSession) RenounceOwnership() (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.RenounceOwnership(&_SmartChefInitializable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.RenounceOwnership(&_SmartChefInitializable.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) StopReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "stopReward")
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_SmartChefInitializable *SmartChefInitializableSession) StopReward() (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.StopReward(&_SmartChefInitializable.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) StopReward() (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.StopReward(&_SmartChefInitializable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.TransferOwnership(&_SmartChefInitializable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.TransferOwnership(&_SmartChefInitializable.TransactOpts, newOwner)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) UpdatePoolLimitPerUser(opts *bind.TransactOpts, _hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "updatePoolLimitPerUser", _hasUserLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) UpdatePoolLimitPerUser(_hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.UpdatePoolLimitPerUser(&_SmartChefInitializable.TransactOpts, _hasUserLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _hasUserLimit, uint256 _poolLimitPerUser) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) UpdatePoolLimitPerUser(_hasUserLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.UpdatePoolLimitPerUser(&_SmartChefInitializable.TransactOpts, _hasUserLimit, _poolLimitPerUser)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) UpdateRewardPerBlock(opts *bind.TransactOpts, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "updateRewardPerBlock", _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.UpdateRewardPerBlock(&_SmartChefInitializable.TransactOpts, _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.UpdateRewardPerBlock(&_SmartChefInitializable.TransactOpts, _rewardPerBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) UpdateStartAndEndBlocks(opts *bind.TransactOpts, _startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "updateStartAndEndBlocks", _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.UpdateStartAndEndBlocks(&_SmartChefInitializable.TransactOpts, _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.UpdateStartAndEndBlocks(&_SmartChefInitializable.TransactOpts, _startBlock, _bonusEndBlock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.Withdraw(&_SmartChefInitializable.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_SmartChefInitializable *SmartChefInitializableTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _SmartChefInitializable.Contract.Withdraw(&_SmartChefInitializable.TransactOpts, _amount)
}

// SmartChefInitializableAdminTokenRecoveryIterator is returned from FilterAdminTokenRecovery and is used to iterate over the raw logs and unpacked data for AdminTokenRecovery events raised by the SmartChefInitializable contract.
type SmartChefInitializableAdminTokenRecoveryIterator struct {
	Event *SmartChefInitializableAdminTokenRecovery // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableAdminTokenRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableAdminTokenRecovery)
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
		it.Event = new(SmartChefInitializableAdminTokenRecovery)
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
func (it *SmartChefInitializableAdminTokenRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableAdminTokenRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableAdminTokenRecovery represents a AdminTokenRecovery event raised by the SmartChefInitializable contract.
type SmartChefInitializableAdminTokenRecovery struct {
	TokenRecovered common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAdminTokenRecovery is a free log retrieval operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterAdminTokenRecovery(opts *bind.FilterOpts) (*SmartChefInitializableAdminTokenRecoveryIterator, error) {

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "AdminTokenRecovery")
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableAdminTokenRecoveryIterator{contract: _SmartChefInitializable.contract, event: "AdminTokenRecovery", logs: logs, sub: sub}, nil
}

// WatchAdminTokenRecovery is a free log subscription operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchAdminTokenRecovery(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableAdminTokenRecovery) (event.Subscription, error) {

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "AdminTokenRecovery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableAdminTokenRecovery)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "AdminTokenRecovery", log); err != nil {
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

// ParseAdminTokenRecovery is a log parse operation binding the contract event 0x74545154aac348a3eac92596bd1971957ca94795f4e954ec5f613b55fab78129.
//
// Solidity: event AdminTokenRecovery(address tokenRecovered, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseAdminTokenRecovery(log types.Log) (*SmartChefInitializableAdminTokenRecovery, error) {
	event := new(SmartChefInitializableAdminTokenRecovery)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "AdminTokenRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SmartChefInitializable contract.
type SmartChefInitializableDepositIterator struct {
	Event *SmartChefInitializableDeposit // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableDeposit)
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
		it.Event = new(SmartChefInitializableDeposit)
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
func (it *SmartChefInitializableDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableDeposit represents a Deposit event raised by the SmartChefInitializable contract.
type SmartChefInitializableDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*SmartChefInitializableDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableDepositIterator{contract: _SmartChefInitializable.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableDeposit)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseDeposit(log types.Log) (*SmartChefInitializableDeposit, error) {
	event := new(SmartChefInitializableDeposit)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the SmartChefInitializable contract.
type SmartChefInitializableEmergencyWithdrawIterator struct {
	Event *SmartChefInitializableEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableEmergencyWithdraw)
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
		it.Event = new(SmartChefInitializableEmergencyWithdraw)
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
func (it *SmartChefInitializableEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableEmergencyWithdraw represents a EmergencyWithdraw event raised by the SmartChefInitializable contract.
type SmartChefInitializableEmergencyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address) (*SmartChefInitializableEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableEmergencyWithdrawIterator{contract: _SmartChefInitializable.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableEmergencyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableEmergencyWithdraw)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseEmergencyWithdraw(log types.Log) (*SmartChefInitializableEmergencyWithdraw, error) {
	event := new(SmartChefInitializableEmergencyWithdraw)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableNewPoolLimitIterator is returned from FilterNewPoolLimit and is used to iterate over the raw logs and unpacked data for NewPoolLimit events raised by the SmartChefInitializable contract.
type SmartChefInitializableNewPoolLimitIterator struct {
	Event *SmartChefInitializableNewPoolLimit // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableNewPoolLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableNewPoolLimit)
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
		it.Event = new(SmartChefInitializableNewPoolLimit)
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
func (it *SmartChefInitializableNewPoolLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableNewPoolLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableNewPoolLimit represents a NewPoolLimit event raised by the SmartChefInitializable contract.
type SmartChefInitializableNewPoolLimit struct {
	PoolLimitPerUser *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPoolLimit is a free log retrieval operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterNewPoolLimit(opts *bind.FilterOpts) (*SmartChefInitializableNewPoolLimitIterator, error) {

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableNewPoolLimitIterator{contract: _SmartChefInitializable.contract, event: "NewPoolLimit", logs: logs, sub: sub}, nil
}

// WatchNewPoolLimit is a free log subscription operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchNewPoolLimit(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableNewPoolLimit) (event.Subscription, error) {

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableNewPoolLimit)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
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

// ParseNewPoolLimit is a log parse operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseNewPoolLimit(log types.Log) (*SmartChefInitializableNewPoolLimit, error) {
	event := new(SmartChefInitializableNewPoolLimit)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableNewRewardPerBlockIterator is returned from FilterNewRewardPerBlock and is used to iterate over the raw logs and unpacked data for NewRewardPerBlock events raised by the SmartChefInitializable contract.
type SmartChefInitializableNewRewardPerBlockIterator struct {
	Event *SmartChefInitializableNewRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableNewRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableNewRewardPerBlock)
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
		it.Event = new(SmartChefInitializableNewRewardPerBlock)
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
func (it *SmartChefInitializableNewRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableNewRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableNewRewardPerBlock represents a NewRewardPerBlock event raised by the SmartChefInitializable contract.
type SmartChefInitializableNewRewardPerBlock struct {
	RewardPerBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRewardPerBlock is a free log retrieval operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterNewRewardPerBlock(opts *bind.FilterOpts) (*SmartChefInitializableNewRewardPerBlockIterator, error) {

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableNewRewardPerBlockIterator{contract: _SmartChefInitializable.contract, event: "NewRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchNewRewardPerBlock is a free log subscription operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchNewRewardPerBlock(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableNewRewardPerBlock) (event.Subscription, error) {

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableNewRewardPerBlock)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
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

// ParseNewRewardPerBlock is a log parse operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseNewRewardPerBlock(log types.Log) (*SmartChefInitializableNewRewardPerBlock, error) {
	event := new(SmartChefInitializableNewRewardPerBlock)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableNewStartAndEndBlocksIterator is returned from FilterNewStartAndEndBlocks and is used to iterate over the raw logs and unpacked data for NewStartAndEndBlocks events raised by the SmartChefInitializable contract.
type SmartChefInitializableNewStartAndEndBlocksIterator struct {
	Event *SmartChefInitializableNewStartAndEndBlocks // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableNewStartAndEndBlocksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableNewStartAndEndBlocks)
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
		it.Event = new(SmartChefInitializableNewStartAndEndBlocks)
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
func (it *SmartChefInitializableNewStartAndEndBlocksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableNewStartAndEndBlocksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableNewStartAndEndBlocks represents a NewStartAndEndBlocks event raised by the SmartChefInitializable contract.
type SmartChefInitializableNewStartAndEndBlocks struct {
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewStartAndEndBlocks is a free log retrieval operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterNewStartAndEndBlocks(opts *bind.FilterOpts) (*SmartChefInitializableNewStartAndEndBlocksIterator, error) {

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableNewStartAndEndBlocksIterator{contract: _SmartChefInitializable.contract, event: "NewStartAndEndBlocks", logs: logs, sub: sub}, nil
}

// WatchNewStartAndEndBlocks is a free log subscription operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchNewStartAndEndBlocks(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableNewStartAndEndBlocks) (event.Subscription, error) {

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableNewStartAndEndBlocks)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
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

// ParseNewStartAndEndBlocks is a log parse operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseNewStartAndEndBlocks(log types.Log) (*SmartChefInitializableNewStartAndEndBlocks, error) {
	event := new(SmartChefInitializableNewStartAndEndBlocks)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SmartChefInitializable contract.
type SmartChefInitializableOwnershipTransferredIterator struct {
	Event *SmartChefInitializableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableOwnershipTransferred)
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
		it.Event = new(SmartChefInitializableOwnershipTransferred)
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
func (it *SmartChefInitializableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableOwnershipTransferred represents a OwnershipTransferred event raised by the SmartChefInitializable contract.
type SmartChefInitializableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SmartChefInitializableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableOwnershipTransferredIterator{contract: _SmartChefInitializable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableOwnershipTransferred)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseOwnershipTransferred(log types.Log) (*SmartChefInitializableOwnershipTransferred, error) {
	event := new(SmartChefInitializableOwnershipTransferred)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableRewardsStopIterator is returned from FilterRewardsStop and is used to iterate over the raw logs and unpacked data for RewardsStop events raised by the SmartChefInitializable contract.
type SmartChefInitializableRewardsStopIterator struct {
	Event *SmartChefInitializableRewardsStop // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableRewardsStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableRewardsStop)
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
		it.Event = new(SmartChefInitializableRewardsStop)
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
func (it *SmartChefInitializableRewardsStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableRewardsStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableRewardsStop represents a RewardsStop event raised by the SmartChefInitializable contract.
type SmartChefInitializableRewardsStop struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsStop is a free log retrieval operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterRewardsStop(opts *bind.FilterOpts) (*SmartChefInitializableRewardsStopIterator, error) {

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableRewardsStopIterator{contract: _SmartChefInitializable.contract, event: "RewardsStop", logs: logs, sub: sub}, nil
}

// WatchRewardsStop is a free log subscription operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchRewardsStop(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableRewardsStop) (event.Subscription, error) {

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableRewardsStop)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "RewardsStop", log); err != nil {
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

// ParseRewardsStop is a log parse operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseRewardsStop(log types.Log) (*SmartChefInitializableRewardsStop, error) {
	event := new(SmartChefInitializableRewardsStop)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "RewardsStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmartChefInitializableWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SmartChefInitializable contract.
type SmartChefInitializableWithdrawIterator struct {
	Event *SmartChefInitializableWithdraw // Event containing the contract specifics and raw log

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
func (it *SmartChefInitializableWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartChefInitializableWithdraw)
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
		it.Event = new(SmartChefInitializableWithdraw)
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
func (it *SmartChefInitializableWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartChefInitializableWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartChefInitializableWithdraw represents a Withdraw event raised by the SmartChefInitializable contract.
type SmartChefInitializableWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*SmartChefInitializableWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &SmartChefInitializableWithdrawIterator{contract: _SmartChefInitializable.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SmartChefInitializableWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SmartChefInitializable.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartChefInitializableWithdraw)
				if err := _SmartChefInitializable.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_SmartChefInitializable *SmartChefInitializableFilterer) ParseWithdraw(log types.Log) (*SmartChefInitializableWithdraw, error) {
	event := new(SmartChefInitializableWithdraw)
	if err := _SmartChefInitializable.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
