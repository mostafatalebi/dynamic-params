package dyanmic_params


import (
	"errors"
	"strconv"
	"strings"
)

type dynamicParamCollection map[string]interface{}

type DyanmicParams struct {
	params dynamicParamCollection
}

func NewCustomParams() *DyanmicParams {
	return &DyanmicParams{
		params: make(dynamicParamCollection, 0),
	}
}


func (c *DyanmicParams) Add(name string, value interface{}) *DyanmicParams {
	if c.params == nil {
		c.params = make(dynamicParamCollection, 0)
	}
	c.params[name] = value
	return c
}

func (c *DyanmicParams) getFromArgs(name string, args []string) string {
	for _, v := range args {
		if strings.Contains(v, "--"+name) {
			spl := strings.Replace(v, "--"+name+"=", "", 1)
			if len(spl) > 0 {
				return spl
			}
		}
	}
	return ""
}

// must be --key=value format and name must NOT include --
func (c *DyanmicParams) AddFromArgsAsInt(name string, args []string) *DyanmicParams {
	argValue := c.getFromArgs(name, args)
	vl, err := strconv.Atoi(argValue)
	if err == nil {
		c.Add(name, vl)
	}
	return c
}

func (c *DyanmicParams) AddFromArgsAsBool(name string, args []string) *DyanmicParams {
	argValue := c.getFromArgs(name, args)
	if argValue == "false" || argValue == "0" {
		c.Add(name, false)
	} else if argValue == "true" || argValue == "1" {
		c.Add(name, true)
	}
	return c
}

func (c *DyanmicParams) AddFromArgsAsStr(name string, args []string) *DyanmicParams {
	argValue := c.getFromArgs(name, args)
	if argValue != "" {
		c.Add(name, argValue)
	}
	return c
}




func (c *DyanmicParams) Get(name string) interface{} {
	if c.params == nil {
		return ""
	} else if val, ok := c.params[name]; ok {
		return val
	}
	return ""
}



func (c *DyanmicParams) GetAsString(name string) (string, error) {
	if c.params == nil {
		return "", errors.New("not fond")
	} else if val, ok := c.params[name]; ok {
		if v, ok := val.(string); ok {
			return v, errors.New("not fond")
		} else if v, ok := val.(*string); ok {
			return *v, errors.New("not fond")
		}
	}
	return "", errors.New("not fond")
}

func (c *DyanmicParams) GetAsInt(name string) (int, error) {
	if c.params == nil {
		return 0, errors.New("not fond")
	} else if val, ok := c.params[name]; ok {
		if v, ok := val.(int); ok {
			return v, nil
		} else if v, ok := val.(*int); ok {
			return *v, nil
		}
	}
	return 0, errors.New("not fond")
}

func (c *DyanmicParams) GetAsBool(name string) (bool, error) {
	if c.params == nil {
		return false, errors.New("not fond")
	} else if val, ok := c.params[name]; ok {
		if v, ok := val.(bool); ok {
			return v, nil
		} else if v, ok := val.(*bool); ok {
			return *v, nil
		}
	}
	return false, errors.New("not fond")
}


func (c *DyanmicParams) Has(name string) bool {
	if c.params == nil {
		return false
	} else if _, ok := c.params[name]; ok {
		return true
	}
	return false
}
