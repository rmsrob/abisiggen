// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SablierV2LockupLinear

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

// Broker is an auto generated low-level Go binding around an user-defined struct.
type Broker struct {
	Account common.Address
	Fee     *big.Int
}

// LockupAmounts is an auto generated low-level Go binding around an user-defined struct.
type LockupAmounts struct {
	Deposited *big.Int
	Withdrawn *big.Int
	Refunded  *big.Int
}

// LockupCreateAmounts is an auto generated low-level Go binding around an user-defined struct.
type LockupCreateAmounts struct {
	Deposit   *big.Int
	BrokerFee *big.Int
}

// LockupLinearCreateWithDurations is an auto generated low-level Go binding around an user-defined struct.
type LockupLinearCreateWithDurations struct {
	Sender       common.Address
	Recipient    common.Address
	TotalAmount  *big.Int
	Asset        common.Address
	Cancelable   bool
	Transferable bool
	Durations    LockupLinearDurations
	Broker       Broker
}

// LockupLinearCreateWithTimestamps is an auto generated low-level Go binding around an user-defined struct.
type LockupLinearCreateWithTimestamps struct {
	Sender       common.Address
	Recipient    common.Address
	TotalAmount  *big.Int
	Asset        common.Address
	Cancelable   bool
	Transferable bool
	Timestamps   LockupLinearTimestamps
	Broker       Broker
}

// LockupLinearDurations is an auto generated low-level Go binding around an user-defined struct.
type LockupLinearDurations struct {
	Cliff *big.Int
	Total *big.Int
}

// LockupLinearStreamLL is an auto generated low-level Go binding around an user-defined struct.
type LockupLinearStreamLL struct {
	Sender         common.Address
	Recipient      common.Address
	StartTime      *big.Int
	IsCancelable   bool
	WasCanceled    bool
	Asset          common.Address
	EndTime        *big.Int
	IsDepleted     bool
	IsStream       bool
	IsTransferable bool
	Amounts        LockupAmounts
	CliffTime      *big.Int
}

// LockupLinearTimestamps is an auto generated low-level Go binding around an user-defined struct.
type LockupLinearTimestamps struct {
	Start *big.Int
	Cliff *big.Int
	End   *big.Int
}

// SablierV2LockupLinearMetaData contains all meta data concerning the SablierV2LockupLinear contract.
var SablierV2LockupLinearMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialAdmin\",\"type\":\"address\"},{\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"initialNFTDescriptor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerNotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelegateCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"PRBMath_MulDiv18_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"PRBMath_MulDiv_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"cliffTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"}],\"name\":\"SablierV2LockupLinear_CliffTimeNotLessThanEndTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"cliffTime\",\"type\":\"uint40\"}],\"name\":\"SablierV2LockupLinear_StartTimeNotLessThanCliffTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"}],\"name\":\"SablierV2LockupLinear_StartTimeNotLessThanEndTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_AllowToHookUnsupportedInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_AllowToHookZeroCodeSize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"UD60x18\",\"name\":\"brokerFee\",\"type\":\"uint256\"},{\"internalType\":\"UD60x18\",\"name\":\"maxBrokerFee\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_BrokerFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SablierV2Lockup_DepositAmountZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"blockTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"}],\"name\":\"SablierV2Lockup_EndTimeNotInTheFuture\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_InvalidHookSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_NotTransferable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_Null\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawableAmount\",\"type\":\"uint128\"}],\"name\":\"SablierV2Lockup_Overdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SablierV2Lockup_StartTimeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamCanceled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamDepleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamNotCancelable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamNotDepleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamSettled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_WithdrawAmountZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamIdsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountsCount\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_WithdrawArrayCountsNotEqual\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_WithdrawToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_WithdrawalAddressNotRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"AllowToHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"senderAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"recipientAmount\",\"type\":\"uint128\"}],\"name\":\"CancelLockupStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deposit\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"brokerFee\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structLockup.CreateAmounts\",\"name\":\"amounts\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"cancelable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"cliff\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"indexed\":false,\"internalType\":\"structLockupLinear.Timestamps\",\"name\":\"timestamps\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"}],\"name\":\"CreateLockupLinearStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"RenounceLockupStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"oldNFTDescriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"newNFTDescriptor\",\"type\":\"address\"}],\"name\":\"SetNFTDescriptor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"WithdrawFromLockupStream\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BROKER_FEE\",\"outputs\":[{\"internalType\":\"UD60x18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"allowToHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"streamIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"totalAmount\",\"type\":\"uint128\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"cancelable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint40\",\"name\":\"cliff\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"total\",\"type\":\"uint40\"}],\"internalType\":\"structLockupLinear.Durations\",\"name\":\"durations\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UD60x18\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structBroker\",\"name\":\"broker\",\"type\":\"tuple\"}],\"internalType\":\"structLockupLinear.CreateWithDurations\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createWithDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"totalAmount\",\"type\":\"uint128\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"cancelable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"cliff\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"internalType\":\"structLockupLinear.Timestamps\",\"name\":\"timestamps\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UD60x18\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structBroker\",\"name\":\"broker\",\"type\":\"tuple\"}],\"internalType\":\"structLockupLinear.CreateWithTimestamps\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createWithTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getAsset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getCliffTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"cliffTime\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getDepositedAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"depositedAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getEndTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getRefundedAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"refundedAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getStartTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getStream\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isCancelable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"wasCanceled\",\"type\":\"bool\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isDepleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isStream\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransferable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deposited\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"refunded\",\"type\":\"uint128\"}],\"internalType\":\"structLockup.Amounts\",\"name\":\"amounts\",\"type\":\"tuple\"},{\"internalType\":\"uint40\",\"name\":\"cliffTime\",\"type\":\"uint40\"}],\"internalType\":\"structLockupLinear.StreamLL\",\"name\":\"stream\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getTimestamps\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"cliff\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"internalType\":\"structLockupLinear.Timestamps\",\"name\":\"timestamps\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawnAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"isAllowedToHook\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isCancelable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isCold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isDepleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isStream\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isWarm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextStreamId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftDescriptor\",\"outputs\":[{\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"refundableAmountOf\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"refundableAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"renounce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"newNFTDescriptor\",\"type\":\"address\"}],\"name\":\"setNFTDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"enumLockup.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"streamedAmountOf\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"streamedAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"wasCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawMax\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"withdrawMaxAndTransfer\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"streamIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"}],\"name\":\"withdrawMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"withdrawableAmountOf\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawableAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SablierV2LockupLinearABI is the input ABI used to generate the binding from.
// Deprecated: Use SablierV2LockupLinearMetaData.ABI instead.
var SablierV2LockupLinearABI = SablierV2LockupLinearMetaData.ABI

// SablierV2LockupLinear is an auto generated Go binding around an Ethereum contract.
type SablierV2LockupLinear struct {
	SablierV2LockupLinearCaller     // Read-only binding to the contract
	SablierV2LockupLinearTransactor // Write-only binding to the contract
	SablierV2LockupLinearFilterer   // Log filterer for contract events
}

