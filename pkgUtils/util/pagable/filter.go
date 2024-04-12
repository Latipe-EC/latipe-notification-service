package pagable

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"net/url"
	"regexp"
	"strings"
)

const (
	filterPattern = "filters\\[(.*?)\\]\\[(.*?)\\]=(.*?)(?:&|\\z)"
)

// Compile the regex pattern
var filterRegex = regexp.MustCompile(filterPattern)

// ExpressionFilter is struct for add filtering in slice or array
// ref: https://docs.strapi.io/dev-docs/api/rest/filters-locale-publication#filtering
type Filter struct {
	Field     string      `json:"field"`
	Value     interface{} `json:"value"`
	Operation Operation   `json:"operation"`
}

type FilterMapValue struct {
	Value     []string
	Operation Operation `json:"operation"`
}

func decodeFilterURL(encodedUrl string) (string, error) {
	decodedUrl, err := url.QueryUnescape(encodedUrl)
	if err != nil {
		return "", err
	}

	log.Info("url:%v", decodedUrl)
	return decodedUrl, nil
}

func FilterBinding(uri string) ([]Filter, error) {

	urlDecode, err := decodeFilterURL(uri)
	if err != nil {
		return nil, err
	}

	var filters []Filter
	var keyArr []string
	filterMap := make(map[string]FilterMapValue)

	// Find all matches in the uri
	matches := filterRegex.FindAllStringSubmatch(urlDecode, -1)
	for _, match := range matches {
		comp, err := OperationMapping(match[2])
		if err != nil {

			arrComp, err := OperationMapping(match[2][:len(match[2])-3])
			if err != nil {
				return nil, err
			}
			fieldName := fmt.Sprintf("%v", match[1])
			_, ok := filterMap[fieldName]
			if !ok {
				keyArr = append(keyArr, fieldName)
			}

			newSlice := append(filterMap[fieldName].Value, match[3])
			filterMap[fieldName] = FilterMapValue{
				Value:     newSlice,
				Operation: arrComp,
			}

		} else {

			filter := Filter{
				Field:     match[1],
				Value:     match[3],
				Operation: comp,
			}
			filters = append(filters, filter)
		}
	}

	for _, i := range keyArr {
		filter := Filter{
			Field:     i,
			Value:     filterMap[i].Value,
			Operation: filterMap[i].Operation,
		}
		filters = append(filters, filter)
	}

	return filters, nil
}

func ArrayToString(arr []string) string {
	var strArr []string

	for _, value := range arr {
		strArr = append(strArr, fmt.Sprintf("%v", value))
	}

	return strings.Join(strArr, ",")
}
