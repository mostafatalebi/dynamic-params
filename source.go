package dyanmic_params

import "log"

type ParamsIteratorFn = func(key string, value interface{})

type ParamsSource interface {
	Add(name string, value interface{}) ParamsSource
	Get(name string) interface{}

	Has(name string) bool
	Count() int64

	// It supports regex pattern
	// It doesn't cache any result, therefore, each
	// time scan starts from the beginning of the list
	// if you have many params, be sure not to use it
	// often
	Scan(regex string) map[string]interface{}

	// Iterates through all params and allows you to
	// apply your call back to them. Their values (and not keys)
	// are passed by reference,
	// Hence any direct change on them, mutates the original value
	Iterate(fn ParamsIteratorFn)
}

func NewSource(name string, vars ...interface{}) ParamsSource {
	if name == SrcNameInternal {
		return NewSourceInternal()
	} else if name == SrcNameArgs {
		if len(vars) == 0 {
			log.Fatal("SourceArgs must have a args collection passed to NewSource()")
		}
		return NewSourceArgs(vars[0])
	}
	return nil
}