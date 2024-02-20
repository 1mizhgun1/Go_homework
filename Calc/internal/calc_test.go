package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const DELTA = 1e-8

func TestCalcSuccess(t *testing.T) {
	var testCases = []struct {
		expression     string
		expectedResult float64
	}{
		{"24+17", 24.0 + 17.0},
		{"12-13", 12.0 - 13.0},
		{"21*34", 21.0 * 34.0},
		{"21/8", 21.0 / 8.0},
		{"11*(12-13)", 11.0 * (12.0 - 13.0)},
		{"  ( 53-  4* (3 /(  17+ 27* 4)-  8)+(  ((43 -11)*  2+1 ) -1)/2)  /2+ 8 ",
			(53.0-4.0*(3.0/(17.0+27.0*4.0)-8.0)+(((43.0-11.0)*2.0+1.0)-1.0)/2.0)/2.0 + 8.0},
		{"1/3", 1.0 / 3.0},
		{"1", 1.0},
		{"-1+2", -1.0 + 2.0},
		{"2+(-1)-(-5)*(3+(-1))", 2.0 + (-1.0) - (-5.0)*(3.0+(-1.0))},
	}

	for _, test := range testCases {
		realResult, err := Calc(test.expression)
		assert.InDelta(t, realResult, test.expectedResult, DELTA)
		assert.Nil(t, err)
	}
}

func TestCalcFail(t *testing.T) {
	var testCases = []string{"3*2+", "(2+3)*(/3+2)", "3+2)", "1+a", ")-1+2", "", "-"}

	for _, expression := range testCases {
		realResult, err := Calc(expression)
		assert.InDelta(t, realResult, answerWhenError, DELTA)
		assert.NotNil(t, err)
	}
}
