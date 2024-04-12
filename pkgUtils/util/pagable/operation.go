package pagable

import "errors"

type Operation string

const (
	Equal       Operation = "$eq"
	EqualI                = "$eqi" // (case-insensitive)
	NotEqual              = "$ne"
	LT                    = "$lt"
	LTE                   = "$lte"
	GT                    = "$gt"
	GTE                   = "$gte"
	In                    = "$in"
	NotIn                 = "$notIn"
	Contains              = "$contains"
	NotContains           = "$notContains"
	IsNull                = "$null"
	IsNotNull             = "$notNull"
	StartsWith            = "$startsWith"
	EndsWith              = "$endsWith"
	Search                = "$search"
)

var operationMap = map[string]Operation{
	"$eq":          Equal,
	"$eqi":         EqualI,
	"$ne":          NotEqual,
	"$lt":          LT,
	"$lte":         LTE,
	"$gt":          GT,
	"$gte":         GTE,
	"$in":          In,
	"$notIn":       NotIn,
	"$contains":    Contains,
	"$notContains": NotContains,
	"$null":        IsNull,
	"$notNull":     IsNotNull,
	"$startsWith":  StartsWith,
	"$endsWith":    EndsWith,
	"$search":      Search,
}

func OperationMapping(key string) (Operation, error) {
	operator := operationMap[key]
	if operator == "" {
		return "", OperatorNotSupportErr
	}
	return operator, nil
}

var (
	OperatorNotSupportErr = errors.New("operator not support")
)
