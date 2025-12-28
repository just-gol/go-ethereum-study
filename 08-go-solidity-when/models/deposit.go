package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Deposit struct {
	Dst common.Address
	Wad *big.Int
}