// SablierV2LockupLinearCaller is an auto generated read-only Go binding around an Ethereum contract.
type SablierV2LockupLinearCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierV2LockupLinearTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SablierV2LockupLinearTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierV2LockupLinearFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SablierV2LockupLinearFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierV2LockupLinearSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SablierV2LockupLinearSession struct {
	Contract     *SablierV2LockupLinear // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SablierV2LockupLinearCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SablierV2LockupLinearCallerSession struct {
	Contract *SablierV2LockupLinearCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SablierV2LockupLinearTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SablierV2LockupLinearTransactorSession struct {
	Contract     *SablierV2LockupLinearTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SablierV2LockupLinearRaw is an auto generated low-level Go binding around an Ethereum contract.
type SablierV2LockupLinearRaw struct {
	Contract *SablierV2LockupLinear // Generic contract binding to access the raw methods on
}

// SablierV2LockupLinearCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SablierV2LockupLinearCallerRaw struct {
	Contract *SablierV2LockupLinearCaller // Generic read-only contract binding to access the raw methods on
}

// SablierV2LockupLinearTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SablierV2LockupLinearTransactorRaw struct {
	Contract *SablierV2LockupLinearTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSablierV2LockupLinear creates a new instance of SablierV2LockupLinear, bound to a specific deployed contract.
func NewSablierV2LockupLinear(address common.Address, backend bind.ContractBackend) (*SablierV2LockupLinear, error) {
	contract, err := bindSablierV2LockupLinear(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinear{SablierV2LockupLinearCaller: SablierV2LockupLinearCaller{contract: contract}, SablierV2LockupLinearTransactor: SablierV2LockupLinearTransactor{contract: contract}, SablierV2LockupLinearFilterer: SablierV2LockupLinearFilterer{contract: contract}}, nil
}

// NewSablierV2LockupLinearCaller creates a new read-only instance of SablierV2LockupLinear, bound to a specific deployed contract.
func NewSablierV2LockupLinearCaller(address common.Address, caller bind.ContractCaller) (*SablierV2LockupLinearCaller, error) {
	contract, err := bindSablierV2LockupLinear(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearCaller{contract: contract}, nil
}

// NewSablierV2LockupLinearTransactor creates a new write-only instance of SablierV2LockupLinear, bound to a specific deployed contract.
func NewSablierV2LockupLinearTransactor(address common.Address, transactor bind.ContractTransactor) (*SablierV2LockupLinearTransactor, error) {
	contract, err := bindSablierV2LockupLinear(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearTransactor{contract: contract}, nil
}

// NewSablierV2LockupLinearFilterer creates a new log filterer instance of SablierV2LockupLinear, bound to a specific deployed contract.
func NewSablierV2LockupLinearFilterer(address common.Address, filterer bind.ContractFilterer) (*SablierV2LockupLinearFilterer, error) {
	contract, err := bindSablierV2LockupLinear(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearFilterer{contract: contract}, nil
}

// bindSablierV2LockupLinear binds a generic wrapper to an already deployed contract.
func bindSablierV2LockupLinear(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SablierV2LockupLinearMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SablierV2LockupLinear *SablierV2LockupLinearRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SablierV2LockupLinear.Contract.SablierV2LockupLinearCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SablierV2LockupLinear *SablierV2LockupLinearRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SablierV2LockupLinearTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SablierV2LockupLinear *SablierV2LockupLinearRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SablierV2LockupLinearTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SablierV2LockupLinear.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.contract.Transact(opts, method, params...)
}

// MAXBROKERFEE is a free data retrieval call binding the contract method 0x027b6744.
//
// Solidity: function MAX_BROKER_FEE() view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) MAXBROKERFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "MAX_BROKER_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBROKERFEE is a free data retrieval call binding the contract method 0x027b6744.
//
// Solidity: function MAX_BROKER_FEE() view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) MAXBROKERFEE() (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.MAXBROKERFEE(&_SablierV2LockupLinear.CallOpts)
}

// MAXBROKERFEE is a free data retrieval call binding the contract method 0x027b6744.
//
// Solidity: function MAX_BROKER_FEE() view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) MAXBROKERFEE() (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.MAXBROKERFEE(&_SablierV2LockupLinear.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Admin() (common.Address, error) {
	return _SablierV2LockupLinear.Contract.Admin(&_SablierV2LockupLinear.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) Admin() (common.Address, error) {
	return _SablierV2LockupLinear.Contract.Admin(&_SablierV2LockupLinear.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.BalanceOf(&_SablierV2LockupLinear.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.BalanceOf(&_SablierV2LockupLinear.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetApproved(&_SablierV2LockupLinear.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetApproved(&_SablierV2LockupLinear.CallOpts, tokenId)
}

// GetAsset is a free data retrieval call binding the contract method 0xeac8f5b8.
//
// Solidity: function getAsset(uint256 streamId) view returns(address asset)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetAsset(opts *bind.CallOpts, streamId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getAsset", streamId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAsset is a free data retrieval call binding the contract method 0xeac8f5b8.
//
// Solidity: function getAsset(uint256 streamId) view returns(address asset)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetAsset(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetAsset(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetAsset is a free data retrieval call binding the contract method 0xeac8f5b8.
//
// Solidity: function getAsset(uint256 streamId) view returns(address asset)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetAsset(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetAsset(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetCliffTime is a free data retrieval call binding the contract method 0x780a82c8.
//
// Solidity: function getCliffTime(uint256 streamId) view returns(uint40 cliffTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetCliffTime(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getCliffTime", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCliffTime is a free data retrieval call binding the contract method 0x780a82c8.
//
// Solidity: function getCliffTime(uint256 streamId) view returns(uint40 cliffTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetCliffTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetCliffTime(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetCliffTime is a free data retrieval call binding the contract method 0x780a82c8.
//
// Solidity: function getCliffTime(uint256 streamId) view returns(uint40 cliffTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetCliffTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetCliffTime(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetDepositedAmount is a free data retrieval call binding the contract method 0xa80fc071.
//
// Solidity: function getDepositedAmount(uint256 streamId) view returns(uint128 depositedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetDepositedAmount(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getDepositedAmount", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositedAmount is a free data retrieval call binding the contract method 0xa80fc071.
//
// Solidity: function getDepositedAmount(uint256 streamId) view returns(uint128 depositedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetDepositedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetDepositedAmount(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetDepositedAmount is a free data retrieval call binding the contract method 0xa80fc071.
//
// Solidity: function getDepositedAmount(uint256 streamId) view returns(uint128 depositedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetDepositedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetDepositedAmount(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetEndTime is a free data retrieval call binding the contract method 0x9067b677.
//
// Solidity: function getEndTime(uint256 streamId) view returns(uint40 endTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetEndTime(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getEndTime", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEndTime is a free data retrieval call binding the contract method 0x9067b677.
//
// Solidity: function getEndTime(uint256 streamId) view returns(uint40 endTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetEndTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetEndTime(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetEndTime is a free data retrieval call binding the contract method 0x9067b677.
//
// Solidity: function getEndTime(uint256 streamId) view returns(uint40 endTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetEndTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetEndTime(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetRecipient is a free data retrieval call binding the contract method 0x6d0cee75.
//
// Solidity: function getRecipient(uint256 streamId) view returns(address recipient)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetRecipient(opts *bind.CallOpts, streamId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getRecipient", streamId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecipient is a free data retrieval call binding the contract method 0x6d0cee75.
//
// Solidity: function getRecipient(uint256 streamId) view returns(address recipient)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetRecipient(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetRecipient(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetRecipient is a free data retrieval call binding the contract method 0x6d0cee75.
//
// Solidity: function getRecipient(uint256 streamId) view returns(address recipient)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetRecipient(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetRecipient(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetRefundedAmount is a free data retrieval call binding the contract method 0xd4dbd20b.
//
// Solidity: function getRefundedAmount(uint256 streamId) view returns(uint128 refundedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetRefundedAmount(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getRefundedAmount", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRefundedAmount is a free data retrieval call binding the contract method 0xd4dbd20b.
//
// Solidity: function getRefundedAmount(uint256 streamId) view returns(uint128 refundedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetRefundedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetRefundedAmount(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetRefundedAmount is a free data retrieval call binding the contract method 0xd4dbd20b.
//
// Solidity: function getRefundedAmount(uint256 streamId) view returns(uint128 refundedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetRefundedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetRefundedAmount(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetSender is a free data retrieval call binding the contract method 0xb971302a.
//
// Solidity: function getSender(uint256 streamId) view returns(address sender)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetSender(opts *bind.CallOpts, streamId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getSender", streamId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSender is a free data retrieval call binding the contract method 0xb971302a.
//
// Solidity: function getSender(uint256 streamId) view returns(address sender)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetSender(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetSender(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetSender is a free data retrieval call binding the contract method 0xb971302a.
//
// Solidity: function getSender(uint256 streamId) view returns(address sender)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetSender(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.GetSender(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetStartTime is a free data retrieval call binding the contract method 0xbc2be1be.
//
// Solidity: function getStartTime(uint256 streamId) view returns(uint40 startTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetStartTime(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getStartTime", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartTime is a free data retrieval call binding the contract method 0xbc2be1be.
//
// Solidity: function getStartTime(uint256 streamId) view returns(uint40 startTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetStartTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetStartTime(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetStartTime is a free data retrieval call binding the contract method 0xbc2be1be.
//
// Solidity: function getStartTime(uint256 streamId) view returns(uint40 startTime)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetStartTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetStartTime(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns((address,address,uint40,bool,bool,address,uint40,bool,bool,bool,(uint128,uint128,uint128),uint40) stream)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetStream(opts *bind.CallOpts, streamId *big.Int) (LockupLinearStreamLL, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getStream", streamId)

	if err != nil {
		return *new(LockupLinearStreamLL), err
	}

	out0 := *abi.ConvertType(out[0], new(LockupLinearStreamLL)).(*LockupLinearStreamLL)

	return out0, err

}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns((address,address,uint40,bool,bool,address,uint40,bool,bool,bool,(uint128,uint128,uint128),uint40) stream)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetStream(streamId *big.Int) (LockupLinearStreamLL, error) {
	return _SablierV2LockupLinear.Contract.GetStream(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns((address,address,uint40,bool,bool,address,uint40,bool,bool,bool,(uint128,uint128,uint128),uint40) stream)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetStream(streamId *big.Int) (LockupLinearStreamLL, error) {
	return _SablierV2LockupLinear.Contract.GetStream(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetTimestamps is a free data retrieval call binding the contract method 0x57404b12.
//
// Solidity: function getTimestamps(uint256 streamId) view returns((uint40,uint40,uint40) timestamps)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetTimestamps(opts *bind.CallOpts, streamId *big.Int) (LockupLinearTimestamps, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getTimestamps", streamId)

	if err != nil {
		return *new(LockupLinearTimestamps), err
	}

	out0 := *abi.ConvertType(out[0], new(LockupLinearTimestamps)).(*LockupLinearTimestamps)

	return out0, err

}

// GetTimestamps is a free data retrieval call binding the contract method 0x57404b12.
//
// Solidity: function getTimestamps(uint256 streamId) view returns((uint40,uint40,uint40) timestamps)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetTimestamps(streamId *big.Int) (LockupLinearTimestamps, error) {
	return _SablierV2LockupLinear.Contract.GetTimestamps(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetTimestamps is a free data retrieval call binding the contract method 0x57404b12.
//
// Solidity: function getTimestamps(uint256 streamId) view returns((uint40,uint40,uint40) timestamps)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetTimestamps(streamId *big.Int) (LockupLinearTimestamps, error) {
	return _SablierV2LockupLinear.Contract.GetTimestamps(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetWithdrawnAmount is a free data retrieval call binding the contract method 0xd511609f.
//
// Solidity: function getWithdrawnAmount(uint256 streamId) view returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) GetWithdrawnAmount(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "getWithdrawnAmount", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawnAmount is a free data retrieval call binding the contract method 0xd511609f.
//
// Solidity: function getWithdrawnAmount(uint256 streamId) view returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) GetWithdrawnAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetWithdrawnAmount(&_SablierV2LockupLinear.CallOpts, streamId)
}

// GetWithdrawnAmount is a free data retrieval call binding the contract method 0xd511609f.
//
// Solidity: function getWithdrawnAmount(uint256 streamId) view returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) GetWithdrawnAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.GetWithdrawnAmount(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsAllowedToHook is a free data retrieval call binding the contract method 0x303acc85.
//
// Solidity: function isAllowedToHook(address recipient) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsAllowedToHook(opts *bind.CallOpts, recipient common.Address) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isAllowedToHook", recipient)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToHook is a free data retrieval call binding the contract method 0x303acc85.
//
// Solidity: function isAllowedToHook(address recipient) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsAllowedToHook(recipient common.Address) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsAllowedToHook(&_SablierV2LockupLinear.CallOpts, recipient)
}

// IsAllowedToHook is a free data retrieval call binding the contract method 0x303acc85.
//
// Solidity: function isAllowedToHook(address recipient) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsAllowedToHook(recipient common.Address) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsAllowedToHook(&_SablierV2LockupLinear.CallOpts, recipient)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsApprovedForAll(&_SablierV2LockupLinear.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsApprovedForAll(&_SablierV2LockupLinear.CallOpts, owner, operator)
}

// IsCancelable is a free data retrieval call binding the contract method 0x4857501f.
//
// Solidity: function isCancelable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsCancelable(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isCancelable", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCancelable is a free data retrieval call binding the contract method 0x4857501f.
//
// Solidity: function isCancelable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsCancelable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsCancelable(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsCancelable is a free data retrieval call binding the contract method 0x4857501f.
//
// Solidity: function isCancelable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsCancelable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsCancelable(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsCold is a free data retrieval call binding the contract method 0x8f69b993.
//
// Solidity: function isCold(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsCold(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isCold", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCold is a free data retrieval call binding the contract method 0x8f69b993.
//
// Solidity: function isCold(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsCold(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsCold(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsCold is a free data retrieval call binding the contract method 0x8f69b993.
//
// Solidity: function isCold(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsCold(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsCold(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsDepleted is a free data retrieval call binding the contract method 0x425d30dd.
//
// Solidity: function isDepleted(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsDepleted(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isDepleted", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepleted is a free data retrieval call binding the contract method 0x425d30dd.
//
// Solidity: function isDepleted(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsDepleted(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsDepleted(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsDepleted is a free data retrieval call binding the contract method 0x425d30dd.
//
// Solidity: function isDepleted(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsDepleted(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsDepleted(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsStream is a free data retrieval call binding the contract method 0xb8a3be66.
//
// Solidity: function isStream(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsStream(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isStream", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStream is a free data retrieval call binding the contract method 0xb8a3be66.
//
// Solidity: function isStream(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsStream(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsStream(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsStream is a free data retrieval call binding the contract method 0xb8a3be66.
//
// Solidity: function isStream(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsStream(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsStream(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsTransferable is a free data retrieval call binding the contract method 0xb2564569.
//
// Solidity: function isTransferable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsTransferable(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isTransferable", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferable is a free data retrieval call binding the contract method 0xb2564569.
//
// Solidity: function isTransferable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsTransferable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsTransferable(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsTransferable is a free data retrieval call binding the contract method 0xb2564569.
//
// Solidity: function isTransferable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsTransferable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsTransferable(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsWarm is a free data retrieval call binding the contract method 0x1c1cdd4c.
//
// Solidity: function isWarm(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) IsWarm(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "isWarm", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWarm is a free data retrieval call binding the contract method 0x1c1cdd4c.
//
// Solidity: function isWarm(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) IsWarm(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsWarm(&_SablierV2LockupLinear.CallOpts, streamId)
}

// IsWarm is a free data retrieval call binding the contract method 0x1c1cdd4c.
//
// Solidity: function isWarm(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) IsWarm(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.IsWarm(&_SablierV2LockupLinear.CallOpts, streamId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Name() (string, error) {
	return _SablierV2LockupLinear.Contract.Name(&_SablierV2LockupLinear.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) Name() (string, error) {
	return _SablierV2LockupLinear.Contract.Name(&_SablierV2LockupLinear.CallOpts)
}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) NextStreamId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "nextStreamId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) NextStreamId() (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.NextStreamId(&_SablierV2LockupLinear.CallOpts)
}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) NextStreamId() (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.NextStreamId(&_SablierV2LockupLinear.CallOpts)
}

// NftDescriptor is a free data retrieval call binding the contract method 0x44267570.
//
// Solidity: function nftDescriptor() view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) NftDescriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "nftDescriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftDescriptor is a free data retrieval call binding the contract method 0x44267570.
//
// Solidity: function nftDescriptor() view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) NftDescriptor() (common.Address, error) {
	return _SablierV2LockupLinear.Contract.NftDescriptor(&_SablierV2LockupLinear.CallOpts)
}

// NftDescriptor is a free data retrieval call binding the contract method 0x44267570.
//
// Solidity: function nftDescriptor() view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) NftDescriptor() (common.Address, error) {
	return _SablierV2LockupLinear.Contract.NftDescriptor(&_SablierV2LockupLinear.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.OwnerOf(&_SablierV2LockupLinear.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupLinear.Contract.OwnerOf(&_SablierV2LockupLinear.CallOpts, tokenId)
}

// RefundableAmountOf is a free data retrieval call binding the contract method 0x1400ecec.
//
// Solidity: function refundableAmountOf(uint256 streamId) view returns(uint128 refundableAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) RefundableAmountOf(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "refundableAmountOf", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundableAmountOf is a free data retrieval call binding the contract method 0x1400ecec.
//
// Solidity: function refundableAmountOf(uint256 streamId) view returns(uint128 refundableAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) RefundableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.RefundableAmountOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// RefundableAmountOf is a free data retrieval call binding the contract method 0x1400ecec.
//
// Solidity: function refundableAmountOf(uint256 streamId) view returns(uint128 refundableAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) RefundableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.RefundableAmountOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// StatusOf is a free data retrieval call binding the contract method 0xad35efd4.
//
// Solidity: function statusOf(uint256 streamId) view returns(uint8 status)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) StatusOf(opts *bind.CallOpts, streamId *big.Int) (uint8, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "statusOf", streamId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// StatusOf is a free data retrieval call binding the contract method 0xad35efd4.
//
// Solidity: function statusOf(uint256 streamId) view returns(uint8 status)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) StatusOf(streamId *big.Int) (uint8, error) {
	return _SablierV2LockupLinear.Contract.StatusOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// StatusOf is a free data retrieval call binding the contract method 0xad35efd4.
//
// Solidity: function statusOf(uint256 streamId) view returns(uint8 status)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) StatusOf(streamId *big.Int) (uint8, error) {
	return _SablierV2LockupLinear.Contract.StatusOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// StreamedAmountOf is a free data retrieval call binding the contract method 0x4869e12d.
//
// Solidity: function streamedAmountOf(uint256 streamId) view returns(uint128 streamedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) StreamedAmountOf(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "streamedAmountOf", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StreamedAmountOf is a free data retrieval call binding the contract method 0x4869e12d.
//
// Solidity: function streamedAmountOf(uint256 streamId) view returns(uint128 streamedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) StreamedAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.StreamedAmountOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// StreamedAmountOf is a free data retrieval call binding the contract method 0x4869e12d.
//
// Solidity: function streamedAmountOf(uint256 streamId) view returns(uint128 streamedAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) StreamedAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.StreamedAmountOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SablierV2LockupLinear.Contract.SupportsInterface(&_SablierV2LockupLinear.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SablierV2LockupLinear.Contract.SupportsInterface(&_SablierV2LockupLinear.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Symbol() (string, error) {
	return _SablierV2LockupLinear.Contract.Symbol(&_SablierV2LockupLinear.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) Symbol() (string, error) {
	return _SablierV2LockupLinear.Contract.Symbol(&_SablierV2LockupLinear.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 streamId) view returns(string uri)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) TokenURI(opts *bind.CallOpts, streamId *big.Int) (string, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "tokenURI", streamId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 streamId) view returns(string uri)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) TokenURI(streamId *big.Int) (string, error) {
	return _SablierV2LockupLinear.Contract.TokenURI(&_SablierV2LockupLinear.CallOpts, streamId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 streamId) view returns(string uri)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) TokenURI(streamId *big.Int) (string, error) {
	return _SablierV2LockupLinear.Contract.TokenURI(&_SablierV2LockupLinear.CallOpts, streamId)
}

// WasCanceled is a free data retrieval call binding the contract method 0xf590c176.
//
// Solidity: function wasCanceled(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) WasCanceled(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "wasCanceled", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasCanceled is a free data retrieval call binding the contract method 0xf590c176.
//
// Solidity: function wasCanceled(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) WasCanceled(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.WasCanceled(&_SablierV2LockupLinear.CallOpts, streamId)
}

// WasCanceled is a free data retrieval call binding the contract method 0xf590c176.
//
// Solidity: function wasCanceled(uint256 streamId) view returns(bool result)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) WasCanceled(streamId *big.Int) (bool, error) {
	return _SablierV2LockupLinear.Contract.WasCanceled(&_SablierV2LockupLinear.CallOpts, streamId)
}

// WithdrawableAmountOf is a free data retrieval call binding the contract method 0xd975dfed.
//
// Solidity: function withdrawableAmountOf(uint256 streamId) view returns(uint128 withdrawableAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCaller) WithdrawableAmountOf(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupLinear.contract.Call(opts, &out, "withdrawableAmountOf", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableAmountOf is a free data retrieval call binding the contract method 0xd975dfed.
//
// Solidity: function withdrawableAmountOf(uint256 streamId) view returns(uint128 withdrawableAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) WithdrawableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.WithdrawableAmountOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// WithdrawableAmountOf is a free data retrieval call binding the contract method 0xd975dfed.
//
// Solidity: function withdrawableAmountOf(uint256 streamId) view returns(uint128 withdrawableAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearCallerSession) WithdrawableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupLinear.Contract.WithdrawableAmountOf(&_SablierV2LockupLinear.CallOpts, streamId)
}

// AllowToHook is a paid mutator transaction binding the contract method 0x406887cb.
//
// Solidity: function allowToHook(address recipient) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) AllowToHook(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "allowToHook", recipient)
}

