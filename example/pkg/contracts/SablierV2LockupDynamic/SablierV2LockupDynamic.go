// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SablierV2LockupDynamic

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

// LockupDynamicCreateWithDurations is an auto generated low-level Go binding around an user-defined struct.
type LockupDynamicCreateWithDurations struct {
	Sender       common.Address
	Recipient    common.Address
	TotalAmount  *big.Int
	Asset        common.Address
	Cancelable   bool
	Transferable bool
	Segments     []LockupDynamicSegmentWithDuration
	Broker       Broker
}

// LockupDynamicCreateWithTimestamps is an auto generated low-level Go binding around an user-defined struct.
type LockupDynamicCreateWithTimestamps struct {
	Sender       common.Address
	Recipient    common.Address
	TotalAmount  *big.Int
	Asset        common.Address
	Cancelable   bool
	Transferable bool
	StartTime    *big.Int
	Segments     []LockupDynamicSegment
	Broker       Broker
}

// LockupDynamicSegment is an auto generated low-level Go binding around an user-defined struct.
type LockupDynamicSegment struct {
	Amount    *big.Int
	Exponent  uint64
	Timestamp *big.Int
}

// LockupDynamicSegmentWithDuration is an auto generated low-level Go binding around an user-defined struct.
type LockupDynamicSegmentWithDuration struct {
	Amount   *big.Int
	Exponent uint64
	Duration *big.Int
}

// LockupDynamicStreamLD is an auto generated low-level Go binding around an user-defined struct.
type LockupDynamicStreamLD struct {
	Sender         common.Address
	Recipient      common.Address
	StartTime      *big.Int
	EndTime        *big.Int
	IsCancelable   bool
	WasCanceled    bool
	Asset          common.Address
	IsDepleted     bool
	IsStream       bool
	IsTransferable bool
	Amounts        LockupAmounts
	Segments       []LockupDynamicSegment
}

// LockupDynamicTimestamps is an auto generated low-level Go binding around an user-defined struct.
type LockupDynamicTimestamps struct {
	Start *big.Int
	End   *big.Int
}

// SablierV2LockupDynamicMetaData contains all meta data concerning the SablierV2LockupDynamic contract.
var SablierV2LockupDynamicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialAdmin\",\"type\":\"address\"},{\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"initialNFTDescriptor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSegmentCount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerNotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelegateCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"PRBMath_MulDiv18_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"PRBMath_MulDiv_Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PRBMath_SD59x18_Div_InputTooSmall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"SD59x18\",\"name\":\"x\",\"type\":\"int256\"},{\"internalType\":\"SD59x18\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"PRBMath_SD59x18_Div_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"SD59x18\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMath_SD59x18_Exp2_InputTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"SD59x18\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMath_SD59x18_IntoUint256_Underflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"SD59x18\",\"name\":\"x\",\"type\":\"int256\"}],\"name\":\"PRBMath_SD59x18_Log_InputTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PRBMath_SD59x18_Mul_InputTooSmall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"SD59x18\",\"name\":\"x\",\"type\":\"int256\"},{\"internalType\":\"SD59x18\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"PRBMath_SD59x18_Mul_Overflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"depositAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"segmentAmountsSum\",\"type\":\"uint128\"}],\"name\":\"SablierV2LockupDynamic_DepositAmountNotEqualToSegmentAmountsSum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"SablierV2LockupDynamic_SegmentCountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SablierV2LockupDynamic_SegmentCountZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"previousTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"currentTimestamp\",\"type\":\"uint40\"}],\"name\":\"SablierV2LockupDynamic_SegmentTimestampsNotOrdered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"firstSegmentTimestamp\",\"type\":\"uint40\"}],\"name\":\"SablierV2LockupDynamic_StartTimeNotLessThanFirstSegmentTimestamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_AllowToHookUnsupportedInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_AllowToHookZeroCodeSize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"UD60x18\",\"name\":\"brokerFee\",\"type\":\"uint256\"},{\"internalType\":\"UD60x18\",\"name\":\"maxBrokerFee\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_BrokerFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SablierV2Lockup_DepositAmountZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"blockTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"}],\"name\":\"SablierV2Lockup_EndTimeNotInTheFuture\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_InvalidHookSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_NotTransferable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_Null\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawableAmount\",\"type\":\"uint128\"}],\"name\":\"SablierV2Lockup_Overdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SablierV2Lockup_StartTimeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamCanceled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamDepleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamNotCancelable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamNotDepleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_StreamSettled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_WithdrawAmountZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamIdsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountsCount\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_WithdrawArrayCountsNotEqual\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"SablierV2Lockup_WithdrawToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SablierV2Lockup_WithdrawalAddressNotRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"AllowToHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"senderAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"recipientAmount\",\"type\":\"uint128\"}],\"name\":\"CancelLockupStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deposit\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"brokerFee\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structLockup.CreateAmounts\",\"name\":\"amounts\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"cancelable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"UD2x18\",\"name\":\"exponent\",\"type\":\"uint64\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"indexed\":false,\"internalType\":\"structLockupDynamic.Segment[]\",\"name\":\"segments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"indexed\":false,\"internalType\":\"structLockupDynamic.Timestamps\",\"name\":\"timestamps\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"broker\",\"type\":\"address\"}],\"name\":\"CreateLockupDynamicStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"RenounceLockupStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"oldNFTDescriptor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"newNFTDescriptor\",\"type\":\"address\"}],\"name\":\"SetNFTDescriptor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"WithdrawFromLockupStream\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BROKER_FEE\",\"outputs\":[{\"internalType\":\"UD60x18\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SEGMENT_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"allowToHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"streamIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"totalAmount\",\"type\":\"uint128\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"cancelable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"UD2x18\",\"name\":\"exponent\",\"type\":\"uint64\"},{\"internalType\":\"uint40\",\"name\":\"duration\",\"type\":\"uint40\"}],\"internalType\":\"structLockupDynamic.SegmentWithDuration[]\",\"name\":\"segments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UD60x18\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structBroker\",\"name\":\"broker\",\"type\":\"tuple\"}],\"internalType\":\"structLockupDynamic.CreateWithDurations\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createWithDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"totalAmount\",\"type\":\"uint128\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"cancelable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"UD2x18\",\"name\":\"exponent\",\"type\":\"uint64\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structLockupDynamic.Segment[]\",\"name\":\"segments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"UD60x18\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structBroker\",\"name\":\"broker\",\"type\":\"tuple\"}],\"internalType\":\"structLockupDynamic.CreateWithTimestamps\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createWithTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getAsset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getDepositedAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"depositedAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getEndTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getRefundedAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"refundedAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getSegments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"UD2x18\",\"name\":\"exponent\",\"type\":\"uint64\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structLockupDynamic.Segment[]\",\"name\":\"segments\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getStartTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getStream\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"endTime\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isCancelable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"wasCanceled\",\"type\":\"bool\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDepleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isStream\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTransferable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deposited\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"withdrawn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"refunded\",\"type\":\"uint128\"}],\"internalType\":\"structLockup.Amounts\",\"name\":\"amounts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"UD2x18\",\"name\":\"exponent\",\"type\":\"uint64\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structLockupDynamic.Segment[]\",\"name\":\"segments\",\"type\":\"tuple[]\"}],\"internalType\":\"structLockupDynamic.StreamLD\",\"name\":\"stream\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getTimestamps\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"internalType\":\"structLockupDynamic.Timestamps\",\"name\":\"timestamps\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getWithdrawnAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"isAllowedToHook\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isCancelable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isCold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isDepleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isStream\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isTransferable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isWarm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextStreamId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftDescriptor\",\"outputs\":[{\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"refundableAmountOf\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"refundableAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"renounce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISablierV2NFTDescriptor\",\"name\":\"newNFTDescriptor\",\"type\":\"address\"}],\"name\":\"setNFTDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"enumLockup.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"streamedAmountOf\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"streamedAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"wasCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawMax\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"withdrawMaxAndTransfer\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawnAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"streamIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"}],\"name\":\"withdrawMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"withdrawableAmountOf\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"withdrawableAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SablierV2LockupDynamicABI is the input ABI used to generate the binding from.
// Deprecated: Use SablierV2LockupDynamicMetaData.ABI instead.
var SablierV2LockupDynamicABI = SablierV2LockupDynamicMetaData.ABI

