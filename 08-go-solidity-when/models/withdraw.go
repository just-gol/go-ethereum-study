package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Withdraw struct {
	Src common.Address
	Wad *big.Int
}
