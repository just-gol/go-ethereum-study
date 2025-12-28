package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Approval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
}