// SablierV2LockupDynamic is an auto generated Go binding around an Ethereum contract.
type SablierV2LockupDynamic struct {
	SablierV2LockupDynamicCaller     // Read-only binding to the contract
	SablierV2LockupDynamicTransactor // Write-only binding to the contract
	SablierV2LockupDynamicFilterer   // Log filterer for contract events
}

// SablierV2LockupDynamicCaller is an auto generated read-only Go binding around an Ethereum contract.
type SablierV2LockupDynamicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierV2LockupDynamicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SablierV2LockupDynamicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierV2LockupDynamicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SablierV2LockupDynamicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierV2LockupDynamicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SablierV2LockupDynamicSession struct {
	Contract     *SablierV2LockupDynamic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SablierV2LockupDynamicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SablierV2LockupDynamicCallerSession struct {
	Contract *SablierV2LockupDynamicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// SablierV2LockupDynamicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SablierV2LockupDynamicTransactorSession struct {
	Contract     *SablierV2LockupDynamicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SablierV2LockupDynamicRaw is an auto generated low-level Go binding around an Ethereum contract.
type SablierV2LockupDynamicRaw struct {
	Contract *SablierV2LockupDynamic // Generic contract binding to access the raw methods on
}

// SablierV2LockupDynamicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SablierV2LockupDynamicCallerRaw struct {
	Contract *SablierV2LockupDynamicCaller // Generic read-only contract binding to access the raw methods on
}

// SablierV2LockupDynamicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SablierV2LockupDynamicTransactorRaw struct {
	Contract *SablierV2LockupDynamicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSablierV2LockupDynamic creates a new instance of SablierV2LockupDynamic, bound to a specific deployed contract.
func NewSablierV2LockupDynamic(address common.Address, backend bind.ContractBackend) (*SablierV2LockupDynamic, error) {
	contract, err := bindSablierV2LockupDynamic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamic{SablierV2LockupDynamicCaller: SablierV2LockupDynamicCaller{contract: contract}, SablierV2LockupDynamicTransactor: SablierV2LockupDynamicTransactor{contract: contract}, SablierV2LockupDynamicFilterer: SablierV2LockupDynamicFilterer{contract: contract}}, nil
}

// NewSablierV2LockupDynamicCaller creates a new read-only instance of SablierV2LockupDynamic, bound to a specific deployed contract.
func NewSablierV2LockupDynamicCaller(address common.Address, caller bind.ContractCaller) (*SablierV2LockupDynamicCaller, error) {
	contract, err := bindSablierV2LockupDynamic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicCaller{contract: contract}, nil
}

// NewSablierV2LockupDynamicTransactor creates a new write-only instance of SablierV2LockupDynamic, bound to a specific deployed contract.
func NewSablierV2LockupDynamicTransactor(address common.Address, transactor bind.ContractTransactor) (*SablierV2LockupDynamicTransactor, error) {
	contract, err := bindSablierV2LockupDynamic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicTransactor{contract: contract}, nil
}

// NewSablierV2LockupDynamicFilterer creates a new log filterer instance of SablierV2LockupDynamic, bound to a specific deployed contract.
func NewSablierV2LockupDynamicFilterer(address common.Address, filterer bind.ContractFilterer) (*SablierV2LockupDynamicFilterer, error) {
	contract, err := bindSablierV2LockupDynamic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicFilterer{contract: contract}, nil
}

// bindSablierV2LockupDynamic binds a generic wrapper to an already deployed contract.
func bindSablierV2LockupDynamic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SablierV2LockupDynamicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SablierV2LockupDynamic *SablierV2LockupDynamicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SablierV2LockupDynamic.Contract.SablierV2LockupDynamicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SablierV2LockupDynamic *SablierV2LockupDynamicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SablierV2LockupDynamicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SablierV2LockupDynamic *SablierV2LockupDynamicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SablierV2LockupDynamicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SablierV2LockupDynamic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.contract.Transact(opts, method, params...)
}

// MAXBROKERFEE is a free data retrieval call binding the contract method 0x027b6744.
//
// Solidity: function MAX_BROKER_FEE() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) MAXBROKERFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "MAX_BROKER_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBROKERFEE is a free data retrieval call binding the contract method 0x027b6744.
//
// Solidity: function MAX_BROKER_FEE() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) MAXBROKERFEE() (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.MAXBROKERFEE(&_SablierV2LockupDynamic.CallOpts)
}

// MAXBROKERFEE is a free data retrieval call binding the contract method 0x027b6744.
//
// Solidity: function MAX_BROKER_FEE() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) MAXBROKERFEE() (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.MAXBROKERFEE(&_SablierV2LockupDynamic.CallOpts)
}

// MAXSEGMENTCOUNT is a free data retrieval call binding the contract method 0x9188ec84.
//
// Solidity: function MAX_SEGMENT_COUNT() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) MAXSEGMENTCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "MAX_SEGMENT_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSEGMENTCOUNT is a free data retrieval call binding the contract method 0x9188ec84.
//
// Solidity: function MAX_SEGMENT_COUNT() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) MAXSEGMENTCOUNT() (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.MAXSEGMENTCOUNT(&_SablierV2LockupDynamic.CallOpts)
}

// MAXSEGMENTCOUNT is a free data retrieval call binding the contract method 0x9188ec84.
//
// Solidity: function MAX_SEGMENT_COUNT() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) MAXSEGMENTCOUNT() (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.MAXSEGMENTCOUNT(&_SablierV2LockupDynamic.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Admin() (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.Admin(&_SablierV2LockupDynamic.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) Admin() (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.Admin(&_SablierV2LockupDynamic.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.BalanceOf(&_SablierV2LockupDynamic.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.BalanceOf(&_SablierV2LockupDynamic.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetApproved(&_SablierV2LockupDynamic.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetApproved(&_SablierV2LockupDynamic.CallOpts, tokenId)
}

