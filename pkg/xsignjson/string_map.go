package xsignjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func interfaceToString(val interface{}) (string, error) {
	var str string
	switch val.(type) {
	case int:
		str = strconv.Itoa(val.(int))
	case float64:
		str = strconv.FormatFloat(val.(float64), 'f', 2, 64)
		if strings.HasSuffix(str, ".00") {
			str = strings.TrimSuffix(str, ".00")
		}
	case bool:
		if val.(bool) {
			str = "1"
		} else {
			str = "0"
		}
	case string:
		str = val.(string)
	default:
		strByte, err := json.Marshal(val)
		if err != nil {
			return "", err
		}
		str = string(strByte)
	}
	return str, nil
}

func jsonStructToStringMap(data interface{}) (map[string]string, error) {
	result := make(map[string]string)
	typ := reflect.TypeOf(data)
	if typ.Kind() != reflect.Struct {
		return nil, errors.New(fmt.Sprintf("data error kind %v", typ.Kind()))
	}

	val := reflect.ValueOf(data)
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			resultMap, err := jsonStructToStringMap(val.Field(i).Interface())
			if err != nil {
				return nil, err
			}
			for km, vm := range resultMap {
				result[km] = vm
			}
			continue
		}

		k := typ.Field(i).Tag.Get("json")
		if k == "-" || k == "" {
			continue
		}
		if strings.Contains(k, ",") {
			idx := strings.Index(k, ",")
			k = k[0:idx]
		}

		v, err := interfaceToString(val.Field(i).Interface())
		if err != nil {
			return nil, err
		}
		result[k] = v

	}

	return result, nil
}

func jsonMapToStringMap(data map[string]interface{}) (map[string]string, error) {
	result := make(map[string]string)

	for k, v := range data {
		val, err := interfaceToString(v)
		if err != nil {
			return nil, err
		}
		result[k] = val
	}

	return result, nil
}
