// (c) 2024-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package types

import (
	ethtypes "github.com/ava-labs/libevm/core/types"
)

type isMultiCoin bool

var (
	extras = ethtypes.RegisterExtras[
		HeaderExtra, *HeaderExtra,
		BlockExtra, *BlockExtra,
		BodyExtra, *BodyExtra,
		isMultiCoin]()
	IsMultiCoinPayloads = extras.StateAccount
)

func IsMultiCoin(s ethtypes.StateOrSlimAccount) bool {
	return bool(extras.StateAccount.Get(s))
}

func HeaderExtras(h *Header) *HeaderExtra {
	return extras.Header.Get(h)
}

func WithHeaderExtras(h *Header, extra *HeaderExtra) *Header {
	extras.Header.Set(h, extra)
	return h
}

func BlockExtras(b *Block) *BlockExtra {
	return extras.Block.Get(b)
}

func WithBlockExtras(b *Block, version uint32, extdata *[]byte) *Block {
	extra := &BlockExtra{
		version: version,
		extdata: extdata,
	}
	extras.Block.Set(b, extra)
	return b
}