// GetAsset is a free data retrieval call binding the contract method 0xeac8f5b8.
//
// Solidity: function getAsset(uint256 streamId) view returns(address asset)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetAsset(opts *bind.CallOpts, streamId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getAsset", streamId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAsset is a free data retrieval call binding the contract method 0xeac8f5b8.
//
// Solidity: function getAsset(uint256 streamId) view returns(address asset)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetAsset(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetAsset(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetAsset is a free data retrieval call binding the contract method 0xeac8f5b8.
//
// Solidity: function getAsset(uint256 streamId) view returns(address asset)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetAsset(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetAsset(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetDepositedAmount is a free data retrieval call binding the contract method 0xa80fc071.
//
// Solidity: function getDepositedAmount(uint256 streamId) view returns(uint128 depositedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetDepositedAmount(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getDepositedAmount", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositedAmount is a free data retrieval call binding the contract method 0xa80fc071.
//
// Solidity: function getDepositedAmount(uint256 streamId) view returns(uint128 depositedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetDepositedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetDepositedAmount(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetDepositedAmount is a free data retrieval call binding the contract method 0xa80fc071.
//
// Solidity: function getDepositedAmount(uint256 streamId) view returns(uint128 depositedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetDepositedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetDepositedAmount(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetEndTime is a free data retrieval call binding the contract method 0x9067b677.
//
// Solidity: function getEndTime(uint256 streamId) view returns(uint40 endTime)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetEndTime(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getEndTime", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEndTime is a free data retrieval call binding the contract method 0x9067b677.
//
// Solidity: function getEndTime(uint256 streamId) view returns(uint40 endTime)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetEndTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetEndTime(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetEndTime is a free data retrieval call binding the contract method 0x9067b677.
//
// Solidity: function getEndTime(uint256 streamId) view returns(uint40 endTime)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetEndTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetEndTime(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetRecipient is a free data retrieval call binding the contract method 0x6d0cee75.
//
// Solidity: function getRecipient(uint256 streamId) view returns(address recipient)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetRecipient(opts *bind.CallOpts, streamId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getRecipient", streamId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecipient is a free data retrieval call binding the contract method 0x6d0cee75.
//
// Solidity: function getRecipient(uint256 streamId) view returns(address recipient)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetRecipient(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetRecipient(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetRecipient is a free data retrieval call binding the contract method 0x6d0cee75.
//
// Solidity: function getRecipient(uint256 streamId) view returns(address recipient)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetRecipient(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetRecipient(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetRefundedAmount is a free data retrieval call binding the contract method 0xd4dbd20b.
//
// Solidity: function getRefundedAmount(uint256 streamId) view returns(uint128 refundedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetRefundedAmount(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getRefundedAmount", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRefundedAmount is a free data retrieval call binding the contract method 0xd4dbd20b.
//
// Solidity: function getRefundedAmount(uint256 streamId) view returns(uint128 refundedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetRefundedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetRefundedAmount(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetRefundedAmount is a free data retrieval call binding the contract method 0xd4dbd20b.
//
// Solidity: function getRefundedAmount(uint256 streamId) view returns(uint128 refundedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetRefundedAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetRefundedAmount(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetSegments is a free data retrieval call binding the contract method 0xb637b865.
//
// Solidity: function getSegments(uint256 streamId) view returns((uint128,uint64,uint40)[] segments)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetSegments(opts *bind.CallOpts, streamId *big.Int) ([]LockupDynamicSegment, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getSegments", streamId)

	if err != nil {
		return *new([]LockupDynamicSegment), err
	}

	out0 := *abi.ConvertType(out[0], new([]LockupDynamicSegment)).(*[]LockupDynamicSegment)

	return out0, err

}

