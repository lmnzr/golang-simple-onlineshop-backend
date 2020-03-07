package model

import (
	"encoding/json"
	"errors"

	// "fmt"
	"reflect"
)

//StructMap : Map Struct Properties
type StructMap struct {
	Map     map[string]interface{}
	Fields  []string
	Values  []interface{}
	Indexes []int
}

//NewStructMap : Create Struct Map Instances
func NewStructMap(model interface{}) StructMap {
	var fields []string
	var values []interface{}
	var indexes []int

	structmap := make(map[string]interface{})

	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	for i := 0; i < t.NumField(); i++ {
		if swtch, ok := t.Field(i).Tag.Lookup("hidden"); !(ok || swtch == "true") {
			indexes = append(indexes, i)

			if val, found := t.Field(i).Tag.Lookup("json"); found {
				fields = append(fields, val)
			} else {
				fields = append(fields, t.Field(i).Name)
			}

			val := v.Field(i).Interface()
			values = append(values, val)

			structmap[fields[i]] = values[i]
		}
	}

	return StructMap{
		Map:     structmap,
		Fields:  fields,
		Values:  values,
		Indexes: indexes,
	}
}

//PrintModelJSON : Print Struct Or Map Into JSON formatted String, Hidden Field Omitted
func PrintModelJSON(model interface{}) (string, error) {
	v := reflect.ValueOf(model)
	var res []byte
	var err error

	if v.Kind() == reflect.Struct {
		structmap := NewStructMap(model)
		for i := 0; i < len(structmap.Fields); i++ {
			var idx []int
			idx = append(idx, structmap.Indexes[i])
			structmap.Map[structmap.Fields[i]] = v.FieldByIndex(idx).Interface()
		}
		res, err = json.Marshal(structmap.Map)
	} else if v.Kind() == reflect.Map {
		res, err = json.Marshal(model)
	} else {
		res = nil
		err = errors.New("Model is not a Struct nor a Map")
	}

	return string(res), err
}

//IsFieldExist : Check if Field Exist in Tag
func IsFieldExist(model interface{}, fieldname string, tag string) bool {
	fieldexist := false

	t := reflect.TypeOf(model)

	for i := 0; i < t.NumField(); i++ {
		if tag == "" {
			if t.Field(i).Name == fieldname {
				fieldexist = true
			}
		} else {
			if t.Field(i).Tag.Get(tag) == fieldname {
				fieldexist = true
			}
		}
	}

	return fieldexist
}