// AllowToHook is a paid mutator transaction binding the contract method 0x406887cb.
//
// Solidity: function allowToHook(address recipient) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) AllowToHook(recipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.AllowToHook(&_SablierV2LockupLinear.TransactOpts, recipient)
}

// AllowToHook is a paid mutator transaction binding the contract method 0x406887cb.
//
// Solidity: function allowToHook(address recipient) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) AllowToHook(recipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.AllowToHook(&_SablierV2LockupLinear.TransactOpts, recipient)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Approve(&_SablierV2LockupLinear.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Approve(&_SablierV2LockupLinear.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) Burn(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "burn", streamId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Burn(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Burn(&_SablierV2LockupLinear.TransactOpts, streamId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) Burn(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Burn(&_SablierV2LockupLinear.TransactOpts, streamId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) Cancel(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "cancel", streamId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Cancel(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Cancel(&_SablierV2LockupLinear.TransactOpts, streamId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) Cancel(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Cancel(&_SablierV2LockupLinear.TransactOpts, streamId)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x8659c270.
//
// Solidity: function cancelMultiple(uint256[] streamIds) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) CancelMultiple(opts *bind.TransactOpts, streamIds []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "cancelMultiple", streamIds)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x8659c270.
//
// Solidity: function cancelMultiple(uint256[] streamIds) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) CancelMultiple(streamIds []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.CancelMultiple(&_SablierV2LockupLinear.TransactOpts, streamIds)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x8659c270.
//
// Solidity: function cancelMultiple(uint256[] streamIds) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) CancelMultiple(streamIds []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.CancelMultiple(&_SablierV2LockupLinear.TransactOpts, streamIds)
}