// GetSegments is a free data retrieval call binding the contract method 0xb637b865.
//
// Solidity: function getSegments(uint256 streamId) view returns((uint128,uint64,uint40)[] segments)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetSegments(streamId *big.Int) ([]LockupDynamicSegment, error) {
	return _SablierV2LockupDynamic.Contract.GetSegments(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetSegments is a free data retrieval call binding the contract method 0xb637b865.
//
// Solidity: function getSegments(uint256 streamId) view returns((uint128,uint64,uint40)[] segments)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetSegments(streamId *big.Int) ([]LockupDynamicSegment, error) {
	return _SablierV2LockupDynamic.Contract.GetSegments(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetSender is a free data retrieval call binding the contract method 0xb971302a.
//
// Solidity: function getSender(uint256 streamId) view returns(address sender)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetSender(opts *bind.CallOpts, streamId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getSender", streamId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSender is a free data retrieval call binding the contract method 0xb971302a.
//
// Solidity: function getSender(uint256 streamId) view returns(address sender)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetSender(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetSender(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetSender is a free data retrieval call binding the contract method 0xb971302a.
//
// Solidity: function getSender(uint256 streamId) view returns(address sender)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetSender(streamId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.GetSender(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetStartTime is a free data retrieval call binding the contract method 0xbc2be1be.
//
// Solidity: function getStartTime(uint256 streamId) view returns(uint40 startTime)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetStartTime(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getStartTime", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartTime is a free data retrieval call binding the contract method 0xbc2be1be.
//
// Solidity: function getStartTime(uint256 streamId) view returns(uint40 startTime)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetStartTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetStartTime(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetStartTime is a free data retrieval call binding the contract method 0xbc2be1be.
//
// Solidity: function getStartTime(uint256 streamId) view returns(uint40 startTime)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetStartTime(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetStartTime(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns((address,address,uint40,uint40,bool,bool,address,bool,bool,bool,(uint128,uint128,uint128),(uint128,uint64,uint40)[]) stream)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetStream(opts *bind.CallOpts, streamId *big.Int) (LockupDynamicStreamLD, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getStream", streamId)

	if err != nil {
		return *new(LockupDynamicStreamLD), err
	}

	out0 := *abi.ConvertType(out[0], new(LockupDynamicStreamLD)).(*LockupDynamicStreamLD)

	return out0, err

}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns((address,address,uint40,uint40,bool,bool,address,bool,bool,bool,(uint128,uint128,uint128),(uint128,uint64,uint40)[]) stream)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetStream(streamId *big.Int) (LockupDynamicStreamLD, error) {
	return _SablierV2LockupDynamic.Contract.GetStream(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns((address,address,uint40,uint40,bool,bool,address,bool,bool,bool,(uint128,uint128,uint128),(uint128,uint64,uint40)[]) stream)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetStream(streamId *big.Int) (LockupDynamicStreamLD, error) {
	return _SablierV2LockupDynamic.Contract.GetStream(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetTimestamps is a free data retrieval call binding the contract method 0x57404b12.
//
// Solidity: function getTimestamps(uint256 streamId) view returns((uint40,uint40) timestamps)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetTimestamps(opts *bind.CallOpts, streamId *big.Int) (LockupDynamicTimestamps, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getTimestamps", streamId)

	if err != nil {
		return *new(LockupDynamicTimestamps), err
	}

	out0 := *abi.ConvertType(out[0], new(LockupDynamicTimestamps)).(*LockupDynamicTimestamps)

	return out0, err

}

// GetTimestamps is a free data retrieval call binding the contract method 0x57404b12.
//
// Solidity: function getTimestamps(uint256 streamId) view returns((uint40,uint40) timestamps)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetTimestamps(streamId *big.Int) (LockupDynamicTimestamps, error) {
	return _SablierV2LockupDynamic.Contract.GetTimestamps(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetTimestamps is a free data retrieval call binding the contract method 0x57404b12.
//
// Solidity: function getTimestamps(uint256 streamId) view returns((uint40,uint40) timestamps)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetTimestamps(streamId *big.Int) (LockupDynamicTimestamps, error) {
	return _SablierV2LockupDynamic.Contract.GetTimestamps(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetWithdrawnAmount is a free data retrieval call binding the contract method 0xd511609f.
//
// Solidity: function getWithdrawnAmount(uint256 streamId) view returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) GetWithdrawnAmount(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "getWithdrawnAmount", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawnAmount is a free data retrieval call binding the contract method 0xd511609f.
//
// Solidity: function getWithdrawnAmount(uint256 streamId) view returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) GetWithdrawnAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetWithdrawnAmount(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// GetWithdrawnAmount is a free data retrieval call binding the contract method 0xd511609f.
//
// Solidity: function getWithdrawnAmount(uint256 streamId) view returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) GetWithdrawnAmount(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.GetWithdrawnAmount(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsAllowedToHook is a free data retrieval call binding the contract method 0x303acc85.
//
// Solidity: function isAllowedToHook(address recipient) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsAllowedToHook(opts *bind.CallOpts, recipient common.Address) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isAllowedToHook", recipient)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToHook is a free data retrieval call binding the contract method 0x303acc85.
//
// Solidity: function isAllowedToHook(address recipient) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsAllowedToHook(recipient common.Address) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsAllowedToHook(&_SablierV2LockupDynamic.CallOpts, recipient)
}

// IsAllowedToHook is a free data retrieval call binding the contract method 0x303acc85.
//
// Solidity: function isAllowedToHook(address recipient) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsAllowedToHook(recipient common.Address) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsAllowedToHook(&_SablierV2LockupDynamic.CallOpts, recipient)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsApprovedForAll(&_SablierV2LockupDynamic.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsApprovedForAll(&_SablierV2LockupDynamic.CallOpts, owner, operator)
}

// IsCancelable is a free data retrieval call binding the contract method 0x4857501f.
//
// Solidity: function isCancelable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsCancelable(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isCancelable", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCancelable is a free data retrieval call binding the contract method 0x4857501f.
//
// Solidity: function isCancelable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsCancelable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsCancelable(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsCancelable is a free data retrieval call binding the contract method 0x4857501f.
//
// Solidity: function isCancelable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsCancelable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsCancelable(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsCold is a free data retrieval call binding the contract method 0x8f69b993.
//
// Solidity: function isCold(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsCold(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isCold", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCold is a free data retrieval call binding the contract method 0x8f69b993.
//
// Solidity: function isCold(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsCold(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsCold(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsCold is a free data retrieval call binding the contract method 0x8f69b993.
//
// Solidity: function isCold(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsCold(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsCold(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsDepleted is a free data retrieval call binding the contract method 0x425d30dd.
//
// Solidity: function isDepleted(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsDepleted(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isDepleted", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepleted is a free data retrieval call binding the contract method 0x425d30dd.
//
// Solidity: function isDepleted(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsDepleted(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsDepleted(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsDepleted is a free data retrieval call binding the contract method 0x425d30dd.
//
// Solidity: function isDepleted(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsDepleted(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsDepleted(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsStream is a free data retrieval call binding the contract method 0xb8a3be66.
//
// Solidity: function isStream(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsStream(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isStream", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStream is a free data retrieval call binding the contract method 0xb8a3be66.
//
// Solidity: function isStream(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsStream(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsStream(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsStream is a free data retrieval call binding the contract method 0xb8a3be66.
//
// Solidity: function isStream(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsStream(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsStream(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsTransferable is a free data retrieval call binding the contract method 0xb2564569.
//
// Solidity: function isTransferable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsTransferable(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isTransferable", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferable is a free data retrieval call binding the contract method 0xb2564569.
//
// Solidity: function isTransferable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsTransferable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsTransferable(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsTransferable is a free data retrieval call binding the contract method 0xb2564569.
//
// Solidity: function isTransferable(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsTransferable(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsTransferable(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsWarm is a free data retrieval call binding the contract method 0x1c1cdd4c.
//
// Solidity: function isWarm(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) IsWarm(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "isWarm", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWarm is a free data retrieval call binding the contract method 0x1c1cdd4c.
//
// Solidity: function isWarm(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) IsWarm(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsWarm(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// IsWarm is a free data retrieval call binding the contract method 0x1c1cdd4c.
//
// Solidity: function isWarm(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) IsWarm(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.IsWarm(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Name() (string, error) {
	return _SablierV2LockupDynamic.Contract.Name(&_SablierV2LockupDynamic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) Name() (string, error) {
	return _SablierV2LockupDynamic.Contract.Name(&_SablierV2LockupDynamic.CallOpts)
}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) NextStreamId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "nextStreamId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) NextStreamId() (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.NextStreamId(&_SablierV2LockupDynamic.CallOpts)
}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) NextStreamId() (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.NextStreamId(&_SablierV2LockupDynamic.CallOpts)
}

// NftDescriptor is a free data retrieval call binding the contract method 0x44267570.
//
// Solidity: function nftDescriptor() view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) NftDescriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "nftDescriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftDescriptor is a free data retrieval call binding the contract method 0x44267570.
//
// Solidity: function nftDescriptor() view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) NftDescriptor() (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.NftDescriptor(&_SablierV2LockupDynamic.CallOpts)
}

// NftDescriptor is a free data retrieval call binding the contract method 0x44267570.
//
// Solidity: function nftDescriptor() view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) NftDescriptor() (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.NftDescriptor(&_SablierV2LockupDynamic.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.OwnerOf(&_SablierV2LockupDynamic.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SablierV2LockupDynamic.Contract.OwnerOf(&_SablierV2LockupDynamic.CallOpts, tokenId)
}

// RefundableAmountOf is a free data retrieval call binding the contract method 0x1400ecec.
//
// Solidity: function refundableAmountOf(uint256 streamId) view returns(uint128 refundableAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) RefundableAmountOf(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "refundableAmountOf", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundableAmountOf is a free data retrieval call binding the contract method 0x1400ecec.
//
// Solidity: function refundableAmountOf(uint256 streamId) view returns(uint128 refundableAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) RefundableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.RefundableAmountOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// RefundableAmountOf is a free data retrieval call binding the contract method 0x1400ecec.
//
// Solidity: function refundableAmountOf(uint256 streamId) view returns(uint128 refundableAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) RefundableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.RefundableAmountOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// StatusOf is a free data retrieval call binding the contract method 0xad35efd4.
//
// Solidity: function statusOf(uint256 streamId) view returns(uint8 status)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) StatusOf(opts *bind.CallOpts, streamId *big.Int) (uint8, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "statusOf", streamId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// StatusOf is a free data retrieval call binding the contract method 0xad35efd4.
//
// Solidity: function statusOf(uint256 streamId) view returns(uint8 status)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) StatusOf(streamId *big.Int) (uint8, error) {
	return _SablierV2LockupDynamic.Contract.StatusOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// StatusOf is a free data retrieval call binding the contract method 0xad35efd4.
//
// Solidity: function statusOf(uint256 streamId) view returns(uint8 status)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) StatusOf(streamId *big.Int) (uint8, error) {
	return _SablierV2LockupDynamic.Contract.StatusOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// StreamedAmountOf is a free data retrieval call binding the contract method 0x4869e12d.
