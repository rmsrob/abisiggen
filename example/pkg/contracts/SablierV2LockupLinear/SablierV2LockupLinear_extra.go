// Code generated by abisiggen.sh for SablierV2LockupLinear; DO NOT EDIT.
package SablierV2LockupLinear

import (
	"github.com/ethereum/go-ethereum/common"
)

// SablierV2LockupLinearTopics contains event signatures for SablierV2LockupLinear
type SablierV2LockupLinearTopics struct {
	AllowToHook              common.Hash `json:"AllowToHook"`              // 0xb4378d4e289cb3f40f4f75a99c9cafa76e3df1c4dc31309babc23dc91bd72801
	Approval                 common.Hash `json:"Approval"`                 // 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925
	ApprovalForAll           common.Hash `json:"ApprovalForAll"`           // 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31
	BatchMetadataUpdate      common.Hash `json:"BatchMetadataUpdate"`      // 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c
	CancelLockupStream       common.Hash `json:"CancelLockupStream"`       // 0x5edb27d6c1a327513b90a792050debf074b7194444885e3144d4decc5caaaa50
	CreateLockupLinearStream common.Hash `json:"CreateLockupLinearStream"` // 0x0edca1b6565301e70a0a79732f82c688b70562d471f5606395690f3ec3366ac3
	MetadataUpdate           common.Hash `json:"MetadataUpdate"`           // 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7
	RenounceLockupStream     common.Hash `json:"RenounceLockupStream"`     // 0x0eb069207093cd3e51cd1370d2d369770057fbe29947e577e5fb428c6c6fc78f
	SetNFTDescriptor         common.Hash `json:"SetNFTDescriptor"`         // 0xa2548bd4b805e907c1558a47b5858324fe8bb4a2e1ddfca647eecbf65610eebc
	Transfer                 common.Hash `json:"Transfer"`                 // 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
	TransferAdmin            common.Hash `json:"TransferAdmin"`            // 0xbdd36143ee09de60bdefca70680e0f71189b2ed7acee364b53917ad433fdaf80
	WithdrawFromLockupStream common.Hash `json:"WithdrawFromLockupStream"` // 0x40b88e5c41c5a97ffb7b6ef88a0a2d505aa0c634cf8a0275cb236ea7dd87ed4d
}

// SablierV2LockupLinearMethods contains method signatures for SablierV2LockupLinear
type SablierV2LockupLinearMethods struct {
	AllowToHook            []byte `json:"AllowToHook"`            // 0x406887cb (allowToHook(address recipient))
	Approve                []byte `json:"Approve"`                // 0x095ea7b3 (approve(address to, uint256 tokenId))
	Burn                   []byte `json:"Burn"`                   // 0x42966c68 (burn(uint256 streamId))
	Cancel                 []byte `json:"Cancel"`                 // 0x40e58ee5 (cancel(uint256 streamId))
	CancelMultiple         []byte `json:"CancelMultiple"`         // 0x8659c270 (cancelMultiple(uint256[] streamIds))
	CreateWithDurations    []byte `json:"CreateWithDurations"`    // 0xab167ccc (createWithDurations((address,address,uint128,address,bool,bool,(uint40,uint40))
	CreateWithTimestamps   []byte `json:"CreateWithTimestamps"`   // 0x53b15727 (createWithTimestamps((address,address,uint128,address,bool,bool,(uint40,uint40,uint40))
	Renounce               []byte `json:"Renounce"`               // 0x7de6b1db (renounce(uint256 streamId))
	SafeTransferFrom       []byte `json:"SafeTransferFrom"`       // 0x42842e0e (safeTransferFrom(address from, address to, uint256 tokenId))
	SafeTransferFrom0      []byte `json:"SafeTransferFrom0"`      // 0xb88d4fde (safeTransferFrom(address from, address to, uint256 tokenId, bytes data))
	SetApprovalForAll      []byte `json:"SetApprovalForAll"`      // 0xa22cb465 (setApprovalForAll(address operator, bool approved))
	SetNFTDescriptor       []byte `json:"SetNFTDescriptor"`       // 0x7cad6cd1 (setNFTDescriptor(address newNFTDescriptor))
	TransferAdmin          []byte `json:"TransferAdmin"`          // 0x75829def (transferAdmin(address newAdmin))
	TransferFrom           []byte `json:"TransferFrom"`           // 0x23b872dd (transferFrom(address from, address to, uint256 tokenId))
	Withdraw               []byte `json:"Withdraw"`               // 0xfdd46d60 (withdraw(uint256 streamId, address to, uint128 amount))
	WithdrawMax            []byte `json:"WithdrawMax"`            // 0xea5ead19 (withdrawMax(uint256 streamId, address to))
	WithdrawMaxAndTransfer []byte `json:"WithdrawMaxAndTransfer"` // 0xc156a11d (withdrawMaxAndTransfer(uint256 streamId, address newRecipient))
	WithdrawMultiple       []byte `json:"WithdrawMultiple"`       // 0x4cc55e11 (withdrawMultiple(uint256[] streamIds, uint128[] amounts))
}

