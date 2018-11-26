package bloom_test

import (
	"bytes"
	"testing"

	"github.com/btcsuite/btcd/wire"
	btcbloom "github.com/btcutil/bloom"
	"github.com/sammy00/bloom"
)

const (
	Tweak = 0x00000005
	C     = 0xfba4c795
)

func TestFilter_Add(t *testing.T) {
	testCases := []struct {
		data   []byte
		expect []byte // the expected bits array in snapshot
	}{
		{bloom.Hexlify("99108ad8ed9bb6274d3980bab5a85c048f0950c8"), nil},
		{bloom.Hexlify("19108ad8ed9bb6274d3980bab5a85c048f0950c8"), nil},
		{bloom.Hexlify("b5a2c786d9ef4658287ced5914b37a1b4aa32eee"), nil},
		{bloom.Hexlify("b9300670b4c5366e95b2699e8b18bc75e5f729c5"), nil},
	}

	for i, c := range testCases {
		filter := bloom.New(3, C, Tweak, 0.01)
		btc := btcbloom.NewFilter(3, Tweak, 0.01, wire.BloomUpdateAll)

		if err := filter.Add(c.data); nil != err {
			t.Fatalf("#%d unexpected error: %v", i, err)
		}

		//btc.Add(c.data)

		if got := filter.Snapshot().Bits; !bytes.Equal(got, c.expect) {
			t.Fatalf("#%d invalid bits: got %v, expect %v", i, got, c.expect)
		}
	}
}