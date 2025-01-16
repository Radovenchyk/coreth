// (c) 2019-2020, Ava Labs, Inc.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package types contains data types related to Ethereum consensus.
package types

import (
	"math/big"
)

// Body is a simple (mutable, non-safe) data container for storing and moving
// a block's data contents (transactions and uncles) together.
type Body struct {
	Transactions []*Transaction
	Uncles       []*Header
	Version      uint32
	ExtData      *[]byte `rlp:"nil"`
}

// BlockBody returns the body of the block containing extras.
// Note the returned data is not an independent copy.
func BlockBody(b *Block) *Body {
	return &Body{
		Transactions: b.Transactions(),
		Uncles:       b.Uncles(),
		Version:      BlockVersion(b),
		ExtData:      BlockExtras(b).extdata,
	}
}

func BlockGasCost(b *Block) *big.Int {
	if HeaderExtras(b.Header()).BlockGasCost == nil {
		return nil
	}
	return new(big.Int).Set(HeaderExtras(b.Header()).BlockGasCost)
}
