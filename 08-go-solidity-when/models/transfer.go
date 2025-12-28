package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
}
