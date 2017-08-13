package esm

import (
	"io"
	"fmt"
	"reflect"
)


type FieldHeader struct {
	_type char4
	dataSize uint16
}



type Field struct {
	FieldHeader
	record *Record
	readerAt io.ReaderAt
	readerSize int64
	//dataSectionReader *io.SectionReader
	sr *io.SectionReader
	off int64
	//headerOffset int64
	//data []byte
	//dataBuf []byte
	dataBuf readBuf
	data interface{}
}


// calculates the size of all the things (header and all data)
func (f *Field) Size() int64 {
	return int64(fieldHeaderLen) + int64(f.dataSize)
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

	f.dataSize = b.uint16()

	f.dataBuf = make([]byte, f.dataSize)

	//if _, err := io.ReadFull(r, f.dataBuf); err != nil {
	//	return err
	//}

	return nil
}



func (f *Field) String() string {
	//str := fmt.Sprintf("Field[%s](%d): buff: %s", f._type, f.dataSize, f.dataBuf)
	str := fmt.Sprintf("Field[%s](%d): data: %s", f.Type(), f.dataSize, f.data)
	return str
}

func (f *Field) readData(reader io.ReaderAt) error {

	sr := io.NewSectionReader(reader, f.off + fieldHeaderLen, f.Size() - fieldHeaderLen)
	_ = sr

	//rs := io.NewSectionReader(reader, f.off + fieldHeaderLen, int64(f.dataSize))
	//if _, err := f.sr.Seek(0, io.SeekStart); err != nil {
	//	return err
	//}

	//buf := bufio.NewReader(rs)


	if _, err := io.ReadFull(sr, f.dataBuf); err != nil {
		return err
	}

	// if type is HEDR, read struct
	data, err := f.getFieldStructure()
	if err != nil && err != ErrUnimplementedField {
		return err
	}
	fmt.Printf("%s.%s: %#v\n", f.RecordType(), f.Type(), data)

	f.data = data

	return nil
}

func (f *Field) getFieldStructure() (out interface{}, err error) {

	recordTypeStr := f.RecordType()
	fieldTypeStr := f.Type()

	zeroValue := FieldsStructLookup[recordTypeStr][fieldTypeStr]

	t := reflect.TypeOf(zeroValue)

	if t == nil {
		err = ErrUnimplementedField
		msg := fmt.Sprintf("### Unimplemented field %s.%s", recordTypeStr, fieldTypeStr)
		return msg, err
	}

	v := reflect.New(t)
	f.dataBuf.readType(t, v.Elem())

	return v.Elem(), nil
}

func (f *Field) Type() (string) {
	return fmt.Sprintf("%s", f._type)
}
func (f *Field) RecordType() (string) {
	return f.record.Type()
}












