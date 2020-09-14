package dyanmic_params

import (
	"errors"
	"strings"
	"sync"
	"time"
)

type DynamicParams struct {
	Mx *sync.RWMutex
	source ParamsSource
}

const (
	ErrNotFound = "not found"
)

// Returns a  new instance of DynamicParams
//
// source is the name of the source to use, available sources are:
// - SrcNameInternal
// - SrcNameArgs
//
// If you want to have DynamicParams concurrent safe, you must pass
// a *sync.Mutex{} as first argument in vars...
//
// Example: To create a new instance with SrnNameArgs, do this:
// NewDynamicParams(SrcNameArgs, &mx, os.Args)
// vars... it is a list of extra parameters that a source might need. For example,
// for SrcNameArgs, you need to pass os.Args (or an array of string you want to treat
// as list of arguments)
func NewDynamicParams(source string, vars ...interface{}) *DynamicParams {
	return createDP(source, vars...)
}

func createDP(source string, vars ...interface{}) *DynamicParams {
	var mx *sync.RWMutex
	var varsNew []interface{}
	if len(vars) > 0  {
		if v, ok := vars[0].(*sync.RWMutex); ok {
			mx = v
			varsNew = make([]interface{}, len(vars)-1)
			for i := 0; i < len(vars); i++ {
				if i == 0 {
					continue
				}
				varsNew = append(varsNew, vars[i])
			}
		} else {
			varsNew = vars
		}
	}

	return &DynamicParams{
		Mx: mx,
		source: NewSource(source, varsNew...),
	}
}


// adds a key and value to the active underlying source
func (c *DynamicParams) Add(name string, value interface{}) *DynamicParams {
	if c.Mx != nil {
		c.Mx.Lock()
		defer c.Mx.Unlock()
	}
	c.source.Add(name, value)
	return c
}

// Checks to see if the value exists in the underlying source
func (c *DynamicParams) Has(name string) bool {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	return c.source.Has(name)
}

// returns the raw value, if exists, and if not found, returns nil
// this function is useful for storing struct and custom compound types
func (c *DynamicParams) Get(name string) interface{} {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	return c.source.Get(name)
}

func (c *DynamicParams) Scan(regex string) map[string]interface{} {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	return c.source.Scan(regex)
}
func (c *DynamicParams) Count() int64 {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	return c.source.Count()
}

// This function is not concurrent safe
// If you need to have concurrency and do locking over this usage,
// you must use your own function for iteration and handle that there.
func (c *DynamicParams) Iterate(fn func(key string, value interface{})) {
	c.source.Iterate(fn)
}

// Casts the existing value to string and then returns it
// It tries to convert value of interface{} type to string
// and if it fails, it returns error
func (c *DynamicParams) GetAsString(name string) (string, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return "", errors.New(ErrNotFound)
	}
	return convertToString(v)
}

// this method removes any surrounding quotation marks (only surrounding)
func (c *DynamicParams) GetAsQuotedString(name string) (string, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return "", errors.New(ErrNotFound)
	}
	s, err := convertToString(v)
	if err != nil {
		return "", err
	}
	return strings.Trim(strings.Trim(s, "'"), "\""), nil
}


func (c *DynamicParams) GetAsInt(name string) (int, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return 0, errors.New(ErrNotFound)
	}
	return convertToInt(v)
}

func (c *DynamicParams) GetStringAsInt(name string) (int, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return 0, errors.New(ErrNotFound)
	}
	return convertNumericStrToInt(v)
}

// parses a string using time.ParseDuration() function
func (c *DynamicParams) GetStringAsTimeDuration(name string) (*time.Duration, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return nil, errors.New(ErrNotFound)
	}
	vs, err := convertToString(v)
	if err != nil {
		return nil, err
	}
	vd, err := time.ParseDuration(vs)
	if err != nil {
		return nil, errors.New(ErrCnvFailed)
	}
	return &vd, nil
}

// if string has these values: 0, 1, true or false,
// then this method converts them bool type and then returns
// the value
func (c *DynamicParams) GetStringAsBool(name string) (bool, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return false, errors.New(ErrNotFound)
	}
	return convertNumericStrToBool(v)
}

func (c *DynamicParams) GetAsInt32(name string) (int32, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return 0, errors.New(ErrNotFound)
	}
	return convertToInt32(v)
}

func (c *DynamicParams) GetAsInt64(name string) (int64, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return 0, errors.New(ErrNotFound)
	}
	return convertToInt64(v)
}

func (c *DynamicParams) GetAsInt8(name string) (int8, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return 0, errors.New(ErrNotFound)
	}
	return convertToInt8(v)
}
func (c *DynamicParams) GetAsInt16(name string) (int16, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return 0, errors.New(ErrNotFound)
	}
	return convertToInt16(v)
}

func (c *DynamicParams) GetAsTimeDuration(name string) (*time.Duration, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return nil, errors.New(ErrNotFound)
	}
	return convertToTimeDuration(v)
}

func (c *DynamicParams) GetAsBool(name string) (bool, error) {
	if c.Mx != nil {
		c.Mx.RLock()
		defer c.Mx.RUnlock()
	}
	v := c.source.Get(name)
	if v == nil {
		return false, errors.New(ErrNotFound)
	}
	return convertToBool(v)
}


