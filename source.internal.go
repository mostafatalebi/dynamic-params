package dyanmic_params

import "regexp"

const SrcNameInternal = "source.internal"

type internalParamCollection map[string]interface{}

type SourceInternal struct {

	storage internalParamCollection
}

func NewSourceInternal() *SourceInternal {
	return &SourceInternal{
		storage: make(internalParamCollection, 0),
	}
}

func (s *SourceInternal) Add(name string, value interface{}) ParamsSource {
	s.storage[name] = value
	return s
}

func (s *SourceInternal) Get(name string) interface{} {
	if s.storage == nil {
		return ""
	} else if val, ok := s.storage[name]; ok {
		return val
	}
	return nil
}


func (s *SourceInternal) Scan(regex string) map[string]interface{} {
	if s.Count() > 0 {
		mp := make(map[string]interface{}, 0)
		rg, err := regexp.Compile(regex)
		if err != nil {
			return nil
		}
		for k, v := range s.storage {
			if rg.MatchString(k) {
				mp[k] = v
			}
		}

		return mp
	}
	return nil
}
func (s *SourceInternal) Iterate(fn func(k string, v interface{})) {
	if s.Count() > 0 {
		for k, v := range s.storage {
			fn(k, v)
		}
	}
	return
}

func (s *SourceInternal) Has(name string) bool {
	if s.storage == nil {
		return false
	} else if _, ok := s.storage[name]; ok {
		return true
	}
	return false
}

func (s *SourceInternal) Count() int64 {
	if s.storage == nil {
		return 0
	}
	return int64(len(s.storage))
}
