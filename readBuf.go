package esm

import (
	"reflect"
	"fmt"
)

func (b *readBuf) readType(t reflect.Type, v reflect.Value) (error) {
	//fmt.Println(t.Kind())
	switch t.Kind() {
	case reflect.Map:
		v.Set(reflect.MakeMap(t))
	case reflect.Struct:
		fieldCount := v.NumField()
		for i := 0; i < fieldCount; i+=1 {
			structFieldType := t.Field(i)
			structFieldValue := v.Field(i)
			b.readType(structFieldType.Type, structFieldValue)
		}
	default:

		var rv reflect.Value

		switch v.Interface().(type) {

		case char:
			rv = reflect.ValueOf(b.char())

		case wchar:
			rv = reflect.ValueOf(b.wchar())

		case uint8:
			rv = reflect.ValueOf(b.uint8())

		case uint16:
			rv = reflect.ValueOf(b.uint16())

		case uint32:
			rv = reflect.ValueOf(b.uint32())

		case uint64:
			rv = reflect.ValueOf(b.uint64())

		case int8:
			rv = reflect.ValueOf(b.int8())

		case int16:
			rv = reflect.ValueOf(b.int16())

		case int32:
			rv = reflect.ValueOf(b.int32())

		case int64:
			rv = reflect.ValueOf(b.int64())

		case float32:
			rv = reflect.ValueOf(b.float32())
			//subVRv := reflect.ValueOf(subV)
			//subRrv = reflect.Value(subVRv)

		case float64:
			rv = reflect.ValueOf(b.float64())

		case zstring:
			rv = reflect.ValueOf(b.zstring())

		default:
			panic(fmt.Errorf("cannot decode type '%s'", v.Type()))
		}

		v.Set(reflect.Indirect(rv))

		return nil
	}
	return nil
}

