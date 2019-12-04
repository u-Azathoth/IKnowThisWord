package model_test

import (
	"github.com/stretchr/testify/assert"
	"iKnowThisWord/internal/model"
	"testing"
)

type validateCardTest struct {
	word    string
	meaning string
	isValid bool
}

var validateCardTests = []validateCardTest{
	{"consider", "deem to be", true},
	{"", "deem to be", false},
	{"consider", "", false},
	{"min", "infinitely or immeasurably small", true},
	{"mi", "infinitely or immeasurably small", false},
	{"m", "infinitely or immeasurably small", false},
}

func TestCard_Validate(t *testing.T) {
	for i := range validateCardTests {
		tc := &validateCardTests[i]
		c := &model.Card{Word: tc.word, Meaning: tc.meaning}

		err := c.Validate()

		if tc.isValid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
