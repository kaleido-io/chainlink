package adapters

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONBigFloatUnmarshalFloat64(t *testing.T) {
	tests := []struct {
		payload float64
		exp     *big.Float
	}{
		{-1, big.NewFloat(-1)},
		{100, big.NewFloat(100)},
		{3.146, big.NewFloat(3.146)},
	}

	for _, tc := range tests {
		var j JSONBigFloat
		buf, err := json.Marshal(tc.payload)
		require.NoError(t, err)
		err = json.Unmarshal([]byte(buf), &j)
		require.NoError(t, err)
		assert.Equal(t, tc.exp.String(), j.Value().String())
	}
}

func TestJSONBigFloatUnmarshalString(t *testing.T) {
	tests := []struct {
		payload string
		exp     *big.Float
	}{
		{"-1", big.NewFloat(-1)},
		{"100", big.NewFloat(100)},
		{"3.146", big.NewFloat(3.146)},
	}

	for _, tc := range tests {
		var j JSONBigFloat
		buf, err := json.Marshal(tc.payload)
		require.NoError(t, err)
		err = json.Unmarshal([]byte(buf), &j)
		require.NoError(t, err)
		assert.Equal(t, tc.exp.String(), j.Value().String())
	}
}