//
// Solidity: function streamedAmountOf(uint256 streamId) view returns(uint128 streamedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) StreamedAmountOf(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "streamedAmountOf", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StreamedAmountOf is a free data retrieval call binding the contract method 0x4869e12d.
//
// Solidity: function streamedAmountOf(uint256 streamId) view returns(uint128 streamedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) StreamedAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.StreamedAmountOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// StreamedAmountOf is a free data retrieval call binding the contract method 0x4869e12d.
//
// Solidity: function streamedAmountOf(uint256 streamId) view returns(uint128 streamedAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) StreamedAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.StreamedAmountOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SablierV2LockupDynamic.Contract.SupportsInterface(&_SablierV2LockupDynamic.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SablierV2LockupDynamic.Contract.SupportsInterface(&_SablierV2LockupDynamic.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Symbol() (string, error) {
	return _SablierV2LockupDynamic.Contract.Symbol(&_SablierV2LockupDynamic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) Symbol() (string, error) {
	return _SablierV2LockupDynamic.Contract.Symbol(&_SablierV2LockupDynamic.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 streamId) view returns(string uri)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) TokenURI(opts *bind.CallOpts, streamId *big.Int) (string, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "tokenURI", streamId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 streamId) view returns(string uri)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) TokenURI(streamId *big.Int) (string, error) {
	return _SablierV2LockupDynamic.Contract.TokenURI(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 streamId) view returns(string uri)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) TokenURI(streamId *big.Int) (string, error) {
	return _SablierV2LockupDynamic.Contract.TokenURI(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// WasCanceled is a free data retrieval call binding the contract method 0xf590c176.
//
// Solidity: function wasCanceled(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) WasCanceled(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "wasCanceled", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasCanceled is a free data retrieval call binding the contract method 0xf590c176.
//
// Solidity: function wasCanceled(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) WasCanceled(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.WasCanceled(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// WasCanceled is a free data retrieval call binding the contract method 0xf590c176.
//
// Solidity: function wasCanceled(uint256 streamId) view returns(bool result)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) WasCanceled(streamId *big.Int) (bool, error) {
	return _SablierV2LockupDynamic.Contract.WasCanceled(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// WithdrawableAmountOf is a free data retrieval call binding the contract method 0xd975dfed.
//
// Solidity: function withdrawableAmountOf(uint256 streamId) view returns(uint128 withdrawableAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCaller) WithdrawableAmountOf(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SablierV2LockupDynamic.contract.Call(opts, &out, "withdrawableAmountOf", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableAmountOf is a free data retrieval call binding the contract method 0xd975dfed.
//
// Solidity: function withdrawableAmountOf(uint256 streamId) view returns(uint128 withdrawableAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) WithdrawableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawableAmountOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// WithdrawableAmountOf is a free data retrieval call binding the contract method 0xd975dfed.
//
// Solidity: function withdrawableAmountOf(uint256 streamId) view returns(uint128 withdrawableAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicCallerSession) WithdrawableAmountOf(streamId *big.Int) (*big.Int, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawableAmountOf(&_SablierV2LockupDynamic.CallOpts, streamId)
}

// AllowToHook is a paid mutator transaction binding the contract method 0x406887cb.
//
// Solidity: function allowToHook(address recipient) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) AllowToHook(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "allowToHook", recipient)
}

// AllowToHook is a paid mutator transaction binding the contract method 0x406887cb.
//
// Solidity: function allowToHook(address recipient) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) AllowToHook(recipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.AllowToHook(&_SablierV2LockupDynamic.TransactOpts, recipient)
}

// AllowToHook is a paid mutator transaction binding the contract method 0x406887cb.
//
// Solidity: function allowToHook(address recipient) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) AllowToHook(recipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.AllowToHook(&_SablierV2LockupDynamic.TransactOpts, recipient)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Approve(&_SablierV2LockupDynamic.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Approve(&_SablierV2LockupDynamic.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) Burn(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "burn", streamId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Burn(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Burn(&_SablierV2LockupDynamic.TransactOpts, streamId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) Burn(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Burn(&_SablierV2LockupDynamic.TransactOpts, streamId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) Cancel(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "cancel", streamId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Cancel(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Cancel(&_SablierV2LockupDynamic.TransactOpts, streamId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) Cancel(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Cancel(&_SablierV2LockupDynamic.TransactOpts, streamId)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x8659c270.
//
// Solidity: function cancelMultiple(uint256[] streamIds) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) CancelMultiple(opts *bind.TransactOpts, streamIds []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "cancelMultiple", streamIds)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x8659c270.
//
// Solidity: function cancelMultiple(uint256[] streamIds) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) CancelMultiple(streamIds []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.CancelMultiple(&_SablierV2LockupDynamic.TransactOpts, streamIds)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x8659c270.
//
// Solidity: function cancelMultiple(uint256[] streamIds) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) CancelMultiple(streamIds []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.CancelMultiple(&_SablierV2LockupDynamic.TransactOpts, streamIds)
}

// CreateWithDurations is a paid mutator transaction binding the contract method 0x54c02292.
//
// Solidity: function createWithDurations((address,address,uint128,address,bool,bool,(uint128,uint64,uint40)[],(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) CreateWithDurations(opts *bind.TransactOpts, params LockupDynamicCreateWithDurations) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "createWithDurations", params)
}

// CreateWithDurations is a paid mutator transaction binding the contract method 0x54c02292.
//
// Solidity: function createWithDurations((address,address,uint128,address,bool,bool,(uint128,uint64,uint40)[],(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) CreateWithDurations(params LockupDynamicCreateWithDurations) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.CreateWithDurations(&_SablierV2LockupDynamic.TransactOpts, params)
}

// CreateWithDurations is a paid mutator transaction binding the contract method 0x54c02292.
//
// Solidity: function createWithDurations((address,address,uint128,address,bool,bool,(uint128,uint64,uint40)[],(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) CreateWithDurations(params LockupDynamicCreateWithDurations) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.CreateWithDurations(&_SablierV2LockupDynamic.TransactOpts, params)
}

// CreateWithTimestamps is a paid mutator transaction binding the contract method 0x31df3d48.
//
// Solidity: function createWithTimestamps((address,address,uint128,address,bool,bool,uint40,(uint128,uint64,uint40)[],(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) CreateWithTimestamps(opts *bind.TransactOpts, params LockupDynamicCreateWithTimestamps) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "createWithTimestamps", params)
}

// CreateWithTimestamps is a paid mutator transaction binding the contract method 0x31df3d48.
//
// Solidity: function createWithTimestamps((address,address,uint128,address,bool,bool,uint40,(uint128,uint64,uint40)[],(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) CreateWithTimestamps(params LockupDynamicCreateWithTimestamps) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.CreateWithTimestamps(&_SablierV2LockupDynamic.TransactOpts, params)
}

// CreateWithTimestamps is a paid mutator transaction binding the contract method 0x31df3d48.
//
// Solidity: function createWithTimestamps((address,address,uint128,address,bool,bool,uint40,(uint128,uint64,uint40)[],(address,uint256)) params) returns(uint256 streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) CreateWithTimestamps(params LockupDynamicCreateWithTimestamps) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.CreateWithTimestamps(&_SablierV2LockupDynamic.TransactOpts, params)
}

// Renounce is a paid mutator transaction binding the contract method 0x7de6b1db.
//
// Solidity: function renounce(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) Renounce(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "renounce", streamId)
}

// Renounce is a paid mutator transaction binding the contract method 0x7de6b1db.
//
// Solidity: function renounce(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Renounce(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Renounce(&_SablierV2LockupDynamic.TransactOpts, streamId)
}

// Renounce is a paid mutator transaction binding the contract method 0x7de6b1db.
//
// Solidity: function renounce(uint256 streamId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) Renounce(streamId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Renounce(&_SablierV2LockupDynamic.TransactOpts, streamId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SafeTransferFrom(&_SablierV2LockupDynamic.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SafeTransferFrom(&_SablierV2LockupDynamic.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SafeTransferFrom0(&_SablierV2LockupDynamic.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SafeTransferFrom0(&_SablierV2LockupDynamic.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SetApprovalForAll(&_SablierV2LockupDynamic.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SetApprovalForAll(&_SablierV2LockupDynamic.TransactOpts, operator, approved)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address newNFTDescriptor) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) SetNFTDescriptor(opts *bind.TransactOpts, newNFTDescriptor common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "setNFTDescriptor", newNFTDescriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address newNFTDescriptor) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) SetNFTDescriptor(newNFTDescriptor common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SetNFTDescriptor(&_SablierV2LockupDynamic.TransactOpts, newNFTDescriptor)
}

