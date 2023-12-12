package mempool

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
)

func (c *MempoolClient) GetRawTransaction(txHash *chainhash.Hash) (*wire.MsgTx, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/tx/%s/raw", txHash.String()), nil)
	if err != nil {
		return nil, err
	}

	tx := wire.NewMsgTx(wire.TxVersion)
	if err := tx.Deserialize(bytes.NewReader(res)); err != nil {
		return nil, err
	}
	return tx, nil
}

type Fee struct {
	fastestFee   int    `json:"fastestFee"`
	halfHourFee   int    `json:"halfHourFee"`
	hourFee   int    `json:"hourFee"`
	economyFee   int    `json:"economyFee"`
	minimumFee   int    `json:"minimumFee"`
}

func (c *MempoolClient) GetFeeRate() (*int, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/v1/fees/recommended"), nil)
	if err != nil {
		return nil, err
	}

	var fee Fee
	err = json.Unmarshal(res, &fee)
	if err != nil {
		return nil, err
	}

	return &fee.fastestFee, nil
}

func (c *MempoolClient) BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, err
	}

	res, err := c.request(http.MethodPost, "/tx", strings.NewReader(hex.EncodeToString(buf.Bytes())))
	if err != nil {
		return nil, err
	}

	txHash, err := chainhash.NewHashFromStr(string(res))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to parse tx hash, %s", string(res)))
	}
	return txHash, nil
}
