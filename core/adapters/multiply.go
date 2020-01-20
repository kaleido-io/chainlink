// +build !sgx_enclave

package adapters

import (
	"encoding/json"
	"fmt"
	"math/big"

	"chainlink/core/store"
	"chainlink/core/store/models"
)

// Multiply holds the a number to multiply the given value by.
type Multiply struct {
	Times *big.Float `json:"-"`
}

type jsonMultiply struct {
	Times *JSONBigFloat `json:"times"`
}

// MarshalJSON implements the json.Marshal interface.
func (ma *Multiply) MarshalJSON() ([]byte, error) {
	jsonObj := jsonMultiply{(*JSONBigFloat)(ma.Times)}
	return json.Marshal(jsonObj)
}

// UnmarshalJSON implements the json.Unmarshal interface.
func (ma *Multiply) UnmarshalJSON(buf []byte) error {
	var jsonObj jsonMultiply
	err := json.Unmarshal(buf, &jsonObj)
	if err != nil {
		return err
	}
	ma.Times = jsonObj.Times.Value()
	return nil
}

// Perform returns the input's "result" field, multiplied times the adapter's
// "times" field.
//
// For example, if input value is "99.994" and the adapter's "times" is
// set to "100", the result's value will be "9999.4".
func (ma *Multiply) Perform(input models.RunInput, _ *store.Store) models.RunOutput {
	val := input.Result()
	i, ok := (&big.Float{}).SetString(val.String())
	if !ok {
		return models.NewRunOutputError(fmt.Errorf("cannot parse into big.Float: %v", val.String()))
	}

	if ma.Times != nil {
		i.Mul(i, ma.Times)
	}
	return models.NewRunOutputCompleteWithResult(i.String())
}
