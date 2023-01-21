package util

import "reflect"

func MergeStructs(in, out interface{}) {
	valIn := reflect.ValueOf(in)
	valOut := reflect.ValueOf(out)
	indirectedOut := reflect.Indirect(valOut)
	indirectedIn := reflect.Indirect(valIn)

	for i := 0; i < indirectedOut.NumField(); i++ {
		fieldOut := indirectedOut.Field(i)
		fieldName := indirectedOut.Type().Field(i).Name

		fieldIn := indirectedIn.FieldByName(fieldName)
		if !fieldIn.IsValid() {
			continue
		}

		if fieldOut.Type() == fieldIn.Type() && fieldOut.CanSet() {
			fieldOut.Set(fieldIn)
		}
	}
}
