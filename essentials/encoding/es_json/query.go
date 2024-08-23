package es_json

import (
	"errors"
	"github.com/itchyny/gojq"
)

var (
	ErrorQuerySingleNotFound = errors.New("query result not found")
)

func QuerySingle(obj interface{}, query *gojq.Query) (interface{}, error) {
	iter := query.Run(obj)
	for {
		v, ok := iter.Next()
		if !ok {
			return nil, ErrorQuerySingleNotFound
		}
		if err, ok := v.(error); ok {
			var hErr *gojq.HaltError
			if errors.As(err, &hErr) && hErr.Value() == nil {
				return nil, ErrorQuerySingleNotFound
			}
			return nil, err
		}
		return v, nil
	}
}
