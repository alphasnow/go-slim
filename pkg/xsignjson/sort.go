package xsignjson

import (
	"sort"
	"strings"
)

func getSortMapKey(data map[string]string) []string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func stringMapToSortMap(dataStr map[string]string) string {
	sortKeys := getSortMapKey(dataStr)

	dataArr := make([]string, 0)
	for _, k := range sortKeys {
		val := dataStr[k]
		dataArr = append(dataArr, k+"="+val)
	}

	resStr := strings.Join(dataArr, "&")
	return resStr
}
