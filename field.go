package esm

import (
	"io"
	"fmt"
	"reflect"
)


type FieldHeader struct {
	_type char4
	dataSize uint32
}



type Field struct {
	FieldHeader
	parentRecord *Record
	sr           *io.SectionReader
	off          int64
	dataBuf      readBuf
	data         interface{}
}


// calculates the size of all the things (header and all data)
func (f *Field) Size() int64 {
	return int64(fieldHeaderLen) + int64(f.dataSize)
}

func (f *Field) Root() *Root {
	return f.parentRecord.Root()
}

func (f *Field) PrevField() *Field {
	fieldCount := len(f.parentRecord.fields)

	if fieldCount == 0 {
		return nil
	}

	prevField := f.parentRecord.fields[fieldCount - 1]

	if f == prevField {
		if fieldCount == 1 {
			return nil
		}
		return f.parentRecord.fields[fieldCount - 2]
	}
	return prevField
}


func (f *Field) readHeader(sr io.SectionReader) error {
	buf := make([]byte, fieldHeaderLen)
	//fmt.Println(sr.Size())
	if _, err := sr.Read(buf); err != nil {
		return err
	}

	b := readBuf(buf[:])

	// TODO: validate signature is in the list of allowed Record header types?

	f._type = char4{byte(b.char()), byte(b.char()), byte(b.char()), byte(b.char())}

	f.dataSize = uint32(b.uint16())

	// if size is 0, check if prev field was XXXX
	if f.dataSize == 0 {
		prevField := f.PrevField()
		if prevField != nil && prevField.Type() == "XXXX" {
			f.dataSize = prevField.data.(uint32)
		}
	}

	//f.dataBuf = make([]byte, f.dataSize)

	return nil
}



func (f *Field) String() string {
	//str := fmt.Sprintf("Field[%s](%d): buff: %s", f._type, f.dataSize, f.dataBuf)
	str := fmt.Sprintf("Field[%s](%d): data: %#v", f.Type(), f.dataSize, f.data)
	return str
}

func (f *Field) Data() interface{} {
	if f == nil {
		return nil
	}
	if f.data == nil {
		return nil
	}
	return f.data
}

func (f *Field) readData(reader io.ReaderAt) error {

	sr := io.NewSectionReader(reader, f.off + fieldHeaderLen, f.Size() - fieldHeaderLen)

	f.dataBuf = make([]byte, f.dataSize)

	if n, err := io.ReadFull(sr, f.dataBuf); err != nil {
		err2 := fmt.Errorf("Unexpected EOF on f.readData() (%s). Read %d bytes, expected %d bytes.", f.Type(), n, len(f.dataBuf))
		return err2
	}

	// if type is HEDR, read struct
	data, err := f.getFieldStructure()
	if err != nil && err != ErrUnimplementedField {
		return err
	}
	//fmt.Printf("%s.%s: %#v\n", f.RecordType(), f.Type(), data)

	f.data = data

	if len(f.dataBuf) > 0 {
		// if its a skip field, skip processing of it
		if _, ok := f.data.(Skip); !ok {
			if _, ok := f.data.(Unknown); !ok {
				if err != nil && err != ErrUnimplementedField {
					err = fmt.Errorf("Leftover data for field read on %s.%s\nCollected: %#v\nRemaining: %#v", f.RecordType(), f.Type(), f.data, f.dataBuf)
					return err
				}
			}
		}
	}

	return nil
}

func (f *Field) getFieldStructure() (out interface{}, err error) {

	recordTypeStr := f.RecordType()
	fieldTypeStr := f.Type()

	zeroValue := FieldsStructLookup[recordTypeStr][fieldTypeStr]

	// if its a skip field, skip processing of it
	if _, ok := zeroValue.(Skip); ok{
		return SkipZero, nil
	}
	// if its a manually marked as unknown field, skip processing of it
	if _, ok := zeroValue.(Unknown); ok{
		return UnknownZero, nil
	}

	t := reflect.TypeOf(zeroValue)

	if t == nil {
		err = ErrUnimplementedField
		msg := fmt.Sprintf("### Unimplemented field %s.%s", recordTypeStr, fieldTypeStr)


		// dumping human readable dumps literally doubles the cpu time
		//msg := fmt.Sprintf("### Unimplemented field %s.%s: %s", recordTypeStr, fieldTypeStr, f.dataBuf.Human())
		LogUnimplementedField(recordTypeStr, fieldTypeStr, f.dataBuf)
		return msg, err
	}

	if t.Kind() == reflect.Func {
		// todo: allow the return of a reflect.Type instead of interface{} and have it just continue thru
		var is_value bool
		var is_type bool
		var args []reflect.Value

		switch zero_type := zeroValue.(type) {
		case func(readBuf, Field) interface{}:
			is_value = true
			args = []reflect.Value{reflect.ValueOf(f.dataBuf), reflect.ValueOf(*f)}
		case func(readBuf) interface{}:
			is_value = true
			args = []reflect.Value{reflect.ValueOf(f.dataBuf)}

		case func(readBuf, Record) interface{}:
			is_value = true
			args = []reflect.Value{reflect.ValueOf(f.dataBuf), reflect.ValueOf(*f.parentRecord)}

		// TYPES!
		case func(readBuf, Field) reflect.Type:
			is_type = true
			args = []reflect.Value{reflect.ValueOf(f.dataBuf), reflect.ValueOf(*f)}
		case func(readBuf) reflect.Type:
			is_type = true
			args = []reflect.Value{reflect.ValueOf(f.dataBuf)}

		case func(readBuf, Record) reflect.Type:
			is_type = true
			args = []reflect.Value{reflect.ValueOf(f.dataBuf), reflect.ValueOf(*f.parentRecord)}




		default:
			msg := fmt.Errorf("Unimplemented field.getFieldStructure() Func type for %s.%s: '%#v'", recordTypeStr, fieldTypeStr, zero_type)
			return nil, msg
		}

		if is_value {
			out := reflect.ValueOf(zeroValue).Call(args)[0].Interface()
			return out, nil
		}
		if is_type {
			t = reflect.ValueOf(zeroValue).Call(args)[0].Interface().(reflect.Type)
		}

		// otherwise, let it pass thru bc its a reflect.Type!
	}

	v := reflect.New(t)
	f.dataBuf.readType(t, v.Elem(), f)

	out_value := v.Elem().Interface()
	_ = out_value

	return out_value, nil
}

func (f *Field) Type() (string) {
	return fmt.Sprintf("%s", f._type)
}
func (f *Field) RecordType() (string) {
	return f.parentRecord.Type()
}












