package tests

import (
	dp "github.com/mostafatalebi/dynamic-params"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicParams_AddInt(t *testing.T) {
	p := dp.NewDynamicParams(dp.SrcNameInternal)
	p.Add("sample-int", 55)
	v, err := p.GetAsInt("sample-int")
	assert.NoError(t, err)
	assert.Equal(t, 55, v)
}

func TestDynamicParams_AddBool(t *testing.T) {
	p := dp.NewDynamicParams(dp.SrcNameInternal)
	p.Add("sample-bool", true)
	v, err := p.GetAsBool("sample-bool")
	assert.NoError(t, err)
	assert.Equal(t, true, v)
}

func TestDynamicParams_GetFromArgs(t *testing.T) {
	p := dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key=someValue"})
	v, err := p.GetAsString("key")
	assert.NoError(t, err)
	assert.Equal(t, "someValue", v)
}

func TestDynamicParams_GetFromArgs_castToBool(t *testing.T) {
	p := dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key=true"})
	v, err := p.GetStringAsBool("key")
	assert.NoError(t, err)
	assert.Equal(t, true, v)
}


func TestDynamicParams_GetFromArgsNumericAsInt(t *testing.T) {
	p := dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key=123456"})
	v, err := p.GetStringAsInt("key")
	assert.NoError(t, err)
	assert.Equal(t, 123456, v)
}

func TestDynamicParams_GetFromArgsNumericAsInt_mustFail(t *testing.T) {
	p := dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key=0123456"})
	v, err := p.GetStringAsInt("key")
	assert.Error(t, err)
	assert.Equal(t, 0, v)
}


func TestDynamicParams_Struct(t *testing.T) {
	type CustomType struct {
		Name string
	}
	p := dp.NewDynamicParams(dp.SrcNameInternal)
	p.Add("key", &CustomType{Name: "Robert"})
	v := p.Get("key")
	assert.NotNil(t, v)
	r, ok := v.(*CustomType)
	assert.True(t, ok)
	if r != nil {
		assert.Equal(t, "Robert", r.Name)
	}
}



func TestDynamicParams_GetFromArgs_AsQuotedString(t *testing.T) {
	p := dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key=\"someValue\""})
	v, err := p.GetAsQuotedString("key")
	assert.NoError(t, err)
	assert.Equal(t, "someValue", v)

	p = dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key='someValue'"})
	v, err = p.GetAsQuotedString("key")
	assert.NoError(t, err)
	assert.Equal(t, "someValue", v)

	p = dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key='someValue'"})
	v, err = p.GetAsString("key")
	assert.NoError(t, err)
	assert.Equal(t, "'someValue'", v)
}