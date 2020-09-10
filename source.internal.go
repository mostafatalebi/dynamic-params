package dyanmic_params

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

func (s *SourceInternal) Has(name string) bool {
	if s.storage == nil {
		return false
	} else if _, ok := s.storage[name]; ok {
		return true
	}
	return false
}