// CreateWithDurations is a paid mutator transaction binding the contract method 0xab167ccc.
//
// Solidity: function createWithDurations((address,address,uint128,address,bool,bool,(uint40,uint40),(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) CreateWithDurations(opts *bind.TransactOpts, params LockupLinearCreateWithDurations) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "createWithDurations", params)
}

// CreateWithDurations is a paid mutator transaction binding the contract method 0xab167ccc.
//
// Solidity: function createWithDurations((address,address,uint128,address,bool,bool,(uint40,uint40),(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) CreateWithDurations(params LockupLinearCreateWithDurations) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.CreateWithDurations(&_SablierV2LockupLinear.TransactOpts, params)
}

// CreateWithDurations is a paid mutator transaction binding the contract method 0xab167ccc.
//
// Solidity: function createWithDurations((address,address,uint128,address,bool,bool,(uint40,uint40),(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) CreateWithDurations(params LockupLinearCreateWithDurations) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.CreateWithDurations(&_SablierV2LockupLinear.TransactOpts, params)
}

// CreateWithTimestamps is a paid mutator transaction binding the contract method 0x53b15727.
//
// Solidity: function createWithTimestamps((address,address,uint128,address,bool,bool,(uint40,uint40,uint40),(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) CreateWithTimestamps(opts *bind.TransactOpts, params LockupLinearCreateWithTimestamps) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "createWithTimestamps", params)
}

