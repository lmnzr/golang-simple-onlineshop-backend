package database

import (
	// "fmt"
	"reflect"

	"github.com/lmnzr/simpleshop/cmd/simpleshop/models"
)

//StructMap : Map struct to arrays
type StructMap struct {
	Fields  []string
	Values  []interface{}
	Indexes []int
}

type tag struct {
	field  string
	value  string
	remove bool
	unique bool
}

func maptype(types []interface{}, t reflect.Type, i int) []interface{} {
	fieldtype := "string"

	if val, found := t.Field(i).Tag.Lookup("type"); found {
		fieldtype = val
	}

	switch fieldtype {
	case "int":
		types = append(types, new(models.NullInt))
	case "float":
		types = append(types, new(models.NullFloat))
	case "boolean":
		types = append(types, new(models.NullBool))
	case "datetime":
		types = append(types, new(models.NullTime))
	default:
		types = append(types, new(models.NullString))
	}

	return types
}

func mapvalue(fields []string, values []interface{}, indexes []int, tag string, v reflect.Value, t reflect.Type, i int) ([]string, []interface{}, []int) {
	fieldtype := "string"

	if val, found := t.Field(i).Tag.Lookup("type"); found {
		fieldtype = val
	}

	val := v.Field(i).Interface()

	switch fieldtype {
	case "int":
		if val.(models.NullInt).Valid {
			fields = mapfield(fields, tag, t, i)
			values = append(values, val)
			indexes = mapindex(indexes, i)
		}
	case "float":
		if val.(models.NullString).Valid {
			fields = mapfield(fields, tag, t, i)
			values = append(values, val)
			indexes = mapindex(indexes, i)
		}
	case "boolean":
		if val.(models.NullBool).Valid {
			fields = mapfield(fields, tag, t, i)
			values = append(values, val)
			indexes = mapindex(indexes, i)
		}
	case "datetime":
		if val.(models.NullTime).Valid {
			fields = mapfield(fields, tag, t, i)
			values = append(values, val)
			indexes = mapindex(indexes, i)
		}
	case "string":
		if val.(models.NullString).Valid {
			fields = mapfield(fields, tag, t, i)
			values = append(values, val)
			indexes = mapindex(indexes, i)
		}
	}

	return fields, values, indexes
}

func mapfield(fields []string, tag string, t reflect.Type, i int) []string {
	if val, found := t.Field(i).Tag.Lookup(tag); found {
		fields = append(fields, val)
	} else {
		fields = append(fields, t.Field(i).Name)
	}

	return fields
}

func mapindex(indexes []int, i int) []int {
	return append(indexes, i)
}

func mapping(model interface{}, tagfield string, flag tag, hidden bool, cmdtype string) StructMap {
	var fields []string
	var values []interface{}
	var indexes []int

	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	for i := 0; i < t.NumField(); i++ {
		swtch, ok := t.Field(i).Tag.Lookup(flag.field)
		match := ok && swtch == flag.value

		var cond bool

		if flag.remove {
			cond = !match || !hidden
		} else if flag.unique {
			cond = match
		}
		
		if cond {

			switch cmdtype {
			case "SELECT":
				indexes = mapindex(indexes, i)
				fields = mapfield(fields, tagfield, t, i)
				values = maptype(values, t, i)

			case "INSERT", "UPDATE", "DELETE":
				fields, values, indexes = mapvalue(fields, values, indexes, tagfield, v, t, i)
			}
		}
	}

	return StructMap{
		Fields:  fields,
		Values:  values,
		Indexes: indexes,
	}
}

//MapModel : Create Mapping of a struct
func MapModel(model interface{}, cmdtype string) (structmap StructMap) {
	var tagselector tag

	switch cmdtype {
	case "SELECT":
		tagselector = tag{
			field:  "hidden",
			value:  "true",
			remove: true,
			unique: false,
		}
	case "INSERT":
		tagselector = tag{
			field:  "increment",
			value:  "auto",
			remove: true,
		}
	case "UPDATE":
		tagselector = tag{
			field:  "pkey",
			value:  "true",
			remove: true,
			unique: false,
		}
	case "DELETE":
		tagselector = tag{
			field:  "pkey",
			value:  "true",
			remove: false,
			unique: true,
		}
	}

	return mapping(model, "field", tagselector, true, cmdtype)
}
