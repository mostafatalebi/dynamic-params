package tests

import (
	dyanmic_params "dynamicparams"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDyanmicParams_AddInt(t *testing.T) {
	p := dyanmic_params.NewCustomParams()
	p.Add("sample-int", 55)
	v, err := p.GetAsInt("sample-int")
	assert.NoError(t, err)
	assert.Equal(t, 55, v)
}

func TestDyanmicParams_AddBool(t *testing.T) {
	p := dyanmic_params.NewCustomParams()
	p.Add("sample-bool", true)
	v, err := p.GetAsBool("sample-bool")
	assert.NoError(t, err)
	assert.Equal(t, true, v)
}
