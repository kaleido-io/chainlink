package adapters_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"chainlink/core/adapters"
	"chainlink/core/internal/cltest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiply_Unmarshal(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		exp     adapters.Multiply
	}{
		{
			"w/ value",
			`{"Times": 5}`,
			adapters.Multiply{Times: big.NewFloat(5)},
		},
		{
			"w/o value",
			`{}`,
			adapters.Multiply{Times: nil},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var m adapters.Multiply
			err := json.Unmarshal([]byte(tc.payload), &m)
			require.NoError(t, err)
			require.Equal(t, tc.exp, m)
		})
	}
}

func TestMultiply_Perform(t *testing.T) {
	tests := []struct {
		name    string
		Times   *big.Float
		json    string
		want    string
		errored bool
	}{
		{"by 100", big.NewFloat(100), `{"result":"1.23"}`, "123", false},
		{"float", big.NewFloat(100), `{"result":1.23}`, "123", false},
		{"negative", big.NewFloat(-5), `{"result":"1.23"}`, "-6.15", false},
		{"no times parameter", nil, `{"result":"3.14"}`, "3.14", false},
		{"object", big.NewFloat(100), `{"result":{"foo":"bar"}}`, "", true},
		{"zero", big.NewFloat(0), `{"result":"1.23"}`, "0", false},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			input := cltest.NewRunInputWithString(t, test.json)
			adapter := adapters.Multiply{Times: test.Times}
			result := adapter.Perform(input, nil)

			if test.errored {
				require.Error(t, result.Error())
				return
			}

			require.NoError(t, result.Error())
			assert.Equal(t, test.want, result.Result().String())
		})
	}
}