// CreateWithTimestamps is a paid mutator transaction binding the contract method 0x53b15727.
//
// Solidity: function createWithTimestamps((address,address,uint128,address,bool,bool,(uint40,uint40,uint40),(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) CreateWithTimestamps(params LockupLinearCreateWithTimestamps) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.CreateWithTimestamps(&_SablierV2LockupLinear.TransactOpts, params)
}

// CreateWithTimestamps is a paid mutator transaction binding the contract method 0x53b15727.
//
// Solidity: function createWithTimestamps((address,address,uint128,address,bool,bool,(uint40,uint40,uint40),(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) CreateWithTimestamps(params LockupLinearCreateWithTimestamps) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.CreateWithTimestamps(&_SablierV2LockupLinear.TransactOpts, params)
}

// Renounce is a paid mutator transaction binding the contract method 0x7de6b1db.
//
// Solidity: function renounce(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) Renounce(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "renounce", streamId)
}

// Renounce is a paid mutator transaction binding the contract method 0x7de6b1db.
//
// Solidity: function renounce(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Renounce(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Renounce(&_SablierV2LockupLinear.TransactOpts, streamId)
}

// Renounce is a paid mutator transaction binding the contract method 0x7de6b1db.
//
// Solidity: function renounce(uint256 streamId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) Renounce(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Renounce(&_SablierV2LockupLinear.TransactOpts, streamId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SafeTransferFrom(&_SablierV2LockupLinear.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SafeTransferFrom(&_SablierV2LockupLinear.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SafeTransferFrom0(&_SablierV2LockupLinear.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SafeTransferFrom0(&_SablierV2LockupLinear.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SetApprovalForAll(&_SablierV2LockupLinear.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SetApprovalForAll(&_SablierV2LockupLinear.TransactOpts, operator, approved)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address newNFTDescriptor) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) SetNFTDescriptor(opts *bind.TransactOpts, newNFTDescriptor common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "setNFTDescriptor", newNFTDescriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address newNFTDescriptor) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) SetNFTDescriptor(newNFTDescriptor common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SetNFTDescriptor(&_SablierV2LockupLinear.TransactOpts, newNFTDescriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address newNFTDescriptor) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) SetNFTDescriptor(newNFTDescriptor common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.SetNFTDescriptor(&_SablierV2LockupLinear.TransactOpts, newNFTDescriptor)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.TransferAdmin(&_SablierV2LockupLinear.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.TransferAdmin(&_SablierV2LockupLinear.TransactOpts, newAdmin)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.TransferFrom(&_SablierV2LockupLinear.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.TransferFrom(&_SablierV2LockupLinear.TransactOpts, from, to, tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfdd46d60.
//
// Solidity: function withdraw(uint256 streamId, address to, uint128 amount) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) Withdraw(opts *bind.TransactOpts, streamId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "withdraw", streamId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfdd46d60.
//
// Solidity: function withdraw(uint256 streamId, address to, uint128 amount) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) Withdraw(streamId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Withdraw(&_SablierV2LockupLinear.TransactOpts, streamId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfdd46d60.
//
// Solidity: function withdraw(uint256 streamId, address to, uint128 amount) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) Withdraw(streamId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.Withdraw(&_SablierV2LockupLinear.TransactOpts, streamId, to, amount)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xea5ead19.
//
// Solidity: function withdrawMax(uint256 streamId, address to) returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) WithdrawMax(opts *bind.TransactOpts, streamId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "withdrawMax", streamId, to)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xea5ead19.
//
// Solidity: function withdrawMax(uint256 streamId, address to) returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) WithdrawMax(streamId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.WithdrawMax(&_SablierV2LockupLinear.TransactOpts, streamId, to)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xea5ead19.
//
// Solidity: function withdrawMax(uint256 streamId, address to) returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) WithdrawMax(streamId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.WithdrawMax(&_SablierV2LockupLinear.TransactOpts, streamId, to)
}