// SetNFTDescriptor is a paid mutator transaction binding the contract method 0x7cad6cd1.
//
// Solidity: function setNFTDescriptor(address newNFTDescriptor) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) SetNFTDescriptor(newNFTDescriptor common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.SetNFTDescriptor(&_SablierV2LockupDynamic.TransactOpts, newNFTDescriptor)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.TransferAdmin(&_SablierV2LockupDynamic.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address newAdmin) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.TransferAdmin(&_SablierV2LockupDynamic.TransactOpts, newAdmin)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.TransferFrom(&_SablierV2LockupDynamic.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.TransferFrom(&_SablierV2LockupDynamic.TransactOpts, from, to, tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfdd46d60.
//
// Solidity: function withdraw(uint256 streamId, address to, uint128 amount) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) Withdraw(opts *bind.TransactOpts, streamId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "withdraw", streamId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfdd46d60.
//
// Solidity: function withdraw(uint256 streamId, address to, uint128 amount) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) Withdraw(streamId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Withdraw(&_SablierV2LockupDynamic.TransactOpts, streamId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfdd46d60.
//
// Solidity: function withdraw(uint256 streamId, address to, uint128 amount) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) Withdraw(streamId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.Withdraw(&_SablierV2LockupDynamic.TransactOpts, streamId, to, amount)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xea5ead19.
//
// Solidity: function withdrawMax(uint256 streamId, address to) returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) WithdrawMax(opts *bind.TransactOpts, streamId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "withdrawMax", streamId, to)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xea5ead19.
//
// Solidity: function withdrawMax(uint256 streamId, address to) returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) WithdrawMax(streamId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawMax(&_SablierV2LockupDynamic.TransactOpts, streamId, to)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xea5ead19.
//
// Solidity: function withdrawMax(uint256 streamId, address to) returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) WithdrawMax(streamId *big.Int, to common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawMax(&_SablierV2LockupDynamic.TransactOpts, streamId, to)
}

// WithdrawMaxAndTransfer is a paid mutator transaction binding the contract method 0xc156a11d.
//
// Solidity: function withdrawMaxAndTransfer(uint256 streamId, address newRecipient) returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) WithdrawMaxAndTransfer(opts *bind.TransactOpts, streamId *big.Int, newRecipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "withdrawMaxAndTransfer", streamId, newRecipient)
}

// WithdrawMaxAndTransfer is a paid mutator transaction binding the contract method 0xc156a11d.
//
// Solidity: function withdrawMaxAndTransfer(uint256 streamId, address newRecipient) returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) WithdrawMaxAndTransfer(streamId *big.Int, newRecipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawMaxAndTransfer(&_SablierV2LockupDynamic.TransactOpts, streamId, newRecipient)
}

// WithdrawMaxAndTransfer is a paid mutator transaction binding the contract method 0xc156a11d.
//
// Solidity: function withdrawMaxAndTransfer(uint256 streamId, address newRecipient) returns(uint128 withdrawnAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) WithdrawMaxAndTransfer(streamId *big.Int, newRecipient common.Address) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawMaxAndTransfer(&_SablierV2LockupDynamic.TransactOpts, streamId, newRecipient)
}

// WithdrawMultiple is a paid mutator transaction binding the contract method 0x4cc55e11.
//
// Solidity: function withdrawMultiple(uint256[] streamIds, uint128[] amounts) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactor) WithdrawMultiple(opts *bind.TransactOpts, streamIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.contract.Transact(opts, "withdrawMultiple", streamIds, amounts)
}

// WithdrawMultiple is a paid mutator transaction binding the contract method 0x4cc55e11.
//
// Solidity: function withdrawMultiple(uint256[] streamIds, uint128[] amounts) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicSession) WithdrawMultiple(streamIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawMultiple(&_SablierV2LockupDynamic.TransactOpts, streamIds, amounts)
}

// WithdrawMultiple is a paid mutator transaction binding the contract method 0x4cc55e11.
//
// Solidity: function withdrawMultiple(uint256[] streamIds, uint128[] amounts) returns()
func (_SablierV2LockupDynamic *SablierV2LockupDynamicTransactorSession) WithdrawMultiple(streamIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _SablierV2LockupDynamic.Contract.WithdrawMultiple(&_SablierV2LockupDynamic.TransactOpts, streamIds, amounts)
}

