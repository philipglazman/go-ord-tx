package mempool

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func TestGetRawTransaction(t *testing.T) {
	//https://mempool.space/signet/api/tx/b752d80e97196582fd02303f76b4b886c222070323fb7ccd425f6c89f5445f6c/hex
	client := NewClient(&chaincfg.SigNetParams)
	txId, _ := chainhash.NewHashFromStr("b752d80e97196582fd02303f76b4b886c222070323fb7ccd425f6c89f5445f6c")
	transaction, err := client.GetRawTransaction(txId)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(transaction.TxHash().String())
	}
}

func TestGetFee(t *testing.T) {
	// https://mempool.space/docs/api/rest#get-recommended-fees
	client := NewClient(&chaincfg.MainNetParams)
	fee, err := client.GetFeeRate()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(*fee)
	}
}
