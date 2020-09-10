package dyanmic_params

import "log"

type ParamsSource interface {
	Add(name string, value interface{}) ParamsSource
	Get(name string) interface{}
	Has(name string) bool
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