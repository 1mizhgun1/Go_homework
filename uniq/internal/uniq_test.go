package uniq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqSuccess(t *testing.T) {
	var testCases = []struct {
		arguments      Arguments
		input          []string
		expectedOutput []string
	}{
		{ // no arguments
			Arguments{c: false, d: false, u: false, i: false, num: 0, chars: 0, input: "", output: ""},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "", "I love music of Kartik.", "Thanks.", "I love music of Kartik."},
		},
		{ // only "-c" argument
			Arguments{c: true, d: false, u: false, i: false, num: 0, chars: 0, input: "", output: ""},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks.", "2 I love music of Kartik."},
		},
		{ // only "-d" argument
			Arguments{c: false, d: true, u: false, i: false, num: 0, chars: 0, input: "", output: ""},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"I love music.", "I love music of Kartik.", "I love music of Kartik."},
		},
		{ // only "-u" argument
			Arguments{c: false, d: false, u: true, i: false, num: 0, chars: 0, input: "", output: ""},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
			[]string{"", "Thanks."},
		},
		{ // only "-i" argument
			Arguments{c: false, d: false, u: false, i: true, num: 0, chars: 0, input: "", output: ""},
			[]string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks.", "I love music of Kartik.", "I love MuSIC of Kartik."},
			[]string{"I LOVE MUSIC.", "", "I love MuSIC of Kartik.", "Thanks.", "I love music of Kartik."},
		},
		{ // only "-f num" argument with num = 1
			Arguments{c: false, d: false, u: false, i: false, num: 1, chars: 0, input: "", output: ""},
			[]string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"We love music.", "", "I love music of Kartik.", "Thanks."},
		},
		{ // only "-s chars" argument with chars = 1
			Arguments{c: false, d: false, u: false, i: false, num: 0, chars: 1, input: "", output: ""},
			[]string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			[]string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		},
		{ // input is empty
			Arguments{c: false, d: false, u: false, i: false, num: 0, chars: 0, input: "", output: ""},
			[]string{},
			nil,
		},
	}

	for _, test := range testCases {
		realOutput, err := Uniq(test.input, test.arguments)
		assert.Equal(t, realOutput, test.expectedOutput)
		assert.Nil(t, err)
	}
}

func TestUniqFail(t *testing.T) {
	var testCases = []struct {
		arguments Arguments
		input     []string
	}{
		{ // "-c" and "-d" arguments concurrently
			Arguments{c: true, d: true, u: false, i: false, num: 0, chars: 0, input: "", output: ""},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
		},
		{ // "-f num" arguments less than zero
			Arguments{c: false, d: false, u: false, i: false, num: -1, chars: 0, input: "", output: ""},
			[]string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks.", "I love music of Kartik.", "I love music of Kartik."},
		},
	}

	for _, test := range testCases {
		realOutput, err := Uniq(test.input, test.arguments)
		assert.Nil(t, realOutput)
		assert.NotNil(t, err)
	}
}
