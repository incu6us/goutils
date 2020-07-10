package assert

import (
	"reflect"
)

func Equals(expected, actual interface{}, skippedFields ...string) bool {
	if isPrimitiveType(expected) || isPrimitiveType(actual) {
		return reflect.DeepEqual(expected, actual)
	}

	actualStruct := parseStruct(actual, skippedFields)
	expectedStruct := parseStruct(expected, skippedFields)

	return reflect.DeepEqual(expectedStruct, actualStruct)
}

func EqualsWithDiffFunc(expected, actual interface{}, skippedFields []string, diffFn func(expected, actual interface{}) bool) bool {
	if isPrimitiveType(expected) || isPrimitiveType(actual) {
		return diffFn(expected, actual)
	}

	actualStruct := parseStruct(actual, skippedFields)
	expectedStruct := parseStruct(expected, skippedFields)

	return diffFn(expectedStruct, actualStruct)
}

func isPrimitiveType(t interface{}) bool {
	typ := reflect.TypeOf(t)
	switch typ.Kind() {
	case reflect.Struct:
		return false
	case reflect.Ptr:
		if typ.Elem().Kind() == reflect.Struct {
			return false
		}

		return true
	default:
		return true
	}
}

func parseStruct(strct interface{}, skippedFields []string) interface{} {
	fn := func(i interface{}) (reflect.Type, reflect.Value, reflect.Value) {
		typ := reflect.TypeOf(i)
		switch typ.Kind() {
		case reflect.Struct:
			return typ, reflect.ValueOf(strct), reflect.Indirect(reflect.New(typ))
		case reflect.Ptr:
			return typ.Elem(), reflect.ValueOf(strct).Elem(), reflect.New(typ.Elem())
		default:
			panic("bad type")
		}
	}

	typ, val, newVal := fn(strct)

	for i := 0; i < typ.NumField(); i++ {
		strctField := typ.Field(i)
		if fieldContains(strctField.Name, skippedFields...) {
			continue
		}

		copyValue(i, newVal, val)
	}

	return newVal.Interface()
}

func fieldContains(s string, fields ...string) bool {
	for _, field := range fields {
		if s == field {
			return true
		}
	}

	return false
}

func copyValue(index int, newVal, val reflect.Value) {
	switch newVal.Kind() {
	case reflect.Struct:
		newVal.Field(index).Set(val.Field(index))
	case reflect.Ptr:
		newVal.Elem().Field(index).Set(val.Field(index))
	}
}