// SablierV2LockupLinearMethodNames contains method names with lowercase first letter
type SablierV2LockupLinearMethodNames struct {
	AllowToHook            string `json:"AllowToHook"`
	Approve                string `json:"Approve"`
	Burn                   string `json:"Burn"`
	Cancel                 string `json:"Cancel"`
	CancelMultiple         string `json:"CancelMultiple"`
	CreateWithDurations    string `json:"CreateWithDurations"`
	CreateWithTimestamps   string `json:"CreateWithTimestamps"`
	Renounce               string `json:"Renounce"`
	SafeTransferFrom       string `json:"SafeTransferFrom"`
	SafeTransferFrom0      string `json:"SafeTransferFrom0"`
	SetApprovalForAll      string `json:"SetApprovalForAll"`
	SetNFTDescriptor       string `json:"SetNFTDescriptor"`
	TransferAdmin          string `json:"TransferAdmin"`
	TransferFrom           string `json:"TransferFrom"`
	Withdraw               string `json:"Withdraw"`
	WithdrawMax            string `json:"WithdrawMax"`
	WithdrawMaxAndTransfer string `json:"WithdrawMaxAndTransfer"`
	WithdrawMultiple       string `json:"WithdrawMultiple"`
}

func InitSablierV2LockupLinearTopics() SablierV2LockupLinearTopics {
	return SablierV2LockupLinearTopics{
		AllowToHook:              common.HexToHash("0xb4378d4e289cb3f40f4f75a99c9cafa76e3df1c4dc31309babc23dc91bd72801"),
		Approval:                 common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
		ApprovalForAll:           common.HexToHash("0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"),
		BatchMetadataUpdate:      common.HexToHash("0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c"),
		CancelLockupStream:       common.HexToHash("0x5edb27d6c1a327513b90a792050debf074b7194444885e3144d4decc5caaaa50"),
		CreateLockupLinearStream: common.HexToHash("0x0edca1b6565301e70a0a79732f82c688b70562d471f5606395690f3ec3366ac3"),
		MetadataUpdate:           common.HexToHash("0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7"),
		RenounceLockupStream:     common.HexToHash("0x0eb069207093cd3e51cd1370d2d369770057fbe29947e577e5fb428c6c6fc78f"),
		SetNFTDescriptor:         common.HexToHash("0xa2548bd4b805e907c1558a47b5858324fe8bb4a2e1ddfca647eecbf65610eebc"),
		Transfer:                 common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
		TransferAdmin:            common.HexToHash("0xbdd36143ee09de60bdefca70680e0f71189b2ed7acee364b53917ad433fdaf80"),
		WithdrawFromLockupStream: common.HexToHash("0x40b88e5c41c5a97ffb7b6ef88a0a2d505aa0c634cf8a0275cb236ea7dd87ed4d"),
	}
}

func InitSablierV2LockupLinearMethods() SablierV2LockupLinearMethods {
	return SablierV2LockupLinearMethods{
		AllowToHook:            []byte{0x40, 0x68, 0x87, 0xcb},
		Approve:                []byte{0x09, 0x5e, 0xa7, 0xb3},
		Burn:                   []byte{0x42, 0x96, 0x6c, 0x68},
		Cancel:                 []byte{0x40, 0xe5, 0x8e, 0xe5},
		CancelMultiple:         []byte{0x86, 0x59, 0xc2, 0x70},
		CreateWithDurations:    []byte{0xab, 0x16, 0x7c, 0xcc},
		CreateWithTimestamps:   []byte{0x53, 0xb1, 0x57, 0x27},
		Renounce:               []byte{0x7d, 0xe6, 0xb1, 0xdb},
		SafeTransferFrom:       []byte{0x42, 0x84, 0x2e, 0x0e},
		SafeTransferFrom0:      []byte{0xb8, 0x8d, 0x4f, 0xde},
		SetApprovalForAll:      []byte{0xa2, 0x2c, 0xb4, 0x65},
		SetNFTDescriptor:       []byte{0x7c, 0xad, 0x6c, 0xd1},
		TransferAdmin:          []byte{0x75, 0x82, 0x9d, 0xef},
		TransferFrom:           []byte{0x23, 0xb8, 0x72, 0xdd},
		Withdraw:               []byte{0xfd, 0xd4, 0x6d, 0x60},
		WithdrawMax:            []byte{0xea, 0x5e, 0xad, 0x19},
		WithdrawMaxAndTransfer: []byte{0xc1, 0x56, 0xa1, 0x1d},
		WithdrawMultiple:       []byte{0x4c, 0xc5, 0x5e, 0x11},
	}
}

func InitSablierV2LockupLinearMethodNames() SablierV2LockupLinearMethodNames {
	return SablierV2LockupLinearMethodNames{
		AllowToHook:            "allowToHook",
		Approve:                "approve",
		Burn:                   "burn",
		Cancel:                 "cancel",
		CancelMultiple:         "cancelMultiple",
		CreateWithDurations:    "createWithDurations",
		CreateWithTimestamps:   "createWithTimestamps",
		Renounce:               "renounce",
		SafeTransferFrom:       "safeTransferFrom",
		SafeTransferFrom0:      "safeTransferFrom0",
		SetApprovalForAll:      "setApprovalForAll",
		SetNFTDescriptor:       "setNFTDescriptor",
		TransferAdmin:          "transferAdmin",
		TransferFrom:           "transferFrom",
		Withdraw:               "withdraw",
		WithdrawMax:            "withdrawMax",
		WithdrawMaxAndTransfer: "withdrawMaxAndTransfer",
		WithdrawMultiple:       "withdrawMultiple",
	}
}
