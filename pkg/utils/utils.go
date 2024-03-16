package utils

import (
	"net/url"
	"strconv"
)

func ConvertArray2Str[T int64 | string](list []T) string {
	var arrayStr string
	for i, v := range list {
		if i != 0 {
			arrayStr += ","
		}

		switch t := any(v).(type) {
		case string:
			arrayStr += string(t)
		case int64:
			arrayStr += strconv.FormatInt(t, 10)
		}

	}

	return arrayStr
}

func SetPagination(values *url.Values, limit int, page int) {
	values.Add("limit", strconv.Itoa(limit))
	values.Add("page_no", strconv.Itoa(page))
}