// WithdrawMaxAndTransfer is a paid mutator transaction binding the contract method 0xc156a11d.
//
// Solidity: function withdrawMaxAndTransfer(uint256 streamId, address newRecipient) returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) WithdrawMaxAndTransfer(opts *bind.TransactOpts, streamId *big.Int, newRecipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "withdrawMaxAndTransfer", streamId, newRecipient)
}

// WithdrawMaxAndTransfer is a paid mutator transaction binding the contract method 0xc156a11d.
//
// Solidity: function withdrawMaxAndTransfer(uint256 streamId, address newRecipient) returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) WithdrawMaxAndTransfer(streamId *big.Int, newRecipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.WithdrawMaxAndTransfer(&_SablierV2LockupLinear.TransactOpts, streamId, newRecipient)
}

// WithdrawMaxAndTransfer is a paid mutator transaction binding the contract method 0xc156a11d.
//
// Solidity: function withdrawMaxAndTransfer(uint256 streamId, address newRecipient) returns(uint128 withdrawnAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) WithdrawMaxAndTransfer(streamId *big.Int, newRecipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.WithdrawMaxAndTransfer(&_SablierV2LockupLinear.TransactOpts, streamId, newRecipient)
}

// WithdrawMultiple is a paid mutator transaction binding the contract method 0x4cc55e11.
//
// Solidity: function withdrawMultiple(uint256[] streamIds, uint128[] amounts) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactor) WithdrawMultiple(opts *bind.TransactOpts, streamIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.contract.Transact(opts, "withdrawMultiple", streamIds, amounts)
}

// WithdrawMultiple is a paid mutator transaction binding the contract method 0x4cc55e11.
//
// Solidity: function withdrawMultiple(uint256[] streamIds, uint128[] amounts) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearSession) WithdrawMultiple(streamIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.WithdrawMultiple(&_SablierV2LockupLinear.TransactOpts, streamIds, amounts)
}

// WithdrawMultiple is a paid mutator transaction binding the contract method 0x4cc55e11.
//
// Solidity: function withdrawMultiple(uint256[] streamIds, uint128[] amounts) returns()
func (_SablierV2LockupLinear *SablierV2LockupLinearTransactorSession) WithdrawMultiple(streamIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupLinear.Contract.WithdrawMultiple(&_SablierV2LockupLinear.TransactOpts, streamIds, amounts)
}

