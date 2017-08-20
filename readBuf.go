package esm

import (
	"reflect"
	"fmt"
	"math"
	"encoding/binary"

	"regexp"
)


type readBuf []byte






//func (b *readBuf) uint16() uint16 {
//	v := binary.LittleEndian.Uint16(*b)
//	*b = (*b)[2:]
//	return v
//}
//
//func (b *readBuf) uint32() uint32 {
//	v := binary.LittleEndian.Uint32(*b)
//	*b = (*b)[4:]
//	return v
//}
//
//func (b *readBuf) uint64() uint64 {
//	v := binary.LittleEndian.Uint64(*b)
//	*b = (*b)[8:]
//	return v
//}
//
//func (b *readBuf) byte() byte {
//	v := (*b)[0]
//	*b = (*b)[1:]
//	return v
//}
//
//func (b *readBuf) skip(size int64) {
//	*b = (*b)[size:]
//	return
//}






func (b *readBuf) char() char { return char(b.uint8()) }

func (b *readBuf) wchar() wchar {panic("Unimplemented readBuf type")}


//func (b *readBuf) bool() bool {
//	x := (*b)[0]
//	*b = (*b)[1:]
//	return x != 0
//}

func (b *readBuf) uint8() uint8 {
	x := (*b)[0]
	*b = (*b)[1:]
	return x
}

func (b *readBuf) uint16() uint16 {
	x := binary.LittleEndian.Uint16((*b)[0:2])
	*b = (*b)[2:]
	return x
}

func (b *readBuf) uint32() uint32 {
	x := binary.LittleEndian.Uint32((*b)[0:4])
	*b = (*b)[4:]
	return x
}

func (b *readBuf) uint64() uint64 {
	x := binary.LittleEndian.Uint64((*b)[0:8])
	*b = (*b)[8:]
	return x
}

func (b *readBuf) int8() int8 { return int8(b.uint8()) }

func (b *readBuf) int16() int16 { return int16(b.uint16()) }

func (b *readBuf) int32() int32 { return int32(b.uint32()) }

func (b *readBuf) int64() int64 { return int64(b.uint64()) }

func (b *readBuf) float32() float32 { return math.Float32frombits(b.uint32()) }

func (b *readBuf) float64() float64 { return math.Float64frombits(b.uint64()) }




// DEPRECATED TYPES
func (b *readBuf) ubyte() {panic("Deprecated readBuf type")}
func (b *readBuf) short() {panic("Deprecated readBuf type")}
func (b *readBuf) ushort() {panic("Deprecated readBuf type")}
func (b *readBuf) long() {panic("Deprecated readBuf type")}
func (b *readBuf) ulong() {panic("Deprecated readBuf type")}
func (b *readBuf) float() {panic("Deprecated readBuf type")}


//// todo: write special parser for this
//func(b *readBuf) vsval() { return []byte }
//
func(b *readBuf) formid() formid { return formid(b.uint32()) }

//func(b *readBuf) iref() { return uint32 }
//func(b *readBuf) hash() { //return8]uint8 }
//func(b *readBuf) filetime() { return uint64 }
//func(b *readBuf) systemtime() { // return6]uint8 }
//func(b *readBuf) rgb() { return uint32 }
//
func(b *readBuf) lstring(root *Root) lstring {
	if root.IsLocalized() {
		lookup := b.uint32()
		return lstring(lookup)
	}
	return lstring(b.zstring())
}

//func(b *readBuf) dlstring() { return string }
//func(b *readBuf) ilstring() { return string }
//func(b *readBuf) bstring() { return string }
//func(b *readBuf) bzstring() { return string }
//func(b *readBuf) wstring() { return string }
//func(b *readBuf) wzstring() { return string }
func(b *readBuf) zstring() zstring {
	i := 0;
	for (*b)[i] != 0 {
		i += 1
	}
	x := zstring((*b)[0:i])
	*b = (*b)[i+1:]
	return x
}

func (b *readBuf) Human() string {
	r := regexp.MustCompile("[^a-zA-Z0-9_ ,/?\\\\+=()\\][&^%$#@!~'\":<>-]")
	out := r.ReplaceAll([]byte(*b), []byte("."))
	return string(out)
}



func (b *readBuf) readType(t reflect.Type, v reflect.Value, f *Field) (error) {
	//fmt.Println(t.Kind())
	switch t.Kind() {
	//case reflect.Func:
	//
	//	swap := func(in []reflect.Value) []reflect.Value {
	//		return []reflect.Value{in[1], in[0]}
	//	}
	//
	//	var makeSwap func()
	//
	//
	//	//args := [reflect.ValueOf(b)]
	//	newFn := reflect.MakeFunc(v, v)
	//
	//	v.Set(newFn)
	//	fmt.Printf("%#v", &v)
	case reflect.Map:
		v.Set(reflect.MakeMap(t))
	case reflect.Struct:
		fieldCount := v.NumField()
		for i := 0; i < fieldCount; i+=1 {
			structFieldType := t.Field(i)
			structFieldValue := v.Field(i)
			b.readType(structFieldType.Type, structFieldValue, f)
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

		case lstring:

			rv = reflect.ValueOf(b.lstring(f.Root()))
		case formid:
			rv = reflect.ValueOf(b.formid())

		//case func(readBuf, Field) interface{}:
		//	fn := v //reflect.ValueOf(v)
		//	fmt.Printf("555 %#v", fn, fn.Interface())
		//	args := []reflect.Value{}
		//	result := fn.Call(args)
		//	rv = reflect.ValueOf(result)

		default:
			panic(fmt.Errorf("cannot decode type '%s'", v.Type()))
		}

		v.Set(reflect.Indirect(rv))

		return nil
	}
	return nil
}

