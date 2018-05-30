package json_codec

import (
	"fmt"
	"errors"
)

// Get and check isExist field which has ANY type.
func GetAny(field string, fields map[string]interface{}) (interface{}, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return nil, errors.New(fmt.Sprintf("Cannot get ANY field '%s'", field))
	}
	return fieldPure, nil
}

// Get and check isExist field which has STRING type.
func GetString(field string, fields map[string]interface{}) (string, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return "", errors.New(fmt.Sprintf("Cannot get STRING field '%s'", field))
	}
	fieldValue, okFieldValue := fieldPure.(string)
	if !okFieldValue {
		return "", errors.New(fmt.Sprintf("Cannot convert field '%s' to STRING", field))
	}
	return fieldValue, nil
}

// Get and check isExist field which has ARRAY STRING type.
func GetArrayString(field string, fields map[string]interface{}) ([]string, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return nil, errors.New(fmt.Sprintf("Cannot get ARRAY STRING field '%s'", field))
	}
	fieldValueArray, okFieldValueArray := fieldPure.([]interface{})
	if !okFieldValueArray {
		return nil, errors.New(fmt.Sprintf("Cannot convert field '%s' to ARRAY STRING", field))
	}
	fieldValueArrayString := []string{}
	for _, v := range fieldValueArray {
		temp, okTemp := v.(string)
		if !okTemp {
			return nil, errors.New(fmt.Sprintf("Cannot convert field '%s' to ARRAY STRING", field))
		}
		fieldValueArrayString = append(fieldValueArrayString, temp)
	}

	return fieldValueArrayString, nil
}

// Get and check isExist field which has JSON type.
func GetJson(field string, fields map[string]interface{}) (map[string]interface{}, error) {
	fieldPure, okPureField := fields[field]
	if !okPureField {
		return nil, errors.New(fmt.Sprintf("Cannot get JSON field '%s'", field))
	}
	fieldValue, okFieldValue := fieldPure.(map[string]interface{})
	if !okFieldValue {
		return nil, errors.New(fmt.Sprintf("Cannot convert field '%s' to JSON", field))
	}
	return fieldValue, nil
}