// SablierV2LockupDynamicAllowToHookIterator is returned from FilterAllowToHook and is used to iterate over the raw logs and unpacked data for AllowToHook events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicAllowToHookIterator struct {
	Event *SablierV2LockupDynamicAllowToHook // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicAllowToHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicAllowToHook)
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
		it.Event = new(SablierV2LockupDynamicAllowToHook)
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
func (it *SablierV2LockupDynamicAllowToHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicAllowToHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicAllowToHook represents a AllowToHook event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicAllowToHook struct {
	Admin     common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowToHook is a free log retrieval operation binding the contract event 0xb4378d4e289cb3f40f4f75a99c9cafa76e3df1c4dc31309babc23dc91bd72801.
//
// Solidity: event AllowToHook(address indexed admin, address recipient)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterAllowToHook(opts *bind.FilterOpts, admin []common.Address) (*SablierV2LockupDynamicAllowToHookIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "AllowToHook", adminRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicAllowToHookIterator{contract: _SablierV2LockupDynamic.contract, event: "AllowToHook", logs: logs, sub: sub}, nil
}

// WatchAllowToHook is a free log subscription operation binding the contract event 0xb4378d4e289cb3f40f4f75a99c9cafa76e3df1c4dc31309babc23dc91bd72801.
//
// Solidity: event AllowToHook(address indexed admin, address recipient)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchAllowToHook(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicAllowToHook, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "AllowToHook", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicAllowToHook)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "AllowToHook", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseAllowToHook(log types.Log) (*SablierV2LockupDynamicAllowToHook, error) {
	event := new(SablierV2LockupDynamicAllowToHook)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "AllowToHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicApprovalIterator struct {
	Event *SablierV2LockupDynamicApproval // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicApproval)
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
		it.Event = new(SablierV2LockupDynamicApproval)
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
func (it *SablierV2LockupDynamicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicApproval represents a Approval event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SablierV2LockupDynamicApprovalIterator, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicApprovalIterator{contract: _SablierV2LockupDynamic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicApproval)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseApproval(log types.Log) (*SablierV2LockupDynamicApproval, error) {
	event := new(SablierV2LockupDynamicApproval)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicApprovalForAllIterator struct {
	Event *SablierV2LockupDynamicApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicApprovalForAll)
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
		it.Event = new(SablierV2LockupDynamicApprovalForAll)
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
func (it *SablierV2LockupDynamicApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicApprovalForAll represents a ApprovalForAll event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SablierV2LockupDynamicApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicApprovalForAllIterator{contract: _SablierV2LockupDynamic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicApprovalForAll)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseApprovalForAll(log types.Log) (*SablierV2LockupDynamicApprovalForAll, error) {
	event := new(SablierV2LockupDynamicApprovalForAll)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicBatchMetadataUpdateIterator struct {
	Event *SablierV2LockupDynamicBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicBatchMetadataUpdate)
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
		it.Event = new(SablierV2LockupDynamicBatchMetadataUpdate)
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
func (it *SablierV2LockupDynamicBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*SablierV2LockupDynamicBatchMetadataUpdateIterator, error) {

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicBatchMetadataUpdateIterator{contract: _SablierV2LockupDynamic.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicBatchMetadataUpdate)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseBatchMetadataUpdate(log types.Log) (*SablierV2LockupDynamicBatchMetadataUpdate, error) {
	event := new(SablierV2LockupDynamicBatchMetadataUpdate)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicCancelLockupStreamIterator is returned from FilterCancelLockupStream and is used to iterate over the raw logs and unpacked data for CancelLockupStream events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicCancelLockupStreamIterator struct {
	Event *SablierV2LockupDynamicCancelLockupStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicCancelLockupStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicCancelLockupStream)
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
		it.Event = new(SablierV2LockupDynamicCancelLockupStream)
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
func (it *SablierV2LockupDynamicCancelLockupStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicCancelLockupStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicCancelLockupStream represents a CancelLockupStream event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicCancelLockupStream struct {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterCancelLockupStream(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, asset []common.Address) (*SablierV2LockupDynamicCancelLockupStreamIterator, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "CancelLockupStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicCancelLockupStreamIterator{contract: _SablierV2LockupDynamic.contract, event: "CancelLockupStream", logs: logs, sub: sub}, nil
}

// WatchCancelLockupStream is a free log subscription operation binding the contract event 0x5edb27d6c1a327513b90a792050debf074b7194444885e3144d4decc5caaaa50.
//
// Solidity: event CancelLockupStream(uint256 streamId, address indexed sender, address indexed recipient, address indexed asset, uint128 senderAmount, uint128 recipientAmount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchCancelLockupStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicCancelLockupStream, sender []common.Address, recipient []common.Address, asset []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "CancelLockupStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicCancelLockupStream)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "CancelLockupStream", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseCancelLockupStream(log types.Log) (*SablierV2LockupDynamicCancelLockupStream, error) {
	event := new(SablierV2LockupDynamicCancelLockupStream)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "CancelLockupStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicCreateLockupDynamicStreamIterator is returned from FilterCreateLockupDynamicStream and is used to iterate over the raw logs and unpacked data for CreateLockupDynamicStream events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicCreateLockupDynamicStreamIterator struct {
	Event *SablierV2LockupDynamicCreateLockupDynamicStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicCreateLockupDynamicStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicCreateLockupDynamicStream)
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
		it.Event = new(SablierV2LockupDynamicCreateLockupDynamicStream)
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
func (it *SablierV2LockupDynamicCreateLockupDynamicStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicCreateLockupDynamicStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicCreateLockupDynamicStream represents a CreateLockupDynamicStream event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicCreateLockupDynamicStream struct {
	StreamId     *big.Int
	Funder       common.Address
	Sender       common.Address
	Recipient    common.Address
	Amounts      LockupCreateAmounts
	Asset        common.Address
	Cancelable   bool
	Transferable bool
	Segments     []LockupDynamicSegment
	Timestamps   LockupDynamicTimestamps
	Broker       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateLockupDynamicStream is a free log retrieval operation binding the contract event 0x33eb09bbf19ea3fb22c760d5164234f8bf62ca07dcf5a437ad389e96b0bd6443.
//
// Solidity: event CreateLockupDynamicStream(uint256 streamId, address funder, address indexed sender, address indexed recipient, (uint128,uint128) amounts, address indexed asset, bool cancelable, bool transferable, (uint128,uint64,uint40)[] segments, (uint40,uint40) timestamps, address broker)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterCreateLockupDynamicStream(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, asset []common.Address) (*SablierV2LockupDynamicCreateLockupDynamicStreamIterator, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "CreateLockupDynamicStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicCreateLockupDynamicStreamIterator{contract: _SablierV2LockupDynamic.contract, event: "CreateLockupDynamicStream", logs: logs, sub: sub}, nil
}

// WatchCreateLockupDynamicStream is a free log subscription operation binding the contract event 0x33eb09bbf19ea3fb22c760d5164234f8bf62ca07dcf5a437ad389e96b0bd6443.
//
// Solidity: event CreateLockupDynamicStream(uint256 streamId, address funder, address indexed sender, address indexed recipient, (uint128,uint128) amounts, address indexed asset, bool cancelable, bool transferable, (uint128,uint64,uint40)[] segments, (uint40,uint40) timestamps, address broker)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchCreateLockupDynamicStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicCreateLockupDynamicStream, sender []common.Address, recipient []common.Address, asset []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "CreateLockupDynamicStream", senderRule, recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicCreateLockupDynamicStream)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "CreateLockupDynamicStream", log); err != nil {
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

// ParseCreateLockupDynamicStream is a log parse operation binding the contract event 0x33eb09bbf19ea3fb22c760d5164234f8bf62ca07dcf5a437ad389e96b0bd6443.
//
// Solidity: event CreateLockupDynamicStream(uint256 streamId, address funder, address indexed sender, address indexed recipient, (uint128,uint128) amounts, address indexed asset, bool cancelable, bool transferable, (uint128,uint64,uint40)[] segments, (uint40,uint40) timestamps, address broker)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseCreateLockupDynamicStream(log types.Log) (*SablierV2LockupDynamicCreateLockupDynamicStream, error) {
	event := new(SablierV2LockupDynamicCreateLockupDynamicStream)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "CreateLockupDynamicStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicMetadataUpdateIterator struct {
	Event *SablierV2LockupDynamicMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicMetadataUpdate)
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
		it.Event = new(SablierV2LockupDynamicMetadataUpdate)
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
func (it *SablierV2LockupDynamicMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicMetadataUpdate represents a MetadataUpdate event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*SablierV2LockupDynamicMetadataUpdateIterator, error) {

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicMetadataUpdateIterator{contract: _SablierV2LockupDynamic.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicMetadataUpdate)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseMetadataUpdate(log types.Log) (*SablierV2LockupDynamicMetadataUpdate, error) {
	event := new(SablierV2LockupDynamicMetadataUpdate)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicRenounceLockupStreamIterator is returned from FilterRenounceLockupStream and is used to iterate over the raw logs and unpacked data for RenounceLockupStream events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicRenounceLockupStreamIterator struct {
	Event *SablierV2LockupDynamicRenounceLockupStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicRenounceLockupStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicRenounceLockupStream)
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
		it.Event = new(SablierV2LockupDynamicRenounceLockupStream)
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
func (it *SablierV2LockupDynamicRenounceLockupStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicRenounceLockupStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicRenounceLockupStream represents a RenounceLockupStream event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicRenounceLockupStream struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRenounceLockupStream is a free log retrieval operation binding the contract event 0x0eb069207093cd3e51cd1370d2d369770057fbe29947e577e5fb428c6c6fc78f.
//
// Solidity: event RenounceLockupStream(uint256 indexed streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterRenounceLockupStream(opts *bind.FilterOpts, streamId []*big.Int) (*SablierV2LockupDynamicRenounceLockupStreamIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "RenounceLockupStream", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicRenounceLockupStreamIterator{contract: _SablierV2LockupDynamic.contract, event: "RenounceLockupStream", logs: logs, sub: sub}, nil
}

// WatchRenounceLockupStream is a free log subscription operation binding the contract event 0x0eb069207093cd3e51cd1370d2d369770057fbe29947e577e5fb428c6c6fc78f.
//
// Solidity: event RenounceLockupStream(uint256 indexed streamId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchRenounceLockupStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicRenounceLockupStream, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "RenounceLockupStream", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicRenounceLockupStream)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "RenounceLockupStream", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseRenounceLockupStream(log types.Log) (*SablierV2LockupDynamicRenounceLockupStream, error) {
	event := new(SablierV2LockupDynamicRenounceLockupStream)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "RenounceLockupStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicSetNFTDescriptorIterator is returned from FilterSetNFTDescriptor and is used to iterate over the raw logs and unpacked data for SetNFTDescriptor events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicSetNFTDescriptorIterator struct {
	Event *SablierV2LockupDynamicSetNFTDescriptor // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicSetNFTDescriptorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicSetNFTDescriptor)
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
		it.Event = new(SablierV2LockupDynamicSetNFTDescriptor)
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
func (it *SablierV2LockupDynamicSetNFTDescriptorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicSetNFTDescriptorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicSetNFTDescriptor represents a SetNFTDescriptor event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicSetNFTDescriptor struct {
	Admin            common.Address
	OldNFTDescriptor common.Address
	NewNFTDescriptor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetNFTDescriptor is a free log retrieval operation binding the contract event 0xa2548bd4b805e907c1558a47b5858324fe8bb4a2e1ddfca647eecbf65610eebc.
//
// Solidity: event SetNFTDescriptor(address indexed admin, address oldNFTDescriptor, address newNFTDescriptor)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterSetNFTDescriptor(opts *bind.FilterOpts, admin []common.Address) (*SablierV2LockupDynamicSetNFTDescriptorIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "SetNFTDescriptor", adminRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicSetNFTDescriptorIterator{contract: _SablierV2LockupDynamic.contract, event: "SetNFTDescriptor", logs: logs, sub: sub}, nil
}

// WatchSetNFTDescriptor is a free log subscription operation binding the contract event 0xa2548bd4b805e907c1558a47b5858324fe8bb4a2e1ddfca647eecbf65610eebc.
//
// Solidity: event SetNFTDescriptor(address indexed admin, address oldNFTDescriptor, address newNFTDescriptor)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchSetNFTDescriptor(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicSetNFTDescriptor, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "SetNFTDescriptor", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicSetNFTDescriptor)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseSetNFTDescriptor(log types.Log) (*SablierV2LockupDynamicSetNFTDescriptor, error) {
	event := new(SablierV2LockupDynamicSetNFTDescriptor)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "SetNFTDescriptor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicTransferIterator struct {
	Event *SablierV2LockupDynamicTransfer // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicTransfer)
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
		it.Event = new(SablierV2LockupDynamicTransfer)
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
func (it *SablierV2LockupDynamicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicTransfer represents a Transfer event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SablierV2LockupDynamicTransferIterator, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicTransferIterator{contract: _SablierV2LockupDynamic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicTransfer)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseTransfer(log types.Log) (*SablierV2LockupDynamicTransfer, error) {
	event := new(SablierV2LockupDynamicTransfer)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicTransferAdminIterator is returned from FilterTransferAdmin and is used to iterate over the raw logs and unpacked data for TransferAdmin events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicTransferAdminIterator struct {
	Event *SablierV2LockupDynamicTransferAdmin // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicTransferAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicTransferAdmin)
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
		it.Event = new(SablierV2LockupDynamicTransferAdmin)
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
func (it *SablierV2LockupDynamicTransferAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicTransferAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicTransferAdmin represents a TransferAdmin event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicTransferAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferAdmin is a free log retrieval operation binding the contract event 0xbdd36143ee09de60bdefca70680e0f71189b2ed7acee364b53917ad433fdaf80.
//
// Solidity: event TransferAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterTransferAdmin(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*SablierV2LockupDynamicTransferAdminIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "TransferAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicTransferAdminIterator{contract: _SablierV2LockupDynamic.contract, event: "TransferAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferAdmin is a free log subscription operation binding the contract event 0xbdd36143ee09de60bdefca70680e0f71189b2ed7acee364b53917ad433fdaf80.
//
// Solidity: event TransferAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchTransferAdmin(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicTransferAdmin, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "TransferAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicTransferAdmin)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "TransferAdmin", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseTransferAdmin(log types.Log) (*SablierV2LockupDynamicTransferAdmin, error) {
	event := new(SablierV2LockupDynamicTransferAdmin)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "TransferAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierV2LockupDynamicWithdrawFromLockupStreamIterator is returned from FilterWithdrawFromLockupStream and is used to iterate over the raw logs and unpacked data for WithdrawFromLockupStream events raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicWithdrawFromLockupStreamIterator struct {
	Event *SablierV2LockupDynamicWithdrawFromLockupStream // Event containing the contract specifics and raw log

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
func (it *SablierV2LockupDynamicWithdrawFromLockupStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierV2LockupDynamicWithdrawFromLockupStream)
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
		it.Event = new(SablierV2LockupDynamicWithdrawFromLockupStream)
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
func (it *SablierV2LockupDynamicWithdrawFromLockupStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierV2LockupDynamicWithdrawFromLockupStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierV2LockupDynamicWithdrawFromLockupStream represents a WithdrawFromLockupStream event raised by the SablierV2LockupDynamic contract.
type SablierV2LockupDynamicWithdrawFromLockupStream struct {
	StreamId *big.Int
	To       common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFromLockupStream is a free log retrieval operation binding the contract event 0x40b88e5c41c5a97ffb7b6ef88a0a2d505aa0c634cf8a0275cb236ea7dd87ed4d.
//
// Solidity: event WithdrawFromLockupStream(uint256 indexed streamId, address indexed to, address indexed asset, uint128 amount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) FilterWithdrawFromLockupStream(opts *bind.FilterOpts, streamId []*big.Int, to []common.Address, asset []common.Address) (*SablierV2LockupDynamicWithdrawFromLockupStreamIterator, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.FilterLogs(opts, "WithdrawFromLockupStream", streamIdRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &SablierV2LockupDynamicWithdrawFromLockupStreamIterator{contract: _SablierV2LockupDynamic.contract, event: "WithdrawFromLockupStream", logs: logs, sub: sub}, nil
}

// WatchWithdrawFromLockupStream is a free log subscription operation binding the contract event 0x40b88e5c41c5a97ffb7b6ef88a0a2d505aa0c634cf8a0275cb236ea7dd87ed4d.
//
// Solidity: event WithdrawFromLockupStream(uint256 indexed streamId, address indexed to, address indexed asset, uint128 amount)
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) WatchWithdrawFromLockupStream(opts *bind.WatchOpts, sink chan<- *SablierV2LockupDynamicWithdrawFromLockupStream, streamId []*big.Int, to []common.Address, asset []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SablierV2LockupDynamic.contract.WatchLogs(opts, "WithdrawFromLockupStream", streamIdRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierV2LockupDynamicWithdrawFromLockupStream)
				if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "WithdrawFromLockupStream", log); err != nil {
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
func (_SablierV2LockupDynamic *SablierV2LockupDynamicFilterer) ParseWithdrawFromLockupStream(log types.Log) (*SablierV2LockupDynamicWithdrawFromLockupStream, error) {
	event := new(SablierV2LockupDynamicWithdrawFromLockupStream)
	if err := _SablierV2LockupDynamic.contract.UnpackLog(event, "WithdrawFromLockupStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
