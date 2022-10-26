package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type Environment map[string]interface{}

func NewEnvironment() Environment {
	return make(Environment)
}

func (e Environment) Consume(namespace string, o Output) Environment {
	t := e

	t[namespace] = o

	return t
}

func (e Environment) GetByPath(path string) (interface{}, error) {
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		return nil, errors.New("not found")
	}

	for i, part := range parts {
		res, ok := e[part]
		if !ok {
			return nil, errors.New("not found")
		}

		if i == len(parts)-1 {
			return res, nil
		}

		switch t := res.(type) {
		case map[string]interface{}:
			e = t
		case Output:
			e = Environment(t)
		case url.Values:
			for k, v := range t {
				if len(v) == 1 {
					e[k] = v[0]
				} else {
					e[k] = v
				}
			}
		default:
			panic(t)
		}
	}

	return nil, errors.New("not found")
}

var templateRegex = regexp.MustCompile(`{{\w+[.\w+]+}}`)

func (e Environment) Render(input string) (string, error) {
	return templateRegex.ReplaceAllStringFunc(input, func(s string) string {
		s = strings.TrimSuffix(strings.TrimPrefix(s, "{{"), "}}")
		r, err := e.GetByPath(s)
		if err != nil {
			panic(err)
		}

		switch rr := r.(type) {
		case string:
			return rr
		case map[string]interface{}:
			jsn, err := json.Marshal(rr)
			if err != nil {
				panic(err)
			}

			return string(jsn)
		}

		return fmt.Sprintf("%v", r)
	}), nil
}
