package dyanmic_params

import (
	"strings"
	"sync"
)

type DynamicParams struct {
	Mx *sync.RWMutex
	source ParamsSource
}


// Returns a  new instance of DynamicParams
//
// source is the name of the source to use, available sources are:
// - SrcNameInternal
// - SrcNameArgs
//
// mx if true, does all operations with mutex
//
// vars... it is a list of extra parameters that a source might need. For example,
// for SrcNameArgs, you need to pass os.Args (or an array of string you want to treat
// as list of arguments)
func NewDynamicParams(source string, mx bool, vars ...interface{}) *DynamicParams {
	return createDP(source, mx, vars...)
}

func createDP(source string, mx bool, vars ...interface{}) *DynamicParams {
	return &DynamicParams{
		Mx: &sync.RWMutex{},
		source: NewSource(source, vars...),
	}
}


// adds a key and value to the active underlying source
func (c *DynamicParams) Add(name string, value interface{}) *DynamicParams {
	c.Mx.Lock()
	defer c.Mx.Unlock()
	c.source.Add(name, value)
	return c
}

// Checks to see if the value exists in the underlying source
func (c *DynamicParams) Has(name string) bool {
	c.Mx.Lock()
	defer c.Mx.Unlock()
	return c.source.Has(name)
}

// returns the raw value, if exists, and if not found, returns nil
// this function is useful for storing struct and custom compound types
func (c *DynamicParams) Get(name string) interface{} {
	return c.source.Get(name)
}


// Casts the existing value to string and then returns it
// It tries to convert value of interface{} type to string
// and if it fails, it returns error
func (c *DynamicParams) GetAsString(name string) (string, error) {
	v := c.source.Get(name)
	return convertToString(v)
}

// this method removes any surrounding quotation marks (only surrounding)
func (c *DynamicParams) GetAsQuotedString(name string) (string, error) {
	v := c.source.Get(name)
	s, err := convertToString(v)
	if err != nil {
		return "", err
	}
	return strings.Trim(strings.Trim(s, "'"), "\""), nil
}


func (c *DynamicParams) GetAsInt(name string) (int, error) {
	v := c.source.Get(name)
	return convertToInt(v)
}

func (c *DynamicParams) GetStringAsInt(name string) (int, error) {
	v := c.source.Get(name)
	return convertNumericStrToInt(v)
}

// if string has these values: 0, 1, true or false,
// then this method converts them bool type and then returns
// the value
func (c *DynamicParams) GetStringAsBool(name string) (bool, error) {
	v := c.source.Get(name)
	return convertNumericStrToBool(v)
}

func (c *DynamicParams) GetAsInt32(name string) (int32, error) {
	v := c.source.Get(name)
	return convertToInt32(v)
}

func (c *DynamicParams) GetAsInt64(name string) (int64, error) {
	v := c.source.Get(name)
	return convertToInt64(v)
}

func (c *DynamicParams) GetAsInt8(name string) (int8, error) {
	v := c.source.Get(name)
	return convertToInt8(v)
}
func (c *DynamicParams) GetAsInt16(name string) (int16, error) {
	v := c.source.Get(name)
	return convertToInt16(v)
}

func (c *DynamicParams) GetAsBool(name string) (bool, error) {
	v := c.source.Get(name)
	return convertToBool(v)
}


