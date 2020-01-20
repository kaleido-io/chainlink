package adapters

import (
	"encoding/json"
	"math/big"
)

// JSONBigFloat accepts both string and float JSON values.
type JSONBigFloat big.Float

// UnmarshalJSON implements the json.Unmarshal interface.
func (j *JSONBigFloat) UnmarshalJSON(buf []byte) error {
	var f float64
	if err := json.Unmarshal(buf, &f); err == nil {
		*j = JSONBigFloat(*big.NewFloat(f))
		return nil
	}
	var bf big.Float
	if err := json.Unmarshal(buf, &bf); err != nil {
		return err
	}
	*j = JSONBigFloat(bf)
	return nil
}

// Value returns the big.Float value.
func (j *JSONBigFloat) Value() *big.Float {
	return (*big.Float)(j)
}
