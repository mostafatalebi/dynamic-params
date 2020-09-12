package dyanmic_params

import (
	"log"
	"regexp"
	"strings"
)

const SrcNameArgs = "source.args"

type argsParamCollection map[string]interface{}

type SourceArgs struct {
	storage argsParamCollection
}

// args must be in this format: --key=value
func NewSourceArgs(args interface{}) *SourceArgs {
	var argsTyped, v = args.([]string)
	if !v {
		log.Println("Args param must be in []string type")
		return nil
	}
	return &SourceArgs{
		storage: createMapFromArgs(argsTyped),
	}
}


// converts arguments in --name=value format to a map of names and values,
// with names saved without --
func createMapFromArgs(args []string)  argsParamCollection {
	mc := make(argsParamCollection, 0)
	if len(args) > 0 {
		rg := regexp.MustCompile(`^\-\-[a-zA-Z\-]+\=.+$`)
		for _, v := range args {
			if m := rg.Match([]byte(v)); m == true{
				spl := strings.SplitN(v, "=", 2)
				name := strings.Replace(spl[0], "--", "", 1)
				mc[name] = spl[1]
			}
		}
	}
	return mc
}

func (s *SourceArgs) Add(name string, value interface{}) ParamsSource {
	s.storage[name] = value
	return s
}

func (s *SourceArgs) Get(name string) interface{} {
	if s.storage == nil {
		return ""
	} else if val, ok := s.storage[name]; ok {
		return val
	}
	return nil
}

func (s *SourceArgs) Scan(regex string) map[string]interface{} {
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

func (s *SourceArgs) Iterate(fn func(k string, v interface{})) {
	if s.Count() > 0 {
		for k, v := range s.storage {
			fn(k, v)
		}
	}
	return
}


func (s *SourceArgs) Has(name string) bool {
	if s.storage == nil {
		return false
	} else if _, ok := s.storage[name]; ok {
		return true
	}
	return false
}


func (s *SourceArgs) Count() int64 {
	if s.storage == nil {
		return 0
	}
	return int64(len(s.storage))
}
