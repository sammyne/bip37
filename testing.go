package bip37

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"

	btcwire "github.com/btcsuite/btcd/wire"
)

func NewOutPoint(hash []byte, index uint32) *btcwire.OutPoint {
	var chash chainhash.Hash
	copy(chash[:], hash)

	return btcwire.NewOutPoint(&chash, index)
}

func Unhexlify(str string) []byte {
	out, _ := hex.DecodeString(str)
	return out
}

func ReadBlock(t *testing.T) *btcwire.MsgBlock {
	fd, err := os.Open(filepath.Join("..", "testdata", "block.json"))
	if nil != err {
		t.Fatal(err)
	}
	defer fd.Close()

	msgBlock := new(btcwire.MsgBlock)
	if err := json.NewDecoder(fd).Decode(msgBlock); nil != err {
		t.Fatal(err)
	}

	return msgBlock
}
