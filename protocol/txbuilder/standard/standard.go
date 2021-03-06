// Package standard implements standard txvm contracts for Sequence
// transactions.
package standard

import (
	"i10r.io/protocol/txvm/asm"
	"i10r.io/protocol/txvm/op"
	"i10r.io/protocol/txvm/txvmutil"
)

// VerifyTxID returns a program that verifies the txid matches the
// given value.
func VerifyTxID(txid [32]byte) []byte {
	var b txvmutil.Builder
	b.Op(op.TxID)
	b.PushdataBytes(txid[:])
	b.Op(op.Eq).Op(op.Verify)
	return b.Build()
}

func mustAssemble(src string) []byte {
	res, err := asm.Assemble(src)
	if err != nil {
		panic(err)
	}
	return res
}