// SablierV2LockupLinearAllowToHookIterator is returned from FilterAllowToHook and is used to iterate over the raw logs and unpacked data for AllowToHook events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearAllowToHookIterator struct {
	Event *SablierV2LockupLinearAllowToHook // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearAllowToHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearAllowToHook)
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
		it.Event = new(SablierV2LockupLinearAllowToHook)
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
func (it *SablierV2LockupLinearAllowToHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearAllowToHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearAllowToHook represents a AllowToHook event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearAllowToHook struct {
	Admin     common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowToHook is a free log retrieval operation binding the contract event 0xb4378d4e289cb3f40f4f75a99c9cafa76e3df1c4dc31309babc23dc91bd72801.
//
// Solidity: event AllowToHook(address indexed admin, address recipient)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterAllowToHook(opts *bind.FilterOpts, admin []common.Address) (*SablierV2LockupLinearAllowToHookIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "AllowToHook", adminRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearAllowToHookIterator{contract: _SablierV2LockupLinear.contract, event: "AllowToHook", logs: logs, sub: sub}, nil
}

// WatchAllowToHook is a free log subscription operation binding the contract event 0xb4378d4e289cb3f40f4f75a99c9cafa76e3df1c4dc31309babc23dc91bd72801.
//
// Solidity: event AllowToHook(address indexed admin, address recipient)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchAllowToHook(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearAllowToHook, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "AllowToHook", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearAllowToHook)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "AllowToHook", log); err != nil {
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

// ParseAllowToHook is a log parse operation binding the contract event 0xb4378d4e289cb3f40f4f75a99c9cafa76e3df1c4dc31309babc23dc91bd72801.
//
// Solidity: event AllowToHook(address indexed admin, address recipient)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseAllowToHook(log types.Log) (*SablierV2LockupLinearAllowToHook, error) {
	event := new(SablierV2LockupLinearAllowToHook)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "AllowToHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearApprovalIterator struct {
	Event *SablierV2LockupLinearApproval // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearApproval)
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
		it.Event = new(SablierV2LockupLinearApproval)
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
func (it *SablierV2LockupLinearApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearApproval represents a Approval event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SablierV2LockupLinearApprovalIterator, error) {

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

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearApprovalIterator{contract: _SablierV2LockupLinear.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearApproval)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseApproval(log types.Log) (*SablierV2LockupLinearApproval, error) {
	event := new(SablierV2LockupLinearApproval)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearApprovalForAllIterator struct {
	Event *SablierV2LockupLinearApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearApprovalForAll)
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
		it.Event = new(SablierV2LockupLinearApprovalForAll)
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
func (it *SablierV2LockupLinearApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearApprovalForAll represents a ApprovalForAll event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SablierV2LockupLinearApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearApprovalForAllIterator{contract: _SablierV2LockupLinear.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearApprovalForAll)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseApprovalForAll(log types.Log) (*SablierV2LockupLinearApprovalForAll, error) {
	event := new(SablierV2LockupLinearApprovalForAll)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearBatchMetadataUpdateIterator struct {
	Event *SablierV2LockupLinearBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearBatchMetadataUpdate)
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
		it.Event = new(SablierV2LockupLinearBatchMetadataUpdate)
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
func (it *SablierV2LockupLinearBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*SablierV2LockupLinearBatchMetadataUpdateIterator, error) {

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearBatchMetadataUpdateIterator{contract: _SablierV2LockupLinear.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearBatchMetadataUpdate)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseBatchMetadataUpdate(log types.Log) (*SablierV2LockupLinearBatchMetadataUpdate, error) {
	event := new(SablierV2LockupLinearBatchMetadataUpdate)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearCancelLockupStreamIterator is returned from FilterCancelLockupStream and is used to iterate over the raw logs and unpacked data for CancelLockupStream events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearCancelLockupStreamIterator struct {
	Event *SablierV2LockupLinearCancelLockupStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearCancelLockupStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearCancelLockupStream)
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
		it.Event = new(SablierV2LockupLinearCancelLockupStream)
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
func (it *SablierV2LockupLinearCancelLockupStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearCancelLockupStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearCancelLockupStream represents a CancelLockupStream event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearCancelLockupStream struct {
	StreamId        *big.Int
	Sender          common.Address
	Recipient       common.Address
	Asset           common.Address
	SenderAmount    *big.Int
	RecipientAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelLockupStream is a free log retrieval operation binding the contract event 0x5edb27d6c1a327513b90a792050debf074b7194444885e3144d4decc5caaaa50.
//
// Solidity: event CancelLockupStream(uint256 streamId, address indexed sender, address indexed recipient, address indexed asset, uint128 senderAmount, uint128 recipientAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterCancelLockupStream(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, asset []common.Address) (*SablierV2LockupLinearCancelLockupStreamIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "CancelLockupStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearCancelLockupStreamIterator{contract: _SablierV2LockupLinear.contract, event: "CancelLockupStream", logs: logs, sub: sub}, nil
}

// WatchCancelLockupStream is a free log subscription operation binding the contract event 0x5edb27d6c1a327513b90a792050debf074b7194444885e3144d4decc5caaaa50.
//
// Solidity: event CancelLockupStream(uint256 streamId, address indexed sender, address indexed recipient, address indexed asset, uint128 senderAmount, uint128 recipientAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchCancelLockupStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearCancelLockupStream, sender []common.Address, recipient []common.Address, asset []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "CancelLockupStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearCancelLockupStream)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "CancelLockupStream", log); err != nil {
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

// ParseCancelLockupStream is a log parse operation binding the contract event 0x5edb27d6c1a327513b90a792050debf074b7194444885e3144d4decc5caaaa50.
//
// Solidity: event CancelLockupStream(uint256 streamId, address indexed sender, address indexed recipient, address indexed asset, uint128 senderAmount, uint128 recipientAmount)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseCancelLockupStream(log types.Log) (*SablierV2LockupLinearCancelLockupStream, error) {
	event := new(SablierV2LockupLinearCancelLockupStream)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "CancelLockupStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearCreateLockupLinearStreamIterator is returned from FilterCreateLockupLinearStream and is used to iterate over the raw logs and unpacked data for CreateLockupLinearStream events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearCreateLockupLinearStreamIterator struct {
	Event *SablierV2LockupLinearCreateLockupLinearStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearCreateLockupLinearStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearCreateLockupLinearStream)
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
		it.Event = new(SablierV2LockupLinearCreateLockupLinearStream)
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
func (it *SablierV2LockupLinearCreateLockupLinearStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearCreateLockupLinearStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearCreateLockupLinearStream represents a CreateLockupLinearStream event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearCreateLockupLinearStream struct {
	StreamId     *big.Int
	Funder       common.Address
	Sender       common.Address
	Recipient    common.Address
	Amounts      LockupCreateAmounts
	Asset        common.Address
	Cancelable   bool
	Transferable bool
	Timestamps   LockupLinearTimestamps
	Broker       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateLockupLinearStream is a free log retrieval operation binding the contract event 0x44cb432df42caa86b7ec73644ab8aec922bc44c71c98fc330addc75b88adbc7c.
//
// Solidity: event CreateLockupLinearStream(uint256 streamId, address funder, address indexed sender, address indexed recipient, (uint128,uint128) amounts, address indexed asset, bool cancelable, bool transferable, (uint40,uint40,uint40) timestamps, address broker)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterCreateLockupLinearStream(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, asset []common.Address) (*SablierV2LockupLinearCreateLockupLinearStreamIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "CreateLockupLinearStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearCreateLockupLinearStreamIterator{contract: _SablierV2LockupLinear.contract, event: "CreateLockupLinearStream", logs: logs, sub: sub}, nil
}

// WatchCreateLockupLinearStream is a free log subscription operation binding the contract event 0x44cb432df42caa86b7ec73644ab8aec922bc44c71c98fc330addc75b88adbc7c.
//
// Solidity: event CreateLockupLinearStream(uint256 streamId, address funder, address indexed sender, address indexed recipient, (uint128,uint128) amounts, address indexed asset, bool cancelable, bool transferable, (uint40,uint40,uint40) timestamps, address broker)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchCreateLockupLinearStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearCreateLockupLinearStream, sender []common.Address, recipient []common.Address, asset []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "CreateLockupLinearStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearCreateLockupLinearStream)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "CreateLockupLinearStream", log); err != nil {
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

// ParseCreateLockupLinearStream is a log parse operation binding the contract event 0x44cb432df42caa86b7ec73644ab8aec922bc44c71c98fc330addc75b88adbc7c.
//
// Solidity: event CreateLockupLinearStream(uint256 streamId, address funder, address indexed sender, address indexed recipient, (uint128,uint128) amounts, address indexed asset, bool cancelable, bool transferable, (uint40,uint40,uint40) timestamps, address broker)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseCreateLockupLinearStream(log types.Log) (*SablierV2LockupLinearCreateLockupLinearStream, error) {
	event := new(SablierV2LockupLinearCreateLockupLinearStream)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "CreateLockupLinearStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearMetadataUpdateIterator struct {
	Event *SablierV2LockupLinearMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearMetadataUpdate)
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
		it.Event = new(SablierV2LockupLinearMetadataUpdate)
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
func (it *SablierV2LockupLinearMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearMetadataUpdate represents a MetadataUpdate event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*SablierV2LockupLinearMetadataUpdateIterator, error) {

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearMetadataUpdateIterator{contract: _SablierV2LockupLinear.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearMetadataUpdate)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseMetadataUpdate(log types.Log) (*SablierV2LockupLinearMetadataUpdate, error) {
	event := new(SablierV2LockupLinearMetadataUpdate)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearRenounceLockupStreamIterator is returned from FilterRenounceLockupStream and is used to iterate over the raw logs and unpacked data for RenounceLockupStream events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearRenounceLockupStreamIterator struct {
	Event *SablierV2LockupLinearRenounceLockupStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearRenounceLockupStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearRenounceLockupStream)
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
		it.Event = new(SablierV2LockupLinearRenounceLockupStream)
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
func (it *SablierV2LockupLinearRenounceLockupStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearRenounceLockupStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearRenounceLockupStream represents a RenounceLockupStream event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearRenounceLockupStream struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRenounceLockupStream is a free log retrieval operation binding the contract event 0x0eb069207093cd3e51cd1370d2d369770057fbe29947e577e5fb428c6c6fc78f.
//
// Solidity: event RenounceLockupStream(uint256 indexed streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterRenounceLockupStream(opts *bind.FilterOpts, streamId []*big.Int) (*SablierV2LockupLinearRenounceLockupStreamIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "RenounceLockupStream", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearRenounceLockupStreamIterator{contract: _SablierV2LockupLinear.contract, event: "RenounceLockupStream", logs: logs, sub: sub}, nil
}

// WatchRenounceLockupStream is a free log subscription operation binding the contract event 0x0eb069207093cd3e51cd1370d2d369770057fbe29947e577e5fb428c6c6fc78f.
//
// Solidity: event RenounceLockupStream(uint256 indexed streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchRenounceLockupStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearRenounceLockupStream, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "RenounceLockupStream", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearRenounceLockupStream)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "RenounceLockupStream", log); err != nil {
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

// ParseRenounceLockupStream is a log parse operation binding the contract event 0x0eb069207093cd3e51cd1370d2d369770057fbe29947e577e5fb428c6c6fc78f.
//
// Solidity: event RenounceLockupStream(uint256 indexed streamId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseRenounceLockupStream(log types.Log) (*SablierV2LockupLinearRenounceLockupStream, error) {
	event := new(SablierV2LockupLinearRenounceLockupStream)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "RenounceLockupStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearSetNFTDescriptorIterator is returned from FilterSetNFTDescriptor and is used to iterate over the raw logs and unpacked data for SetNFTDescriptor events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearSetNFTDescriptorIterator struct {
	Event *SablierV2LockupLinearSetNFTDescriptor // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearSetNFTDescriptorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearSetNFTDescriptor)
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
		it.Event = new(SablierV2LockupLinearSetNFTDescriptor)
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
func (it *SablierV2LockupLinearSetNFTDescriptorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearSetNFTDescriptorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearSetNFTDescriptor represents a SetNFTDescriptor event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearSetNFTDescriptor struct {
	Admin            common.Address
	OldNFTDescriptor common.Address
	NewNFTDescriptor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetNFTDescriptor is a free log retrieval operation binding the contract event 0xa2548bd4b805e907c1558a47b5858324fe8bb4a2e1ddfca647eecbf65610eebc.
//
// Solidity: event SetNFTDescriptor(address indexed admin, address oldNFTDescriptor, address newNFTDescriptor)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterSetNFTDescriptor(opts *bind.FilterOpts, admin []common.Address) (*SablierV2LockupLinearSetNFTDescriptorIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "SetNFTDescriptor", adminRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearSetNFTDescriptorIterator{contract: _SablierV2LockupLinear.contract, event: "SetNFTDescriptor", logs: logs, sub: sub}, nil
}

// WatchSetNFTDescriptor is a free log subscription operation binding the contract event 0xa2548bd4b805e907c1558a47b5858324fe8bb4a2e1ddfca647eecbf65610eebc.
//
// Solidity: event SetNFTDescriptor(address indexed admin, address oldNFTDescriptor, address newNFTDescriptor)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchSetNFTDescriptor(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearSetNFTDescriptor, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "SetNFTDescriptor", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearSetNFTDescriptor)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
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

// ParseSetNFTDescriptor is a log parse operation binding the contract event 0xa2548bd4b805e907c1558a47b5858324fe8bb4a2e1ddfca647eecbf65610eebc.
//
// Solidity: event SetNFTDescriptor(address indexed admin, address oldNFTDescriptor, address newNFTDescriptor)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseSetNFTDescriptor(log types.Log) (*SablierV2LockupLinearSetNFTDescriptor, error) {
	event := new(SablierV2LockupLinearSetNFTDescriptor)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearTransferIterator struct {
	Event *SablierV2LockupLinearTransfer // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearTransfer)
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
		it.Event = new(SablierV2LockupLinearTransfer)
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
func (it *SablierV2LockupLinearTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearTransfer represents a Transfer event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SablierV2LockupLinearTransferIterator, error) {

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

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearTransferIterator{contract: _SablierV2LockupLinear.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearTransfer)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseTransfer(log types.Log) (*SablierV2LockupLinearTransfer, error) {
	event := new(SablierV2LockupLinearTransfer)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearTransferAdminIterator is returned from FilterTransferAdmin and is used to iterate over the raw logs and unpacked data for TransferAdmin events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearTransferAdminIterator struct {
	Event *SablierV2LockupLinearTransferAdmin // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearTransferAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearTransferAdmin)
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
		it.Event = new(SablierV2LockupLinearTransferAdmin)
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
func (it *SablierV2LockupLinearTransferAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearTransferAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearTransferAdmin represents a TransferAdmin event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearTransferAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferAdmin is a free log retrieval operation binding the contract event 0xbdd36143ee09de60bdefca70680e0f71189b2ed7acee364b53917ad433fdaf80.
//
// Solidity: event TransferAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterTransferAdmin(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*SablierV2LockupLinearTransferAdminIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "TransferAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearTransferAdminIterator{contract: _SablierV2LockupLinear.contract, event: "TransferAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferAdmin is a free log subscription operation binding the contract event 0xbdd36143ee09de60bdefca70680e0f71189b2ed7acee364b53917ad433fdaf80.
//
// Solidity: event TransferAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchTransferAdmin(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearTransferAdmin, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "TransferAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearTransferAdmin)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "TransferAdmin", log); err != nil {
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

// ParseTransferAdmin is a log parse operation binding the contract event 0xbdd36143ee09de60bdefca70680e0f71189b2ed7acee364b53917ad433fdaf80.
//
// Solidity: event TransferAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseTransferAdmin(log types.Log) (*SablierV2LockupLinearTransferAdmin, error) {
	event := new(SablierV2LockupLinearTransferAdmin)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "TransferAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupLinearWithdrawFromLockupStreamIterator is returned from FilterWithdrawFromLockupStream and is used to iterate over the raw logs and unpacked data for WithdrawFromLockupStream events raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearWithdrawFromLockupStreamIterator struct {
	Event *SablierV2LockupLinearWithdrawFromLockupStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupLinearWithdrawFromLockupStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupLinearWithdrawFromLockupStream)
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
		it.Event = new(SablierV2LockupLinearWithdrawFromLockupStream)
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
func (it *SablierV2LockupLinearWithdrawFromLockupStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupLinearWithdrawFromLockupStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupLinearWithdrawFromLockupStream represents a WithdrawFromLockupStream event raised by the SablierV2LockupLinear contract.
type SablierV2LockupLinearWithdrawFromLockupStream struct {
	StreamId *big.Int
	To       common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFromLockupStream is a free log retrieval operation binding the contract event 0x40b88e5c41c5a97ffb7b6ef88a0a2d505aa0c634cf8a0275cb236ea7dd87ed4d.
//
// Solidity: event WithdrawFromLockupStream(uint256 indexed streamId, address indexed to, address indexed asset, uint128 amount)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) FilterWithdrawFromLockupStream(opts *bind.FilterOpts, streamId []*big.Int, to []common.Address, asset []common.Address) (*SablierV2LockupLinearWithdrawFromLockupStreamIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.FilterLogs(opts, "WithdrawFromLockupStream", streamIdRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupLinearWithdrawFromLockupStreamIterator{contract: _SablierV2LockupLinear.contract, event: "WithdrawFromLockupStream", logs: logs, sub: sub}, nil
}

// WatchWithdrawFromLockupStream is a free log subscription operation binding the contract event 0x40b88e5c41c5a97ffb7b6ef88a0a2d505aa0c634cf8a0275cb236ea7dd87ed4d.
//
// Solidity: event WithdrawFromLockupStream(uint256 indexed streamId, address indexed to, address indexed asset, uint128 amount)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) WatchWithdrawFromLockupStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupLinearWithdrawFromLockupStream, streamId []*big.Int, to []common.Address, asset []common.Address) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _SablierV2LockupLinear.contract.WatchLogs(opts, "WithdrawFromLockupStream", streamIdRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupLinearWithdrawFromLockupStream)
				if err := _SablierV2LockupLinear.contract.UnpackLog(event, "WithdrawFromLockupStream", log); err != nil {
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

// ParseWithdrawFromLockupStream is a log parse operation binding the contract event 0x40b88e5c41c5a97ffb7b6ef88a0a2d505aa0c634cf8a0275cb236ea7dd87ed4d.
//
// Solidity: event WithdrawFromLockupStream(uint256 indexed streamId, address indexed to, address indexed asset, uint128 amount)
func (_SablierV2LockupLinear *SablierV2LockupLinearFilterer) ParseWithdrawFromLockupStream(log types.Log) (*SablierV2LockupLinearWithdrawFromLockupStream, error) {
	event := new(SablierV2LockupLinearWithdrawFromLockupStream)
	if err := _SablierV2LockupLinear.contract.UnpackLog(event, "WithdrawFromLockupStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
